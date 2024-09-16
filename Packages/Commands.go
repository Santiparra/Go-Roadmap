/* 
Common Sense:
1. Hide internal logic
2. Don’t change APIs changing exported function’s signatures
3. Don’t export functions from the main package
4. Packages shouldn't know about dependents
*/


//go mod init initializes an go.mod module file 
go mod init {REMOTE}/{USERNAME}/mystrings
//i.e inits
module example.com/username/hellogo

go 1.23.0

require (
	example.com/username/mystrings v0.0.0
)

//go run .... Compile and run x go package
go run main.go

//go build compiles a single executable
go build
./hellogo

//go install installs package/s in local environment
go install
hellogo

//go get package in x folder will install the remote package and update the module dependencies
go get github.com/wagslane/go-tinytime