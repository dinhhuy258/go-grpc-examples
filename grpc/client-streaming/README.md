# Sum

1. Generate proto

```
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative sum.proto --go-grpc_out=proto --go-grpc_opt=paths=source_relative
```

2. Run the server

```
go run server/main.go
```

3. Run the client

```
go run client/main.go
```
