# PagerDuty CLI

A CLI for interacting with the PagerDuty API.

## Installation

**This tool assumes that you are a part of an account using
[Advanced Permissions](https://support.pagerduty.com/docs/advanced-permissions).**

1.  Acquire a PagerDuty [developer API token](https://support.pagerduty.com/docs/generating-api-keys#section-generating-a-personal-rest-api-key).
2.  Create a configuration file at `$HOME/pagerduty.yml`.  The file should
have the following information:

```yaml
---
api-token: <your token>
```

## Development

1.  Checkout this repository.
2.  Run `go build .`
3.  Run `./pagerduty-cli`
