@echo off
set GOPATH=%~dp0
cd /d %~dp0
go test censorship
