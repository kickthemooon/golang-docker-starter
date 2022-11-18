# GoLang - Docker Starter

This is just a very small starter kit that enables 
a quick-start to get our hands dirty 
running some go code using docker

## Requirements

- A Shell (Linux, MacOS, MS WSL2, VM)
- Docker

## Start Programming

You can start programming in the `./src/main.go` file.

## Check Go Version

You can run Go commands using the shell script `./go`  
Try running `./go version`

## Compile & Run Your Application

```
./build-and-run
```

## External Dependencies - GOPATH

Before adding and using external packages, you first need to tell your editor where packages are located at.  

In case of this repository, packages will be downloaded and located in this directory under `gopath/pkg`:  
```
/path-to-this-repository/gopath/pkg/
```
After substituting `path-to-this-repository` You can instruct your editor to use this path as the GOPATH  

This way, all the instances of this repository will have their own directory of dependencies.

Related links:  
vscode: https://github.com/golang/vscode-go/blob/master/docs/gopath.md#different-gopaths-for-different-projects  
jetbrains: https://www.jetbrains.com/help/go/go-gopath.html

## Add a Dependency

For a list of interesting Go packages, check out this:  
https://github.com/avelino/awesome-go

To add a new package to your project, use `./get [package-location]`  

Example dependency for generating UUID's:
```
./get github.com/jakehl/goid
```

To install a specific dependency version use `@`:  
```
./get github.com/jakehl/goid@v1.1.0
```

## Tidy Up

To tidy up and remove unused dependencies, run:  
```
./tidy
```
