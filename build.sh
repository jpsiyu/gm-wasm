#!/bin/bash

GOARCH=wasm GOOS=js go build -o main.wasm .