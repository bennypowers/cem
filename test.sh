#!/bin/bash
go test -json ./... "$@" | go tool tparse -all
