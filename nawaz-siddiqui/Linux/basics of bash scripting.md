# Bash Shell Scripting

**#!/bin/bash**

**find . -name "*.c" -ls**

The first line of the script, which starts with **#!**, contains the full path of the command interpreter (in this case **/bin/bash**) that is to be used on the file.

All shell scripts generate a return value upon finishing execution, which can be explicitly set with the **exit** statement.

### Viewing Return Values

As a script executes, one can check for a specific value or condition and return success or failure as the result. By convention, success is returned as **0**, and failure is returned as a non-zero value. The return value is stored in the environment variable represented by **$?**:

**$ ls /etc/logrotate.conf**

**/etc/logrotate.conf**

**$ echo $?**

**0**

In this example, the system is able to locate the file **/etc/logrotate.conf** and **ls** returns a value of **0** to indicate success. When run on a non-existing file, it returns **2**.

## Syntax

Scripts require you to follow a standard language syntax. The table lists some special character usages within bash scripts:

| Character | Description |
| --- | --- |
| # | Used to add a comment, except when used as \#
, or as #! when starting a script |
| \ | Used at the end of a line to indicate continuation on to the next line |
| ; | Used to interpret what follows as a new command to be executed next. |
| $ | Indicates what follows is an environment variable |
| > | Redirect output |
| >> | Append output |
| < | Redirect input |
| | | Used to pipe the result into the next command. |

### Built-in shell commands

Shell scripts execute sequences of commands and other types of statements. These commands can be:

- Compiled applications
- Built-in bash commands
- Shell scripts or scripts from other interpreted languages, such as perl and python.

Shell scripts always have access to applications such as **rm**, **ls**, **df**, **vi**, and **gzip**, which are programs compiled from lower level programming languages such as C.

In addition, bash has many built-in commands, which can only be used to display the output within a terminal shell or shell script. Sometimes, these commands have the same name as executable programs on the system, such as **echo**, which can lead to subtle problems. bash built-in commands include **cd**, **pwd**, **echo**, **read**, **logout**, **printf**, **let**, and **ulimit**. Thus, slightly different behavior can be expected from the built-in version of a command such as **echo** as compared to **/bin/echo**. A complete list of bash built-in commands can be found in the bash man page, or by simply typing **help.**
