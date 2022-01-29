# Constructs in Bash Shell Scripting

### If Statement

In compact form, the syntax of an **if** statement is :

```jsx
if [ -f "$1" ]
then
    echo file "$1 exists" 
else
    echo file "$1" does not exist
fi
```

You can use the **if** statement to compare strings using the operator **==** (two equal signs). The syntax is as follows:

```jsx
if [ string1 == string2 ] ; then
   ACTION
fi
```

## Elif Statement

You can use the **elif** statement to perform more complicated tests, and take action appropriate actions. The basic syntax is:

```jsx
if [ sometest ] ; then
    echo Passed test1 
elif [ somothertest ] ; then
    echo Passed test2 
fi
```

## Testing for Files

```jsx
if [ -x /etc/passwd ] ; then
    ACTION
fi
```

the **if** statement checks if the file **/etc/passwd** is executable, which it is not. Note the very common practice of putting:

**; then**

on the same line as the **if** statement.

| Condition | Meaning |
| --- | --- |
| -e file | Checks if the file exists. |
| -d file | Checks if the file is a directory. |
| -f file | Checks if the file is a regular file (i.e. not a symbolic link, device node, directory, etc.) |
| -s file | Checks if the file is of non-zero size. |
| -g file | Checks if the file has sgid set. |
| -u file | Checks if the file has suid set. |
| -r file | Checks if the file is readable. |
| -w file | Checks if the file is writable. |
| -x file | Checks if the file is executable. |

## Boolean Expression

Boolean expressions return either TRUE or FALSE. For example, to check if a file exists, use the following conditional test:

**[ -e <filename> ]**

Similarly, to check if the value of **number1** is greater than the value of **number2**, use the following conditional test:

**[ $number1 -gt $number2 ]**

The operator **gt** returns TRUE if **number1** is greater than **number2**.

### Numerical Tests

You can use specially defined operators with the **if** statement to compare numbers. The various operators that are available are listed in the table:

| Operator | Meaning |
| --- | --- |
| -ne | Not equal to |
| -eq | Equal to |
| -gt | Greater than |
| -lt | Less than |
| -ge | Greater than or equal to |
| -le | Less than or equal to |

The syntax for comparing numbers is as follows:

**exp1 -op exp2**

## Arithmetic Expressions

Arithmetic expressions can be evaluated in the following three ways (spaces are important!):

- Using the **expr** utility
    
    **expr** is a standard but somewhat deprecated program. The syntax is as follows:
    
    **expr 8 + 8echo $(expr 8 + 8)**
    
- Using the **$((...))** syntax
    
    This is the built-in shell format. The syntax is as follows:
    
    **echo $((x+1))**
    
- Using the built-in shell command **let**. The syntax is as follows:
    
    **let x=( 1 + 2 ); echo $x**
    

In modern shell scripts, the use of **expr** is better replaced with **var=$((...))**.
