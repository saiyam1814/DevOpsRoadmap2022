## What is Docker?
Docker is a tool or platform design to simplify the process of creating, deploying, and packaging and shipping out applications along with its parts such as libraries and other dependencies. Its primary purpose is to automate the application deployment process and operating-system-level virtualization on Linux. 

It allows multiple containers to run on the same hardware and provides high productivity, along with maintaining isolated applications and facilitating seamless configuration.

## Why do we need Docker?
- Compatibilty/Dependency
- Long setup time
- Different Dev/Test/Prod environments

| Web Server | Database | Messaging | Orchestration |
| ---------- | -------- | --------- | ------------- |
| NodeJS Express | mongoDB | Redis | Kubernetes |


## What can Docker do?
- Containerize Applications
- Run each service with its own dependecies in separate containers

| Container | Container | Container | Container |
| --------- | --------- | --------- | --------- |
| Web Server | Database | Messaging | Orchestration |

## What are Containers?
Containers are completely isolated environment as they can have their own processes for services, their own network interfaces, they all share the same OS kernel.

Docker utilizes LXc containers

## Docker Images
The concept of Images and Container is like class and object in which object is instance of class and class is blue print of object. <br/>

Images are different in Virtual Machines and Docker, in virtual machines images are just snapshots of running virtual machine at different point of times but Docker images are little bit different from them and most important and major difference is that docker images are immutable that is they can not be changed.

**Real World Example:**
It happens a lot that a software works on one computer but it does not works on others due to different environments, this issue is completely solved by docker images and using this, application will work same on everyone’s PC . <br/>

Every developer on a team will have exact same development instance. Each testing instance is exactly same as development instance. Your production instance is exactly same as testing instance. <br/>

Also Developers around World can share their Docker Images on a Platform called Docker HUB.

## Docker Container:
If a Docker image is a map of house, then Docker Container is actual build house or in other words we can call it as instance of image.

Containers are runnable instance of an image. You can create, start, stop, move, or delete a container using Docker API or CLI. You can connect a container to one or more networks, attach storage to it, or even create a new image based on its current state.

An application runs using a cluster of containers which are self isolated from one another and also from host machine where they are running.

## Docker Containers V/S Virtual Machines
![image](https://user-images.githubusercontent.com/74575612/155124921-fd003638-3585-450d-9890-4b14fad0cdfb.png)

## Docker Image V/S Docker Container

| Docker Image | Docker Container |
| ------------ | ---------------- |
| It is blueprint of the container | It is instance of the image |
| Image is a logical entity | Container is a real world entity |
| Image is created only once | Containers are created any number of times using image |
| Images are immutable | Containers changes only if old image is deleted and new is used to build the container |
| Images does not require computing resource to work | Containers requires computing resources to run as they run as Docker Virtual Machine |
| To make a Docker Image you have to write script in Dockerfile | To make container form image, you have to run "docker build ." command |
| Docker Images are used to package up applications and preconfigured server environments | Containers use server information and file system provided by image in order to operate |
| There is no such running state of Docker Image | Containers uses RAM when created and in running state |

## Docker Commands With Example

- `docker version` - finding the version
- `docker run nginx` - start a container
- `docker pull` - only pull image and not run container
- `docker images` - list all the docker images pulled on the system
- `docker run` - run the docker image
- `docker ps` - list all the docker containers that are running
- `docker ps -a` - list all the containers running/exited/stopped
- `docker exec` - access the docker container and run commands inside the container
- `docker rm` - remove the docker container
- `docker rmi` - removing image
- `docker restart` - restart the docker container
- `docker stop` - stop a container
- `docker start` - starting docker
- `docker kill` - stops the container immediately
- `docker commit` - save a new docker image
- `docker hub` - login into docker hub
- `docker push` - upload a docker image with the image name
- `docker network` - lists the details of all the network in the cluster
- `docker info` - get detailed information about docker installed
- `docker cp` - copy a file from a docker container to the local system
- `docker history` - checking history
- `docker logs` - checking logs
- `docker search` - searching image
- `docker update` - updating configuration
- `docker volume` - creating volume
- `docker plugin` - installing plugin
- `docker logout` - logging out

**NOTE:** A container only lives as long as the process inside the container is alive.

# Docker

### Docker Environment Variables
In applications we often write something than we want to configure every time we run the application. Generally the good practice is to set that as the env variable. <br/>
We will use `-e` with docker run to pass the environment variable. 

### Docker Images
- You need to create an image if the image you are looking for is not available on the dockerhub. If you decided to containerize the application that you're developing then also you need to create docker image. 

```
1. OS-Ubuntu // Configure OS
2. Update apt repository
3. Install the dependencies using apt
4. Install the python dependencies using pip
5. Copy source code to /opt folder 
6. Run webserver using flask command
```

- Same instruction we will have to write in terms of Dockerfile
```
FROM Ubuntu

RUN apt-get update
RUN apt-get install python

RUN pip install flask
RUN pip install flask-mysql
COPY . /opt/source-code
ENTRYPOINT FLASK_APP=/opt/source-code/app.py flask run
```
- Once we have written the Dockerfile we will build the image with docker build. <br/>
  `docker build -t <container-name>:<tag-name> .`
- To push the container image to docker hub we use the command <br/>
  `docker push <container-name>:<tag-name>`

### Dockerfile
- Dockerfile is a text file written in a specific file that docker can understand. 
- It consist of two things
  1. Instructions
  2. Arguments

- Everything written in the UPPERCASE on the left most side is an instruction.
- Docker builds the image in layered architecture. 
- If we wish to see the layer then you can `docker history <image-name>`
- Rebuilding of images do not start from the absolute beginning. 
- If we update something in our application then only the layer that is added needs to be rebuild.

