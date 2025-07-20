package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path"
	"strconv"
)

func main() {
	vdfData := VdfTemplate{
		AppId:       getRequiredEnv("PLUGIN_APP_ID"),
		ContentRoot: getRequiredEnv("PLUGIN_CONTENT_ROOT"),
	}

	description, hasDescription := os.LookupEnv("PLUGIN_DESCRIPTION")

	if hasDescription {
		vdfData.Description = description
	} else {
		vdfData.Description = fmt.Sprintf("Pushed from CI at commit %s", os.Getenv("CI_COMMIT_SHA"))
	}

	setLive, hasSetLive := os.LookupEnv("PLUGIN_SET_LIVE")

	if hasSetLive {
		vdfData.HasSetLive = true
		vdfData.SetLive = setLive
	}

	// Get Depot Id
	depotId := getRequiredEnv("PLUGIN_DEPOT_ID")
	parsed := make(map[string]json.RawMessage)
	err := json.Unmarshal([]byte(depotId), &parsed)

	if err == nil {
		matrix := getRequiredEnv("PLUGIN_MATRIX")

		var depot_id int
		parsedValue, ok := parsed[matrix]

		if !ok {
			panic(fmt.Sprintf("You have specified %s in your matrix, but field doesn't exist in depot_id", matrix))
		}

		err := json.Unmarshal(parsedValue, &depot_id)
		check(err)

		fmt.Printf("Found DepotID %d for matrix %s\n", depot_id, matrix)

		vdfData.DepotId = strconv.Itoa(depot_id)
	} else {
		vdfData.DepotId = depotId
	}

	// Get Local Path
	localPath := getRequiredEnv("PLUGIN_LOCAL_PATH")
	parsed = make(map[string]json.RawMessage)
	err = json.Unmarshal([]byte(localPath), &parsed)

	if err == nil {
		matrix := getRequiredEnv("PLUGIN_MATRIX")

		var local_path string
		parsedValue, ok := parsed[matrix]

		if !ok {
			panic(fmt.Sprintf("You have specified %s in your matrix, but field doesn't exist in local_path", matrix))
		}

		err := json.Unmarshal(parsedValue, &local_path)
		check(err)

		fmt.Printf("Found LocalPath %s for matrix %s\n", local_path, matrix)

		vdfData.LocalPath = local_path
	} else {
		vdfData.LocalPath = localPath
	}

	tmpl, err := template.New("vdf.tmpl").ParseFS(templates, "templates/vdf.tmpl")
	check(err)

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, vdfData)
	check(err)

	err = os.WriteFile(path.Join(os.Getenv("CI_WORKSPACE"), "__upload__.vdf"), buf.Bytes(), os.ModeAppend)
	check(err)

	fmt.Println("Uploading on Steam using VDF:")
	fmt.Println(buf.String())

	// Start the Steam Command
	steamAuthVdf := getRequiredEnv("PLUGIN_STEAM_AUTH_VDF")
	base64Decoded, hasError := base64Decode(steamAuthVdf)

	if hasError {
		panic("Error while decoding STEAM_AUTH_VDF base64")
	}

	err = os.WriteFile("/root/Steam/config/config.vdf", base64Decoded, 0777)
	check(err)

	steamUsername := getRequiredEnv("PLUGIN_STEAM_USERNAME")

	err = runSteamCommand([]string{"+login", steamUsername, "+quit"})
	check(err)

	err = runSteamCommand([]string{"+login", steamUsername, "+run_app_build", path.Join(os.Getenv("CI_WORKSPACE"), "__upload__.vdf"), "+quit"})
	check(err)

	fmt.Println("Everything went fine!")
}

func getRequiredEnv(envName string) string {
	result, isPresent := os.LookupEnv(envName)

	if !isPresent {
		panic(fmt.Sprintf("Variable %s is not present but required", envName))
	}

	return result
}

func base64Decode(str string) ([]byte, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return []byte{}, true
	}
	return data, false
}

func runSteamCommand(params []string) error {
	cmd := exec.Command("/home/steam/steamcmd.sh", params...)

	cmd.Dir = "/woodpecker"
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
