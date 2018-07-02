# MongoDB Stitch CLI

## Documentation

https://docs.mongodb.com/stitch/import-export/stitch-cli-reference/

## Linting

provided by gometalinter

```go
gometalinter --exclude=vendor --vendor --config=.linter.config ./...
```

## Testing

Run all tests:

```go
go test -v $(go list github.com/10gen/stitch-cli/...)
```
