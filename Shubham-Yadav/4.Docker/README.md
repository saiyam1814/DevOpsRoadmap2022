# docker images commands

- Shows the help for the `docker images` command.
``` 
docker images --help 
``` 

- pulls an image from a registry.
```
 docker pull [image] `
```  

- Pulls an image from a registry and tag it.
```
 docker pull [image]:[tag] 
``` 

- Lists all images in the local registry.
```
 docker images 
``` 

- list image id only
```
 docker images -q 
``` 

- list dangling images
```
 docker images -f "dangling=true" 
``` 

- Runs an image in a new container.
```
 docker run [image] 
``` 

- Lists all running containers.
```
 docker ps 
 ```  

- Runs an image in a new container and names it.
```
 docker run --name [container] [image] 
``` 

- Runs an image in a new interactive container.
```
 docker run -it [image] 
``` 

- Runs an image in a new interactive container and names it.
```
 docker run -it --name [container] [image] 
``` 

- Inspects a container.
```
 docker inspect [container] 
``` 

- Removes an image from the local registry.
```
 docker rmi [image] 
``` 

- Removes an image from the local registry and force remove it.
```
 docker rmi -f [image] 
``` 

- Shows the history of an image.
```
 docker history [image] 
``` 

# docker container commands

- Lists all running containers.
```
 docker ps 
``` 

- Runs an image in a new container.
```
 docker run [image] 
``` 

- Starts a container.
```
 docker start [container-name] / [container-id] 
``` 

- Stops a container.
```
 docker stop [container-name] / [container-id] 
``` 

- Pauses a container.
```
 docker pause [container-name] / [container-id] 
``` 

- Unpauses a container.
```
 docker unpause [container-name] / [container-id] 
``` 

- Restarts a container.
```
 docker restart [container-name] / [container-id] 
 ``` 
- remove all stopped containers
```
 docker container prune
```

- Shows the running processes of a container.
```
 docker top [container-name] / [container-id] 
 ``` 

- Shows the resource usage of a container.
```
 docker stats [container-name] / [container-id] 
 ``` 

- Inspects a container.
```
 docker inspect [container-name] / [container-id] 
 ``` 

- Attaches to a container.
```
 docker attach [container-name] / [container-id] 
 ``` 
docker attach usage : Use docker attach to attach your terminal’s standard input, output, and error (or any combination of the three) to a running container using the container’s ID or name. This allows you to view its ongoing output or to control it interactively, as though the commands were running directly in your terminal.

- Kills a container.
```
 docker kill [container-name] / [container-id] 
 ``` 

- Removes a container.
```
 docker rm [container-name] / [container-id] 
 ``` 

- Shows the history of a container.
```
 docker history [container-name] / [container-id] 
 ``` 


# dockerfile commands

- Builds an image from a Dockerfile.
```
 docker build [path] 
 ``` 

- Builds an image from a Dockerfile and tag it.
```
 docker build -t [image] [path] 
 ``` 

# docker compose commands

- Shows the configuration of a service.
```
 docker compose config [service] 
 ``` 

- Starts services.
```
 docker-compose up [service] 
 ``` 

- Starts services and keep them running.
```
 docker-compose up -d [service] 
 ``` 

- Stops services.
```
 docker-compose down [service] 
 ``` 

- Lists all running services.
```
 docke compose ps 
 ``` 

## docker volume commands

- get information
``` 
docker volume 
```

- create a volume
```
docker volume create [name]
```

- remove a volume
```
docker volume rm
```

- list all volumes
```
docker volume ls
```

- inspect a volume
```
docker volume inspect [volume-name]
```

- remove all unused volumes
```
docker volume prune
```