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

### read-notifications

It can optionally ignore certain reasons in certain repositories. To achieve this, set up the
environment variables:

- `RESTRICTED_REPOSITORIES_PATTERN`
- `RESTRICTED_REPOSITORIES_ALLOWED_REASONS`

These values will be interpreted as regular expressions.
