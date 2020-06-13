# mattermost-sandbox

## How to run Incoming Webhook

1. Start Mattermost server
    ```
    $ make serve
    ```
1. Open http://localhost:8065/
1. Add Incoming Webhook manually
    * see https://docs.mattermost.com/developer/webhooks-incoming.html
1. Change Integration Management setting via System Console
    ```
    Enable integrations to override usernames: true
    Enable integrations to override profile picture icons: true
    ```
1. Modify Incoming Webhook URL
    ```golang
    $ vi incoming/main.go
    ...
    const (
        WEBHOOK_URL = "http://localhost:8065/hooks/ftcxcu4aypfutqby89yhpe7kze"
    )
    ...
    ```
1. Run bot program
    ```
    $ go run incoming/main.go
    ```
