
# Networking

Internet is collection of Computer Network that interconnect computing devices. Networking the computers in your home or office so that they can share information with one another and connecting your computer to the worldwide Internet are two separate but related tasks.

> *World Wide Web is the is first website. It is information system where documents and other resources are identified by URL, which are interlinked by hyperlinks.*

### Terminology

| Term |  Description|
|-------|-------|
|Server | The network computer that contains the hard drives, printers, and other resources that are shared with other network computers is a server |
| TCP (Transmission Control Protocol) | It ensures that the data is not corrupted while reaching its destination  |
| UDP (User Datagram Protocol) | It provides a mechanism to detect corrupt data in packets, but it does not attempt to solve other problems that arise with packets, such as lost or out of order packets.|
| HTTP | *Allows client and server programs on different system  to exchange messages.* It is Stateless Protocol, servers do not maintain any information about the client |
| Packets | *A packet is a message that's sent over the network from one node to another node.* The packet includes the address of the node that sent the packet, the address of the node the packet is being sent to, and data. |
| ISP | Internet service provider (ISP) refers to a company that provides access to the Internet to both personal and business customers.|
 |DiffServ (Differentiated services) |  In network it allows you to directly configure the relevant parameters on the switches and routers rather than using a resource reservation protocol. |
| DSCP (Differentiated Services Code Point ) | The six most significant bits of the DiffServ field is called as the DSCP.  |
| Ports |  A virtual point where network connections start and end. Ports are software-based and managed by a computer's operating system. |
|  Domain Name System (DNS)|  It is the phonebook of the Internet. Humans access information online through domain names, like google.com|

### Ports


##### What are the different port numbers?

There are 65,535 possible port numbers, although not all are in common use. Some of the most commonly used ports, along with their associated networking protocol, are:

- `Ports 20 and 21`: File Transfer Protocol (FTP). FTP is for transferring files between a client and a server.
- `Port 22`: Secure Shell (SSH). SSH is one of many tunneling protocols that create secure network connections.
- `Port 25`: Simple Mail Transfer Protocol (SMTP). SMTP is used for email.
- `Port 53`: Domain Name System (DNS). DNS is an essential process for the modern Internet; it matches human-readable domain names to machine-readable IP addresses, enabling users to load websites and applications without memorizing a long list of IP addresses.
- `Port 80`: Hypertext Transfer Protocol (HTTP). HTTP is the protocol that makes the World Wide Web possible.
- `Port 123`: Network Time Protocol (NTP). NTP allows computer clocks to sync with each other, a process that is essential for encryption.
- `Port 179`: Border Gateway Protocol (BGP). BGP is essential for establishing efficient routes between the large networks that make up the Internet (these large networks are called autonomous systems). Autonomous systems use BGP to broadcast which IP addresses they control.
- `Port 443`: HTTP Secure (HTTPS). HTTPS is the secure and encrypted version of HTTP. All HTTPS web traffic goes to port 443. Network services that use HTTPS for encryption, such as DNS over HTTPS, also connect at this port.
- `Port 500`: Internet Security Association and Key Management Protocol (ISAKMP), which is part of the process of setting up secure IPsec connections.
- `Port 3389`: Remote Desktop Protocol (RDP). RDP enables users to remotely connect to their desktop computers from another device.