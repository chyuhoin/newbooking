go env -w CGO_ENABLED=0
go env -w GOOS=linux
go build -o ./target/server-0.0.1
copy .\database_config.json target\
go env -w CGO_ENABLED=1
go env -w GOOS=windows
