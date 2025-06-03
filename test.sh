#!/bin/bash
go test -json ./... | tparse -all
