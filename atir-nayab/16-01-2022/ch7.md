# Command Lint Operations

**main Highlights**
* Use the command line to perform operations in linux.
* Search for files.
* Create and manage files.
* Install and update software.

## Introduction to the Command Line
There is a saying, *"graphical user interface make easy tasks easier, while command line interfaces make difficult tasks possible".*
* No GUI overhead is incurred.
* Virtually any and every task can be accomplished while sitting at the command line.
* You can implement scripts for often-used (or easy=to=forget) tasks and series of procedures.
* You can sign into remote machines anywhere on the internet.
* You can initiate graphical applications directly from the command line instead of hunting through menus.
* While graphical tools amy vary among Linux distributions, the command line interface does not.

By default, on GNOME desktop environments, the gnome-terminal applications is used to emulate a text-mode terminal in a windows.
  * xterm
  * lnsole (default on KDE)
  * terminatior

## Some Basic Utilities

* cat: used to type out a  file (or combine files)
* head: used to show the first few lines of a file.
* tail: used to show the last few lines of a file.
* man: used to view documentation

## The Command Line

Most input lines entered at the shell prompt have three basic elements.
* Command
* Options
* Arguments

Options ususally start with one or two dashes.

## Sudo
sudo allows users to run programs usign the security privileges of another user, generally root (superuser)

## Steps for setting Up and Running sudo
If your system does not already have sudo set up and enabled,
1. At the command line prompt type su and press Enter, type your password. You should end up with a different looking prompt, often ending '#' eg
**su password:**
**#**
2. Now create a configuration file to enable your user account to use sudo. Typically this file is created in the `/etc/sudoers.d` directory with the name of the file the same as your username.
eg
**# echo "student All=(ALL) ALL" > /etc/sudoers.d/student**
3. Finally, some Linux distributions will complain if you do not also change permissions on the file by doing:
**# chmod 440 /etc/sudoers/d/student**

## Switching Between the GUI and the Command Line

The customizable nature of linux allows you to drop the graphical interface (temporarily or permanently) or to start it up after the system has been running

## Virtual Terminals (VT) 
VT are console sessions that use the entire display and keyboard outside of a graphical environment. Such terminals are considered "virtual" because, although there can be multiple active terminals, only one terminal remains visible at a time. A VT is not quite the same as a command line terminal window.

![Switching between Virtual Terminals](https://courses.edx.org/assets/courseware/v1/cce9159be8b08390567dc02f1043cf92/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch06_screen07.jpg)

## Turning On or Off the Graphical Desktop

Linux distributions can start and stop teh graphical desktop in various ways, 

**sudo systemctl stop gdm**

and restart it 

**sudo systemctl start gdm**