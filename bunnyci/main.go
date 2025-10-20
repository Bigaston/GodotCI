package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

var client *bunnystorage.Client

func main() {
	cfg := &bunnystorage.Config{
		StorageZone: getRequiredEnv("PLUGIN_BUNNY_STORAGE_ZONE"),
		Key:         getRequiredEnv("PLUGIN_BUNNY_STORAGE_KEY"),
		Endpoint:    bunnystorage.EndpointFalkenstein,
	}

	var err error
	client, err = bunnystorage.NewClient(cfg)
	check(err)

	if getEnvOrDefault("PLUGIN_BUNNY_CLEAN_STORAGE", "false") == "true" {
		exploreAndDelete(getEnvOrDefault("PLUGIN_BUNNY_PATH", ""))
	}

	if getEnvOrDefault("PLUGIN_ZIP", "false") == "true" {
		zipFolder(path.Join(getRequiredEnv("CI_WORKSPACE"), getRequiredEnv("PLUGIN_PATH")), "temp.zip")

		uploadFile("temp.zip", getEnvOrDefault("PLUGIN_BUNNY_PATH", ""), getRequiredEnv("PLUGIN_BUNNY_FILENAME"))

		fmt.Println("File uploaded!")
	} else {
		exploreAndUpload(path.Join(getRequiredEnv("CI_WORKSPACE"), getRequiredEnv("PLUGIN_PATH")), getEnvOrDefault("PLUGIN_BUNNY_PATH", ""))
	}

	if getEnvOrDefault("PLUGIN_BUNNY_CLEAR_CACHE", "false") != "true" {
		os.Exit(0)
	}

	url := fmt.Sprintf("https://api.bunny.net/pullzone/%s/purgeCache", getRequiredEnv("PLUGIN_BUNNY_PULL_ZONE"))

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("AccessKey", getRequiredEnv("PLUGIN_BUNNY_PULL_ZONE_KEY"))

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 204 {
		panic(res.Status)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	fmt.Println("Cache cleared!")
}

func exploreAndUpload(path string, bunnyPath string) {
	files, err := os.ReadDir(path)
	check(err)

	for _, file := range files {
		if file.IsDir() {
			exploreAndUpload(path+"/"+file.Name(), bunnyPath+"/"+file.Name())
		} else {
			uploadFile(path+"/"+file.Name(), bunnyPath, file.Name())
		}
	}
}

func exploreAndDelete(path string) {
	fmt.Printf("Exploring %s...\n", path)

	objects, _, err := client.List(context.Background(), path)
	check(err)

	fmt.Printf("Found %d objects\n", len(objects))

	for _, obj := range objects {
		if obj.IsDirectory {
			exploreAndDelete(path + "/" + obj.ObjectName)
			fmt.Printf("Deleting %s...\n", obj.ObjectName)
			client.Delete(context.Background(), path, obj.ObjectName)
		} else {
			fmt.Printf("Deleting %s...\n", obj.ObjectName)
			client.Delete(context.Background(), path, obj.ObjectName)
		}
	}
}

func uploadFile(path string, bunnyPath string, filename string) {
	fmt.Printf("Upload %s...\n", path)

	file, err := os.Open(path)
	check(err)

	_, err = client.Upload(context.Background(), bunnyPath, filename, "", file)
	check(err)

	fmt.Println("Uploaded!")
}

func zipFolder(source, target string) error {
	file, err := os.Create(target)
	if err != nil {
		return err
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	baseFolder := filepath.Base(source)

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == source {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.Join(baseFolder, strings.TrimPrefix(path, source))

		if info.IsDir() {
			header.Name += "/"
			_, err = zipWriter.CreateHeader(header)
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)
		return err
	})
}

func getRequiredEnv(envName string) string {
	result, isPresent := os.LookupEnv(envName)

	if !isPresent {
		panic(fmt.Sprintf("Variable %s is not present but required", envName))
	}

	return result
}

func getEnvOrDefault(envName string, def string) string {
	result, isPresent := os.LookupEnv(envName)

	if !isPresent {
		return def
	}

	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
