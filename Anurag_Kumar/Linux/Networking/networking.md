When two or more devices connect to each other via cables or wireless medium they form a network. Devices connected to a network have their own IP address which is essential for routing packets of information through the network. We need networking to share resources and communicate across the network with other users

Two types of IP addresses are there IPv4 and IPv6 

IPv4 - uses 32 bits and only 4.3 billion unique addresses

IPv6 - uses 128 bits and allows $3.8*10^8$ unique addresses

`ip -4 addr` will give you the IPv4 IP address

`ip -6 addr` will give you the IPv6 IP address

`netstat` will give you all the details about your network system

`netstat -at` will show you only tcp connections

`netstat -l` all the active ports 

`netstat -u` all the UDP connections

## curl

`curl URL` will send a get request to the mentioned URL

`curl -X POST URL` will send a post request to the mentioned URL

A very good image explaining all the important CURL commands that one should know

![curl commands](../../codeimages/linux/curl.jpg)
