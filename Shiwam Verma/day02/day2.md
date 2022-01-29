# Learnt about all these things on day 2 of Devops.

Took about 3 hours to cover this from [kunal's video](https://www.youtube.com/watch?v=IPvYjXCsTg8)

Not hard topics, but very important in learning about how internet works.

Also learnt about markdown file commands, like adding links. 

---

> OSI Model (Open Systems Interconnection Model)
- Standard way for two computers to communicate .

___

- Layers
    - Application : 
        - Implemented in software
        - As name suggests , it is the app that user interacts with
        - Sends message to presentation layer 

    - Presentation
        - Changes data received from application layer to machine readable binary format
        - From ASCII -> EBCDIC ( Translation )
        - Encoding , Encryption happen at this layer
        - Provides abstraction
        - Data is compressed
        - Medium between application and session layer

    - Session
        - Helps in setting up and managing connections
        - Authentication , Authorisation takes place
        - Medium between  session layer and Transport layer

    - Transport
        - Makes sure that data is transported easily
        - Uses protocols such as TCP or UDP
        - Segmentation is done here 
            - Every segment contains sequence number
        - Flow control
        - Error control
        - Adds checksum to each data segment.
        - Medium between Transport layer and network layer

    - Network
        - Transmission of received data segments 
        - Router lives in this layer
        - Logical addressing ( IP addressing done in network layer )
            - Assigns sender’s and receiver’s IP address to every segment 
            - Forms an IP packet
        - Performs routing
        - Determines best path to send data
        - Load balancing happens here
        - Medium between network layer and data link layer

    - Data Link
        - Directly communicate with computers and hosts
        - Physical addressing 
            - MAC address of sender and receiver are assigned to form a frame 
                - MAC address : 12 digit alpha-numeric number of network interface of a device

    - Physical
        - Contains hardware
        - Transports bits as electrical signals , light signal etc
        - When data is received
            - Physical layer converts signal to bits and pass it to data link layer

___

- Each layer on sender’s device thinks that it is talking to same layer on receiver’s device
    - Example : Session layer of sender’s device thinks that it is sending data to receiver’s session layer, But in reality , Data goes from
	 session (S) -> Transport (S) -> Network (S) -> Data Link (S) -> Physical (S) -> Physical (R) -> Data Link (R) -> Network (R) -> Transport (R) -> Session (R)
___

> TCP/IP model
- Practically used , unlike OSI model which is theoretical
- Layers
    - Application
    - Transport
    - Network
    - Data Link
    - Physical
___

> Application Layer
- Users Interact with this layer
- Present on devices
- Protocols on application layer
    - TCP/IP
        - HTTP
            - Methods
            - Status Codes
            - Cookies
                - Unique string
                - Stored in browser
                - Third party cookies
                    - Cookies set for url that you have not visited
            - DNS
        - DHCP
        - FTP
        - SMTP
            - To send emails 
        - POP3 , IMAC
        - SSH
        - VNC
    - TELNET
    - UDP
- Client Server Architecture
    - Ping time
- Peer to peer architecture
    - Every device in client and server 
    - Decentralised
- Program 
    - Application such as messages
- Process
    - Running instance of a program
    - Example : send a message
- Thread
    - Lighter version of a process
    - Example : Set up page , take keyboard input
- Socket
    - The interface between process and internet
- Ports
    - Tell us which application we are working with
    - Ephemeral Ports
___

- DNS (Domain name system)
    - Root DNS server
        - Top Level domain
            - .io , .org, .com
___

> Transport Layer
- Within a device , transport layer mediates data between application and network Layer.
- Multiplexing , Demultiplexing takes place here 
- Takes care of congestion control
- Congestion control algorithm built in TCP
- CheckSum
    - Used to ensure data packet is received in correct order
- Timers
- Sequence numbers
- Protocols
    - UDP
        - User datagram protocol
            - Data may or may not be delivered 
            - Data may change
            - Data may not be in order
        - Connection less protocol
        - Uses Checksums
        - Segment contains 
            - Header 
                - Service port number
                - Length of data given 
                - Destination port number
                - checksums
            - Data
        - Use cases
        - Since it is very fast
            - Video conferencing
            - Gaming
            -  DNS uses UDP

    - TCP - Transmission Control Protocol

        - App Layer sends lots of data
        - TCP segments this data
        - Add headers,
        - Collect data network layer 
            - Checks data 
        - Congestion Control
        - Takes care when 
            - Data does not arrive 
        - Maintains order of data
        - Features
            - Connection oriented
            - Error control
            - Congestion control
            - Full Duplex
            - 1 TCP connection between 2 computers

___

> Network Layer

- Every device has its own network address
- Hop by Hop forwarding
    - Routing table
        - Forwarding table
- Control Plane 
    - Used to build routing table 
    - Routers -> nodes
    - Links -> Edges
- Protocol
    - Internet Protocol (IP)
        - IPV4
            - 32 bit number with 4 words
        - IPV6
            - 128 bit number
        - Class of IP address 
    - Subnet ID
    - Class of IP addresses
        - Subnet mask
        - Variable length subnet
    - Reserved address
        - Loopback address
            - Example : localhost
- Packet
    - Header is 20 
- Middle boxes
- NAT ( network address translation)

___

> Data Link Layer

- Sends packets from network layer to physical address
- Data link layer address also known as MAC address
- ARP cache
- Frame
- Works very closely with physical layer.
---
