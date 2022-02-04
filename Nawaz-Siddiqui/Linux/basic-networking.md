# Basics of Networking

Internet is the largest network in the world. If you are wondering what is a network?

A network is a connection of devices connected with each other through cables and wireless media.

Now, there are millions of machines over the internet, what makes your system unique?

It's the IP address.  It’s an address needed to send network information and packets to your system or for sharing information from one machine to another.

Previously there is IP version 4 type of address is available commonly called IPV4, now the newer version is called IPV6. The limitation if IPV4 is 32 bits unique address which can only be assigned to 4.3 billion machines, but its inadequate. IPV6 uses 128-bits for address meaning 3.4 X 10^38 unique addresses.

A 32-bit IPv4 address is divided into four 8-bit sections called [octets](https://en.wikipedia.org/wiki/Octet_(computing)).

Example:IP address →            172  .          16  .          31  .         46

Bit format →                       10101100.00010000.00011111.00101110

Network addresses are divided into five classes: A, B, C, D, and E. Classes A, B, and C are classified into two parts: Network addresses (Net ID) and Host address (Host ID).

### Class A

Class A addresses use the first octet of an IP address as their Net ID and use the other three octets as the Host ID**.** Each Class A network can have up to 16.7 million unique hosts on its network. The range of host address is from **1.0.0.0** to **127.255.255.255**.

### Class B

Class B addresses use the first two octets of the IP address as their Network ID and the last two octets as the Host ID. Each Class B network can support a maximum of 65,536 unique hosts on its network. The range of host addresses is from **128.0.0.0** to **191.255.255.255**.

### Class C

Class C addresses use the first three octets of the IP address as their Net ID and the last octet as their Host ID. Each Class C network can support up to 256 (8-bits) unique hosts. The range of host addresses is from **192.0.0.0** to **223.255.255.255**.

## Name Resolution

Name Resolution is used to convert numerical IP address values into a human-readable format known as the hostname. For example, **104.95.85.15** is the numerical IP address that refers to the hostname whitehouse.gov. Hostnames are much easier to remember!
