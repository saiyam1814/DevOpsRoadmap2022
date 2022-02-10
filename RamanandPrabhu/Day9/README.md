#Day9
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
 # To be contd