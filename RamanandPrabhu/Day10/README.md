#Day10
* Contd Linux Foundation's Intoduction to Linux and learnt about.
  * Linux Command Line 
  - sudo : This command allows us to run programs using the security privileges of another user, usually root. The functionality of sudo is similar to 'run as' in Windows
	How to setup SUDO for a user:
	- We need root credentials to do this. Type su to swith user to root and then type the root password when prompted. The promt will change to #
	- We need a configuration file in /etc/sudoers.d/ directory with filename as username. eg. if username is johndoe then file should be /etc/sudoers.d/johndoe
	- We create file by doing # echo "johndoe ALL=(ALL) ALL" > /etc/sudoers.d/johndoe. This will create the file for the  user johndoe. 
  * Login and Logout
	- An available text terminal will prompt for username and password. Once the session is started you can also connect to remote systems using SecureShell(SSH) utility.
  * Rebooting and Shutting Down
    - 'shutdown' command is used to reboot or halt the system. 'shutdown -r' will reboot the system and 'shutdown -h' will halt the system.
	- Both rebooting and shutdown from the command line requires superuser(root) access.
	- We can send notification prior to halting in a multiuser system by 'sudo shutdown -h 10:00 "Shutting down for scheduled maintenance."
  * Locating applications
    - Programs and software packages are generally installed in /bin, /usr/bin, /sbin, /usr/sbin directories or under /opt
	- To locate program we can use 'which' or 'whereis" utility. eg 'which java' or 'whereis java'
  * Accessing Directories
    - When you first login into system or open a terminal the default directory should be your home directory. You can know the exact path by typing echo $HOME.
	- Some commands, pwd (present working directoy), cd ~ or cd (Change to your home directory), cd .. (Change to parent directory), cd - (Change to previous directory)

  * sudo telinit 3 is used to turnoff Graphical desktop on an RPM-based system.
  
  * Absolute and Relative path :
	- Absolute Path starts with a /(from root)
	- Relative path starts from the present working directory. (. present directory, .. parent directory, ~ home directory)
  
  * Exploring files : 'ls' lists the contents of the pwd, 'ls -a' even lists the hidden files and directories. 'tree' gives birds eye view of the directory. 'tree -d' list the directory and suppress the file names. On CentOS, 'yum install tree' to install tree.
  
  * Hard and Soft(aka Symbolic) links
    - 'ln' command is used to create hard link. 'ln -s' will give a soft link. 
	$ ln file1 file2 - will create hard link from file1 to file2
	$ ls -il - this will show the inode in the first column and it will be same for both the files.
	
	- We can not create a soft link with same name as hard link.
	  ln -s file1 file2 will fail saying link already exists
	- Symbolic links take no extra space on the filesystem (unless their names are very long). They are extremely convenient as they can easily be modified to point to different places. An easy way to create a shortcut from your home directory to long pathnames is to create a symbolic link.
	- Unlike hard links, soft links can point to objects even on different filesystems (or partitions) which may or may not be currently available or even exist. In the case where the link does not point to a currently available or existing object, you obtain a dangling link.
	- Hard links are very useful and they save space, but you have to be careful with their use, sometimes in subtle ways. For one thing if you remove either file1 or file2 in the example on the previous screen, the inode object (and the remaining file name) will remain, which might be undesirable as it may lead to subtle errors later if you recreate a file of that name.
	- If you edit one of the files, exactly what happens depends on your editor; most editors including vi and gedit will retain the link by default but it is possible that modifying one of the names may break the link and result in the creation of two objects.
	
	
    
	
 # To be cont