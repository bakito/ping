[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/bakito/ping/Build%20Image?logo=github)](https://github.com/bakito/ping/actions/workflows/publish.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bakito/ping)](https://goreportcard.com/report/github.com/bakito/ping)
[![Releases](https://img.shields.io/github/v/release/bakito/ping?label=Release)](https://github.com/bakito/ping/releases)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
# Ping

A simple ping docker container.

# Configuration

| Env var name  | Description                                                                         |
|---------------|-------------------------------------------------------------------------------------|
| PING_TARGET   | Define the target address to be pinged, if not defined, the default gateway is used |   
| PING_INTERVAL | Set the ping interval (default 1s)                                                  |   


