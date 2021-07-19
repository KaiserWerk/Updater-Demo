go build -o "data/app_v1.0.0.exe" -ldflags "-s -w -X 'main.version=1.0.0'" "cmd/app/main.go"
go build -o "data/app_v1.0.1.exe" -ldflags "-s -w -X 'main.version=1.0.1'" "cmd/app/main.go"
go build -o "data/app_v1.0.2.exe" -ldflags "-s -w -X 'main.version=1.0.2'" "cmd/app/main.go"
go build -o "data/app_v1.0.3.exe" -ldflags "-s -w -X 'main.version=1.0.3'" "cmd/app/main.go"

go build -o "data/launcher_v1.0.0.exe" -ldflags "-s -w -X 'main.version=1.0.0'" "cmd/launcher/main.go"
go build -o "data/launcher_v1.0.1.exe" -ldflags "-s -w -X 'main.version=1.0.1'" "cmd/launcher/main.go"
go build -o "data/launcher_v1.0.2.exe" -ldflags "-s -w -X 'main.version=1.0.2'" "cmd/launcher/main.go"
go build -o "data/launcher_v1.0.3.exe" -ldflags "-s -w -X 'main.version=1.0.3'" "cmd/launcher/main.go"