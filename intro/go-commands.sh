# initialize go application
go mod init github.com/rjdecilos/Learn-GoLang/intro/app
# run file
go run app.go
# build app directory then run build file
go build
./app
# install application in your OS (install binary module in $HOME/go/bin/app, or under windows, %USERPROFILE%\go\bin\app.exe)
# very useful if you want to build a CLI tool
go install github.com/rjdecilos/Learn-GoLang/intro/app
app