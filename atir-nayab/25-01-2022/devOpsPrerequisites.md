# DevOps prerequisites

**CARRICULUM**
* LINUX BASICS
* SETUP LAB ENVIRONMENT
* VIRTUAL BOX
* VAGRANT
* LINUX NETWORKING
* JSON/YAML

## Basic Linux Command
* Linux CLI
* VI editor
* Package Management
* Service Management

### Shell Types
```bash
echo $SHELL
```
To print shell type being used
* Bourne Shell (sh Shell)
* C Shell (csh or tcsh)
* Z Shell (zsh)
* Bourne again Shell (bash)

### Basic Commands
- echo - print to screen
- ls - list files & folders
- cd - change directory
- pwd - present working directory
- mkdir -  make directory
- cd new_directory;mkdir www; pwd - multiple commands
 - mkdir -p /tmp/asia/india/delhi - make directory hierarchy one under other
 - rm -r /tmp/my_dir1 - remove directory child with its parent

 #### Working with files
 - touch new_file.txt - ceate a new blank file
 - cat > new_file.txt - add contents to file
 - cat new_file.txt - view contents of file
 - cp newfile.txt copy_file.txt - copy file
 - mv mew_file.txt sample_file.txt - move (rename) file 
 - rm new_file.txt - remove (delete) file

#### User Accounts
- whoami - logged in user name
- id - to get user id, group id
- su aparna - switch to other user
- ssh aparna@192.12.22.23 - ssh into other linux machine
- ls /root - root directory
- sudo ls/root - grant root access to user


#### Download Files
- curl url -O - downlaod file and save to a file
- wget url - also used to download files

#### Check OS Version
- ls /etc/\*release* - to check for current OS
- cat /etc/\*release* - to get more details about OS.

#### Package Managers
RPM(Red Hat Package Manger)

- rpm -i telnel.rpm - install package
- rpm -e telnel.rpm - unistall package
- rpm -q telnel.rpm - query package

#### YUM
A high level package manager that manages all the dependencies of a package

- yum install ansible - install package along with it's dependencies

*sometimes we need update the repositories and existing repositories don't have the software we need*

- yum repolist - shows repo list
- ls /etc/yum.repos.d - show the files where these repose are configured
- yum list ansible - to show all dependencies of a package
- yum remove ansible - to remove a package
- yum --showduplicates list ansible - show duplicate package
- yum install ansible-2.4.2.0 - to package with particular version

#### Services
- service httpd start or systemctl start httpd - to start HTTPD service
- systemctl stop httpd - stop httpd service
- systemctl status httpd - to check httpd service
- systemctl enable httpd - configure httpd to start at startup
- systemctl disbale httpd - configure httpd to not start at startup

**Configure a program to run as a service**
in `/etc/systemd/system` create a file my_app.service
```service
[Unit]
Description=My python web application

[service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py
<!-- run before running application -->
execStartPre = /opt/code/configure_db.sh
<!-- run after running application -->
execStartPost=/opt/code/emai_status.sh
<!-- restart application -->
Restart=always 

<!-- auto start after boot -->
[Install]
WantedBy=multi-user.target
```
`systemctl daemon-reload`
`systemctl start my-app`
`systemctl stop my_app`

#### Working with Files
* VI Editor
`vi index.html` to open file vi editor

* command mode
- mode around: with arrow keys or H,J,K,L
- delete : x or dd
- copy and paste: yy and p
- scroll up/down: ctrl+u /ctrl+d
- command: :(colon)
- save: :w or :w (filename)
- Quit(discard) :q
- save+quit: :wq 
- find: /of (example word to find) n(to go to next word)

#### IP Address
- ip addr show - to show ip address
- ip addr add 192.168.1.10/24 dev eth0 to add ip address

#### start ssh service
- service sshd status - to check status of ssh daemon
- service sshd start - to start ssh service

#### Networking
**Objectives**
* Networking in VirtualBox
* Networking Adapters
* NAT
* Bridge
* Host Only
* Internet Connectivity

- IP Address
It is assigned to interface
`id addr show` to display ip addresses.

- Host only network
It creates a virtual network inside a machine and each virtual box is assigned an ip address. But they can't reach out to outer world.

- NAT (Network Address Translation)
It allows vms to reach out the world.
It basically changes ip of vms to its own if request is sent and when request is received it changes the ip back to vm ip.

The only difference between NAT and NAT network is in NAT vms are isolated that means they can talk to each other but they can reach outer world with the help of host. Whereas in NAT network each vms can communicate to each other.

- Bridge Network
When external resources want to reach out to vms they can do with the help of bridge network. Vm will assigned an ip in the range of LAN. It looks like a another machine in same network like host machine.

- Internet Connectivity
  - NAT - possible
  - Bridged - possible
  - Host-Only Network - not possible, but it possible through port forwading that redirecting traffic from one network to other network. Basically out laptop/machine will act as router

- Port Forwading
It allow us to map a port on host to the port on vm eg. port 80 on host can be mapped with port 80 on vm. It means all traffic on port 80 on host will be forwarded to post 80 on vm

#### VB multiple VMs
types of clone

- full clone
exact replica will made, but with separate storage space.So each vm will occupy same amount of space.

- linked clone
used disk space of existing vm and only use extra space for the changes made in new vm.

#### Virtual Box Networking

|VM can reach internet/other systems in the network|vms can reach each other|Host can reach vm (without port forwarding|other system in the network can react vm|solution|
|:-:|:-:|:-:|:-:|:-:|
|Yes|No|No|No|NAT|
|Yes|Yes|No|No|NAT Network|
|No|Yes|Yes|No|Host Network|
|Yes|Yes|Yes|No|Host Network + NAT|
|Yes|Yes|Yes|Yes|Bridged|

#### VM snapshot
It used to take a snapshot of a vm and resotre to that snapshot later point in that time.
We can even clone a vm from snapshot

#### Vagrant
- Download
- create VM
- Create Networks
- Configure Networking
- Configure Port Forwarding
- Boot up VM

Vagrant help us to automate all manual task from downloading  to port forwarding n all.

#### Vagrant init
After downloading `vagrant init centos/7` to deloy. eg centos/7 box
what is a box?
It is a packaged format of vagrant environment. It contains os image as well script requried to configure os.

`ls`
Vagrantfile

`vagrant up` 
to start a vm

Vagrant commands
- init - inititalizes a new vagrant environment by creating a vagrantfile
- up - starts and provisions the vagrant evironment
- suspend - suspends the machine
- resume - resume a suspended vagrant machine
- halt - stops the vagrant machine
- destroy - stops and deletes all traces of the vagrant machine
- status - outputs status of the vagrant machine
- reload - restarts vagrant machine, loads new vagrantfile configuration
- snapshot -  manages snapshots: saving, restoring etc.


#### Vagrantfile
vagrant config file

```
Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"
  config.vm.network "forwarded_port, guest: 80, host:8080
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "1024"
  end
  config.vm.provision "shell", inlint: <<-SHELL
    apt-get update
    apt-get install -y apache2
  SHELL
end
```
Vagrant Providers
- VirtualBox
- VMware
- Hyper-V
- Docker
- Custom

## DNS
adding name for IP can be by adding to `/etc/hosts` file

```
cat >> /etc/hosts
192.168.1.11  db
```
when we try to ping or ssh through it look host file and find it's ip
Translating IP address this way is called name resolution 
Managing the host name centrally is through DNS server
Every host has the ip of dns server at `/etc/resolv.conf` 
A host is able to map name to ip address locally and through dns server

## Domain name
eg www.google.com

- Root .(dot)
- top level domain name  - .com
- goole
- subdomain - mail, drive, www, maps, apps

## DNS loopkup hierarchy

1. Org dns - it can also cache the ip for future use
1. root dns
1. .com dns
1. googel dns
1. sends back the ip

you can also add search in `/etc/resolv.conf` so that it can appended to every name we search for

## Record Types

- A - namping name to ipv4
- AAAA - maping name ipv6
- CNAME - mapping name to other name or can be multiple aliases

## nslookup
ping may not be the right tool to test dns resolution. We can use nslookup to query a host name from a dns server. It will not consider the entry in our local host file. nslookup only queries only the dns server

## dig
It is another tool to test dns resolution

## Networking Pre-Requisites
- Switching
- Routing
- Default Gateway
- DNS Configuraions on linux

### Switch
It deliver packets between two devices within a network

### Router
It help connect two network. So that device in one network connects with device on  other network

### Gateway
When a device wants to connect with a device on other network through router, but that device coudn't know where the router is it can be any where as it is just a device connected to the network. That's where we configure a network with a gateway or a route. A gateway is like a door where as network is like a room. A door to outside world

To view router config
`route`

To add gateway
`ip route add 192.168.2.0/24 via 192.168.1.1` 
run `route` again you can see route config

or you can all unknow network to which the network don't know the route to use route as default gateway
`ip route add default via 192.168.2.1`
instead of default 0.0.0.0 is for default

### Make linux act as a router 
A ---- B ----- c

A wants to connect with C

A will default gateway so will c
In B we have configure so that packets from one interface is forwaded to other interface

`/proc/sys/net/ipv4/ip_forward` this file contains config for portfording
default is 0
set it to 1 and A can ping C
However this change is temporary to make it permanent we have to do changes in `/etc/sysctl.conf`
```
net.ipv4.ip_forward = 1
```

## Take Aways
- `ip link` - is to list and modify interfaces on the host
- `ip addr` - is used to see ip addresses assigned to those addresses
- `ip addr add 192.168.1.10/24 dev eth0` is used add ip to interfaces
-`ip route` or `route` - used to view routing table
- `ip route add 192.168.1.10/24 via 192.168.2.1` - add entries in routing table
- `cat /proc/sys/net/ipv4/ip_forward` - used to check if route forwarding is enabled or not

