### Basic Linux Commands
- `ls`            => List files and Folders
- `echo Hi`       => Prints Hi to screen
- `cd my_dir`     => Change directory to my_dir
- `pwd`           => Shows present working directory
- `mkdir new_dir` => Makes new_directory
- `mkdir -p tmo/asia/india/bangalore` => make directory hierarchy
- `rm -r /tmp/my_dir` => remove directory
- `cp -r my_dir /tmp/my_dir`  => copy directory

**Run multiple commands => seperate them by semicolon eg: `cd new_dir; mkdir www; pwd`**
### Working with files
- `touch new_file.txt`  => create new file
- `cat > new_file.txt`  => add contents to file
- `cat new_file.txt`    => view contents of file
- `cp new_file.txt copy_file.txt` => copy file
- `mv new_file.txt move_file.txt` => move (rename) file
- `rm new_file.txt`     => remove (delete) file
### Users
- `whoami`  => find which user we are
- `id`      => more information about the user
- `su aparna` => switch user to aparna
- `ssh aparna@192.168.1.2` => ssh into another system with different user  
- `sudo`    => prefix to elevate permission level to perform tasks of root user
- `curl http://www.some-site.com/some-file/txt -O` => download a file and save it using -O otherwise it will just print it
- `wget http://www.some-site.com/some-file/txt -O some-file.txt`  => also to download a file
- `ls /etc/*release*` => check release files to know what OS we are on
- `cat /etc/*release*`  => check OS version in more detail
### Package manager
- `rpm -i telnet.rpm` => install package telnet
- `rpm -e telnet.rpm` => unistall package
- `rpm -q telnet.rpm` => query package
- `yum install ansible` => installs ansible and all of its dependencies too!
- `yum repolist`  => see repos available on system
- `yum list ansible`  => list installed and available ansible packages
- `yum remove ansible`  => remove package
- `yum --showduplicates list ansible` => list all available packages for ansible
- `yum install ansible-2.4.2.0` => install specific version of ansible package
### Service
**Service in linux help us configure software to run in the background even when system is rebooted and make sure they
follow right order of startup**
- `service httpd start` => Start HTTPD service
- `systemctl start httpd` => newer method to start service on systemD managed server
- `systemctl stop httpd`  => Stop httpd service
- `systemctl status httpd`  => check httpd service status
- `systemctl enable httpd`  => configure https to start at startup
- `systemctl disable httpd` => configure httpd to not start at startup
#### Create new Service
.service files generally go in `/etc/systemd/system`

`my_app.service`
```
[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py # command to run
```
- `systemctl daemon-reload` => let systemd know we have a new service
- `systemctl start my_app`  => systemctl start <name of service
- `systemctl status my_app` => check status of our service
- `systemctl stop my_app`   => stop service

#### More options in service file
```
[Unit]
Description=My python web application           #some metadata

[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py  # command to run
ExecStartPre=/opt/code/configure_db.sh          # dependency to run before ExecStart
ExecStartPost=/opt/code/email_status.sh         # dependency to run after ExecStart
Restart=always                                  # always restart when it crashes

[Install]
WantedBy=multi-usser.target                     # run after the multi-user.target run level is started
```

## Vi editor
- `vi index.html` => open a file with vi
- `i`             => switch to insert mode
- `esc`           => switch to command mode

#### Command mode:
- `H J K L` or `arrow keys` => move around
- `X`             => delete character
- `dd`            => delete entire line
- `yy`            => copy a line
- `p`             => paste
- `CTRL+u` `CTRL+d` => scroll page
- `:`             => open prompt
- `:w` or `:w filename` => save
- `:q`            => discard changes and quit
- `:wq`           => save changes and quit
- `/of`           => find the word "of" in file
- `n`             => move cursor to next occurence of word "of"
  
  
