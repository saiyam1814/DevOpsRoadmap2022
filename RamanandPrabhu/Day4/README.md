#Day4
* Contd Linux Foundation's Intoduction to Linux and learnt about.
  * Linux FileSystem -(Embodiment of method of storing and organizing arbitrary collection of data in human readable form.)
	-	Conventional disk filesystem : ext2, ext3, ext4, NTFS, XFS, Btrfs, JFS, etc.
	-	Flash storage system : ubifs, JFFS2, YAFFS, etc.
	-	Database filesystem.
	-	Special purpose filesystem : procfs, sysfs, tmpfs, YAFFS, etc.
  * Partitions and FileSystem
	-	Partition : A logical part on the disk.
	-	Filesystem : A method of storing/finding files on hard disk usually in a partition.By way of analogy, we can think of filesystems as being like family trees that show descendants and their relationships, while the partitions are like different families (each of which has its own tree).
	-	Linux system stores important filesystem according to a standared called FHS or Filesystem Hieararchy Standard. This standard ensures that the user can move between the distros without having to re-learn how the system is organized.
	-	Linux uses '/' chanaracter as path seprator.
	-	Many distributions distinguish between core utilities needed for system operation and other program, and place the latter in directories under /usr
  * The Boot Process : The procedure to initialize the system.
	-	Power On-> BIOS-> Master Boot Record(MBR, aka First sector of Hard Disk)-> Boot Loader(eg GRUB)-> Kernel-> Initial RAM disk-> /sbin/init(parentprocess)-> CommandShell-> X Windows System
	-	BIOS initializes the hardware like screen, keyboard and tests the main memory, this is called POST. Power on slef test. BIOS is stored on a ROM chip on motherboard, remainder of the boot process is handled by the OS.
	-	MBR  : Boot loader is usualy stored either in boot sector on hard disk or EFI(Extensible Firmware Interface)partition. Thereafter information on date, time and most important pheriperals are loaded from CMOS(Complimentary metal-oxide semiconductor) values. Most common boot loaders are GRUB(GRand Unified Bootloader)and ISOLINUX(for booting from removable media). While booting Linux, te bootloader is responsible for loading the kernel image and the initial Ram disk(which contains critical files and device drivers) into memory.
	-	Boot Loader: It is 2 disinct stages: 
	First stage: For systems using BIOS/MBR, the bootloader resides in the first sector of HD known as MBR. Boot loader finds the bootable partition by examining the partition table. Once bootable partition is found it searches for the second stage boot loader GRUB and load it in RAM.
	For systems using EFI/UEFI systems, UEFI firmware reads its Boot Manager data to determine which UEFI application is to be launched and from where (i.e., from which disk and partition the EFI partition can be found). The firmware then launches the UEFI application, for example, GRUB, as defined in the boot entry in the firmware's boot manager.

	Second stage: The second stage boot loader resides under /boot. A splash screen gives option to choos the OS to boot. The boot loader loads the selected kernel image and passes control to it.
	-	Linux Kernel: The boot loader loads kernel and an initial RAM filesystem (initramfs)into memory so that it can be used by kernel directly. Once kernel is loaded in the memory it initializes and configures compuer's memory and also configures the hardware attached to the computer.
	- Initial RAM Disk: The initramfs filesystem image contains programs and binary files that perform all actions needed to mount the proper root filesystem, like providing kernel functionality for the needed filesystem and device drivers for mass storage controllers with a facility called udev (for User Device) which is responsible for figuring out which devices are present, locating the drivers they need to operate properly, and loading them. After the root filesystem has been found, it is checked for errors and mounted. Mounting is done by mount command. If this is successful, the initramfs is cleared from RAM and the init program on the root filesystem (/sbin/init) is executed. init handles the mounting and pivoting over to the final real root filesystem. If special hardware drivers are needed before the mass storage can be accessed, they must be in the initramfs image.
	-	/sbin/init: Once the kernel has set up all its hardware and mounted the root filesystem, the kernel runs the /sbin/init program. This then becomes the initial process, which then starts other processes to get the system running. Most other processes on the system trace their origin ultimately to init; the exceptions are kernel processes, started by the kernel directly for managing internal operating system details. Besides starting the system, init is responsible for keeping the system running and for shutting it down cleanly. It acts as the "manager of last resort" for all non-kernel processes, cleaning up after them when necessary, and restarts user login services as needed when users log in and out.
	-	Text mode: Near the end of the boot process,  init starts a number of text-mode login prompts (done by a program called getty). These enable you to type your username, followed by your password, and to eventually get a command shell.
	-	X Windows System: A service called the display manager keeps track of the displays being provided, and loads the X server (so-called because it provides graphical services to applications, sometimes called X clients). The display manager also handles graphical logins, and starts the appropriate desktop environment after a user logs in. A desktop environment consists of a session manager, which starts and maintains the components of the graphical session, and the window manager, which controls the placement and movement of windows, window title-bars, and controls.
	Although these can be mixed, generally a set of utilities, session manager, and window manager are used together as a unit, and together provide a seamless desktop environment.

	If the display manager is not started by default in the default runlevel, you can start X a different way, after logging on to a text-mode console, by running startx from the command line.
	
 # To be contd