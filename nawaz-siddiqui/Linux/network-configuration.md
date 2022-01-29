# Network Configuration

For Debian family configurations, the basic network configuration files could be found under **/etc/network/**, while for Red Hat and SUSE family systems one needed to inspect **/etc/sysconfig/network**.

## Network Interface

Network interfaces are a connection channels between a device and a network. Physically, network interfaces can proceed through a network interface card (NIC), or can be more abstractly implemented as software.

## Route

A network requires the connection of many nodes. Data moves from source to destination by passing through a series of routers and potentially across multiple networks. Servers maintain routing tables ****containing the addresses of each node in the network. The IP routing protocols enable routers to build up a forwarding table that correlates final destinations with the next hop addresses.

## Traceroute

**traceroute** is used to inspect the route which the data packet takes to reach the destination host, which makes it quite useful for troubleshooting network delays and errors.

| Command | Function |
| --- | --- |
| ip addr show | To view the IP address |
| ip route show | To view the routing information |
| ifconfig | show the Information about a particular network interface or all network interfaces |
| ping host-name | To check the status of the remote host |
| route –n
 or ip route | Show current routing table |
| route add -net address
 or ip route add | Add static route |
| route del -net address
 or ip route del | Delete static route |
| traceroute <address> | Prints the route taken by the packet to reach the network host |
| ethtool | Queries network interfaces and can also set various parameters such as the speed |
| netstat | Displays all active connections and routing tables; useful for monitoring performance and troubleshooting |
| nmap | Scans open ports on a network; important for security analysis |
| tcpdump | Dumps network traffic for analysis |
| iptraf | Monitors network traffic in text mode |
| mtr | Combines functionality of ping and traceroute and gives a continuously updated display |
| dig | Tests DNS workings; a good replacement for host and nslookup |
| sudo systemctl restart NetworkManager | If your device was up but had no IP
 address, this command should have helped fix it |
| sudo systemctl restart network | If your device was up but had no IP
 address, this command should have helped fix it |
| sudo service NetworkManager restart | If your device was up but had no IP
 address, this command should have helped fix it |
| sudo service network restart | If your device was up but had no IP
 address, this command should have helped fix it |
| sudo dhclient eth0 | If your device was up but had no IP
 address, you can try to get a fresh address |
| hostname | If your interface is up and running with an assigned IP
 address and you still can not reach google.com
, we should make sure you have a valid hostname assigned to your machine, with hostname |
| sudo apt-get install links2 | Installs a command line browser named links2. |
| links2 google.com | Connects to google.com from terminal |
