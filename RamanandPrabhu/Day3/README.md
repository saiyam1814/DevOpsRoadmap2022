#Day3
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
 # To be contd.