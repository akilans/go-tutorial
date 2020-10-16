# Golang

Golang is a statically typed, high performance, and simple language

### Features

- Simplicity
- In-built Concurrency
- Speed of Compilation
- Staticaly typed
- Garbage collection

### Environment Setup

- There is some specific way to organize your go projects.Follow the folder structure as below
- Install Go and set GOPATH
- Create folder[go-lib] to keep external packages. Create subfolders [bin,pkg] folders
- Create folder[go-workspace] to keep your go projects. Create subfolders [bin,pkg,src] folders. Under src folder create folders like below [ Replace with your github account]

```
go-workspace/
├── bin
├── pkg
└── src
    └── github.com
        └── akilans
            └── go-tutorial
                ├── 01-hello
                │   └── main.go
                └── readme.md

```

```bash
tar -C /usr/local -xzf go1.15.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
export GOPATH=/home/akilan/go-lib
export PATH=$PATH:$GOPATH/bin
export GOPATH=$GOPATH:/home/akilan/go-workspace
go version
```

### Execution

There are 3 ways to run/compile your go application

```bash
# Without creating binary it executes
go run /home/akilan/go-workspace/src/github.com/akilans/go-tutorial/01-hello/main.go
# Creates binary under root directory [bin,src,pkg,$BINARY]
go build /home/akilan/go-workspace/src/github.com/akilans/go-tutorial/01-hello/
/home/akilan/go-workspace/01-hello
# Creates binary under bin directory [bin/$BINARY,src,pkg]
go install /home/akilan/go-workspace/src/github.com/akilans/go-tutorial/01-hello/
/home/akilan/go-workspace/bin/01-hello
```

```
go-workspace/
├── 01-hello
├── bin
│   └── 01-hello
├── pkg
└── src
    └── github.com
        └── akilans
            └── go-tutorial
                ├── 01-hello
                │   └── main.go
                └── readme.md

```
