#!/bin/bash

cd "$( dirname "${BASH_SOURCE[0]}" )"
protoc --go_out=. --go_opt=paths=source_relative proto/*.proto
