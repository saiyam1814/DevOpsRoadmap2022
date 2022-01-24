## Started Watching [DevOps Prerequisite Course](https://www.youtube.com/watch?v=Wvf0mBNGjXY)

### Linux Basic Commands

- `cd` is used for changing directory
- `echo` is used to print anything on the terminal
- `clear` is used to clear the terminal
- `mkdir` is used to create a new directory
- `rmdir` is used to remove a directory
- `mv` is used to move the file _(same as copy)_
- `rm` is used to remove files
- `pwd` is used to print the current working directory
- `whoami` gives the current user 
- `sort` will sort the file
- `cat` is used to display the content of the file
- `touch` is used to create a file
- `cp srcfile destfile` is used to copy the file to destination directory
- `man` command shows the manual page of the command.
- `su` is used to switch to the root user
- `sudo` is used to access the privileges
- `sudo bash` is used to switch to the root user
- `sudo useradd demo` is used to create a user with name demo
- `sudo passwd demo` is used to set a password for the demo user
- `sudo userdel demo` is used to delete the user demo
- `sudo groupadd geeks` will create a group, a group can have multiple users.
- `sudo groupdel geeks` will delete the group
- `grep words filename` to search texts in a particular file 

## Few More Linux Commands
- `chown` to change the ownership of the file/directory
- `chmod` to change the access permissions of the files/directories
   - `0` --> no permissions
   - `1` --> executes permissions
   - `2` --> writes permission
   - `4` --> reads permission
   - `7` --> reads, writes and executes all three
- `chmod 777 demo.txt` will give the readm write and execute permission to user, groups any other users.
- `lsof` list of open files and details associated with them
- `lsof -u username` to see the list of the files opened by a particular user
- `id` is used to find out user and the group names and numeric IDs of the current user or any other other
- `tar -cvf filename source-folder` is used to zip of `.tar` format
- `tar -xvf tarfile` to unzip a file.
