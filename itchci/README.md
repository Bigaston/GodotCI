# ItchCI
```
ghcr.io/bigaston/itchci:latest
```

A simple Debian image with butler already installed.  
The Butler executable is located at /bin/butler and can be accessed with `butler`.  
For information, if you include some OS name after the : at username/game, Itch.io will automaticaly tag them on your page.

- [Butler Documentation](https://itch.io/docs/butler/)

## Example:
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

  - name: Upload Itchio
    image: ghcr.io/bigaston/itchci:latest
    commands:
      - butler -V
      - butler push .build/${TARGET} username/game:${TARGET}
    environment:
      BUTLER_API_KEY:
        from_secret: BUTLER_API_KEY
```