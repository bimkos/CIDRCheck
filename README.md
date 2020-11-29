# cidrCheck
Simple script for bulk check CIDR

# Install 
```
go get github.com/alexflint/go-arg 
```

# Test

Crate file with content:
```
185.217.95.2
185.217.94.2
```
and run
```
go run main.go --file <filename> --ranges 185.217.94.1/24
```