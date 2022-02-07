## Basic Linux Commands

| Command | Function |
| --- | --- |
| echo hi | Prints the word “hi” |
| ls | list all the current files and folders of a directory |
| cd  | changes directory |
| cd - | changes the position to the last directory |
| pwd | prints the present working directory |
| mkdir | stands for make directory. Creates a new directory. |
| cd Pictures/; mkdir www; pwd | To run multiple commands we use ; . In the mentioned command the directory is changed to ‘Pictures’ and a new directory ‘www’ is created and the present working directory is printed. |
| mkdir -p Pictures/new/fruits/mango | mkdir -p creates a hierarchical directory. |
| rm file.txt | removes a file |
| rm -r | To recursively remove a directory |
| cp -r Pictures/new/fruits/mango | copies a directory |
| touch file.txt | A new file with a .txt extension is created |
| cat >> file.txt | Adds content in the file. >> symbol is the input symbol. Whatever we type after the mentioned command will be added to the file.txt until the CTRL+D button is pressed. |
| cat file.txt | To view the content of the file |
| cp source_path&name t arget_path&name | copies the content from source path to target path |
| mv source_path&name  target_path&name | moves a file |
| whoami | Displays the current user |
| su other_user | switches to other users, after you enter the password |
| ssh user@host-address | Remote login to the host computer |
| sudo | some tasks can be performed by the root user or superuser. You need to type the command sudo and enter the password to perform the task |
| curl “URL” -O | Downloads the file content from the URL. O is used to save the result in a file locally. |
| rpm -i package.rpm | installs a package.  |
| rpm -e package.rpm | Uninstall package |
| rpm -q package.rpm | Query package |
| yum repolist | To see the list of repositories available on the system |
| yum list package-name | To see a list of installed/available packages |
| yum remove package-name | To remove a package |
| yum —showduplicates list package-name | Lists the available versions of the package |
| service httpd start | Start HTTPD service |
| systemctl start httpd | Start HTTPD service |
| systemctl stop httpd | Stop HTTPD service |
| systemctl status httpd | check HTTPD service status |
| systemctl enable httpd | configure HTTPD to start at startup |
| systemctl disable httpd | configure HTTPD to not start at startup |
| systemctl daemon-reload | To let systemd know a new service is configured  |
| vi file-name | Open a file in the vi editor |

### RPM

RPM stands for red-hat package manager and it belongs to the red-hat distributions. It doesn't install dependencies with the package. For example, if you are installing ansible and if it needs python as its dependency, but if you don't have python or let’s say any other dependent library, it doesn't bother to install.

### YUM

YUM is a high-level package manager which uses RPM. It’s somewhat a solution to RPM where it queries the package, finds its location and installs all dependencies with the package.

### Services

A Linux service is an application (or set of applications) that runs in the background waiting to be used or carrying out essential tasks. Service command uses systemctl command underneath.

### Systemctl

The systemctl command is a utility that is responsible for examining and controlling the systemd system and service manager. It is a collection of system management libraries, utilities, and daemons that function as a successor to the System V init daemon. The new systemctl commands have proven quite useful in managing a server's services. It provides detailed information about specific systemd services and others that have server-wide utilization.

### Configuring program as Systemd services

We know systemctl command line utility is used to manage the systemd services. In order to do that we must configure our program as systemd services. The files of systemd are located mostly at

**/etc/systemd/system** 

example of a configuration file named my_app.service

```jsx
[Unit]
Description=my python web application
#additional metadata

[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py
#executes the program
ExecStartPre=/opt/configure_db.sh
ExecStartPost=/opt/email_status.sh
Restart=always
#ExecStartPre is the dependency to start the script before the start of application
#ExecStopPre dependency is needed after the start of application
#in case of crash, the application will restart again

[install]
WantedBy=multi-user.target
#configure service to start during bootup
```

Now run **systemctl daemon-reload** to let systemd know a new service is configured.
