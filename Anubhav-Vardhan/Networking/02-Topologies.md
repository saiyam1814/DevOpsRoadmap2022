## BUS

- Single backbone and all the devices are connected to it.
```
  C1    C2    C3    C4
  |     |     |     |
----------------------------
    |     |     |     |
    C5    C6    C7    C8
```
- The problem is if the main link (backbone) is broken, entire network is spoiled.
- Since everything is being transferred by a single cable, only one person can send data at a time

## Ring

- Every system communicates with one another
```
  C1----C2----C3
  |           |
  C8          C4
  |           |
  C7----C6----C5
```
- If we send data from C1 to C4, it will have to go through C2 and C3. Hence, lots of unnecessary calls are being made
- If one of the links break, we won't be able to transfer data

## Star

- One central device connected to all the computers
```
      C1
      |
C2----D----C3
      |
      C4
```
- If C1 wants to send data to C4, it will have to communicate via D
- If central device i.e D fails, network will go down

## Tree

- Combination of BUS and Star topologies
- Many Star topologies connected to a BUS
```
    A4            B4            C4
    |             |             |
A2--A1--A3    B2--B1--B3    C2--C1--C3
    |             |             |
------------------------------------------
           |             |
       D2--D1--D3    E2--E1--E3
           |             |
           D4            E4
```

# Mesh

- Every device is connected to every device
```
         C1
        /  \
       /    \
      /      \
     C2-------C3

Similarly with any number of devices
```
- It is most expensive
- Scalability issues as we have to connect new device to every other existing device
