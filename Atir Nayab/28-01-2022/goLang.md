# GO

## go mod init \<module path>
- initialized a go.mod file
- Describes the module: with name/module path and go version in the program

**All our code must belong a package**\
**The first statement in Go file must be "package..."**

## go packages
- Go prgram are organized into packages
- Go's standard library, provides different core packages for us to use.
- A package is a collection of source files

## Variables
- Variables are used to store values
- Link containers for values
- Give the variable a name and reference everywhere in the app.

## Keywords
- keywords are reserved words in every language

*Go compile Errors to enforce better code quality, to avoid possible dead code, code that is never used in the code.*  

## Contant
- Constant are like varibles, except their values cannot be changed

## Print formatted data
- It takes a template string that contains the text that needs to be formatted
- Plus some annotation verbs (or placeholder) that tells the fmt functions how to format the data passed in

*go is syntactically typed language*


## "fmt" package
Formatted Input and output (i/o)
- print messages
- collect user input
- Write into a file

## pointer
- A pointer is a variable that points to memory address of other variable

## Arrays
- fixed size
`var variable_name [size]variable_type`

## Slices
- slices is an abstraction of an array
- slices are more flexible and powerful:
  variable-length or get an sub-array of its own 
- slices are also indexed-base and have a size, but is resized when needed

## append
- Adds the element at the end of the slice
- Grows the slice if greater capacity is needed and returns the updated slice value

## Loops
- A loop statement allows us to execute code multiple time, in a loop
- you only have for loop

## Range
- Range iterates over elements of different data structures (so not only array and slices)