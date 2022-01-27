### Docker Commands

| Command | Function |
| --- | --- |
| docker ps | shows all the running containers |
| docker images | checks all the existing docker images |
| docker run image-name | runs the image |
| Ctrl+C | Terminates a running container |
| docker run -d image-name | runs a container in a detached mode |
| docker stop container-ID | stops a container |
| docker start container-ID | starts a stopped container |
| docker ps -a | Lists running and stopped container |
| docker run image:version | Pulls and run the different version of the already installed image. |
| docker run -p(host_port):container_port | binds host port and container port |
| docker logs container_ID | Shows the log of the particular container or you can also use the container name in place of ID.  |
| docker run —name new-name older-name | Changes the name of the container and runs it |
| docker exec -it container_ID path-name | Takes you to the path of the mentioned container ID |
| exit | Exits the terminal |
| docker network ls | Lists docker network |
| docker network create network-name | creates a new docker network |
| docker logs container_ID -f | Streams the logs. |
