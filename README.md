[![Build Image](https://github.com/bakito/ping/actions/workflows/publish.yml/badge.svg)](https://github.com/bakito/ping/actions/workflows/publish.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bakito/ping)](https://goreportcard.com/report/github.com/bakito/ping)

# Ping

A simple ping docker container.

# Configuration

| Env var name  | Description                                                                         |
|---------------|-------------------------------------------------------------------------------------|
| PING_TARGET   | Define the target address to be pinged, if not defined, the default gateway is used |   
| PING_INTERVAL | Set the ping interval (default 1s)                                                  |   


