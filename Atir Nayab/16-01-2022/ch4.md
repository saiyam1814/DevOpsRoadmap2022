# Graphical Interface

**Main Highlighes**
* Manage graphical interface sessions.
* Perform basic operations using the graphical interface.
* Change the graphical desktop to suit your needs.

## Graphical Desktop
You use either a Command Line Interface (CLI) or a Graphical User Interface (GUI) when using linux.

## X Window System
In a Linux desktop system, the X Window System is loaded as one of the final steps in the boot process. If is often just called X.

A service called the Display Manager keeps track of the displays being provided and loads the X server. The display manager also handles graphical logins and starts the appropriate desktop environment after a user logs in.

![Display Manager](https://courses.edx.org/assets/courseware/v1/44717c86868ff7e9edc71c5747bb84ab/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch03_screen28.jpg)

## More About X
If the display manager is not started by default in the default runlevel, you can start the graphical desktop different way, after loggin on to a text-mode console, by running startx from the command line. Or, you can start the display manager (gdm, lightdm, kdm, xdm etc)

![Desktop Environment](https://courses.edx.org/assets/courseware/v1/c4a2925d0a2d22c238c9f1d91f71635b/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch03_screen29.jpg)

## GNOME Desktop Environment
GNOME is a popular desktop environment with an easy-to-use graphical user interface. It is bundled as the default desktop environment for most Linux distributions, including Red Hat Enterprise Linux (RHEL), Fedora, CentOs, SUSE Linux Enterprise, Ubuntu and Debian.

## gnome-tweaks
There is a standard utility, gnome-tweaks, which exposes many more setting options. It also permits you to easily install extensions by external parties.

![gnome-tweaks](https://courses.edx.org/assets/courseware/v1/b9aeb9e063eda9567443ab77501286d3/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/gnometweaktool.png)

**Fing the lastest modified file in /var/log**

**trash files are stored in `~/.local/share/Trash`**