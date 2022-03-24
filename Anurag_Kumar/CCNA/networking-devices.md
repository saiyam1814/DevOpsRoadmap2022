## Building blocks of networking 
  - Devices - these have their unique mac addresses
  - connections - How are devices connected to the internet? coaxial cables, wifi, ethernet or what?
  - switches
  - routers
  - network servers


## IP address
- It is composed of network, host ID and port. 
- Host ID is something like the building in the city and port is like different ways to get into that building.
- Each part of IP address is an octet and each octet contains 8 bits and combined together it's 32 bits. 
- NAT- network address translation 
- IP addresss to you is provided by your ISP providers. Your IP is dynamic. 
- Ports - it is associated with an IP address and identifies an application or service running on networked device. 
- Different port 
- May be you have a seperate entry for employees, seperate entries for shipping and seperate entries for students. 
- There are 65,535 ports in total. 0-1023 ports are reserved. 
- SMTP - 25, DNS - 53, HTTP-80, HTTPS-443

- Think of old days when we used to assemble our computer
  - First we used to buy PC cabinet
  - Then motherboard
  - them RAM
  - then processor

- Now days when we buy motherboard then we get ethernet port along with that. It means that it is already having the NIC card. 
- NIC card have mac address 
- To surf the internet we need two addresses which are
  - mac address -> mac address is provided by NIC card. 
  - ip address

- Just like in case of our home when we need to insert multiple chargers in our plug then we bring power extension.
- Similarly in case if we need to connect more than more than one devices then we need a network hub. 
- We don't use network hub anymore. 
- When we used to send data over a network then it used to send to every computer.
- network hub is also called dummy device as this is not smart. 
- It broadcasts the message. Broadcast means it is communicating with each device. 
- In today's laptop we have 3 NIC.


## Switch

- In place of hub now we use switch
- Switch is not a dummy device. It's an intelligent device.
- Switch learns the mac address
- First time broadcast. It creates a table and then next time it reaches directly to that address. 
- First broadcast and then unicast
- It's like the courier man who when first time comes then it calls everyone who's this person? But once it delivers to you then next time it remembers you and your home address.
- Hubs and switches are only confined to the LAN. It is not used to exchange data outside their own network. To exchange data outside of your network, a device needs to read IP address.

## Router
- Router connects two or more networks
- Router routes the data from one network to another based on their IP address. This was not possible with switches as they used to rely on mac addresses for communication. 
- Router is a gateway for a local area network
