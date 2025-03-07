# protobuf example

https://protobuf.dev/getting-started/gotutorial/

## Instalation

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc -I="." --go_out=. --proto_path "." .\address_book.proto
go get google.golang.org/protobuf
```

## twirp example

```bash
protoc -I="." --go_out=. --twirp_out=. --proto_path "." .\twirp_example.proto
```