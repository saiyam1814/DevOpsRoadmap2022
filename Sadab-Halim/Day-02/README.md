# 02 January 2022

## Completed Chapter 3: Linux Basics and Startup System

**The Boot Process:** The Linux Boot Process is the procedure for initializing the system. It consists of everything that happens from when the computer power is first switched on until the user interface is fully operation.

1. Power ON
2. BIOS
3. Master Boot Record (MBR) also known as first sector of hard disk
4. Boot Loader (eg GRUB)
5. Kernal (linux OS)
6. Initial RAM Disk - initramfs image
7. /sbin/init (parent process)
8. Command Shell using getty
9. X Window System(GUI)

**BOIS:** The First Step Starting x86 based Linux System involves a number of steps. When the computer is power on the Basic Input/Output System (BIOS) initializes the hardware, including server and keyboard and tests the main memory. The process is called POST(Power On Self Test)

1. Power ON
2. BIOS (Basic Input/Output System)
3. Initializes screen keyboard and tests main memory

The Bios is stored on ROM chip on Mother Board. After this the remainder of the boot process is controlled by the Operating System(OS)

Master Boot Record (MBR) and Boot Loader Once the post is complete the system control passes from the BIOS to the Boot Loader. The boot loader is usually stored on one of the hardisks in the system, either in the boot sector or EFI parition.
* MBR
  - Partition 1 Searches for GRUB and loads into RAM
  - Partition 2
  - Partition 3
