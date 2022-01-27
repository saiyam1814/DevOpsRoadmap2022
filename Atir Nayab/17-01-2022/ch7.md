## Basic Operations

![Basic Operations](https://courses.edx.org/assets/courseware/v1/678d889dcb1112024ef10815f9210a07/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch06_screen11.jpg)

## Rebooting and Shutting Down
The halt and poweroff commands issue shutdown -h to halt the system; reboot issues shutdown -r and causes the machine to reboot instead of just shuting down. Both rebooting and shutting down fromthe command line requires superuser (root) access.

**$ sudo shutdown -h 10:00 "Shutting down for scheduled maintenance."**

## Locating Application
In general, executing programs and  scripts should live in the `/bin, /usr/bin, /sbin, /usr/sbin`  directories or somewhere under  `/opt`. They can also appear in `/usr/local/bin` and `/usr/local/sbin`, or in a directory in a user's account space such as `/home/student/bin`.

One way to locate program is to employ which utility.

**which diff**
`/usr/bin/diff`

if which does not find the program, whereis is a good alternative because is looks for packages in a broader range of system directories.

![which and whereis utilities](https://courses.edx.org/assets/courseware/v1/716532f91059bd5a66899f8ef6e07c31/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/whereis.png)

## Accessing Directories

|command|Result|
|:---:|:---:|
|pwd| Displays the present working directory|
|cd ~ or cd|Change to your home directory|
|cd ..|Change to parent (..)|
|cd ~|Change to previous directory (- (minus))|

## Understanding Absolute and Relative Paths

* Absolute pathname
An absolute pathname begins with the root directory and follows the tree, branch by branch until it reaches the desired directory or file.
eg:  `cd /usr/bin`

* Relative pathname
A relative pathname starts from the present working directory. Relative paths never start with `/`.
eg `../../usr/bin`

Multiple slashes (/) between directories and files are allowed.

`///usr/bin` ---> `/usr/bin`

## Exploring Filesystem

|Command|Usage|
|:---:|:---:|
|cd /|Changes your current directory to the root (/) directory|
|ls|list the contents of the present working directory|
|ls -a|List all files, including hiddne files and directories|
|tree|Displays a tree view of the filesystem|

## Hard Links

the ln utility is used to create hard links and -s option for soft link, also known as symbolic links or symlinks.

**ln file1 file2**

![Hard Links](https://courses.edx.org/assets/courseware/v1/aefe6c7fa6a198680e110ceae5c95c11/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/lnubuntu.png)

## Soft (Symbolic) Links

Soft links are created with the -s option

**ln -s file1 file3**
**ln -li file1 file3**

*file3 no longer appears to be a regular file, and it clrearly points to file1 and has different inode number*

Symbolic links take no extra space on the filesystem (unless their names are very long). They are extrmely convenientm as they can easily be modified to point to different places.

Unlike hark links, soft links can point to objects even on different filesystems, paritions, and/or disks and other media, which may or may not be currently available or even exist.

![Soft (Symbolic Links)](https://courses.edx.org/assets/courseware/v1/cea407ef8cfd36b34ede2a154959a98f/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/lnsubuntu.png)

## Naivgating the Directory History

The cd command remembers where you were last, and lets you get back there with cd -. For remembering more than just the last directroy visited use pushd to change the directory.instead of cd; this pushes your starting directory onto a list. Using popd will then send you back to those directories

![Navigating Through Directory History](https://courses.edx.org/assets/courseware/v1/319814cbd06ee587a78854e88478c5b0/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/pushdfedora.png)

## Viewing Files

|Command|Usage|
|:-:|:-:|
|cat|Used for viewing files that are not very long; it does not provide and scroll back|
|tac|Used to look at a file backwards, starting with the last line|
|less|Used to view larger files because it is a paging program. It pauses at each screen full of text, provides scoll-back capabilities,and lets you search and navigate withing the file. *use / to search for a pattern in the forward direction and ? for a pattern in the backward direction.*|
|tail|Used to print the last 10 lines of a files by default. You can change the number of lines by doing -n 15 of just -15 if you wanted to look at the last 15 lines instead of the default.|
|head|The opposite of tail, by default it prints the first 10 lines of a file.|

## touch

touch is often used to set or update the access, change, and modify times of files. By defauly, it resets a file's timestamp to match the current time.

You can also create an empty file using touch

**touch <filename>**

-t options allows you to set the date and timestamp of the file to a specific value.

**touch -t 12091600 myfile**

![touch](https://courses.edx.org/assets/courseware/v1/d29a554d4187aae729d4ed40e42a0146/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/touch.png)

## mkdir and rmdir

mkdir is used to create a directory.

rmdir is used to remove a directory but the folder must be empty or the command will fail.

To remove a directory and all of its contents you have to do rm -rf.

![mkdir](https://courses.edx.org/assets/courseware/v1/72f578cd278d2bd6bd48d63efbfe589e/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/mkdir.png)

## Moving, Renaming or Removing a File

|Command|Usage|
|:-:|:-:|
|mv|Rename a file|
|rm|Remove a file|
|rm -f|Forcefully remove a file|
|rm -i|Interactively remove a file|

## Modifyign the Command Line Prompt

The PS1 variable is the character string that is displayed as the prompt on the command line. Most distributions set PS1 to a known default value.

$ echo $PS1
\$

## Standard File Streams

|Name|Symbolic Name|Value|Example|
|:-:|:-:|:-:|:-:|
|standard input|stdin|0|keyboard|
|standard outpul|stdout|1|terminal|
|standard error|stderr|2|log file|

Usually stdin is your keyboard and stdout and stderr are printed on your terminal. stderr is often redirected to an error logging file, while stdin is supplied by directing input to come from a file.

## I/O Redirection

Through the command shell, we can redirect the three standard file streams so that we can get input from either a file or another command, instead of from our keyboard, and we can write output and errors to files or use them to provide input for subsequent commands.

**do_something < input-file**
Instead for reading from stdin less than sign (<) followed by the name of the file to be consumed for input data.

**do_something > output-file**
If you want to send the output to a file, use the greater-than sign (>)

If you want to redirect stderr to a separate file, you use strerr'r file descriptor number(2), the greater than sing(>), followed by the name of the file you want to hold everything the running command writes to stderr:
**do_something 2> error-file**

bash permits an easier syntax for the above
**do_something >& all-output-file**

## pipes

**command1 | command2 | command3**

Pipeline allows linux to combine the actions of several commands into one. This is extraordinary effiecient because command2 and command3 do not have to wait for the previous pipeline commands to complete before they can hacking at the data in their input streams; on multiple CPU or core systems, the available computing power is much better utilized and things get done quicker

![Pipeline](https://courses.edx.org/assets/courseware/v1/50bdd18ba2e7d4343c184f5e0e3e058a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/pipeline.png)

## locate
The locate utility program performs a search taking advantage of a perviously contructed database of files and directories on you system

locate utilizes a database created by a related utility, updatedb. Most linux systems run this automatically once a day.

