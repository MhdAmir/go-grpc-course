# Linux Step-by-Step Guide for gRPC with Go
## Prerequisites: Install Required Go Packages

Run the following commands to install the necessary Go tools:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Update Your Shell Configuration

Add the following line to your .zshrc or .bashrc to ensure Go binaries are accessible from your PATH:

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

After adding, don't forget to reload the shell configuration:
```
source ~/.zshrc   # For zsh users
# OR
source ~/.bashrc  # For bash users
```

## How to Compile the greet Package

Use the make command to compile the greet package:
```
make greet
```

## Fixing Issues with go.sum or go.mod

If you encounter problems with go.sum or go.mod, clean up and update dependencies using:
```
go mod tidy
```