# Install Go with VS code

1. Download and install go from golang.org/dl/
   1. golang.org/pkg = go language package information
2. Open VS code and add the GO extension
3. Create a file with .go
   1. Notice in the bottom right it says to install some files. click on the install button and install the dependencies

## Go info

1. Typing "go" in terminal will display a list of the available GO commands
2. There are two types of Packages
   1. Executable = generate a file that we can run
   2. Reusable = code used as "helpers". Good place to put reusable logic
3. package main is used to build an executable file
   1. Defines a package that ca be compiled and the executed
4. Other names are considered "helper" code
5. import allows access to other files

## Go CLI

1. go build = compile a bunch of go source
1. go run = compile and execute one or two files
1. go fmt = format all the code in each file in the current directory
1. go install = compiles and "installs" a package
1. go get = downloads the raw source code of someone else's package
1. go test = run any tests associated with the current project

## How to run GO

1. Open terminal
2. navigate to project directory
3. type "go run (file name).go"
