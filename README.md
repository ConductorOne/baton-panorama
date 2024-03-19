![Baton Logo](./docs/images/baton-logo.png)

# `baton-panorama` [![Go Reference](https://pkg.go.dev/badge/github.com/conductorone/baton-panorama.svg)](https://pkg.go.dev/github.com/conductorone/baton-panorama) ![main ci](https://github.com/conductorone/baton-panorama/actions/workflows/main.yaml/badge.svg)

`baton-panorama` is a connector for Baton built using the [Baton SDK](https://github.com/conductorone/baton-sdk). It works with Panorama XML API.

Check out [Baton](https://github.com/conductorone/baton) to learn more about the project in general.

# Prerequisites

Connector requires credentials (username and password) that are used throughout the communication with API. Credentials are the same that you use to log in to Panorama.

Passing credentials to connector can be done by setting `BATON_USERNAME` and `BATON_PASSWORD` or by passing `--username` and `--password`.

Also you can set up host of Panorama. It's can be done by `BATON_PANORAMA_URL` or by passing `--panorama-url`. 

Important to mention that API has to be enabled for user. [See more](https://docs.paloaltonetworks.com/pan-os/10-2/pan-os-panorama-api/get-started-with-the-pan-os-xml-api/enable-api-access).

# Getting Started

## brew

```
brew install conductorone/baton/baton conductorone/baton/baton-panorama

BATON_USERNAME=panorama_user BATON_PASSWORD=pass BATON_PANORAMA_URL=https://123.eu-central-1.compute.amazonaws.com baton-panorama
baton resources
```

## docker

```
docker run --rm -v $(pwd):/out -e BATON_USERNAME=panorama_user BATON_PASSWORD=pass BATON_PANORAMA_URL=https://123.eu-central-1.compute.amazonaws.com ghcr.io/conductorone/baton-panorama:latest -f "/out/sync.c1z"
docker run --rm -v $(pwd):/out ghcr.io/conductorone/baton:latest -f "/out/sync.c1z" resources
```

## source

```
go install github.com/conductorone/baton/cmd/baton@main
go install github.com/conductorone/baton-panorama/cmd/baton-panorama@main

BATON_USERNAME=panorama_user BATON_PASSWORD=pass BATON_PANORAMA_URL=https://123.eu-central-1.compute.amazonaws.com baton-panorama
baton resources
```

# Data Model

`baton-panorama` will fetch information about the following Baton resources:

- Users
- Users groups

# Contributing, Support and Issues

We started Baton because we were tired of taking screenshots and manually building spreadsheets. We welcome contributions, and ideas, no matter how small -- our goal is to make identity and permissions sprawl less painful for everyone. If you have questions, problems, or ideas: Please open a Github Issue!

See [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more details.

# `baton-panorama` Command Line Usage

```
baton-panorama

Usage:
  baton-panorama [flags]
  baton-panorama [command]

Available Commands:
  capabilities       Get connector capabilities
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command

Flags:
      --client-id string         The client ID used to authenticate with ConductorOne ($BATON_CLIENT_ID)
      --client-secret string     The client secret used to authenticate with ConductorOne ($BATON_CLIENT_SECRET)
  -f, --file string              The path to the c1z file to sync with ($BATON_FILE) (default "sync.c1z")
  -h, --help                     help for baton-panorama
      --ignore-bad-certificate   Ignore bad certificate. !This should be used only for testing purposes.
      --log-format string        The output format for logs: json, console ($BATON_LOG_FORMAT) (default "json")
      --log-level string         The log level: debug, info, warn, error ($BATON_LOG_LEVEL) (default "info")
      --panorama-url string      Url of Panorama instance
      --password string          Password
  -p, --provisioning             This must be set in order for provisioning actions to be enabled. ($BATON_PROVISIONING)
      --username string          Username
  -v, --version                  version for baton-panorama

Use "baton-panorama [command] --help" for more information about a command.
```