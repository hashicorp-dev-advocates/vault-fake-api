# Fake Vault API Service

This is a fake Vault API service for a honeypot experiment

## Usage
```go
go run .
```

### Example API call

```shell
curl -X POST -H 'Content-Type: application/json' \
  -d '{"username": "rob", "password": "password"}' \
  localhost:820/login -v

```