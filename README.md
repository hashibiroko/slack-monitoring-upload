# slack-monitoring-upload

- [Docker Hub](https://hub.docker.com/r/hashibiroko/slack-monitoring-upload/)

## Install

require Golang environment and setup GOPATH.

```
$ go get github.com/hashibiroko/slack-monitoring-upload
```

## Usage

Example 1:

```
$ slack-monitoring-upload -token=xxxxxx-xxxxxxxxx -channel=random
```

Example 2: setting environment

```
$ export SLACK_USER_TOKEN="xxxxxx-xxxxxxxxx"
$ export SLACK_CHANNEL_NAME="random"
$ slack-monitoring-upload
```

Example 3: using docker

```
$ docker run -itd --name slack-monitoring-upload -e SLACK_USER_TOKEN=xxxxxx-xxxxxxxxx -e SLACK_CHANNEL_NAME=random hashibiroko/slack-monitoring-upload
```

### Flags

| name | description | default | require |
| :--- | :---------- | :-----: | :-----: |
| token | Set Slack User Token |  | true |
| channel | Set the channel name on which the message is posted | general |  |

### Environments

| name | description |
| :--- | :---------- |
| SLACK_USER_TOKEN | User Slack Token |
| SLACK_CHANNEL_NAME | Slack Channel Name |