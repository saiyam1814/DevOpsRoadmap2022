## More About Priorities

At any given time, many processes are running (i.e. in the run queue) on the system. However, a CPU can actually accomodate only one task at a time. Linux allows you to set and manipulate process priority. Higher priority processes grep preferential access to the CPU.

The priority for a process can be set by specifying a nice value, or niceness, for the process. The lower the nice value, the higher the priority. Low values are assigned to important processes, while high values are assigned to processes that can wait longer.

![nice Output](https://courses.edx.org/assets/courseware/v1/d929bb0d3a611c780e2dcb6dd00cddd5/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/niceout.png)

You can also assign a so-called real-time priority to time-sensitive tasks, such as controlling machines through a computer or collecting incoming data. This is just a very high priority, and has more to do with making sure a job gets completed within a very well defined time window.

![nice Values](https://courses.edx.org/assets/courseware/v1/fcc61556971b7cdefbafabf3d7abab22/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch16_screen08.jpg)

## Load Averages
The load average is the average of the load number for a given period of time. It takes into account processes that are

  * Actively running on a CPU.
  * Considered runnable, but waiting for a CPU to become available.
  * Sleeping: i.e. waiting for some kind of resource (typically I/O) to become available.

  The load average can be viewed by running w, top or uptime. We will explain the numbers on the next page.

  ![Load Averages](https://courses.edx.org/assets/courseware/v1/ca05a14d78d8e3bb26b519fe65047a66/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/wuptimesuse.png)

  ## Interpreting Load Averages

  The load average is displayed using three numbers (0.45, 0.17, 0.12) is the below screenshot.

  * 0.45: For the last minute the system has been 45% utilized on average
  * 0.17: For the last 5 minutes utilization has been 17%
  * 0.12: For the last 15 minutes utilization has been 12%

  ![Interpreting Load Averages](https://courses.edx.org/assets/courseware/v1/5ad78e82ed03efc7777fad630abed5dd/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/woutputrhel.png)

  ## Background and Foreground Processes

Linux supports background and foreground job processing.

You can put a job in the backgorund by suffixing & to the command eg updatedb &.

You can either use CTRL-Z to suspend a foreground job or CTRL-C to terminate a foreground job and can always use the bg and fg commands to run a process in the background and foreground respectively.

![background and Foreground Processes](https://courses.edx.org/assets/courseware/v1/3ff2741d99789599c91efda5c5028150/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/bgfgrhel.png)

## Managing Jobs
The jobs utility displays all jobs running in ground. The display show the job ID, state, and command name, as shown here.

jobs -l provides the same information as jobs, and adds the PID of the background jobs.

The background jobs are connected to the terminal window, so, if you log off, the jobs utility will now show the ones started from that window.

![Managing Jobs](https://courses.edx.org/assets/courseware/v1/8dcff92dcec85e717944d972b96d6fcc/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/jobsrhel.png)

## The ps Command (System V Style)
ps provides information about currently running processes keyed by PID. If you want a repetitive update of this status.

![The ps Command (System V Style)](https://courses.edx.org/assets/courseware/v1/a910c16bb6f18c4d38e9ff123a6f5e02/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/ubuntupsef.png)

![ps Command (BSD Style)](https://courses.edx.org/assets/courseware/v1/8cca52331523da587fab092df4bc7dba/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/psbsdrhel.png)

## The Process Tree
pstree displays the processes running on the system in the form of a tree diagram showing the relationship between a process and its parent process and any other processes that it created. Repeated entries of a process are not displayed, and threads are displayed in curly braces.

![The Process tree](https://courses.edx.org/assets/courseware/v1/bce96558c9c9be5c152a567a0c63d392/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/ubuntupstree.png)

## top
While a static view of what the system is doing is useful, monitoring the system performance live over time is also valuable.

~[top](https://courses.edx.org/assets/courseware/v1/9eaf15c635ff33e0dbd318a36f295925/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/toprhel.png)

## First Line of the top Output 
The first line of the top output displays a quick summary of what is happening is the system, including
  * How long the system has been up
  * How many users are logged on
  * What is the load average.

## Second Line of the top Output
The second line of the top output displays the total number of processes, the number of running, sleeping, stopped, and zombie processes.

## Third line of the top Output
The third line of the top output indicates how the CPU time is being divided between the users (us) and the kernal (sy) by displaying the percentage of CPU time used for each.

## Fourth and Fifth Lines of the top output
The fourth and fifth lines of the top indicate memory usage, which is divided in two categories:
  * Physical memory (RAM) - displayes on line 4.
  * Swap space - displayed on line 5.
Both categories display total memory, used memory and free space.

## Process List of the top Output
Each line in the process list of the top output displays information about a process. By default processes are ordered by highest CPU usage.
  * Process Indentification Number(PID)
  * Process owner (USER)
  * Priority (PR) and nice values (NI)
  * Virtual (VIRT), phydical (RES) and shared memory (SHR)
  * status (S)
  * Percentage of CPU (%CPU) and memory (%MEM) used
  * Execution time (TIME+)
  * Command (COMMAND)

## Interactive Keys with top
|Command|Output|
|:-:|:-:|
|t|Display or hide summary information (rows 2 and 3)|
|m|Display or hide memory information (rows 4 and 5)|
|A|Sort the process list by top resource consumers|
|r|Renice (change the priority of) a specific processes|
|k|kill a specific process|
|f|Enter the top configuration screen|
|o|Interactively select a new sort order in the process list|

## Scheduling Future Processes Using at

![Scheduling Future Processes Using at](https://courses.edx.org/assets/courseware/v1/ec37c00269266c49e55f7a52aab93f9a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/atout.png)

## cron
cron is a time-based scheduling utility program. It can launch routine background jobs at specific times and/or days on an on-going basis cron is driven by a configuration file called `/etc/crontab` (cron table), which contains the various shell commands that need to be run at the properly scheduled times.

typing crontab -e wil open the crontab editor to edit existing jobs or to create new jobs.

|Field|Description|Values|
|:-:|:-:|:-:|
|MIN|Minutes|0 to 59|
|HOUR|Hour field|0 to 23|
|DOM|Day of Month|1-31|
|MON|Month field|1-12|
|DOW|Day of Week|0-6(0=Sunday)|
|CMD|Command|Any command to be executed|

## sleep
sleep suspends execution for at least the specified period of time, which can be given as the number of seconds (the default)

`sleep NUMBER[SUFFIX]`
* s for seconds (the default)
* m for minutes
* h for hours
* d for days

sleep and at quite different; sleep delays execution for a specific period, while at starts execution at a later time