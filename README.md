# slack-monitoring-upload

## Install

require Golang environment and setup GOPATH.

```
$ go get github.com/hashibiroko/slack-monitoring-upload
```

## Usage

Example:

```
$ slack-monitoring-upload -token=xxxxxx-xxxxxxxxx channel=random
```

### Flags

| name | description | default | require |
| :--- | :---------- | :-----: | :-----: |
| token | Set Slack Bot Token |  | true |
| channel | Set the channel name on which the message is posted | general |  |
