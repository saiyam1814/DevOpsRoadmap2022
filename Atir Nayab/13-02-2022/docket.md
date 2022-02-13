# docker introduction

## what is docker

- docker is a company
- container runtime and orchestration engine
- open source
- founder, solomon hykes
- dotcloud was rebranded as docker
- tow editions are available
  - community edition
  - enterprise edition

## need of docker

- saving time and money
- dev/prod parity
- simplifying configuration
- developer productivity
- app isolation

## docker as a solution

- containerize applications
- run each service with its own dependencies

## docker commands

- `docker images` - legacy list of images
- `docker image ls` - list of images
- `docker pull ubuntu` - to pull ubuntu image
- `docker build .` - build docker image from docker file

# Written and developed by Prakhar Srivastav

## what is docker?

an open source project that automated the deployment of software applications inside containers by an additional layer of abstraction and automation of os-level virtualization on linux.

docker is a tool that allows developers, sys-admins to easily deploy their applications in a sandbox (called containers) to run on the host operating syste,

## what are containers

the industry standard today is to use virtual machines (vms) to run software applications. vms run applications inside a guest operating system, which runs on virtual hardware powered by the sreve's host os.

containers take a different approach; by leveraging the low-level mechanics of the host operating system, containers provide most of the isolation of virtual machines at a fraction of the computing power.

## why use containers

containers offer a logical packaging mechanism in which applications can be abstracted from the environment in which they actually run. this decoupling allows containers-based applications to be deployed easily and consistently, regardless of whether the target environment is a private data center, the puclic cloud, or even a developer's personal laptop.

## playing with busybox

`docker pull busybox` the pull command fetches the busybox image from the docker registry and saves it to out system.

## docker run

`docker run busybox echo "hello from busybox"`

## container list

- `docker ps`
- `docker ps -a` - shows a list of all containers that we ran. do notice that the status column shows that these containers exited a few minutes ago.

## run in interactive mode

`docker run -it busybox sh` running the run command with the -it flags attaches us to an interactive tty in the container. now we can run as many commands in the container as we want.

> DANGER ZONE: you can try rm -rf bin in the container. make sure you run this command in the container and not in your laptop/desktop. doing this will make any other commands like ls, uptime not work. once everthing stops working. you can exit the container and then start it up again with the `docker run -it busybox`. since docker creates a new container every time, everything should start working again.

we saw above that we can still see remnants of the container even after we've exited by running `docker ps -a`. you'll run `docker run` multiple times and leaving stray containers will eat up disk space. hence as a rule of thumb, i clean up containers once i'm done with them, i clean up containers will eat up disk space hence as a rule of thumb, i clean up containers once i'm done with them. to do that you can run the docker rm command. just copy the container ids from above and paste them alongside the command

on deletion you should see the ids echoed back to you. if you have a bunch of containers to delete in one go, copy pasting ids can be tedious in that case you can simply run

`docker rm $(docker ps -a -q -f status=exited)`

this command deletes all containers that have a status of exited. in case you're wondering the -q flag, only returns the numeric ids and -f filters output based on conditions. One last thing that'll be useful is the --rm flag that can be passed to `docker run` which automatically deletes the container once it's exited from.

in the laster versions of docker the `docker container prune` command can be used to achieve the same effect.

# terminology

- images - the blueprints of our application which form the basis of containers, in the demo above we used the `docker pull` command to download the busybox image.
- containers - created from docker images and run the actual application. we create a container using docker run which we did using the busybox image that we downloaded.
- docker daemon - the background service running on the host that manages building, running and distributing docker containers. the daemon is the process that runs in the operating system which clients talk to.
- docker client - the command line tool that allows the user to interact with the daemon.
- docker hub - a registry of docker images you can think of the registry as a directory of all available docker images. if requried, one can host their own docker registries and can use them for pulling images

## detached mode

the client is not exposing any ports so we need to re-run the `docker run` command to publish ports. while we're at it, we should also find a way so that out terminal is not attached to the running container. this way you can happily close your teminal and keep the container running this is called detached mode.

`docker run -d -P --name static-site prakhar1989/static-site`

in the above command `-d` will detach out terminal, `-P` will publish all exposed ports to random ports and finally `--name` corresponds to a name we want to give. now we can see the ports by running the `docker port [container]` command

`docker port static-site

you can also specify custom port to which the client will forward connections to the containers

`docker run -p 8888:80 prakher1989/static-site`

to stop a detached container run `docker stop` by giving the container id or name

## docker images

docker images are the basis of containers,

the `tag` refers to a particular snapshot of the image and the `image id` is the corresponding unique identifier for that image.

for simplicity you can think of an image akin to a git repository - images can be commited with changes and have multiple versions. if you don't provide a specific version number, the client defaults to latest. for example,, you can pull a specific version of ubuntu image

`docker pull ubuntu:18.04`

### different between base and child images

- base images are images that have no parent image. usually images with an os like ubuntu, busybox or debain.
- child images are images that build on base images and add additionally functionality.

then there are official and user images, which can be both base and child images.

- official images are images that are officially maintained and supported by the folks at docker. these are typically one word long.
- user images are images created and shared by users like you and me. they build on base images and add additional functionality. typically, these are formatted as `user/image-name`

## dockerfile

a dockerfile is a simple text that contains a list of commands that the docker client calls while creating an image. it's a simple way to automate the image creation process. the best part is that the commands you write in a dockerfile are almost identical to their equivalent linux commands.

we start with specifying our base image. use the `from` keyword to do that

`FROM python:3`

the next step usually is to write the commands of copying the files and installing the dependencies. first set a working directory and then copy all the files for our app.

```
# set a directory for the app
WORKDIR /usr/src/app

# copy all the files to the container
COPY . .
```

now that we have the files we can install the dependencies

`RUN pip install --no-cache-dir -r requirements.txt`

the next thing we need to specify is the port number that needs to be exposed. since our flask app is running on port 5000

`EXPOSE 5000`

the last step is to write the command for running the application, which is simply `python ./app.py`

`docker build -t yourusername/catnip .`
