# Types of shells
* Bourne shell (sh shell)
* C shell (csh shell)
* Z shell (zsh shell)
* Bourne again shell (bash shell)

# linux commands

### basic commands
1. `echo` - print a string
2. `ls` - list files in a directory
3. `cd` - change directory
4. `pwd` - print working directory
5. `mkdir` - make a directory
6. `rmdir` - remove a directory
7. `rm` - remove a file
8. `cp` - copy a file
9. `mv` - move a file
10. `cat` - concatenate files and print on the standard output
11. `grep` - search for a pattern in a file

### User accounts
1. `whoami` - print the name of the current user
2. `id` - print the user ID of the current user
3. `sudo` - run a command as root
4. `su` - run a command as another user

### Download files
1. `curl <url>` - download a file from the internet
2. `wget <url>` - download a file from the internet


### Check OS version
1. `uname -a` - print operating system information
2. `lsb_release -a` - print operating system information

# Package managers
RPM: Red Hat Package Manager

1. `rpm -i <package>` - install a package
2. `rpm -e <package>` - remove a package
3. `rpm -q <package>` - query a package

YUM: Yum Package Manager
1. `yum install <package>` - install a package
2. `yum remove <package>` - remove a package
3. `yum list <package>` - query a package
4. `yum repolist` - list available repositories

# Services
1. `service <service> start` - start a service
2. `systemctl start <service>` - start a service
3. `systemctl stop <service>` - stop a service
4. `systemctl status <service>` - check the status of a service
5. `systemctl enable <service>` - enable a service to start at boot
6. `systemctl disable <service>` - disable a service to start at boot


# VI editor
1. `vi <file>` - open a file in vi editor
2. to switch to insert mode: `i`
3. to switch to command mode: `esc`
4. to switch to visual mode: `v`
5. to switch to replace mode: `r`
6. to delete a character: `x`
7. to delete a line: `dd`
8. to delete a word: `dw`
9. to copy a line: `yy`
10. to paste a line: `p`
11. to save a file: `:wq`
12. to quit vi: `:q`
13. to quit vi without saving: `:wq!`
14. to quit vi and exit: `:q!`

# Networking
1. `ifconfig` - print network interface configuration
2. `route` - print routing table
3. `netstat` - print network statistics
4. `ping <host>` - send ICMP echo request to a host
5. `traceroute <host>` - trace the route to a host
6. `dig <host>` - query DNS information
7. `nslookup <host>` - query DNS information
8. `nmap <host>` - scan a network
9. `ss` - print network statistics

# Routing
1. `route add -net <network> <gateway>` - add a route
2. `route del -net <network>` - delete a route
3. `route add -host <host> <gateway>` - add a route

