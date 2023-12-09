#!/bin/bash
cwd=$(pwd)

sh $cwd/scripts/build_tailwind.sh
echo "Running web server"
go run $cwd/cmd/web_server.go
