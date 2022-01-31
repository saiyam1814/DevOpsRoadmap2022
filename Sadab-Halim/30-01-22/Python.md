# Revised Python from [Python Tutorial for Beginners](https://www.youtube.com/watch?v=t8pPdKYpowI) and from Docs

## Function in Python
a function is a group of related statements that performs a specific task.

Functions help break our program into smaller and modular chunks. As our program grows larger and larger, functions make it more organized and manageable.

It also avoids repition and makes the code reusable and cleaner.

Syntax:
```
def function_name (paramters):
  """docstring"""
  statement(s)
```

Example:
```
def greet(name):
    """
    This function greets to
    the person passed in as
    a parameter
    """
    print("Hello, " + name + ". Good morning!")

greet('Paul')
```
**NOTE:** <br/>
In python, the function definition should always be present before the function call.

**Docstring:**
The first string after the function header is called the docstring and is short for documentation string. It is briefly used to explain what a function does.

**return Statement**
The return statement is used to exit a function and go back to the place from where it was called.

Syntax:
```
return [expression_list]
```

Example:
```
def absolute_value(num):
    """This function returns the absolute
    value of the entered number"""

    if num >= 0:
        return num
    else:
        return -num


print(absolute_value(2))

print(absolute_value(-4))
```

**How function works in Python?**

![image](https://user-images.githubusercontent.com/74575612/151848292-723ec47e-ee77-40a5-aed2-1fc652dca64f.png)

**Types of Functions:**
- Built-in functions: Functions that are built into Python
- User-defined functions: Functions defined by the users themselves.

**Arguments**
In Python, there are other ways to define a function that can take variable number of arguments rather than having a fixed number of arguments.

- **Python Default Arguments:**
  We can provide a default value to an argument by using the assignment operator (=).
  
  Example:
  ```
  def greet(name, msg="Good morning!"):
    """
    This function greets to
    the person with the
    provided message.

    If the message is not provided,
    it defaults to "Good
    morning!"
    """

    print("Hello", name + ', ' + msg)

    greet("Kate")
    greet("Bruce", "How do you do?")
    ```
- **Python Keyword Arguments:**
  Example"
  
  ```
  # 2 keyword arguments
  greet(name = "Bruce",msg = "How do you do?")

  # 2 keyword arguments (out of order)
  greet(msg = "How do you do?",name = "Bruce") 

  1 positional, 1 keyword argument
  greet("Bruce", msg = "How do you do?") 
  ```

- **Python Arbitrary Arguments:**
  Sometimes, we do not know in advance the number of arguments that will be passed into a function. Python allows us to handle this kind of situation through function calls with an arbitrary number of arguments.

  In the function definition, we use an asterisk (*) before the parameter name to denote this kind of argument. 
  
  Example:
  ```
  def greet(*names):
    """This function greets all
    the person in the names tuple."""

    # names is a tuple with arguments
    for name in names:
        print("Hello", name)

  greet("Monica", "Luke", "Steve", "John")
  ```
  
## Recursion
Recursion is the process of defining something in terms of itself.

**Working of Recursive Function**

![image](https://user-images.githubusercontent.com/74575612/151849213-3ff38447-3e2b-440b-896d-82e9f3b399c2.png)

Example:
```
def factorial(x):
    """This is a recursive function
    to find the factorial of an integer"""

    if x == 1:
        return 1
    else:
        return (x * factorial(x-1))

num = 3
print("The factorial of", num, "is", factorial(num))
```

**Lambda Function or Anonymous Function:**
In Python, an anonymous function is a function that is defined without a name.

While normal functions are defined using the def keyword in Python, anonymous functions are defined using the lambda keyword.

Hence, anonymous functions are also called lambda functions.

Syntax of Lambda Function:
lambda arguments: expression

Example:
```
double = lambda x: x * 2
print(double(5))
```

**Use of Lambda Function:**
- We use lambda functions when we require a nameless function for a short period of time.
- we generally use it as an argument to a higher-order function (a function that takes in other functions as arguments).
- Lambda functions are used along with built-in functions like filter(), map() etc.

Example with filter()
The filter() function in Python takes in a function and a list as arguments.

The function is called with all the items in the list and a new list is returned which contains items for which the function evaluates to True.

```
# Program to filter out only the even items from a list
my_list = [1, 5, 4, 6, 8, 11, 3, 12]

new_list = list(filter(lambda x: (x%2 == 0) , my_list))
print(new_list)
```

Output: [4, 6, 8, 12]

Example with map()
The map() function in Python takes in a function and a list.

The function is called with all the items in the list and a new list is returned which contains items returned by that function for each item.

```
my_list = [1, 5, 4, 6, 8, 11, 3, 12]
new_list = list(map(lambda x: x * 2 , my_list))
print(new_list)
```

Output: [2, 10, 8, 12, 16, 22, 6, 24]

## Global, Local and Nonlocal Variables

**Global Variables**
a variable declared outside of the function or in global scope is known as a global variable. This variable cam ne accessed inside or outside of the function.

Example:
```
x = "global"
def foo():
    print("x inside:", x)

foo()
print("x outside:", x)
```

**Local Variables**
A variable declared inside the function's body or in the local scope is known as a local variable.

Example:
```
def foo():
    y = "local"

foo()
print(y)
```

**Nonlocal Variables**
Nonlocal variables are used in nested functions whose local scope is not defined. This means that the variable can be neither in the local nor the global scope.

Example:
```
def outer():
    x = "local"

    def inner():
        nonlocal x
        x = "nonlocal"
        print("inner:", x)

    inner()
    print("outer:", x)

outer()
```

**NOTE:** If we change the value of a local variable, the changes appear in the local variable.

## global Keyword
In Python, global keyword allows us to modify the variable outside of the current scope. 

It is used to create a global variable and make changes to the variable in a local context.

**Rules:**
- When we create a variable inside a function, it is local by default.
- When we define a variable outside of a function, it is global by default. You don't have to use global keyword.
- We use global keyword to read and write a global variable inside a function.
- Use of global keyword outside a function has no effect.

**Use of global keyword:**
Example-1, Accessing global Variable From Inside a Function
```
c = 1 # global variable
def add():
    print(c)
add()
```

Example-2, Modifying Global Variable From Inside the Function
```
c = 1 # global variable
def add():
    c = c + 2 # increment c by 2
    print(c)
add()
```

When we run the above program, the output shows an error:

```
UnboundLocalError: local variable 'c' referenced before assignment
```

This is because we can only access the global variable but cannot modify it from inside the function.

The solution for this is to use the global keyword.

Example-3, Changing Global Variable From Inside a Function using global
```
c = 0 # global variable
def add():
    global c
    c = c + 2 # increment by 2
    print("Inside add():", c)
add()
print("In main:", c)
```

## Modules in Python
Modules refer to a file containing Python statements and definitions.

A file containing Python code, for example: example.py, is called a module, and its module name would be example.

Modules are used to break down large programs into small manageable and organized files. 

Thus, modules provide reusability of code.

Example: filename: example.py
```
# Python Module example

def add(a, b):
   """This program adds two
   numbers and return the result"""

   result = a + b
   return result
```

**How to import modules?**
We use `import` keyword to import modules.

Example:
```
import math as m
print("The value of pi is", m.pi)
```

**from...import Statement:**
We can import specific names from a module without importing the module as a whole.

Example:
```
# import only pi from math module
from math import pi
print("The value of pi is", pi)
```

To import all names:
```
from math import *
print("The value of pi is", pi)
```

**dir() built-in function**
We can use the dir() function to find out names that are defined inside a module.

Example:
```
dir(example)
['__builtins__',
'__cached__',
'__doc__',
'__file__',
'__initializing__',
'__loader__',
'__name__',
'__package__',
'add']
```

## Packages in Python
A python package is a collection of modules. Modules that are related to each other are mainly put in the same package. When a module from an external package is required in a program, that package can be imported and its modules can be put to use.

**Package Module Structure in Python**

![image](https://user-images.githubusercontent.com/74575612/151851949-99bc8d8c-7885-401b-a881-361e308eea73.png)


**Importing module from a package**
We can import modules from packages using the dot (.) operator.

Example:
```
import Game.Level.start
```
