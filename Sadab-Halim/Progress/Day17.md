# Computer Networking (CCNA)

## Topologies

- **Bus Topology:** <br/>
  a. They are connected through a single cable known as backbone cable.<br/>
  b. If one part gets broken entire system will fail. <br/>
  c. Only 1 person at a time can send information. <br/>
  d. It is mainly used in 802.3 (ethernet) and 802.4 standard networks.
  
  ![image](https://user-images.githubusercontent.com/74575612/151399039-ec286d6b-99d3-4b97-ae66-8df9cc0160c7.png)

  
- **Ring Topology:** <br/>
  a. Ring Topology is like bus topology, but with connected ends. <br/>
  b. It's unidirectional, the data flows in one direction. <br/>
  c. It has no terminated ends. <br/>
  d. Every system communicates with one another. <br/>
  e. If one of the cables break we won't be able to send data. <br/>
  f. A lot of unnecessary calls are made. <br/>
  g. The data in a ring topology flow in a clockwise direction.
  
  ![image](https://user-images.githubusercontent.com/74575612/151399571-87d95804-5218-46f7-9e2c-994ac278e012.png)

  
- **Star Topology:** <br/>
  a. It is an arrangement of the network in which every node is connected to the central hub, switch or a central computer. <br/>
  b. There will be one central device that will be connected to all computers <br/>
  c. If central device fails, then the system will go down. <br/>
  d. The central computer is known as a _server_, and the peripheral devices attached to the server are known as _clients._ <br/>
  e. Coaxical cable or RJ-45 cables are used to connect the computers. <br/>
  
  ![image](https://user-images.githubusercontent.com/74575612/151400158-40a2e57e-4865-4d10-97a6-4d305b8bb614.png)

- **Tree Topology:** <br/>
  a. It combines the characterisitcs of bus topology and star topology. <br/>
  b. A tree topology is a type of structure in which all the computers are connected with each other in hierarchical fashion. <br/>
  c. The top most node in the tree topology is known as a root node, and all other nodes are the descendants of the root node. <br/>
  
  ![image](https://user-images.githubusercontent.com/74575612/151400777-7fb41cba-c91a-42d6-8dcc-954ce44ac3f9.png)

- **Mesh Topology:** <br/>
  a. Mesh Topology is an arrangement of the network in which computers are interconnected with each other through various redundant connections. <br/>
  b. Every single computer will be connected to every single computer. <br/>
  c. It is mainly used for WAN implementations where communication failures are a critical concern. It is mainly used for wireless networks. <br/>
  d. Internet is an example of the mesh topology. <br/>
  e. Expensive and Scalability issues <br/>

![image](https://user-images.githubusercontent.com/74575612/151401537-134221c4-f6cd-4940-8470-dff60f78eb5a.png)

- **Hybrid Topology:** <br/>
  a. The combination of various different topologies is known as Hybrid topology. <br/>
  b. It is a connection between different links and nodes to transfer the data. <br/>
  
  ![image](https://user-images.githubusercontent.com/74575612/151401937-2b88d5c6-3bee-4d90-94a3-3e38bf70e9ee.png)

## TCP/IP Model
It is basically known as **Internet Protocol Suite** <br/>

There are 5 layers:
- Application Layer
- Transport Layer
- Network Layer
- Datalink Layer
- Physical Layer

## Structure of Network OSI Model
There are seven layers in OSI Model:

![image](https://user-images.githubusercontent.com/74575612/151403496-2d28ecce-f65f-473a-abdc-6d84ad8aed13.png)

| Layers             | Description |
| ------------------ | ----------- |
| Application Layer  | Implemented in software, it is just the application like browser, chat apps etc. |
| Presentation Layer | It converts those messages, data into binary format. Encoding, Encryption happends, provides Abstraction, Compression and Translation. |
| Session Layer      | It helps in setting up and managing the connections and enables sending and receiving of data followed by terminationo of connected session. Authentication and Authorization takes place. |
| Transport Layer    | Data receieved from session layer is divided into small data units called segment. Every segment has source and destination's port as well as sequence number. Flow Control & Error Control. |
| Network Layer      | The transmission of the received data segments from one computer to another that is located in different network. IP addressing  done here is called Logical Addressing. Routing is performed and Load Balancing is done. |
| Datalink Layer     | Physical addressing is done here. Mac addresses are physical address. Now, these addresses of sender and receiver are assigned to packets called frames. |
| Physical Layer     | Hardwares like Cables, Wires, etc. |

![image](https://user-images.githubusercontent.com/74575612/151405962-3de83226-de93-4582-b653-869fc1a7e4c1.png)

**Mac Address:** It is a 12 digit alphanumeric number of network interface of computer.

## Client-Server Architecture
- A server is basically a system that controls the website you are hosting.
- The application has two parts: Client part and Server part <br/>
  These are known as processes and they communicate through each other.
- Clients are the ones who are using/cosuming these resources like we making a request to google.
- Data Center is a collection of huge number of computers. It may have static IP addresses. They have good internet connection and high upload speed.

Command: `ping google.com`

**ping** measures the round trip time for messages send from the originating host to the destination computer and are echoed back.

![image](https://user-images.githubusercontent.com/74575612/151407044-1bc6b86b-f18d-4b9a-8960-501f9cce76ab.png)

## Peer-To-Peer Architecture:
- Peer-To-Peer network is a network in which all the computers are linked together with equal privileges and resposibilities for processing the data.
- It is useful for small environments, usually up to 10 computers.
- It has no dedicated server.
- Here, every single computer can be termed as a client as well as a server.
- The key advantage is we can scale it rapidly.
- Special permissions are assigned to each computer for sharing the resources, but this can lead to a problem if the computer with the resource is down.

![image](https://user-images.githubusercontent.com/74575612/151407582-043b0a38-7fe6-4818-95c1-cff4c00dd473.png)

Resources:
- 
- [Computer Networking Course](https://www.youtube.com/watch?v=IPvYjXCsTg8)