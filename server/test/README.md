# Test
### One-click Execute All Tests and Generate Reports
```shell
go test -coverprofile=coverage.out ../...
go tool cover -html=coverage.out -o coverage.local.html
```
