# My First GO programms

## Install Go with VS code

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
   1. Defines a package that ca be compiled and then executed
4. Other names are considered "helper" code
5. import allows access to other files
6. Go Playground allows you to test code <https://play.golang.org/>.
7. Packages are GO functions that are builtinto the framework and can be used in your code through imports <https://golang.org/pkg/>
8. Struct is like an object (JS) or dict (python).

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
4. to run multiple files repeat step 3 and add the other files (ex: go run main.go deck.go)

## Things to know

1. fmt.Printf("%+v",(value)) = print the function with the value and field names
2. &variable = give me to the memory address of the value this var is pointing at
   1. Turn value into address with &value
3. *pointer Give me the value this memory address is pointing at
   1. turn address into value with *address
