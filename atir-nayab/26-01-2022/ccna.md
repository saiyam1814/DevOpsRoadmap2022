# CCNA

## What is network
When two or more computers are connected together so that they can share resources is called network.

- Host - any device which gets ip address is an host

## What is networking?
sharing of data and resources in a network is called as networking

- LAN - Local Area Network
- WAN - Wide Area Network
  - pubic - when we use publically availble resources to connect different network
  - private - when we use our own resources to connect different network eg leased line

## Router
When we want to connect two or more different network we need router

## Modem
Modem is used for modulation and demodulation.
It basically converts digital signals to analog and vice versa.

## SFP (Small Form Pluggable adapter)
optical ---- electrical
It is a chip that is used to convert light signals to electrical signals, And we don't need optical model separetely

## Switch
A device which is used to connect to two or more host within a network 

- Unmanageable switch - It redirects traffic to all the host instead of just a particular host. all cheap switches are unmanageable switch
- Manageable switch - It redirects traffic only a to particular host i.e the actual receiver
- L2 - 
- L3 - 

## IP Address classes
N is octet to match
- A - 1-126 - N.H.H.H 
- B - 128-191 - N.N.H.H
- C - 192-223 - N.N.N.H
- 127 is reserved for localhost

## Ping
A tools to check connectivity between connectivity between two host

## Types of IP Address
- IPv4 
  - 32 bit combo
  - 4 groups
  - each group has 8 bits
- IPv6
  - 128 bits combo    
  - 8 groups
  - each group is of 16 bits

## IPv4 classes
IANA - Internet Assigned Numbered Authority 
|128|64|32|16|8|4|2|1|||
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|0|0|0|0|0|0|0|0|0-127|class A|
|1|0|0|0|0|0|0|0|128-191|class B|
|1|1|0|0|0|0|0|0|192-223|class C|
|1|1|1|0|0|0|0|0|224-239|class D|
|1|1|1|1|0|0|0|0|240-255|class E|

## OSI Reference Model
1. Application - the software which interact one human to another works in application layer 
eg  web browsers - HTTP, HTTPS, SMTP, TELNET, POP
1. Presentation - tells about the format of data
eg Image-png, jpeg, Audio-Mp3, wav, Video-mp4 or avi
It is also responsible for Encryption/Decryption and Compression/Decompression
1. Session - create and maintain the session timeframe. Application,Presentation, Session makes data
1. Transport - segements - It breaks the data in segments sequencing i.e. assigning sequence number to each segment. It also take care of re-transmission if packets is lost.
Segemts add port number
All comes under TCP\
  TCP - when we send data with TCP we get the acknowledgement back.\
    Reliable\
    Retransmisson\
    connection oritented\
    more overhead\
  UDP - when we send data with UDP we don't get the acknowledgement back.\
     Non Reliable\
     No Retransmission\
     connectionless\
     less Overhead

1. Network - packets - adds source ip address and destination ip address. Routers work on this layer
1. Data Link - Frame - add mac address of both source and destination. Switches work on this layers 
1. Physical - convert frame into signals (encoding or encapsulation/ decoding or decapsulation)

Top to Down - **All People Seem To Need Data Processing**
Down to Top - **Please do not throw sauces pizza away**

## TCP/IP 
1. Application
1. Transport
1. Network
1. Data
1. Physical 

## IANA (Interner Assigned Numbers Authority)
It is responsbible for assigning ip addresses.
IANA subdivions
- ARIN
- RIPE
- AfriNIC
- LATNIC
- APNIC

global domain like .com, .co etc are controlled by IANA.
.uk, .in etc are controlled by a country

## Routing
Adding my ip to global routing table is called routing. It means telling the world this my server here is its ip address. You have to purchase that ip from IANA

## Subnet Mask
Represents network bits, because computer need extra information that which network it belongs to.


## How to calculate subnet mask

/8 = 255.0.0.0

8 bits is all 1's which equal to 255

for /9 = 8+1 = 255.128.0.0

255 - 8 bits is sum of all bits
128 - 1 bit is on i.e 2^7^

for /10 = 8+2 = 255.192.0.0
255 - 8bits is sum of all bits
192 = 2 bits are on i.e. 2^7^ + 2^6^

## Subnets
Divide a network into smaller parts, to reduce wastage of ips\
ip - 192.150.10.0\
n|h - 24|8\
we have to create 4 network 50 in each\
2^6 - 64\
so\
H - 8-2 = 6\
N - 24+2 = 26
for\
N-26
subnet will be - 255.225.225.192
so ip will be for h-6 2^6 = 64
- category 1 - 0-63
- category 2 - 64-127
- category 3 - 128 - 191
- category 4 - 192 - 25 5


ip - 200.1.1.1/24\
create 8 subnet with any number of host\
N - 24 + 3 = 27  bcuz 2^3=8\
H - 8 - 3 = 5

so ip range will 2^5 = 32
- 0-32
- 33-64
- 65-96
- 97-128 and so on



### Types of Subnet Mask
- FLSM (Fixed Length Subnet Mask) - in which network will same number of host
- VLSM (Variable Length Subnet Mask) - in which each network will have variable number of host.

