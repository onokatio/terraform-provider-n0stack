#!/bin/bash

GOARCH=amd64

GOOS=darwin
go build -o terraform-provider-n0stack-$GOOS-$GOARCH

GOOS=linux
go build -o terraform-provider-n0stack-$GOOS-$GOARCH

GOOS=windows
go build -o terraform-provider-n0stack-$GOOS-$GOARCH
