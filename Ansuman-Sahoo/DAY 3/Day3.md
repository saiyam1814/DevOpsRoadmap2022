## 12-Jan-22:


- Chapter 3: **Linux Basics and System Startup**:
  - *Introduction completed*
  - *The boot process completed:*
    - The Linux boot process is the procedure for initializing the system.
    - It consists of everything that happens from when the computer power is first switched on until the user interface is fully operational. 
    - **BIOS(Basic Input/Output System):**
      - When the computer is powered on, the Basic Input/Output System (BIOS) initializes the hardware, including the screen and keyboard, and tests the main memory.
      - The BIOS software is stored on a ROM chip on the motherboard
    - **MBR(Master Boot Record) and Boot loader**
      - Usually stored on one of the hard disks in the system
      - A number of boot loaders exist for Linux; the most common ones are:
        - GRUB (for GRand Unified Boot loader)
        - ISOLINUX (for booting from removable media)
        - DAS U-Boot (for booting on embedded devices/appliances)
      - The size of the MBR is just 512 bytes.
      - In this stage, the boot loader examines the partition table and finds a bootable partition. 
      - Once it finds a bootable partition, it then searches for the second stage boot loader, for example GRUB, and loads it into RAM (Random Access Memory).
      - The second stage boot loader resides under /boot. A splash screen is displayed, which allows us to choose which operating system (OS) to boot.
    - **Initial Ram Disks:**
      - The initramfs filesystem image contains programs and binary files that perform all actions needed to mount the proper root filesystem.
      - The mount program instructs the operating system that a filesystem is ready for use, and associates it with a particular point in the overall hierarchy of the filesystem (the mount point).
      - init handles the mounting and pivoting over to the final real root filesystem.
    
- Yaml tutorial *(TechWorld with Nana) completed* :
  - Why YAML?
    - Configuration files all written in YAML.
    - widely used format
    - for different DevOps tools and applications
    - Human readable and intuitive
  - What is YAML?
    - It is a data serialization language i.e. standard format to transfer data like JSON and XML
  - YAML stands for YAML Ain't Markup Language.
  - File extensions: .yaml or .yml
  - **Basic Syntax of YAML:**
    - *Key-value pairs*
      ```
      app: user-authentication
      port: 9000
      version: 1.7
      ```
    - *Comments*
      - Using '#' we write comments:
        ```
        #This is a comment
        ```
    - *Objects*
      ```
      Microservice:
       app: user-authentication
       port: 9000
       version: 1.7
      #Microservice is an object which contain attributes named app, port and version
      ```
    - *Lists*
      ```
      Microservice:
       - app: user-authentication
         port: 9000
         version: 1.7
         deployed: yes #boolean
      # List can be made using '-' another way is by '[]'
      ```
    - *REAL KUBERNETES CONFIGURATION EXAMPLE:*
      ```
      #A POD CONFIGURATION
      apiVersion: v1
      kind: Pod
      metadata:
        name: nginx
        labels:
          app: nginx
      spec:
        containers:
        - name: nginx-container
          image: nginx
          ports:
          - containerPort: 80
          volumeMounts:
          - name: nginx-vol
            mountPath: /usr/nginx/html
        - name: sidecar-container
          image: curlimages/curl
          command: ["/bin/sh"]
          args: ["-c","echo Hello from the sidecar container; sleep 300"]
      ```
    - *Multiline strings*
      ```
      multilineStrings: |
        Hi my name 
        is 
        Ansuman
      #Using pipe symbol '|' multiline strings can be written
      ```
    - *Environment variables*
      - Using '$' sign we can access any enviroment variable.
    - *Placeholders*
      - It can be dfined by using "{{}}"
