# Linux

### Graphical User Interface
Wec can use either a Command Line Interface (CLI) or a Graphical User Interface (GUI) when using linux.

### X Window System
In a Linux desktop system, the X Window System is loaded as one of the final steps in the boot process. If is often just called X.

A service called the Display Manager keeps track of the displays being provided and loads the X server. The display manager also handles graphical logins and starts the appropriate desktop environment after a user logs in.

### GNOME Desktop Environment
GNOME is a popular desktop environment with an easy-to-use graphical user interface. It is bundled as the default desktop environment for most Linux distributions, including Red Hat Enterprise Linux (RHEL), Fedora, CentOs, SUSE Linux Enterprise, Ubuntu and Debian.

gnome-tweaks
There is a standard utility, gnome-tweaks, which exposes many more setting options. It also permits you to easily install extensions by external parties.

gnome-tweaks

Find the lastest modified file in `/var/log` <br/>
trash files are stored in `~/.local/share/Trash`

### Linux Package Managers

**Red Hat Package Manager (RPM):**
- `rpm -i telnet.rpm` install package
- `rpm -e telnet.rpm` uninstall package
- `rpm -q telnet.rpm` query package

**YUM:**
YUM Package Manager user RPM underneath.
- `yum isntall ansible` install package

YUM searches software repositories that acts as warehouses that cotains RPM package files. These repositores can be local, available on the internet or in cloud.

When we try to install a package using YUM, YUM searches the repositories finds the required packages and dependencies and installs all of them in the right order.

![image](https://user-images.githubusercontent.com/74575612/150744613-8aec7737-62a7-49d4-bf49-7e40ac0215b7.png)

**How does YUM find where a particular package is located?**
The information about the repository in a configuration file at path `/etc/yum.repos.d` directory.

![image](https://user-images.githubusercontent.com/74575612/150745125-eb82b26c-272e-4aec-b5a2-2f166436ee76.png)

**YUM Repos**
- `yum repolist`
![image](https://user-images.githubusercontent.com/74575612/150745332-cbd40de0-e4bc-4e3f-8824-a4b20c347acc.png)

- `yum list ansible` lists the mentioned package
- `yum remove ansible` removes the mentioned package
- `yum --showduplicates list ansible` to list all available packages
- `yum install ansible-2.4.2.0` to install a particular version of the package

### Linux Services Configuration

- `service httpd start` or `systemctl start httpd` Start HTTPD service
- `systemctl stop httpd` Stop HTTPD service
- `systemctl status httpd` Check HTTP service status
- `systemctl enable httpd` Configure HTTPD to start at startup
- `systemctl disable httpd` Configure HTTPD to not start at startup

Resources:
- [Introductiont to Linux](https://www.edx.org/course/introduction-to-linux)
- [DevOps Prerequisite Course](https://www.youtube.com/watch?v=Wvf0mBNGjXY)