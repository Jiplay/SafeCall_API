#!/bin/sh

go mod download
go build -o api *.go