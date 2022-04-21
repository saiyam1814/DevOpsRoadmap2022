- I don't know should I call it a tool or a kernel functionality. 
- By default any application or process can communicate with syscall interface directly and we can restrict that using seccomp or apparmor. 
- We control the restriction by creating something called profiles. 
- e.g. We can create a appArmor profile for firefox and we can manage what firefox can do and what firefox cannot do. 
- We can create a profile for kubelet running on a node. 
- 

### appArmor 
- appArmor is a kernel hardening tool. 
- Debian based alternative to SELinux which is prominent in centos and redhat linux family. 
- `aa-status` gives you the details of what appArmor is doing right now. 
- To check the status use `systemctl service apparmor` 
- There are three modes in apparmor
  - enforce mode (Process cannot run if it is not having necessary permissions)
  - complain mode (Progarm will run but it will be logged)
  - unconfined (This progarm can do whatever it want)

### Writing your own appArmor profile 
- let's first install apparmor-utils. Run the command `sudo apt-get install apparmor-utils` 
- We will use `aa-genprof` to create a new profile `aa-genprof <appname>` 
- appArmor will look at what the process actually needs and create a profile around it. 
- The actual profiles are located at `/etc/apparmor.d/` 
- from the exam perspective it is very unlikely that you will be told to write your own profile. 
- You may have to install pods and create a profile around it. 

### appArmor w.r.t. kubernetes 
- You will configure the appArmor profile via the annotation. 
- We will use annotation at the pod level not at the deployment level. 


### commands 
- `aa-status` -> show all profiles
- `aa-genprof` -> generate a new profile 
- `aa-profile` -> to put the profile in complain mode 
- `aa-enforce` -> put the profile in enforce mode 
- `aa-logprof` -> update the profile if app produced some more usage logs. 