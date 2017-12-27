#! /bin/bash
# Defines the operational system and the processor architecture
GOOS=windows GOARCH=386 \
# -o Allows to set a different name to the builded app
#   Default: Same name of project folder
# -v Verbose mode
go build -o app_win.exe -v github.com/adrianuf22/learn_golang/build # The path at the end defines the project path, useful when the build occurs in different directory
