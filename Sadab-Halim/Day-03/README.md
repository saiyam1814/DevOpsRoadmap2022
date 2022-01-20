# 03 January 2022

## Completed Chapter 4: Graphical Interface of [Introduction to Linux](https://www.edx.org/course/introduction-to-linux)

### Graphical Desktop
We can use either a Command Line Interface (CLI) or Graphical User Interface (GUI) using Linux.

### X Window System
In a Linux desktop system, the X Window System is loaded as one of the final steps in the boot process. If is often just called X.
A service called the Display Manager keeps track of the displays being provided and loads the X server. The display manager also handles graphical logins and starts the appropriate desktop environment after a user logs in.

![image](https://user-images.githubusercontent.com/74575612/150348457-323f68a9-d081-4480-97dc-0291de8e6256.png)

If the display manager is not started by default in the default runlevel, we can start the graphical desktop different way, after logging on to a text-mode console, by running startx from the command line. Or, you can start the display manager (gdm, lightdm, kdm, xdm etc)

![image](https://user-images.githubusercontent.com/74575612/150348575-f748ead2-9691-446a-8f24-fe5fa6dc0625.png)

### GNOME Desktop Environment
GNOME is a popular desktop environment with an easy-to-use graphical user interface. It is bundled as the default desktop environment for most Linux distributions, including Red Hat Enterprise Linux (RHEL), Fedora, CentOs, SUSE Linux Enterprise, Ubuntu and Debian.

All the trash files are stored in `~/.local/share/Trash`


## Completed Chapter 5: System Configuration from the Graphical Interface of [Introduction to Linux](https://www.edx.org/course/introduction-to-linux)

### Debian Packaging
dpkg is the underlying package manager for these systems. It can instal, remove, and build packages. Unlike higher-level package management systems.

### Red Hat Package Manager
Red Hat Package Manager (RPM) is the other package management system popular linux distributions. It was developed by Red Hat, and adopted by a number of other disttibutions, including SUSE/openSUSE, Mageia, CentOS, Oracle Linux, and others.

### open SUSE's Software Management
The Yet another Setup Tool (YaST) software manager is similar to other graphical package managers. It is an RPM-based application.


## Completed Chapter 6: Common Applications of [Introduction to Linux](https://www.edx.org/course/introduction-to-linux)

### Internet Applications

![image](https://user-images.githubusercontent.com/74575612/150349425-2fe06bf4-29e0-4c66-8163-c91f8174a862.png)

