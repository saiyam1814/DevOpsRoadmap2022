## Linux Intermediate

- To open the terminal use `ctrl + alt + T`
- To make changes in the orientation of the termainal use `superkey + arrows`
- Superkey is just the alternative name of windows key. 
- To list all the subdirectories while using the ls command we will use `ls -lR /path-of-the-directory`. R means recursively.


## grep and piping 
- There are two methods of redirecting the output. Either you can do it using `|` or `>` 
  - The difference between both is simple. `>` will not print anything to the terminal. It will write the output in a file but in case we are using `|` then the output will be shown on the terminal. 

## Distribution and kernel information 
- When you first login in any system then you might want to check which system and user you're logged in. 
- `whoami` - gives you the current user
- `hostname` - gives you the current hostname of the computer 
- If you want to modify your hostname then you can do it by using the command `sudo vim /etc/hostname`
- id - gives you the user ID and group ID 
- groups <username> - gives you all the groups that the user is part of.
- `lsb_release -a` gives you the distribution specific information
- `lscpu` will give you information related to your CPU.
- `uname` information related to kernel and it stands for unix name.  `-r` flag will give you the version of kernel.

## Shells and bash configuration
- To check you current shell `echo $SHELL`
- To check all your shells in system `cat /etc/shells`
- To write all your aliases use `.bash_aliases` file. 

## Disk Usage
- We will use `du` command which stands for disk usage
- If you will type `du` then it will give you weird results and this is not human readable. 
- To make it more human readable we will use `du -sh *` s means summary and h means human readable. * basically means all the files&directories in the current directory. 
- If you want to check all the directories that are crossing more than 1GB then you can grep it with `du -sh * | grep "G"`
- `df` file system disk space usage 
- `whatis df` 
- It basically shows you how much you have used in comparision to available space. 
- `df -ht ext4` this basically shows you only the linux file system. 



## Linking files
- Something similar to windows shortcut option
- `ln <original> <link>` This is used to create hard link
- `ln -s <original> <link>` This is used to create soft link. 

## Passing multiple arguments in linux
- xargs is used to pass multiple arguments in linux.
- `ls | xargs cat` In this case xargs takes input from the ls command and passes that to the cat command. 
- `ls | xargs wc -l` We will get count of lines of all the files present in the current directory. 
