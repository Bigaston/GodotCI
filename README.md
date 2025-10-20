# GodotCI

A set of Docker Container usefull to automate build of Godot games with Woodpecker CI.  
I think it can be usefull for a lot of people so I publish them here. If you find these Docker usefull, feel free to ping me on [BlueSky](https://bsky.app/profile/bigaston.me), [check my website](https://bigaston.me) or [support me on Ko-Fi](https://ko-fi.com/bigaston).  

You can find some examples workflow [here](./examples/).

If you see this repository on Github, known that this is a mirror of the [Codeberg Repository](https://codeberg.org/Bigaston/GodotCI) !

## What is CI/CD?
CI and CD respectively mean Continuous Integration and Continuous Deployment. It's a way to execute a set of actions everytime an action occures. Like every day at 10AM, at every commit on main branch, on each tags... It's used in a lot of big game studio to create games build, execute tests, deploy server... It's removing all the pain of building the game and discover that you can't because reason. And combined with [Itch.io butler](https://itch.io/docs/butler/) and [Steam Cmd](https://developer.valvesoftware.com/wiki/SteamCMD) you can even push your game on Steam and Itch.io.

This set of container is aimed to be used with [Woodpecker CI](https://woodpecker-ci.org/), a lightway CI system that's work with Docker. You may have some basic Docker Container where you execute a list of commands, or a Docker that run a custom binarie and do a specific task (like the Discord Notification).

## How to use?
On each folder, you will find an other Docker image. You can create a new repository, copy the code, including the *.woodpecker* folder that will create automate the Image creation for you on your private Docker Repository. That's how I use it with Forgejo!

But to ease your process, I've published every Docker Container on Github (maybe on Codeberg to if I can have the authorization), so you will have some link to use that to. But it's better to host it yourself. You know, avoid private owned service...

## More documentation
Here some links to documentation that I use when I need to create some new CI for my Godot games.
- [Itch.io butler](https://itch.io/docs/butler/)
- [GameCI's Steam Deploy](https://github.com/game-ci/steam-deploy): I use this as an inspiration to understand how to use SteamCMD in CI
- [Godot CLI](https://docs.godotengine.org/en/stable/tutorials/editor/command_line_tutorial.html)

## Containers
- [discordnotification](./discordnotification/): Send a message to a Discord channel when the build is a success/fail via Discord Webhook
- [itchci](./itchci/): A container with itch.io butler already installed
- [bunnyci](./bunnyci/): A plugin to help you publish your game on Bunny.net
- [steamci](./steamci/): A plugin to help you publish your game on Steam from CI
- [godotci](./godotci/): Contains everything you need to export a Godot game (Godot and export templates in the good folder) for Windows/Linux/Web/Android