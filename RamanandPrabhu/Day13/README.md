#Day13
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
  * Navigating the directories
    - Use pushd to change the directory instead of cd; this pushes your starting directory onto a list.
	- Use popd to pop the last pushed.
	- Use dirs to display the pushed directories.
	
  * Difference btw . and pwd
    - . refers to the current directory used in shell scripting and is relative path.
	- pwd prints the current working directory and it is absolute path.
	
  * dirs command options
    - dirs -p : prints the remembered directories one entry per line.
	- dirs -v : prints the remembered directories one entry per line with order.
	- dirs -c : clears the remembered directories.
	- dirs -l : prints the absolute path for ~ in remembered directories.
	
  * How to shutdown system without rebooting(assuming you have root privileges)
    - halt, poweroff, shutdown -h now
  * Which commands allow you to locate programs?
    - which and whereis
	
  * Standard File Streams
  - What are Streams: In computer programming, standard streams are interconnected input and output communication channels btw a computer program and its environmenet when it begins execution. Originally I/O happened via physically connected devices(monitor, keyboard), but standard streams abstract this. When a command is executed via an interactive shell, the streams are typically connected to the text terminal on which the shell is running, but can be changed with redirection or a pipeline. More generally, a child process inherits the standard streams of its parent process.
  
  Usually users know streams as I/O channels that handle the data coming from an input device, or that write data from application. Streams may be used to chain application, meaning the output from one program can be rediredted to be the input steam for another application. In many OSs this is expressed by pipeline character.
  




 # To be contd