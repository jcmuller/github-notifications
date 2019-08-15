# github-notifications

CLI tool that shows github notifications and another one that allows you
to mark them read.

## Installation
```bash
$ go get github.com/jcmuller/github-notifications/cmd/read-notifications
$ go get github.com/jcmuller/github-notifications/cmd/mark-notifications-read
```

## Usage
```bash
$ read-notifications

```

```bash
$ mark-notifications-read [-duration 24h] [-timestamp 2019-01-01] [-help]

```
The default is `24h`.
