## registries

- docker iamges are stored in registries
- whenever you pull an image, you pull it from the official docker registry (hub.docker.com)
- docker repositories and images can be categorized as secure or insecure.

## what needs to be secured

- docker host
- the docker daemon
- containers
- authentication and communication between the components of the docker platform need to be setup and secured.
- registry security and authenticity needs to be analyzed and understood.

## docker manifest

docker uses manifest file to generate image, it all sha of the layers. `docker image ls digest` gives the sha of manifest slight change in manifest generate new sha.

# all articles on containers

## attach vs exec - what's the difference

![container management](https://iximiuz.com/containers-101-attach-vs-exec/docker-containerd-runc-2000-opt.png)

**three key takeaways**

- container management architecture is layered
- a container is ~a regular process~ an isolated and restricted environment + a process
- there is a shim component in between the container manager and the container.

### what does attach command do

inside a container, there is a regular linux process. as every normal process, it has stdio streams - stdin, stdout, and stderr. but what happens to these streams when a container is started in the detached (i.e., daemon-like) mode:

`docker run -d nginx`

back in the day, when you started a process as a daemon (i.e., detaching it from the starter process), it would be reparented to pid 1, and its stdio streams would be simply closed.

docker came up with a clever idea of putting an extra process between the container and the rest of the system, called a container runtime shim.  
this is the shim process that becomes a daemon - it's reparented to pid 1, and its stdio streams are closed.

![ps axfo pid, ppid, command](https://iximiuz.com/containers-101-attach-vs-exec/ppid-1-2000-opt.png)

however, the shim takes control over the container's stdio streams!

the daemonized shim process reads from the container's stdout and stderr and dumps the read bytes to the log driver. by default, the shim closes the container's stdin stream, but it can keep it open if `-i` was passed to the corresponding `docker run` command.

container runtime shim actually acts as a server! it provides rpc means (for instance, a UNIX socket) to connect to it. and when you do so, it starts streaming the container's stdout and stderr back to your end of the socket. hence, you kind of attach to the container's stdio streams!

so, finally when you run `docker attach <container>`, you basically create a relay.

`terminal <-> docker <-> dockerd <-> shim <-> container's stdio streams`

![docker attach command](https://iximiuz.com/containers-101-attach-vs-exec/docker-attach-2000-opt.png)

### difference between attach and logs

`docker attach` streams the container's logs back to the terminal. however the logs command does a similar thing.

the logs command provides various options to filter the logs while attach in that regard acts as a simple tail. but what's even more important is that the stream established by the logs command is always unidirectional and connected to the container's logs, and the container's stdio streams directly.

the logs command simply streams the content of the container's logs back to your terminal, and that's it. so regardless of how you created your container (interactive or non-interactive, controlled by a psuedo-terminal or not), you cannot accidentally impact the container while using the logs command.

however when attach is used:

- if a container was created in the interactive mode (-i), everything you type in the terminal after `attach`-ing to the container will be sent to its stdin.
- you can (intentionally or accidently) send a signal to the container - for instance, hitting ctrl+c on your end while attached sends SIGNINT to the container.

### what does exec command do

in the case of attach, we were connecting our terminal to an existing container (read, process). however, the exec command starts a totally new container! in other words, exec is form of the run command (which itself is just a shortcut for create + start).

the trick here is that the auxiliary container created by the exec command shares all the isolation boundaries of the target contaienr! i.e., the same net, pid, mount, etc. namespaces, same cgroups hierarchy, etc. so from the outside, it feels like running a command inside of an existing container.

the confusion of attach and exec commands arises because the exec -uted command is also a process with its own stdio streams.so, you can choose whether to exec in detached mode, whether to keep the stdin open,  
also, when `exec`-ing the relay looks quite similarly:

`terminal<->docker cli<->dockerd<->shim<->command's stdio streams`

![docker exec command](https://iximiuz.com/containers-101-attach-vs-exec/docker-exec-2000-opt.png)
