# Recipe Web App [![Build Status](https://travis-ci.org/napalm684/recipes.svg?branch=master)](https://travis-ci.org/napalm684/recipes)

## Description
Just a simple Go (golang) web application for quick lookup of recipes by my wife.  Designing this to be performant/have a small footprint so I can host internally on my home network via Raspberry Pi.

## Install

```
dep ensure
go install ./...
```

## Execution

```
go run cmd/server.go
```
