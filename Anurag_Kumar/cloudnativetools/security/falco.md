## Falco 

- One thing before we start is that falco is a threat detection engine and not a prevention tool. It will not save you from the incoming threats. It will only let you detect that there is something wrong with the system or somebody is trying to break into the system or some malicious is activity is going on which shouldn't happen. 
e.g. Getting into your .ssh directory


- What we have used till now?
  - Firewalls
  - Antivirus 
  - Authentication (multi factor)

### Runtime security 
- Another layer to secure your infrastructure 
- Runtime security works between system and application resources 
- Runtime security is 
  - Detecting when a file was opened for read, write and execute. 
  - Detecting priviledge escalation using priviledge containers
  - Detecting drift in the system binary 


- Falco (runtime threat detection engine)
  - Events
  - Rules
  - Alerts 

### Installing Falco
- Very well written docs - https://falco.org/docs/getting-started/installation/ 


- With falco you need to install two things 
  - Falco driver (this will run in kernel space)
  - Falco engine (this will run in user space)
- The standard configuration file of falco is stored at `/etc/falco/falco.yaml` location. 
- The rules are stored at `/etc/falco/falco_rules.yaml` and `/etc/falco/falco_rules.local.yaml` 
- when you will install falco via helm then falco will be installed as a daemonset and it's obvious why. You want to secure all your nodes not specific nodes only. 
- When you're using a managed kubernetes service then you cannot ssh into a node and install falco on the node directly. You need to use helm and that will install falco as daemonset. 
- One more thing to note here is that falco gets installed in a seperate namespace so when you screw up with your falco system then just delete the namespace. I have done this a lot while practicing :)). 
- In the previous installation we have seen that the driver was module and now the driver is bpf. 
- Now the obvious question is when to use the kernel module and when to use eBPF?
  - Use the eBPF probe when loading a kernel module is not a viable option. eBPF is considered safe and is unable to crash or panic a kernel. The eBPF code is already compiled into a Linux kernel and is enabled by simply using the eBPF program. For example, if you are running Falco on GKE or ARM/Raspberry Pis, it is recommended to use the eBPF probe. But, as a rule of thumb, always use the Kernel Module.
- This is one project where I need to pay so much attention to the installation method only. It's not frustrating but I'm learning about the usecases and which installation method to use in which scenario?
- pdig is used when you don't have the option to tweak around kernel. Basically in environments where kernel modules and eBPF are not an option. 
- Although this approach(pdig) is slower than the others, it enables Falco thread detection in environments where no other instrumentation is allowed.


### What are falco rules? 
- These are written in YAML. 
- One example of a falco rule will be if someone tries to access your private ssh key then falco will flag that rules as dangerous and reports back to you via container logs or UI. 
- Recommended way to use falco is to install it on the host system itself. 
- Majorly consist of three things 
  - Rules
  - Macros
  - lists

### Falco Events 
- You can easily forward falco events to slack, Kafka, AWS lambda and falcosidekick. 