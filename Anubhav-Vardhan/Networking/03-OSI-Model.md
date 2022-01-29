# Open Systems Interconnection (OSI) Model

OSI model was developed to setup a standard of how systems like computers and servers communicate with each other.
It divides the whole process into **7 layers**:

- Application layer
  - Implemented in software like browsers, messaging apps
  - Data from here is sent to Presentation layer
  - Application layer protocols are HTTP, FTP, TELNET etc.
- Presentation layer
  - Presentation layer converts data from Application layer into machine representable binary format.
  - For example from ASCII to something like EBCDIC
  - This is known as translation.
  - Things like encoding and encryption happens
  - It also provides abstraction
  - Data is also compressed to make it easier to transfer over network. This can be lossy or lossless depending on other factors
  - SSL protocol is used here
- Session layer
  - Helps in setting up and managing the connections and enables the sending and recieving of data followed by termination of connected sessions
  - We can do authentication before a session is established (for eg: asking username and password)
  - Session layers also assumes that layer below will do their work
- Transport layer
  - Segmentation: Data recieved from session layer is divided into small data units called segments. 
    Every segment contains source and destinations's port number and sequence number
  - Flow control: Controls the amount of data being transferred.
    It also helps with error control using checksum 
- Network layer
  - Communication happens with other computers
  - Router works at this layer
  - Logical addressing happens here
  - Network layer assigns the sender's and reciever's IP address to every segment and it forms an IP packet (so that every data packet can reach the correct destination)
  - It also performs routing (moving data packet from source to destination)
  - Load balancing also happens at this layer
- Data link layer
  - Allows us to directly communicate with the computer and hosts.
  - Recieves the data packet from network layer which contains ip address of both sender and reciever (and subnet mask)
  - MAC addresses are physical addresses. MAC adresses of sender and reviever are assigned to data link to form a frame
  - MAC adress is 12 digit alphanumeric number of network interface of our computer
  - Frame is basically data unit of the data link layer
- Physical layer
  - Wires, light or electrical signal
  - Physical layer recieves the signal, convert it to bits and pass it to data link layer as a frame.

## How this works

```
  "ME"              "My FRIEND"
    |                   |
Application         Application
    |                   |
Presentation        Presentation
    |                   |
Session             Session
    |                   |
Transport           Transport
    |                   |
Network             Network
    |                   |
Data link           Data link
    |                   |
Physical ---------> Physical
```

each "me" layer assumes it is talking to same layer of "my friend"
