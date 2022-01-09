## Linux Basic Commands 


- `man` command will show you a manual page of the command.
- `clear` command will clear your terminal
- `pwd` will show you the current working directory.
- `cd` is used for changing directory
- `echo` to print anything on the terminal
- `whoami` gives you the current user
- `su` switch to the root user
- `sudo bash` switch to the root user
- `sudo` to access the priviliges
- `sudo useradd demo` will create a user with name demo
- `sudo passwd demo` will set a password for the demo user
- `sudo userdel demo` will delete the user demo
- `sudo groupadd techies` will create a group. A group can have multiple users.
- `sudo groupdel techies` will delete the group. 
- `touch` to create a file
- `cat` to display the content of the file
- `cp filename destination` to copy the file to destination directory
- `mv` to move the file (same as copy) 
- `rm` to remove files
- `mkdir` to create a new directory
- `rmdir` to remove directory
- `grep words filename` to search texts in a particular file
- `sort` will sort the file



## Devops Specific

- `chown` to change the ownership of the file/directory
- `chmod` to change the access permissions of the files/directories
    - `0` means no permissions
    - `1` means execute permissions
    - `2` means write permission
    - `4` means read permission
    - `7` means read, write and execute all three
- `chmod 777 demo.txt` will give the read, write and execute permission to user,groups and any other users.
- `lsof` list of open files and details associated with them
- `lsof -u username` to see the list of the files opened by a particular user
- `id` is used to find out user and the group names and numeric IDs of the current user or any other user
- `tar -cvf filename source-folder` is used to zip of `.tar` format 
- `tar -xvf tarfile` to unzip a file. 
- 

