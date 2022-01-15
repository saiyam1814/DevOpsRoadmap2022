#### Chapter 3: Linux Basics and System Startup

**Main Highlights**
* **Linux filesystems**
* **Difference between partition and filesystem**
* **Boot process**
* **Install Linux**

**The Boot Process**
The linux boot process is the procedure for initializing the system. It consists of everything that happens from when the computer power is first switched on until the user interface is fully operation.

1. Power ON
1. BIOS
1. Master Boot Record(MBR) also known as first sector of hard disk
1. Boot Loader (eg GRUB)
1. Kernal (linux OS)
1. Initial RAM Disk - initramfs image
1. /sbin/init (parent process)
1. Command Shell using getty
1. X Window System(GUI)

**BOIS** - The First Step
Starting x86 based Linux System involves a number of steps. When the computer is power on the Basic Input/Output System (BIOS) initializes the hardware, including server and keyboard and tests the main memory. The process is called POST(Power On Self Test)

1. Power ON
1. BIOS (Basic Input/Output System)
1. Initializes screen keyboard and tests main memory

*The Bios is stored on ROM chip on Mother Board. After this the remainder of the boot process is controlled by the Operating System(OS)*

**Master Boot Record (MBR) and Boot Loader**
Once the post is complete the system control passes from the BIOS to the Boot Loader. The boot loader is usually stored on one of the hardisks in the system, either in the boot sector (for traditional BIOS/MBR systems) or EFI parition. (for more recent (Unified) Extensible Firmware Interface or EFI/UEFI). Up to this stage machine does not access any mass storage media. Thereafter information on data, time and most important peripherals are loaded from the CMOS values.(after a technology used for the battery powered memory store which allows the system to keep track of the data and time even when it is powerd off).
eg *GRUB*, *ISO Linux (for booting removable media)* *DAS U-Boot (for booting on embeded devices/appliances)*

* MBR
  * Partition 1 Searches for GRUB and loads into RAM
  * Partition 2
  * Partition 3