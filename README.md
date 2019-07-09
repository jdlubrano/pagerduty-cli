# PagerDuty CLI

[![Build Status](https://travis-ci.org/jdlubrano/pagerduty-cli.svg?branch=master)](https://travis-ci.org/jdlubrano/pagerduty-cli)

A CLI for interacting with the PagerDuty API.

## Installation

**This tool assumes that you are a part of an account using
[Advanced Permissions](https://support.pagerduty.com/docs/advanced-permissions).**

1.  Acquire a PagerDuty [developer API token](https://support.pagerduty.com/docs/generating-api-keys#section-generating-a-personal-rest-api-key).
2.  Create a configuration file at `$HOME/.pagerduty.yml`.  The file should
have the following information:

```yaml
---
api-token: <your token>
```

3.  Download the [latest binary release](https://github.com/jdlubrano/pagerduty-cli/releases)
and place the binary somewhere in your `$PATH`.

4.  You should be able to run `pagerduty-cli` to see the available commands.

5.  You may wish to create an alias for the CLI as it's full name is admittedly
verbose.  You can add `alias pd="pagerduty-cli"` to your shell's profile if you
so choose.

### Upgrading

Upgrading is still a work in progress, so in the meantime, repeat step 3 from
above with the latest binary.  Eventually I would like this tool to upgrade
itself, but I haven't yet gotten that far.

## Usage

This tool, thanks to the awesome [Cobra package](https://github.com/spf13/cobra)
strives to be self-service and self-documenting.  You can run
`pagerduty-cli help` with no additional commands to see a list of available
commands.  You can see additional help for subcommands by running
`pagerduty-cli help <subcommand>`.

## Development

1.  Checkout this repository.
2.  Run `go build .`
3.  Run `./pagerduty-cli`
