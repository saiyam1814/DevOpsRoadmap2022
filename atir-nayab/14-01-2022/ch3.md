**Boot Loader in action**

The Boot Loader has two distict stages: 

* For system using the BIOS/MBR method, the boot loader resides at the first sector of the hard disk, also know as Master Boot Record (MBR). The size of MBR is 512 Bytes. In this stage, the boot loader examines the partition table and finds a bootable partition. Once it finds a bootable partition, it then searches for the second stage boot loader for example GRUB and loads it into RAM manager data to determine which UEFI application is to be launched and from where the UEFI application, for example GRUB as defined in the boot entry in the firmware's boot manager. This procedure is more complicated, but more versatile.

  * System 
  * Bootable Hard Disk
    * Kernal
    * Inital RAM Disk

* The Second stage boot loader reside under `/boot`. A splash screen is displayed, which allows us to choose which Operating System (OS) to boot. After choosing the OS the boot loader loads the kernal of the selected operating system into RAM and passes control to it kernals are almost always compressed. So, its first job is to uncompress itself. After this, it will check and analyze the system hardware and initialize any hardware device drivers built into the kernal.



