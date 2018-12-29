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
9. Map is like an object (JS) or dict (python)
   1. Keys must be of the same type
   2. Values must be of the same type
10. Struct Vs Map
    1. Map
       1. All keys must be the same type
       2. Used to represent a collection of related properties
       3. All values must be of the same type
       4. Don't need to know all the keys at the compile time
       5. Keys are indexed and can be iterated over
       6. Reference type!
    2. Struct
       1. Values can be of different type
       2. Need to know all the different fields at compile time
       3. Keys dont support indexing
       4. Use to represent a "thing" with a lot of different porperties
       5. Value type!

## Go CLI

1. go build = compile a bunch of go source
2. go run = compile and execute one or two files
3. go fmt = format all the code in each file in the current directory
4. go install = compiles and "installs" a package
5. go get = downloads the raw source code of someone else's package
6. go test = run any tests associated with the current project

## How to run GO

1. Open terminal
2. navigate to project directory
3. type "go run (file name).go"
4. to run multiple files repeat step 3 and add the other files (ex: go run main.go deck.go)

## Things to know

1. fmt.Printf("%+v",(value)) = print the function with the value and field names
2. &variable = give me the memory address of the value this var is pointing at
   1. Turn value into address/pointer with &value
3. *pointer Give me the value this memory address is pointing at
   1. turn address/pointer into value with *address
4. Reference types (dont need pointers to chage value):
   1. slices
   2. maps
   3. channels
   4. pointers
   5. functions
5. Value types (need pointers to change value):
   1. int
   2. float
   3. string
   4. bool
   5. sructs
6. Interfaces are implicit, not generic types, and can link pieces of code together.