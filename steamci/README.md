# SteamCI
```
ghcr.io/bigaston/steamci
```

This is a all contained Woodpecker CI plugins that will generate the .vdf of Steam for you, and then upload your game on the specified depot.

## Settings
- steam_username: The username linked with the auth vdf
- steam_auth_vdf: (More doc coming) A base64 encoded string that contains all the content of config.vdf that you will get when you execute the local +login command. A bit clanky because Steam don't provide easy access token, so you have to login on local to pass SteamGuard, and then copy the content of the file
- app_id: Your Steam App Id
- content_root: The root of the published folder
- depot_id: Can be a single int or a mapping from matrix:depot_id
- local_path: Where your builded game finish. Can be a single string or a mapping from matrix:local_path
- matrix: A parameter to map your depot_id and local_path (Like Windows, Linux...)
- set_live: Optionnal, deploy the build on a Steam branch. Can't be *default* due to security reason
- description: Optionnal, a description that you may find on your Steamworks Dashboard (Default to the commit SHA)

## Example
```yml
when:
  - event: push
    branch: main

matrix:
  TARGET:
    - Windows
    - Linux

steps:
  - name: Build the game for ${TARGET}
    image: ghcr.io/bigaston/godotci:4.5-beta3
    commands:
      - godot --version
      - godot --headless --editor --import
      - mkdir -p .build/${TARGET}
      - godot --headless --export-release ${TARGET}

  - name: Upload Steam
    image: ghcr.io/bigaston/steamci
    settings:
      steam_username: something_deploy
      steam_auth_vdf:
        from_secret: STEAM_AUTH_VDF
      app_id: 000
      content_root: ".build"
      depot_id: 
        Windows: 000
        Linux: 001
      local_path: 
        Windows: "Windows/*"
        Linux: "Linux/*"
      matrix: ${TARGET}
      set_live: "beta"
```