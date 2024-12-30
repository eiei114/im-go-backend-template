## Go Backend Template

## Routes
- http://localhost:8080/swagger/index.html

## Commands
### Format
```bash
go fmt ./...
```

### Run
```bash
go run app/cmd/main.go
```

### Docker Up
```bash
docker compose -f app/compose.yml up -d
```

### Docker Down
```bash
docker compose -f app/compose.yml down
```

### Swagger Init
```bash
swag init -g app/cmd/main.go -o app/docs
```


