## to run
go mod tidy
go build ./main.go

--- 
  
### credits to https://faun.pub/understanding-go-mod-and-go-sum-5fd7ec9bcc34

## go mod init example/m
a go mod file saves the efforts of running the go get command for each dependent module
creates a new module, initializing the go.mod file that describes the module. At the start, it will only add the module path and go version in go mod file
## go mod tidy
ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, if there are some not used dependencies go mod tidy will remove those from go.mod accordingly.
It also adds any missing entries to go.sum and removes unnecessary entries.
## go get 
to install a specific package ex. go get go.mongodb.org/mongo-driver
## .sum file
After running any package building command like go build, go test for
the first time, it will install all the packages with specific versions i.e which are the latest at that moment.
It will also create a go.sum file which maintains the checksum so when you run the project again it will not install all packages again. But use the cache which is stored inside $GOPATH/pkg/mod directory (module cache directory).

## why the _ before package name?

To import a package solely for its side-effects (initialization), use the blank identifier as explicit package name:

import _ "lib/math"

In the case of go-sqlite3, the underscore import is used for the side-effect of registering the sqlite3 driver as a database driver in the init() function, without importing any other functions:

sql.Register("sqlite3", &SQLiteDriver{})
Once it's registered in this way, sqlite3 can be used with the standard library's sql interface in your code like in the example:

db, err := sql.Open("sqlite3", "./foo.db")
