# Fake Vault API Service

This is a fake Vault API service for a honeypot experiment

## Usage
```go
go run .
```

### Example API call

```shell
curl -X POST -H 'Content-Type: application/json' \
  -d '{"password": "password"}' \
  http://localhost:820/v1/auth/userpass/login/rob -v

```