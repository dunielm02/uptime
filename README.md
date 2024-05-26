# Uptime

Uptime is a tool used to monitor all kind of services and **notify the state** of every one of them in **real time** by many notifications channels as Slack, Google Workspace, Microsoft Teams, etc.

This tool can be hosted anywhere because it is a really light tool and can access to protected services via ssh.

- [How to Use](#how-to-use)
  - [Define your services in a config file](#define-your-services-in-a-config-file)
    - [Set up the database](#set-up-database)
      - [InfluxDB](#influxdb)
    - [Set up your services](#set-up-your-services)
      - [Http](#http)
      - [Tcp](#tcp)
      - [Ping](#ping)
    - [Notification channels](#notification-channels)
      - [Slack](#slack)
      - [Microsoft Teams](#microsoft-teams)
      - [Google Workspace](#google-workspace)
      - [Discord](#discord)
      - [Telegram](#telegram)
    - [Ssh tools](#ssh-tools)
      - [Port Forwarding](#portforwarding)
  - [Run the tool](#run-it)
- [Compile](#how-to-compile-it)
  - [Windows](#windows)
  - [Linux](#linux)
- [Contribute](#contribute)

## How to Use

### Define your services in a config file

To set up the services that the tool will be checking we use a **YML** configuration file.

A configuration file Should look like this:

```yaml
database:
  type: influxdb
  spec:
    token: DA8osWZB208kHI7CpQVn1Fz0E_gXsrTfk2oz0_U0XkfKNUwQ224cShAz2_nh_j84TgrvD5Y8mw8X1Ag4uHG7mg==
    url: http://localhost:8086
    org: myOrg
    bucket: firstBucket
    measurement: services
services:
  - name: My Http Service
    type: http
    timeout: 30
    waiting-time: 60
    inverted: false
    notification-channels: ["Telegram"]
    spec:
      url: http://www.google.com
      method: POST
      headers:
        content-type: application/json
        authentication: bearer 
      expected-status: 201
      body: '{"ID": 4524}'
  - name: My Tcp Service
    type: tcp
    timeout: 30
    waiting-time: 60
    inverted: false
    notification-channels: ["Telegram"]
    spec:
      host: www.google.com
      port: 3000
  - name: My Ping Service
    type: ping
    timeout: 30
    waiting-time: 60
    inverted: false
    notification-channels: ["Telegram"]
    spec:
      host: www.google.com
      ping-count: 4
      must-receive: 4
port-forward:
  - server-address: www.host.com
    remote-address: localhost:5432
    local-address: localhost:8000
    username: <YourUsername>
    password: <YourPassword>
notification-channels:
  - name: Telegram
    type: telegram
    spec:
      token:
      chat-id:
```

Lets explain all of this.

#### Set up database

##### Influxdb

```yaml
database:
  # Database Type
  type: influxdb
  # Specifications of this database Type
  spec:
    # Influxdb authentication token
    token: DA8osWZB208kHI7CpQVn1Fz0E_gXsrTfk2oz0_U0XkfKNUwQ224cShAz2_nh_j84TgrvD5Y8mw8X1Ag4uHG7mg==
    # Database Url
    url: http://localhost:8086
    # Influxdb Organization
    org: myOrg
    # Influxdb Bucket
    bucket: firstBucket
    # Influxdb Measurement
    measurement: services
```

#### Set up your services

##### Http

```yaml
services:
  - name: My Http Service
    type: http
    # Specifies the timeout of the request
    timeout: 30
    # Time between every request
    waiting-time: 60
    # If the request is successful then the service is dead
    inverted: false
    # Name of the channel that it will notify through
    notification-channels: ["Telegram"]
    spec:
      # Http Request specifications
      url: http://www.google.com
      method: POST
      headers:
        content-type: application/json
        Authorization: <auth-scheme> <authorization-parameters>
      # Expected status code of the response
      expected-status: 201
      body: '{"ID": 4524}'
```

##### Tcp

```yaml
services:
  - name: My Tcp Service
    type: tcp
    # Tcp dial timeout
    timeout: 30
    # Time between each request
    waiting-time: 60
    # If the request is successful then the service is dead
    inverted: false
    # Name of the channel that it will notify through
    notification-channels: ["Telegram"]
    spec:
      host: www.google.com
      port: 80
```

##### Ping

```yaml
services:
  - name: My Ping Service
    type: ping
    # Timeout of the ping request
    timeout: 30
    # Time between each ping request
    waiting-time: 60
    # If the request is successful then the service is dead
    inverted: false
    # Name of the channel that it will notify through
    notification-channels: ["Telegram"]
    spec:
      # Host of the ping request
      host: www.google.com
      # How many packages will be send
      ping-count: 4
      # How many packages must receive to be alive
      must-receive: 4
```

#### Notification Channels

##### Slack

[Sending messages to Slack using incoming webhooks](https://api.slack.com/messaging/webhooks)

```yaml
notification-channels:
  - name: Slack
    type: slack
    spec:
      webhook: https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
```

##### Microsoft Teams

[Sending messages to Teams using incoming webhooks](https://learn.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook?tabs=newteams%2Cdotnet)

```yaml
notification-channels:
  - name: Microsoft Teams
    type: microsoft-teams
    spec:
      webhook: https://xxxxx.webhook.office.com/xxxxxxxxx
```

##### Google Workspace

[Sending messages to Google Workspace using incoming webhooks](https://developers.google.com/workspace/chat/quickstart/webhooks?hl=es-419)

```yaml
notification-channels:
  - name: Google Workspace
    type: google-workspace
    spec:
      webhook: https://chat.googleapis.com/v1/spaces/SPACE_ID/messages
```

##### Discord

[Sending messages to Discord using incoming webhooks](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks)

```yaml
notification-channels:
  - name: Discord
    type: discord
    spec:
      webhook: https://discord.com/api/webhooks/000000/XXXXXXXXXXXXX
```

##### Telegram

For telegram you need a bot-token and a chat id. [How to get it?](https://gist.github.com/nafiesl/4ad622f344cd1dc3bb1ecbe468ff9f8a)

```yaml
notification-channels:
  - name: Telegram
    type: telegram
    spec:
      # Token of the bot
      token: 63xxxxxx71:AAFoxxxxn0hwA-2TVSxxxNf4c
      # Chat id that you want to notify to
      chat-id: "827317315"
```

#### Ssh Tools

##### PortForwarding

We this we are capable of forward a remote port through ssh.

```yaml
port-forward:
    # Ssh Host
  - server-address: www.host.com
    # remote address to access
    remote-address: localhost:5432
    # Local address 
    local-address: localhost:8000
    # External host Username
    username: <YourUsername>
    # External host Password
    password: <YourPassword>
```

### Run it

To use this tool you can [download the binary file]() or [build it by yourself](#how-to-compile-it). Then Execute the command bellow to run it:

```bash
uptime -f uptime-config.yml
```

## How to Compile it

### Windows

```bash
git clone https://github.com/dunielm02/uptime.git
cd uptime
go mod download
env GOOS=windows GOARCH=amd64 go build -o ./build/uptime.exe
./build/uptime.exe -f uptime-config.yml
```

### Linux

```bash
git clone https://github.com/dunielm02/uptime.git
cd uptime
go mod download
env GOOS=linux GOARCH=arm go build -o ./build/uptime.exe
./build/uptime.exe -f uptime-config.yml
```

## Contribute
