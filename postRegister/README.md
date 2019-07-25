# PostRegister Service

This is the PostRegister service

Generated with

```
micro new micro/postRegister --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.postRegister
- Type: srv
- Alias: postRegister

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./postRegister-srv
```

Build a docker image
```
make docker
```