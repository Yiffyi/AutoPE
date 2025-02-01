cd /d %~dp0
@REM set GOOS=linux
go build -ldflags "-s -w" -x -o bin\playground.exe exe\playground\playground.go
go build -ldflags "-s -w" -x -o bin\autope.exe exe\autope\autope.go
pause