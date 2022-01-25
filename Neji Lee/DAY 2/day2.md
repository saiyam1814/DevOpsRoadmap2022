## DAY 2

      

## 16-01-2022

 1. Completed the Chapter3 of the linux. 
 2. Larned about the Fileystem Hierachy Standard in Linuix

-   / - The root directory of the entire filesystem hierarchy, everything is nestled under this directory.
-   /bin - Essential ready-to-run programs (binaries), includes the most basic commands such as ls and cp.
-   /boot - Contains kernel boot loader files.
-   /dev - Device files.
-   /etc - Core system configuration directory, should hold only configuration files and not any binaries.
-   /home - Personal directories for users, holds your documents, files, settings, etc.
-   /lib - Holds library files that binaries can use.
-   /media - Used as an attachment point for removable media like USB drives.
-   /mnt - Temporarily mounted filesystems.
-   /opt - Optional application software packages.
-   /proc - Information about currently running processes.
-   /root - The root user's home directory.
-   /run - Information about the running system since the last boot.
-   /sbin - Contains essential system binaries, usually can only be ran by root.
-   /srv - Site-specific data which are served by the system.
-   /tmp - Storage for temporary files
-   /usr - This is unfortunately named, most often it does not contain user files in the sense of a home folder. This is meant for user installed software and utilities, however that is not to say you can't add personal directories in there. Inside this directory are sub-directories for /usr/bin, /usr/local, etc.
-   /var - Variable directory, it's used for system logging, user tracking, caches, etc. Basically anything that is subject to change all the time.
