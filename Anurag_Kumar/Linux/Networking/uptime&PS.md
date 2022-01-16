### Uptime

- `uptime` tell how long the system has been running. It gives you a one line display of current time, how long the system has been running, how many users are currently logged on and the system load averages for 1,5 and 15 minutes. 
  - `uptime -p,--pretty` show uptime in pretty format 
  - `uptime -s,--since` system up since

- `cat /etc/os-release` for knowing which system you are in. Helpful when you are in the cloud. 
- `cat /etc/issue` for knowing the VM which you are running 
- `uname` print system information 
  - `uname -s, --kernel-name` print the kernel name 
  - `-v` flag for version 
  - `-p` flag for processor 
  - `-s` for kernel name
  - `-a` for all means it will print everything in order. For me this works best as I don't have to check everything one by one.

### Linux Processes
- A process is a running command, application or any program.
- Two types of linux process are there: 
  - Foreground process - one that runs by default 
  - Background process - It runs in background so that other processes can run in parallel. To run a process in background just add `&` at the end. e.g. `pwd &`

- `ps` stands for process status and it shows you all the processes that is running in that terminal.
- `ps -f` will give you more information about the commands that are running. 
- `ps <id>` will give you information about that PID.
  - UID - User ID
  - PID - Process ID
  - PPID - Parent process ID
  - STIME - process start time
  - TTY - Terminal type associated with the process
  - TIME - CPU time taken by the process
  - CMD - the command that started the process
- Stopping Process
  - For processes running in foreground you can simply stop it by pressing `ctrl+C`
  - For processes running in background you can first run the ps command then get the process ID. After getting it, you can kill the process using `kill <process ID>`. If a process ignores the regular kill command then use `kill -9 <process ID>`

