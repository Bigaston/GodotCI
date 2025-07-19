# GodotCI
```
ghcr.io/bigaston/godotci:4.3-stable
ghcr.io/bigaston/godotci:4.4.1-stable
ghcr.io/bigaston/godotci:4.5-beta3
ghcr.io/bigaston/godotci:android-4.3-stable
ghcr.io/bigaston/godotci:android-4.4.1-stable
ghcr.io/bigaston/godotci:android-4.5-beta3
```

This is the big container. In this image, you will have everything you need to export Godot games for Windows/Linux/Web and Android (on the android version.)  
It's an ubuntu based image where Godot is installed at `/godot` and can be accessed with `godot` in headless mode.  
**/!\ With Godot version prior to 4.5, you may have some random build error when you are building games with custom fonts. /!\**  

You will have multiple variant of this image, for Godot **4.3**, **4.4.1** and **4.5-beta3**.  
You can find the correct image with the image labels **4.3-stable**, **4.4.1-stable** and **4.5-beta3**.  

Random info that can be usefull. The file `/usr/bin/godot.version` will contain the Godot Version of this image.  

Take a look at the [Godot CLI](https://docs.godotengine.org/en/stable/tutorials/editor/command_line_tutorial.html) to see every commands available and what you can do with it!

## Android variant
The android image is bigger because all the tools you will need to create and APK are also included (OpenJDK-17 and Android SDK). But it can be usefull for exporting APK and games for VR Headset like Meta Quest 2.  
A debug.keystore is included, but please replace it by a new keystore that you will create!  

## Example
Just basic build. The game must have export template called Windows, Linux and Web, with build folder `.build/{TARGET}`. But feel free to change it of course!
```yml
when:
  - event: push
    branch: main

matrix:
  TARGET:
    - Windows
    - Linux
    - Web

steps:
  - name: Build the game for ${TARGET}
    image: ghcr.io/bigaston/godotci:4.5-beta3 # You can change the Godot version here
    commands:
      - godot --version

      # This part is not required. But it will create a commit.txt file
      # that you can include on your game export, and read from it
      # to display your commit SHA inside of the menu
      - "> commit.txt"
      - short_commit=$(git rev-parse --short HEAD)
      - echo "$short_commit" > commit.txt

      # Maybe not required now
      - godot --headless --editor --import

      - mkdir -p .build/${TARGET}

      # You can replace --export-release with --export-debug
      - godot --headless --export-release ${TARGET}
```

A basic example with Android build.
```yml
when:
  - event: push
    branch: main

steps:
  - name: Build the game for Android
    image: ghcr.io/bigaston/godotci:android-4.5-beta3
    commands:
      - godot --version

      - godot --headless --editor --import

      - mkdir -p .build/Android

      # Because sometimes you will need the android folder when you have custom Java,
      # this is an example of the command you can use to create it from the export templates
      - cat /usr/bin/godot.version
      - mkdir -p android/build
      - unzip ~/.local/share/godot/export_templates/$(cat /usr/bin/godot.version)/android_source.zip -d android/build
      - touch android/build/.gdignore
      - cat /usr/bin/godot.version > android/.build_version
 
# TODO: Find a way to export as release and not debug :D
      - godot --headless --export-debug Android
```