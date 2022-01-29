## What is Docker?
* Docker is a tool for running applications in an isolated Environment.
* It is similar to Virtual Machine but it's much faster and does not require a lot of memory and an entire operating system to operate.
* The cool thing about docker that your application runs in exact same environment for eg. if it works on my machine it will definitely work on your machine, or if it works on the staging environment it will also work in production environment.

## Containers vs VM
#### CONTAINERS
* containers are an abstraction at the app layer that packages code and dependencies together. Multiple containers can run on the same machine ans share the OS kernel with other containers, each running as isolated process in user space.

### Virtual Machine 
* An application on a VM requires a guest OS and thus an underlying hypervisor to run. Hypervisor is used to create multiple machines on a host operating system and it manages virtual machines. These virtual machines have their own operating system and do not use the host’s operating system. They have some space allocated.
![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/4i78rb96nnjdwd71eumr.jpeg)

# Docker Images
* Image is a template for creating an environment of your choice.It is also an  snapshots.
* Images are immutable. Once built, the files making up an image do not change. Images can be stored locally or remote 
locations like hub.docker.com.

Once you have an image, then containers comes into play. A container is a running instance of an image.

To see the basics commands of docker refer the link below:-

[useful docker commands](https://dev.to/shubh/docker-cheat-sheet-4ob3)



![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/8jxil34zw3wr6d0e4tf2.png)
# Dockerfile
* a docker file allows us to create our own images by creating a file called dockerfile. This files simply containes the list steps to create images, then we can run images created from dockerfile.

# Docker Registries
* This is a server side application that stores Docker images.
* Images can be stored on your local or can be pushed to public registries like Docker Hub.

# Docker Compose
* Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration.It is a declarative way of describing the components of a Docker application.

# Docker Volume and Bind Mounts
* By default all the files inside the container are stored on a writable container layer.
* the data does not persist when the container is no longerr running.
* Docker has two ways to persist data.
* 1. Volumes
* 2. Bind Mounts

### Volumes
- Volumes are a way to persist data between containers.
- volumes are stored in a part of the filesystem which is managed by Docker.
- a gven volume can be mounted into multiple containers.
- when no running container is using the volume, it is still available.

### Bind Mounts
- Bind mount are storesd anywhere on the host system.
- bind mounts are used to mount a volume from a container into another container.
- In Bind Mounts, the file or directory is referenced ny its full path in the host system.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/o7fsdubtyt603h24p7xu.png)




## To see basic commadns like pulling building images etc.[Click Here](https://dev.to/shubh/docker-cheat-sheet-4ob3)

Now that you know what docker is you can now create your application with tons of containers that interact with each other so you can imagine the effort to manage hundreds of container manually, every time a container crash or the communications between them is not working you have to manually fix these problems. Here is where container-orchestration system comes into play, It helps in automating application deployment, scaling, and management. One of the very widely used COE is Kubernetes.

## connect :[Twitter](https://twitter.com/shubhztwt)