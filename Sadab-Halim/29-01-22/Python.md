# Revised Python from [Python Tutorial for Beginners](https://www.youtube.com/watch?v=t8pPdKYpowI) and from Docs

## Introduction

**Key Points:**
- Easy to learn
  - Simple syntax
  - Easy to setup
- Large ecosystem
  - Many libraries
  - Large community
- Flexible
  - Not limited to language specifics

**What is Python used for?**
- Web Development
- Data Science
- Machine Learning
- Artificial Intelligence
- Web Scraping
- Automation
  - Automate DevOps Tasks
  - Automate General Tasks

`Our Source Code` --> `Python Interpreter` --> `Running Code`

**Python Interpreter:**
- knows how to execute Python code
- translates our program into machine readable binary code

Python is a very simple language that uses a very familiar language to code.
It uses the identation to define blocks of code and they need to be consistent throughtout the block.

Python is a case-sensitive language.

**Syntax of Python File:** `.py`

## Python Keywords
Keywords are the reserved words in Python.

We cannot use a keyword as a variable name, function name or any other identifier. They are used to define the syntax and structure of the Python language.

There are 33 keywords in Python <br/>

`False` `await` `else` `import` `pass` <br/>
`None` `break` `except` `in` `raise` <br/>
`True` `class` `finally` `is` `return` <br/>
`and` `continue` `for` `lambda` `try` <br/>
`as` `def` `from` `nonlocal` `while` <br/>
`assert` `del` `global` `not` `with` <br/>
`async` `elif` `if` `or` `yield` <br/>

## Variables in Python

Declaring and assigning value to a variable
```
website = "google.com"
print(website)
```

**NOTE:** Python is a type-inferrred language, so we don't have to expicitly define the variable type.

Changing the value of a variable
```
website = "google.com"
print(website)

website = "youtube.com"
print(website)
```

## Data Types in Python
- **Numbers**
  - int
  - float
  - complex
  
  Example:
  ```
  a = 5
  print(a, "is of type". type(a))
  
  b 2.0
  print(a, "is of type". type(a))
  
  c= 1+2j
  print(a, "is it a complex number?", isinstance(1+2j), complex))
  ```
  
- **List** is an ordered sequence of items. All the items in a list do not need to be of the same type. Lists are mutable. 
  
  It is defined by [ ].
  
  Example:
  ```
  a = [1, 2.2, 'python']
  ```
  
- **Tuple** is an ordered sequence of items same as a list. The only difference is that tuples are immutable. Tuples are used to write-protect data and are usually faster than lists as they cannot change dynamically.

  It is defined within parentheses ()

  Example:
  ```
  t = (5,'program', 1+3j)
  ```
- **Strings** is sequence of Unicode characters. We can use single quotes or double quotes to represent strings. Multi-line strings can be denoted using triple quotes, ''' or """.

  Example:
  ```
  s = "This is a string"
  print(s)
  s = '''A multiline
  string'''
  print(s)
  ```
- **Set** is an unordered collection of unique items. Set is defined by values separated by comma inside braces { }.
  
  We can perform set operations like union, intersection on two sets. Sets have unique values. They eliminate duplicates.
  
  Since, set are unordered collection, indexing has no meaning.
  
  Example-1:
  ```
  a = {5,2,3,1,4}
  print("a = ", a)
  print(type(a))
  ```
  
  Example-2:
  ```
  a = {1,2,2,3,3,3} 
  print(a)
  ```
  
- **Dictionary** is an unordered collection of key-value pairs.

  It is generally used when we have a huge amount of data. Dictionaries are optimized for retrieving data.
  
  It is defined within braces {} with each item being a pair in the form key:value. Key and value can be of any type.

  Example:
  ```
  d = {1:'value','key':2}
  print(type(d))

  print("d[1] = ", d[1])

  print("d['key'] = ", d['key'])

  # Generates error
  print("d[2] = ", d[2])

  ```
**Conversion between Data Types**
We can convert between different data types by using different type conversion functions like int(), float(), str(), etc.

## Type Conversion:
The process of converting the value of one data type to another data type is called type conversion. 

There are two types of conversion:
1. **Implicit Conversion:** In Implicit type conversion, Python automatically converts one data type to another data type. 

    Example:
    ```
    num_int = 123
    num_flo = 1.23

    num_new = num_int + num_flo

    print("datatype of num_int:",type(num_int))
    print("datatype of num_flo:",type(num_flo))

    print("Value of num_new:",num_new)
    print("datatype of num_new:",type(num_new))
    ```

2. **Explicit Conversion:** In Explicit Type Conversion, users convert the data type of an object to required data type. We use the predefined functions like int(), float(), str(), etc to perform explicit type conversion.

    This type of conversion is also called typecasting because the user casts (changes) the data type of the objects.

    Syntax:
    ```
    <datatype>(expression)
    ```

    Example:
    ```
    num_int = 123
    num_str = "456"

    print("Data type of num_int:",type(num_int))
    print("Data type of num_str before Type Casting:",type(num_str))

    num_str = int(num_str)
    print("Data type of num_str after Type Casting:",type(num_str))

    num_sum = num_int + num_str

    print("Sum of num_int and num_str:",num_sum)
    print("Data type of the sum:",type(num_sum))
    ```

**NOTE:**
- Implicit Type Conversion is automatically performed by the Python interpreter.
- In Type Casting, loss of data may occur as we enforce the object to a specific data type.

## Python Input/Output
For Output:
- print()
- Output Formatting: str.format()
  Example-1:
  ```
  print('I love {0} and {1}'.format('bread','butter'))
  print('I love {1} and {0}'.format('bread','butter'))
  ```
  
  Example-2:
  ```
  print('Hello {name}, {greeting}'.format(greeting = 'Goodmorning', name = 'John'))
  ```
  
  Example-3:
  ```
  x = 12.3456789
  print('The value of x is %3.2f' %x)
  ```
 
For Input: input([prompt])
Example:
```
num = input('Enter a number: ')
Enter a number: 10
```

## Import in Python
Import is used to access different packages available in Python. <br/>
It is also used when our program grows bigger, it is a good idea to break it into different modules and import them using `import` keyword.

Example:
```
import math
import numpy as np
import matplotlib as mp
```

## Operators in Python
- **Arithmetic Operators:** +, -, *, /, %, //, **
- **Comparison Operators:** >, <, ==, !=, >=, <=
- **Logical Operators:** and, or, not
- **Bitwise Operators:** &, |, ~, ^, >>, <<
- **Assignment Operators:** =, +=, -=, *=, /=, %=, //=, **=, &=, |=, ^=, >>=, <<=
- **Identity Operators:** is, is not
- **Membership Operators:** in, not in

## Flow Controls

**if Statement Syntax**
```
if test expression:
  statement(s)
```

**if else Statement Syntax**
```
if test expression:
  Body of if
else:
  Body of else
```

**if elif else Statement Syntax:**
```
if test expression
  Body of if
elif test expression:
  Body of elif
else:
  Body of else
```

**for Loop Syntax:**
```
for val in sequence:
  loop body
```

**for Loop with else**
```
for val in sequence:
  statements(s)
else:
  statement(s)
```

**while Loop Syntax:**
```
while test_expression:
  Body of while
```


**while Loop with else Syntax:**
```
while test_expression:
  Body of while
else:
  statement(s)
```

**break Statement Syntax**
```
for var in sequence:
  # codes inside for loop
   if condition:
      break
# codes outside for loop

while test_expression:
  # codes inside for loop
  if condition:
      break
# codes outside while loop
```

**continue Statement Syntax**
```
for var in sequence:
  # codes inside for loop
   if condition:
      continue
   # codes inside for loop
# codes outside for loop

while test_expression:
  # codes inside for loop
  if condition:
      continue
  # codes inside while loop
# codes outside while loop
```

**pass Statement**
In Python, the pass statement is a null statement.
The difference between a comment a pass statement in Python is that while the interpreter ignores a comment entirely, pass is not ignored.

Syntax:
```
pass
```

Example:
```
'''pass is just a placeholder for
functionality to be added later.'''
sequence = {'p', 'a', 's', 's'}
for val in sequence:
    pass
```
