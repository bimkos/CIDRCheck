# cidrCheck
Simple script for bulk check CIDR

# Install 
```
go get github.com/alexflint/go-arg 
```

# Test

Create file with content:
```
185.217.95.2
185.217.94.2
```
and run
```
go run main.go --file <filename> --ranges 185.217.94.1/24
```
```
go run main.go --file <filename> --ranges 185.217.94.1/24 185.217.95.2
```
