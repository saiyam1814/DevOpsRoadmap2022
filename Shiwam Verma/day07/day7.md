#Completed Docker Tutorial 
- Docker is open platform for developing, shipping, and running applications.

###Learnt about containers
- Containers are like a small space inside your machine, in which you can run various applications such as mySQL, mongoDB.
- It saves you from downloading these applications.
- A container is a standard unit of software that packages up code and all its dependencies so the application runs quickly.
- The application runs reliably from one computing environment to another.

###Learnt about Docker Images
- A Docker container image is a lightweight, standalone, executable package of software that includes everything needed to run an application: code, runtime, system tools, system libraries and settings.
- It is like the set of instructions and code required to run an application in container.

###Learnt about Docker architecture
- It can be summed as ```docker CLI -> Rest API -> Daemon server```
- Docker clients communicates with daemon server which communicates with containerd which connects with ```shim``` and further to ```runc``` which finally connects with application.
```Docker CLI -> daemon server -> containerd -> shim -> runc -> app```

-This makes it possible to update the daemon server without interrupting with the running container.  
