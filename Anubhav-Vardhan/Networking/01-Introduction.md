## Protocol
Protocols are simply just some rules defined by the internet society which defines how data is transferred

- **TCP (Transmission Control Protocol)**:  Ensures data reaches its destination and is not corrupted along the way. When we don't want to lose any data we use TCP
- **UDP (User Datagram Protocol)**:         When it doesn't matter if 100% of data is reaching its destination. For eg: Video call
- **HTTP (Hyper Text Transfer Protocol)**:  Used by www, defines the format of data being transferred between web clients and servers
- Others...

**It is impractical to send large files over the internet, data is sent in small chunks called packets** 

## IP Address
All the computers and servers on a network is identified by their unique IP addresses. 
Think of it the same as we all have phone numbers
<br>
For eg: when we use google.com, it is resolved to its IP address 
similar to when we call a person from our contacts, it gets resolved to their phone numbers

```
X.X.X.X
|
This X can have no. from 0 to 255
```

**We can check our IP adress on ifconfig.me or just run:**
``` BASH
$curl ifconfig.me
```
---
```
ISP (Internet Service Provider)
              |
         Modem/Router   -----> Global IP address
        |     |     |
        D1    D2    D3  -----> Local IP address (Same global IP address)
```
Modem is assigning local IP adresses using **DHCP (Dynamic Host Configuration Protocol)**
<br>
If we make a request to google.com, google will only see global IP adress (i.e our modem or router) 
and when response comes back, the router will decide which device requested it using **NAT (Network Address Translation)**
<br>
But IP address decided which device to send data to, but which application on that device needs that data? Here comes **Ports**

So,
<br>
IP address denotes which device to send data to
<br>
Port number denotes which application to send data to

## Ports
Ports are 16 Bit numbers
<br>
So, total possible number of ports are 2^16 i.e around 65000

For eg:
- HTTP: 80
- MongoDB: 27017
- SQL: 1433

- **0 to 1023:** Reserved ports which means we cannot host our thing on these ports as they are already reserved for other purposes like port 80 for HTTP
- **1024 to 49152:** Reserved for applications like mongoDB, MySQL etc
- **Remaining:** Free for our use

## LAN MAN WAN
- **LAN (Local Area Network):** Small houses/office for example with ethernet or wifi
- **MAN (Metropolitan Area Network):** Across a city
- **WAN (Wide Area Network):** Across countries for example with optical fibre cables
  - **SoNET (Synchronous Optical Networking):** Carries data using optical fibre cables covering larger distances
  - **Frame Relay:** Way to connect our local network to a wider network like the internet

LANs are conected together in a MAN and MANs are connected together in a WAN and this is how
the Internet works all over the world

## Modem and Router
- Modem is used to covert digital signals to analog signals and vice versa.
  - For example modem can convert our digital data into an electrical signal so we can transfer it over telephone lines and convert it back to digital data.
- Router is a device that routes the data packet based on their IP address
