## Inital Ram Disk

The initramfs filesystem image contains programs and binary files that performs all actions needed to mount porper filesystem, like providing kernal functionality for the needed filesystem and device drivers for mass storage controllers with a facility call udev (**u**ser **dev**ice), which is responsible for figuring out which devices are present, locating the device drivers they to operate properly and loading them. After the root filesystem has been found, it is checked for errors and mounted.

The mount program instructs the operating system that a filesystem that a filesystem is ready for use, and associate it with a particular point in the overall heirarchy of the filesystem (the mount point). If this is successful, the initramfs is cleared from RAM and init program on the root filesystem (`/sbin/init`) is executed.

Init handles the mounting and pivoting over the final real root fulesystem. If special hardware drivers are needed before mass storage can be accessed, they must be in initramfs image.

![initramfs flow](https://courses.edx.org/assets/courseware/v1/13f8548b13ebe15a19aa1a6c3964fceb/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch03_screen22.jpg)

## Text Mode Login
Near the end of the boot process, init starts a no.of text mode login prompts. These enable you to type your username, followed by your password, and to eventually get a command shell. However, if you are running a system graphical login interface, you will not see this at first

![the GNU Bourne Again Shell](https://courses.edx.org/assets/courseware/v1/e35bea5a8c6b9a41453a0e01c5ca3077/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch03_screen26.jpg)

## The Linux Kernal
The Boot Loader loads both the kernal and initial RAM-based file system (initramfs) into memory, so it can be used directly by the kernal.

![Kernal](https://courses.edx.org/assets/courseware/v1/b953394cd3145a1bd239673dc5c5a5b7/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch03_screen21.jpg)

When the kernal is loaded in RAM, it immidiately initalizes and configures the computer memory and also configures all the hardware attached to the system. This includes all processors I/O subsystems, storage devices etc. The kernal also leads some mecessary user space applications.


## /sbin/init
Once the kernal has setup all its hardware and mounted the root filesystem, the kernal runs /sbin/init. This then becomes the initial process, which then starts other processes to get the system running. Most other processes on the system tree their origin ultimately to **init**, exceptions include the so called kernal processes. These are sorted by the kernal directly and their job is to manage internal operating system details.

Besides the system init is reponsible for keeping the system running and for shutting it down cleanly. 

On of its reponsibilities is to act when necessary as a manager for all non kernal process.

![/sbin/init flow](https://courses.edx.org/assets/courseware/v1/640a31713f9fded06718cb06c468f685/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch03_screen24.jpg)


## Startup Alternatives
* **SysVinit** viewed things as a serial process, divided into a series  of sequestial stages. Each stage required completeness before the next could proceed. Thus, startup did not easily take advantage of the parallel processing that could be done on multiple processors or cores.

* **upstart**
* **systemd**

## systemd Features

Systems with **systemd** start up faster than those with earlier **init** methods. This is largely because it replaces a serialized set of steps with aggressive parallelization techniques, which permits multiple services to be initiated simultaneoulsy.

Complicated startup shell scripts are replaced with simpler configuration files, which enumerate what has to be done before a service is started, how to execute service startup, and what conditions the service should indicate have been accomplished when startup is finished. *One this to note is that `/sbin/init` now just points to `/lib/systemd/systemd`.* i.e. systemd takes over init process.

One systemd command (systemctl) is used for most basic tasks.

* $ sudo systemctl start|stop|restart httpd.service (apache web server)

* Enabling or disbaling a system service form starting up at system boot.
sudo systemctl enable|disable httpd.service

## Linux Filesystem Basics

Different typese of filesystems supported by linux:
* Conventional disk filesystems: ext3, ext4, XFS, Btrfs, JFS, NTFS, vfat, exfat, etc.
* Flash storage filesystems: ubifs, jffs2, yaffs, etc,
* Database filesystems
* Special purpose filesystems: proofs, sysfs, tmpfs, squashfs, debugfs, fuse etc.

## Partitions and Filesystems
A partition is a physically contiguous section of a disk, or what appears to be so in some advanced setups.

A filesystem is a method of storing/finding files on a hard disk (usually in a partition).

One can think of a partition as a container in which a filesystem resides, although in some circumstances, a filesystem can span more than one partition if one uses symbolic links.

A comparison between filesystems in windows and linux 

| | Windows | Linux |
| :----: |  :----: | :----: |
|Partition|Disk 1|/dev/sda1|
|Filesystem type|NTFS/VAT|EXT3/EXT4/XFS/BTRFS|
|Mounting Parameters|DriveLetter|MountPoint|
|Base Folder (where OS is stored)| C:\\ | / |


## The Filesystem Hierarchy Standard

Linux systems store their important files according to a standard layout called the Filesystem Hierarchy standard (FHS).

Linux uses the '/' character to separate paths (unlike Windows, which uses '\\'), and does not have drive letters. Multiple drives and/or partitions are mounted as directories in the single filesystem. Removable media such as USB drives and CDs and DVDs will show up as mounted at `/run/media/yourusername/disklabel` for recent linux systems, or under `/media` for older distributions.

![Linux File Heirarchy](https://courses.edx.org/assets/courseware/v1/66def40e2774fd96011565107706da2d/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/dirtree.jpg)

## More About the Filesystem Hierarchy Standard

All Linux filesystem names are case-sensitive, so `/boot`, `/Boot`, `/BOOT` represent three different directories.

![Filesystem example](https://courses.edx.org/assets/courseware/v1/65256a6c88506b6e45744b97b8875775/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/fstree1.png)