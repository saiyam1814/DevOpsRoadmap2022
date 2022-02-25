### Virtualization Softwares

Also known as hypervisor, it helps you create a different environment and play around with different operating systems and virtual machines which will not hamper your local machine state or performance.

**Type 1 Hypervisor:**

A Type 1 hypervisor **runs directly on the underlying computer's physical hardware**, interacting directly with its CPU, memory, and physical storage. For this reason, Type 1 hypervisors are also referred to as bare metal hypervisors. Example: VMware ESXi and Microsoft Hyper-V.

**Type 2 Hypervisor**

A Type 2 hypervisor, also called a hosted hypervisor, is a virtual machine (VM) manager that is installed as a software application on an existing operating system (OS). Example: Oracle VirtualBox and VMware workstation.

### Vagrant

Vagrant makes the task of maintaining multiple VMs more efficient. It is very useful in case of complex system including multiple VMs.

### DNS

DNS is all about providing a name to a specific IP address. To add local name resolution on your machine go to **/etc/hosts** and add IP with the name you want to associate.ip

### Networking

| Command | Function |
| --- | --- |
| ip link | To list and modify interfaces on the host |
| ip addr | See the IP assigned to the interfaces |
| ip addr add 192.1.1.0 dev eth0 | Set IP addresses on the interfaces. Changes made are only available till restart. To make it persistent add it to /etc/network-interfaces file |
| ip route | view routing table |
| ip route add 192.168.1.0/24 via 192.168.2.1 | Add entries to the routing table |
| cat /proc/sys/net/ipv4/ip_forward | If the output is 1, IP forwarding is on else vice versa |
