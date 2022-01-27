# The pieces of the puzzle
As per my understanding, linux package management is fragmented into : **underlying lower and higher level package management systems**.
- The underlying system is responsible for installing, removing and building local packages and is devoid of any remote capabilities.
- The higher level system utilises the underlying system and adds remote capabilites by having a distribution specific repository from which packages along with their dependecies can be pulled from.


### Debian
- dpkg is the underlying package manager for debian based systems.
- Advanced Package Tool (APT) is the higher-level package management system for debian based systems.
- Generally, distributions create their own user interface on top of it (for example, apt and apt-get, synaptic, gnome-software, Ubuntu Software Center, etc)
- Although apt repositories are generally compatible with each other, the software they contain generally is not. Therefore, most repositories target a particular distribution (like Ubuntu), and often software distributors ship with multiple repositories to support multiple distributions.

### Red Hat & SUSE
- RPM (Red Hat Package Manager ) is the underlying package manager for Red Hat and Suse family of distributions.
- The higher level package management differs in between distributions. Red Hat family (RHEL/CentOS and Fedora) use dnf, while retaining good backwards compatibility with the older yum program.
- SUSE family use zypper interface.
- openSUSE has Yet another Setup Tool (YaST) as a graphical package manager based on RPM.