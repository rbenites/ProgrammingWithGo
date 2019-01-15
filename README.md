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
11. Reader interface
    1. Allows us to to create code that can interface with various functions
12. Channels
    1. Allows us to send and receive valuse with the channel operator, <-.

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
7. Every value has a type
8. Every function has to specify the type of its arguments
9. Control Flow
   1. Iterative/Loop
   2. Conditional
   3. Sequence

## Other Go info

1. A STATEMENT is the smallest standalone element of a program that expresses some action to be carried out.
   1. It is an instruction that commands the computer to perform a specified asction.
   2. They usually take up one line.
2. An EXPRESSION is a combination of one or more explicit values, constants, variables, operators, and functions that the programming language interprets and computes to produce another value.
3. The "scope" of a variable is where you can access the variable, eg, write to it or read the value from it.
4. A "primitive" data TYPE is one that is built into the language AND/OR just a basic data type which is built into the language
   1. Int, String
5. A "composite" data TYPE allows you to compose together values of other data TYPES
6. When a variable is declared in Go using the "var" keyword, and no VALUE is ASSIGNED to that variable, then the compiler assigns a default value to the variable. This is known as the "zero value"
7. Keywords are words that a reserved for use by the Go programming language; they have to be used in a certain way for a certain purpose.
   1. Keywords are sometimes called “reserved words.”
   2. You can’t use a keyword for anything other than its purpose.
   3. Keywords = package, var
8. In “2 + 2” the “+” is the OPERATOR 
9. In “2 + 2” the “2”s are OPERANDS
10. The entry point for all programs is in func main() which needs to be inside package main
11. The short declaration operator can only be used at the block level - between curly braces. The language spec says, "Short variable declarations may appear only inside functions.
12. Good Package names should be Short, Concise, and Evocative
13. https://play.golang.org/ can be used to write (most) Go code and have it run
14. A great place to ask questions is the "golang bridge forum" at https://forum.golangbridge.org/ 
15. When you see something like "fmt.Println()" this is calling the "Println()" function from the "fmt" package.
16. An "identifier" is the name assigned to a variable or a function or a constant.
17. To call a func, variable, or constant from a package, use the "package-dot-identifier" syntax. For example, like this, "fmt.Println()" 
18. Idiomatic code is Go code which conforms to best practices for writing Go code.
19. _ is a characterused to "throw away" code
20. In Go, you cannot have a variable which you do not use.
21. When you see that a func has a parameter of this type "...interface{}" this is called a "variadic parameter" and it means that the func can take as many values of that type as you want to pass in.
22. Every value in Go is also of type "empty interface" which is expressed like this: "interface{}
23. If you wanted to print to a string and then assign that value to a variable, you could use the "func Sprintf()" from the "fmt" package.
24. In Go, you can create your own TYPE
    1. When you create our own TYPE in Go, that TYPE will have an "underlying TYPE".
25. We don't say "casting" in Go, we say "conversion"
26. Go emphasizes ease of programming. If you use the short declaration operator, you do not need to specify the type.
27. how computers work
    1.  Computers run on electricity. Electricity has two discrete states: on & off. We can associate a coding scheme with the state of a circuit. For example, the porch light on Halloween in America: when it is "on" it means "come trick or treat", and when it is "off" it means "go away." If we had two porch lights, we could encode four messages:
   on on = some message
   on off = some message
   off on = some message
   off off = some message

   If we had 3 porch lights, we could encode 8 messages. The formula for figuring out how many messages can be encoded is 2 to the power of N where "N" is the number of porch lights. For instance, 2 to the power of 3, is 8.

   Instead of writing "on off on on off", etcetera, we can have "1" represent "on" and "0" represent "off" and thus more easily write "1 0 1 1 0"
28. A boolean value is one that is either true or false
29. if you have 5 porch lights, how many messages can you encode? 32 or 2 to the 5th power
30. "Bit" is an abbreviation of "binary digit"
    1.  1000 bytes = 1kb
    2.  1000GB = 1TB
31. rune is an alias for int32
    1.  A string is a sequence of bytes that represent Unicode code points, called runes.
32. byte is an alias for uint8
33. If you use type int, then the compiler will choose whether int32 or int64 is used. Another way to say this is that int has implementation-specific sizes
34. As a rule of thumb, for numeric types, you should just use "int" for whole numbers (without decimals) and "float64" for real numbers (with decimals)
35. A string is a sequence of bytes, which is also known as a "slice of bytes"
36. Go source code is always UTF-8.
