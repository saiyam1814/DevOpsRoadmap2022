# Linux

## Setting Environment Variables
By default, variables created within a script are only avaible to the current shell; child processes (sub-shells) will not have access to values that have been set or modified.
Allowing child processes to see the values requires use of the export command.

|Task|Command|
|:-:|:-:|
|Show the value of a specific variables|echo $SHELL|
|Export a new variable value|export VARIABLE=value (or VARIABLE=value; export VARIABLE)|
|Add a variable permanently|Edit `~/.bashrc` and add the line export VARIABLE=value    Type source ~/.bashrc or just . ~/.bashrc (dot ~/.bashrc); or just start a new shell by typing bash|

## The HOME Variable
Home is an environment variable that represents the home (or login) directory of the user. cd without arguments will change the current working directory to the value of HOME. 

`cd $HOME` and `cd ~` are completely equivalent statements.

|Command|Explanation|
|:-:|:-:|
|`echo $HOME`|show the value of the HOME environment variable|
|pwd|where are we? Use print working directory to find out|
|cd|Change directory|
|pwd|...takes us back to HOME, as you can now see|

## The PATH Variable
PATH is an ordered list of directories (the path) which is scanned when a command is given to find to appropriate program or script to run. Each directory in path is separated by colons(:). A null (empty) directory name (or ./) indicates the current directory at any given time.

* :path1:path2 - null directory defore the first colon (:)
* path1::path2 - null directory between path1 and path2

![The PATH Variable](https://courses.edx.org/assets/courseware/v1/b955fc762a2b63221bc0a46e1fb9b419/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/setpath.png)

## The SHELL Variable
The environment variable SHELL point to the user's default command shell and contains the full pathname to the shell.

## The PS1 Variable and the command line prompt
Prompt statement (PS) is used to customize your prompt string in your terminal windows to display the information you want.

PS1 is the primary prompt variable which controls what your command line prompt look like. The following special characters can be included in PS1.

`\u` - User name
`\h` - Host name
`\w` - Current working directory
`!` - History number of this comand
`\d` - Date

They must be surrounded in single quotes when they are used.

![the ps1 variable and the command line prompt](https://courses.edx.org/assets/courseware/v1/0e41f5477fe686acda77ec04fe717d3d/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/ps1.png)


## Recalling Previous Commands
bash keeps track of previously entered commands and statements in a history buffer. You can recall previously used commands simply by using the up and down cursor keys. The information is stored in `~/.bash_history`. If you have multiple terminals open, the commands typed in each session are not saved until the session terminates.

![recalling previous commands](https://courses.edx.org/assets/courseware/v1/0dc7d17f95a156ae24e9ee8a22d98b69/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/debianhistory.png)

## Using history environment variables
* HISTFILE
  The location of the history file.
* HISTFILESIZE
  The maximum number of lines in the history file(default 500)
* HISTSIZE
  The maximum number of commands in the history file.
* HISTCONTROL
  How commands are stored
* HISTIGNORE
  which command line can be unsaved

  ![using history evironment variables](https://courses.edx.org/assets/courseware/v1/529cf8bca6c3fff74ad538be6dbab7f2/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/history.png)

## Finding and Using Previous Commands
  |Key|Usage|
  |:-:|:-:|
  |Up/Down arrow keys|Browse through the list of commands previously executed|
  |!!(pronounced as bang-bang)|Execute the previous command|
  |CTRL-R|search previously used commands. Press enter to run search command|

## Executing Previous Commands
|Syntax|Task|
|:-:|:-:|
|!|Start a history substitution|
|!$|Refer to the last argument in a line|
|!n|Refer to the nth command line|
|!string|Refer to the most recent command starting with string|

## Keyboard Shortbut
|Keyboard shortcut|task|
|:-:|:-:|
|CTRL-L|Clears the screen|
|CTRL-D|Exits the current shell|
|CTRL-Z|Puts the current process into suspended backgrond|
|CTRL-C|kills the current process|
|CTRL-H|works the same as backspace|
|CTRL-A|Goes to the begining of the line|
|CTRL-W|Deletes the word before the cursor|
|CTRL-U|Deletes from begining of line to cursor position|
|CTRL-E|goes to the end of the line|
|TAB|Auto-completes fies, directories, and binaries|

## File Ownership
In Linux and other UNIX-based operating systems, every file associated with a user who is the owner. Every file is also associated with a group which has an interest in the filea and certain rights, or permissions: read, write, and execute.

|command|Usage|
|:-:|:-:|
|chown|used to change user ownership of a file or directory|
|chgrp|used to change group ownership|
|chmod|used to change the permissions on the file, which can be done separately for owner, group and the rest of the world|

## File Permission Modes and chmod
Files have three kinds of permissions: read (r), write(w), execute(x). These are generally represented as in rwx. These permissions affect three groups of owner: user/owner(u), group(g), and others(o).

As a result you have the following three groups of three permissions:

rwx:  rwx:  rwx
u:    g:    o

There are a number of different ways to use chmod.
```
ls -l somefile
-rw-rw-r-- 1 student student 1601 Mar 9 15:04 somefile
chmod uo+x,g-w somefile
ls -l somefile
-rwxr--r-x 1 student student 1601 Mar 9 15:04 somefile
```

where u stands for user (owner), o stands for other(world), and g stands for group.

* 4 if read permission is desired
* 2 if write permission is desired
* 1 if execute permission is desired
Thus, 7 means read/write/execute, 6 means read/write, and 5 means read/execute

When you apply this to the chmod command, you have to give three digits for each degree of freedom

```
chmod 755 somefile
ls -l somefile
```

![file permission modes and chmod](https://courses.edx.org/assets/courseware/v1/5d60930deeaaca887d468867240fc6e0/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/chmodmint.png)

## Example of chown
![chown](https://courses.edx.org/assets/courseware/v1/d99d45386528584f3f861f182577fb1a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/chownrhel7.png)

## Example of chgrp
![chgrp](https://courses.edx.org/assets/courseware/v1/2416814b8977e0048d01804a8319aace/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/chgrouprhel7.png)

## Command Line Tools for manipulating Test files
![command line tool for manipulating text files](https://courses.edx.org/assets/courseware/v1/1003ad62cadceacd75a659e16ec77873/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/cmdlinetext.png)

## cat
cat is short concatenate. It is often used to read and print files, as well as for simply viewing file contents.
`cat <filename>`

the tac command prints the lines of a file in reverse order.

|Command|Usage|
|:-:|:-:|
|cat file1 file2|Concatenate multiple files and display the output; i.e. the entire content of the first is followed by that of the second file|
|cat file1 file2 > newfile|Combine multiple files and save the output into a new file|
|cat file >> axistingfile|Append a file to the end of an existing file|
|cat > file|Any subsequent lines typed will go into the file, until CTRL-D is typed|
|cat >> file|Any subsequent lines are appended to the file, until CTRL-D is typed.|

## Using cat interactively
![using cat](https://courses.edx.org/assets/courseware/v1/2186f2162bb7f8d6f7fac9004f7d4784/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/cateoffedora.png)

Another war to create a file at the terminal is `cat> <filename> <<EOF`. A new file is created and you can type the requried input. To, exit enter EOF at the begining of a line.

*EOF is case sensitive. One can also use another word, such as STOP*

## echo
echo simply displays text. 

`echo string`

echo can be used to display a string on standard output or to place in a new file (using the > operator) or append to an already existing file (using the >> operator).

The -e option, along with the following switches, is used to enable special character sequences, such as the newline character or horizontal tab.
* \n represents newline
* \t represents horizontal tab.

echo is particularly useful for viewing the values of environment variables

|Command|Usage|
|:-:|:-:|
|echo string > newfile|The specified string is placed in a new file|
|echo string >> existringfile|The specified string is appended to the end of an already existing file|
|echo $variable|the contents of the specified environment variable are displayed|

## Working with Large Files
Viewing somefile can be done by typing either of the two following commands
```
less somefile
cat somefile | less
```

By default man pages are sent through the less command. You may have encountered the older more utility which has the same basic function but fewer capabilities i.e. less is more!

## head
head reads the first few lines of each named file (10 by default) and displays it on standard output.

eg
`head -n 5 /etc/default/grub`
you cal also just say:
`head -5 /etc/defaultgrub`
![head](https://courses.edx.org/assets/courseware/v1/a19d53c185f88f384a24c5b0ae6fe03a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/headubuntu.png)

## tail
tail prints the last few lines of each name file and displays the last 10 lines.
tail is especially useful when you are troubleshooting any issue using log fiiles, as you probably want to see the most recent line of output.

![tail](https://courses.edx.org/assets/courseware/v1/e218e0f357f957b78420b93f6dc63aaf/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/tailubuntu.png)

## Viewing compressed files
When working working with compresses files, many standard commands cannot be used directly. For many commonly-used file and text manipulation programs, there is also a version especially designed to work directly with compressed files. These associated utilities have the letter 'z' prefixed to their name. For example we have utility programs such as zcat, zless, zdiff, and zgrep

|Command|Description|
|:-:|:-:|
|zcat compressed-file.txt.gz|to veiw a compressed file|
|zless somefile.gz or zmore somefile.gz|to page through a compresses file|
|zgrep -i less somefile.gz|to search inside a compresses file|
|zdiff file1.txt.gz|to search inside a compresses file|
|zdiff file1.txt.gz file2.txt.gz|to compare two compresses files|

*if you run zless on an uncomresses file, it will still work and ignore the decompression stage. *
*There are also quivalent utility programs for other compression methods besides gzip, for example we have bzcat and bzless associated with bzip2 and xzcat and xzless associated with xz.*

## sed and awk
It is very common to create and then repeatedly edit and/or extract contents from a file. Let's learn how to use sed and awk to easily perform such operations.

## sed
It is a powerful text processing tool and is one of the oldest, earliest and most popular UNIX utilities. It is used to modify the contents of a file or input stream, usually placing the contents into a new file or output stream. 

![sed](https://courses.edx.org/assets/courseware/v1/64bebc2555b3777d251a871e72e873d7/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch12_screen_13.jpg)

sed can filter text as well as perform susbtitutions in data streams.

Data from an input source/file (or stream) is taken and moved to wroking space. The entire list of operations/modifications is applied over the data in the working space and the final contents are moved to the standard output space (or stream).

## sed command syntax
|Command|Usage|
|:-:|:-:|
|sed -e command \<filename>|specify editing commands at the command line, operate on file and put the output on standard out.|
|sed -f scripfile \<filename>|specify a scriptfile containing sed commands operate on file and put output on standard out|
|echo "I hate you" \| sed s/hate/love|use sed to filter standard input, putting output on standard out.|

![sed command syntax](https://courses.edx.org/assets/courseware/v1/43267d11e71aba5dee81d8e780b24579/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/fedorased.png)

## sed basic operations
We can perform multiple editing and filtering operations with sed.
|Command|Usage|
|:-:|:-:|
|sed s/pattern/replace_string/file|substitute first string occurence in every line|
|sed s/pattern/replace_string/g|substitute all string occurences in every line|
|sed 1,3s/pattern/replace_string/g file| susbstitute all string occurences in a range of lines|
|||
|sed -i s/pattern/replace_string/g file|save chages for string substitution in the same file|


# Go

### Conditionals
| Operator | Description             |
| -------- | ----------------------- |
| ==       | Compares two things     |
| !=       | Not equals              |
| <        | Less than               |
| <=       | Less than equal to      |
| >        | Greater than            |
| >=       | Greater than equal to   |

**If Else-if Else Statement**

```
if condition {
   // code to be executed if condition is true
} else if condition {
  // code to be executed if condition is true
else {
  // code to be executed if both conditions are false
}
```

**Boolean**

```
func main() {
    var bVal bool   // default is false
    fmt.Printf("bVal: %v\n", bVal)
}
```

**Switch Statement**

```
func main() {
	
	switch day:=4; day{
	case 1:
	fmt.Println("Monday")
	case 2:
	fmt.Println("Tuesday")
	case 3:
	fmt.Println("Wednesday")
	case 4:
	fmt.Println("Thursday")
	case 5:
	fmt.Println("Friday")
	case 6:
	fmt.Println("Saturday")
	case 7:
	fmt.Println("Sunday")
	default:
	fmt.Println("Invalid")
  }
}
```

Resources:
- [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)