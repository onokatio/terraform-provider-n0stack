#!/bin/bash

GOARCH=amd64
GOOS=darwin
go build -o terraform-provider-n0stack-$GOOS-$GOARCH

GOARCH=amd64
GOOS=linux
go build -o terraform-provider-n0stack-$GOOS-$GOARCH

GOARCH=386
GOOS=linux
go build -o terraform-provider-n0stack-$GOOS-$GOARCH

GOARCH=amd64
GOOS=windows
go build -o terraform-provider-n0stack-$GOOS-$GOARCH.exe

GOARCH=386
GOOS=windows
go build -o terraform-provider-n0stack-$GOOS-$GOARCH.exe
