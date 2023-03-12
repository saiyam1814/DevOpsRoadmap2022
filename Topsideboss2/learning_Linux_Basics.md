# Learning Linux Basics

The main objective is to be comfortable with the administration, configuration and operation of Linux. Spending a lot of time working with Linux so any effort spent there will pay dividends later. As an added bonus, learning how Linux operates takes a lot of mystery out of "how things work".

No need to become an expert but instead beinging comfortable with the tasks here is all that is important. Googling or using the man pages for syntax help is important. The key thing is to know what tool to use.Using any Linux distribution available, though I prefer Debian as these are the ones i will encounter most frequently. Ubuntu is very user friendly btw.

### Skills Objectives
Learned about/how to:
* Install/configure OS
* Using BASH
* Starting/stopping/creating service
* User management
* The Linux filesystem
* Text manipulation(awk, sed, grep, wc, echo, cat, tail, etc...)
* Process monitoring with top, ps, lsof
* Networking and network configuration
* System performance
* Using screen or tmux
* Logging in locally and remotely
* Firewall management
* Install/configure applications using the OS package manager(apt or yum)
* POSIX basics
* File permissions

## Scripting

Scripting is a must-have skill for DevOps and sysadmins alike. Scripting comes in handy when one needs to make repetitive changes across multiple systems. Think about tasks like:
* applying a security patch to 1000 servers to close a security hole
* copying millions of files from one storage location to another
* digging through log files to identify a specific scenario

It set the stage for me to learn programming, which I will eventually do. Even after knowIing a coding language, I'll still use scripting for those quick and dirty tasks.
I chose from a few different options for scripting: BASH, Python, Javascript or Ruby.
I personally recommend BASH: It's guaranteed to be on every system you ever touch. Python is a great choice also, and can double as your programming language when you reach that point. I advise against Javascript. Its asycnhronous nature is not suited for our case. Ruby is fine but not common. I'd only recommend it if you knew you were going to be working in a Ruby shop.

### Skills Objectives
Learned about:
* Making scripts executable
* Capturing input from keyboards and computers
* Reading/writing to filesystems, networks and (in some cases) databases
* Providing script feedback (output status) and reacting accordingly
* Executing scripts on remote computers
* Using conditionals, variables, logic, functions, and proper syntax.

Common Linux commands one should know:

| Commands | Description |
| ------------- | -------------- |
| `where <filename>` | This will show the list of directories in which a file is present |
| `open .` | This will open all the files present in a specific directory |
| ` echo $PATH` | When one types a command to run, the system looks for it in the directorios specified by PATH. It will display the files and folder paths by the difference of colon : It will check the executable command in one of these paths. |
| `echo "Hey" > file.txt` | This will override the text in a file |
| `echo "Hey" >> file.txt` | This will append text in that file |
| `export MY_PATH="Mark"` | This will create another path that will contain the string. But this is not permanent |
| `pwd` | This will print the current working directory |
| `ls` | This will show you the list of all the files present in a specific directory |
| `.` | This represents the current directory |
| `..` | This represents the previous directory |
| `cd` | This will change the path location from one directory to another |
| `cat filename`| This will  print all the content of a file in a standard output |
| `cat > filename` | This will create a new file if it is not present and allow us to enter the text also |
| `tr` | Thus will translate the characters from one string to another string |
| `cat lower.txt` | tr a-z A-Z > upper.txt` | The output of the first command is the input of the second command |
| `man command-name` | This will show the details of a specific command |
| `mkdir directory-name` | This will create a new directory |
| `mkdir -p this/middle/one` | This will create a middle directory between two directories |
| `touch filename` | This will create a new file |
| `cp file.txt copy-file.txt` | Thus wull make a copy of the file.txt |
| `mv file.txt random` | This will move the file.txt to the random folder |
| `mv file.txt new-file.txt` | This will rename the file.txt to new-file.txt |
| `mv test renamedTest` | This will rename the directory |
| `rm file.txt` | This will remove file.txt from the compuster permanently |
| `rm -R directory-name` | This will remove the directory recursively |
| `rm -rf directory-name` | This will forcefully remove the directory |

### Here is a link to my ALX Scripting Directory
* [ALX Scripting Directory](https://github.com/Topsideboss2/alx-system_engineering-devops)
