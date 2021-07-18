go build -o "app.exe" -ldflags "-s -w -X 'main.version=1.0.0'" "cmd/app/main.go"
go build -o "launcher.exe" -ldflags "-s -w -X 'main.version=1.0.0'" "cmd/launcher/main.go"
go build -o "update-server.exe" -ldflags "-s -w" "cmd/update-server/main.go"