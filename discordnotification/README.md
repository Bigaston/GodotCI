# DiscordNotification

A Go application that run on an Alpine based container that can send notification to your Discord when the build Success/Fail.  
Simply copy and paste this at the end of your Woodpecker CI workflow, this is required to have it 2 times, because Woodpecker don't have a system to let your next action now if the pipeline is a success or not.  
And add a [Discord Webhook](https://discord.com/developers/docs/resources/webhook) on a secret in Woodpecker CI.  

```YAML
  - name: Discord Notification Success
    image: ghcr.io/bigaston/discordnotification
    settings:
      result: success
      webhook:
        from_secret: WEBHOOK
    when: 
      - status: ["success"]
  - name: Discord Notification Failure
    image: ghcr.io/bigaston/discordnotification
    settings:
      result: failure
      webhook:
        from_secret: WEBHOOK
    when:
      - status: ["failure"]
```

## Settings
- result (success/failure)
- webhook: Discord Webhook
- matrix: Used when you execute your workflow on different system, give more informations in the message
- build_url: If you push your game somewhere, you can give a public URL here (like a file server)