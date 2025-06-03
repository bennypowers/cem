#!/bin/bash
go test -json $argv ./... | go tool tparse -all
