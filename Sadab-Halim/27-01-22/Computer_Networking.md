# Continued Learning [Computer Networks](https://www.youtube.com/watch?v=IPvYjXCsTg8) from DevOps Bootcamp

## Repeater
A repeater is a dynamic network device used to reproduce the signals when they transmit over a greater distance so that the signal’s strength remains equal. It can be used to create an Ethernet network. A repeater that occurs as the first layer of the OSI layer is the physical layer.

**Features of Repeater:**
- These repeaters are linked to each other at the physical layer.
- It sends the signals for the unsteady areas to enlarge the system signals.
- These repeaters can eliminate the distance between two devices.

**Advantages of Repeater:**
- These repeaters are cost-effective and easy to use.
- The repeater provides the stabiility of the signals.

**Disadvantages of Repeaters:**
- They cannot connect two distinct networks.
- While amplifying the signals, the repeaters also amplify the level of noise in those signals.

## Hub
Hubs are networking devices operating at a physical layer of the OSI Model that are used to connect multiple devices in a network. They are generally used to connect computers in a LAN.

A hub has many ports in it. <br/>

A computer which intends to be connected to the network is plugged in to one of these ports. When a data frame arrives at a port, it is broadcast to every other port, without considering whether it is destined for a particular destination device or not.

**Features of Hub:**
- A hub operates in the physical layer of the OSI model.
- A hub cannot filter data, It only sends message to all ports.
- It primarily broadcasts messages. So, the collision domain of all nodes connected through the hub stays one.
- Collisions may occurs during setup of transmission when more than one computer's place data simultaneously in the corresponding ports.

**Types of Hubs:**
- **Active Hubs:** amplify and regenerate the incoming electrical signals before broadcasting them <br/>
  They have their own power supply and serves both as a repeater as well as connecting centre. Due to their regenerating capabilities, they can extend the maximum distance between nodes, thus increasing the size of LAN.
  
- **Passive Hubs:** connects nodes in a star configuration by collecting wiring from nodes. <br/>
  They broadcast signals onto the network without amplifying or regenerating them. As they cannot extend the distance between nodes, they limit the size of the LAN.

- **Intelligent Hubs:** are active hubs that provide additional network management facilities. <br/>
They can perform a variety of functions of more intelligent network devices like network management, switching, providing flexible data rates etc.

## Bridges
Bridges are used to connect two subnetworks that use interchangeable protocols. It combines two LANs to form an extended LAN.

The main difference between the bridge and repeater is that the bridge has a penetrating efficiency.

**Working of Bridges:**
A bridge accepts all the packets and amplifies all of them to the other side. The bridges are intelligent devices that allow the passing of only selective packets from them. <br/>

A bridge only passes those packets addressed from a node in one network to another node in the other network.

![image](https://user-images.githubusercontent.com/74575612/151418045-6ce1f5ea-1284-488e-b22b-2079edc1c297.png)

- A bridge receives all the packets or frame from both LAN (segment) A and B.
- A bridge builds a table of addresses from which it can identify that the packets are sent from which LAN (or segment) to which LAN.
- The bridge reads the send and discards all packets from LAN A sent to a computer on LAN A and that packets from LAN A send to a computer on LAN B are retransmitted to LAN B.
- The packets from LAN B are considered in the same method.

**Types of Bridges:**
- **Transparent Bridges:** These are the bridges in which the stations are completely unaware of the bridge's existence. i.e. whether or not a bridge is added or deleted from the network, reconfiguration of the stations is unnecessary. <br/>
  These bridges makes use of two processes; bridge forwarding and bridge learning.
 
- **Source Routing Bridges:** In these bridges, routing id performed by source stations and the frame specifies which route to follow. <br/>
  The host can discover frame by sending a special frame called _discovery frame_, which spreads through the entire network using all possible paths to destination.

**Uses of Bridges:**
- Bridges are used to divide large busy networks into multiple smaller and interconnected networks to improve performance.
- Bridges also can increase the physical size of a network.
- Bridges are also used to connect a LAN segment through a synchronous modem relation to another LAN segment at a remote area.

## Switches
Switches are networking devices operating at layer 2 or a data link layer of the OSI model. They connect devices in a network and use packet switching to send, receive or forward data packets or data frames over the network. <br/>

A switch has many ports, to which computers are plugged in. When a data frame arrives at any port of a network switch, it examines the destination address, performs necessary checks and sends the frame to the corresponding device(s).It supports unicast, multicast as well as broadcast communications.

**Features of Switches:**
- A switch operates in the layer 2, i.e. data link layer of the OSI model.
- It is an intelligent network device that can be conceived as a multiport network bridge.
- It uses MAC addresses to send data packets to selected destination ports.
- It uses packet switching technique to receive and forward data packets from the source to the destination device.
- It supports unicase, multicast and broadcast communications.
- Switches are active devices, equipped with network software and network management capabilities.
- Switches can perform some error checking before forwarding data to the destined port.
- The number of ports is higher -24/48.

**Types of Switches:**
- **Unmanages Switch:** These are inexpensive switches commonly used in home networks and small businesses. They can be set up by simply plugging in to the network, after which they instantly start operating.

- **Managed Switch:** These are costly switches that are used in organisations with large and complex networks, since they can be customized to augment the functionalities of a standard switch.

- **LAN Swithch:** LAN switches connects devices in the internal LAN of an organization. They are also referred as Ethernet switches or data switches. These switches are particularly helpful in reducing network congestion or bottlenecks. They allocate bandwidth in a manner so that there is no overlapping of data packets in a network.

- **PoE Switch:** switches are used in PoE Gogabit Ethernets. PoE technology combine data and power transmission over the same cable so that devices connected to it can receive both electricity as well as data over the same line. PoE switches offer greater flexibility and simplifies the cabling connections

## Router
A router is a device like switch that routes data packets based on their IP addresses. It is mainly a network layer device. Routers normally connects LAN and WANs together and have a dynamically updating routing table based on which they make decisions on routing the data packets. 

## Gateway
 gateway is a network node that forms a passage between two networks operating with different transmission protocols. The most common type of gateways, the network gateway operates at layer 3, i.e. network layer of the OSI (open systems interconnection) model. <br/>
 
It acts as the entry – exit point for a network since all traffic that flows across the networks should pass through the gateway. Only the internal traffic between the nodes of a LAN does not pass through the gateway.

**Features of Gateway:**
- Gateway is located at the boundary of a network and manages all data that inflows or outflows from that network.
- It forms a passage between two different networks operating with different transmission protocols.
- A gateway operates as a protocol converter, providing compatibility between the different protocols used in the two different networks.
- It also stores information about the routing paths of the communicating networks.
- It uses packet switching technique to transmit data across the networks.

**Brouter:**
A brouter is a networking device that functions both as a bridge and a router. It can forward data between networks (serving as a bridge), but can also route data to individual systems within a network (serving as a router). <br/>

It operates either at Data link layer or a Network layer. <br/>
It stores routing table when it is configured as a router and stores MAC address when configured as a bridge. <br/>

It works on more than one broadcast domain when it is configured as a router and It works on single broadcast domain when configured as a bridge.    

