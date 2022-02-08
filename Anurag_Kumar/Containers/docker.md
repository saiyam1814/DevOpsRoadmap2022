- Credits - [KodeKloud course](https://youtu.be/zJ6WbK9zFpI)


### Docker commands
- docker run nginx - this will run a container from an image. If the image is not present then it will be pulled from dockerhub.
- `docker ps` - list the running containers
- `docker ps -a` - list all the containers those which are running as well as those which are not runnning.
- `docker rm` - remove a container
- `docker images` - list images
- `docker rmi nginx` - to remove nginx image
- `docker pull nginx` - for pulling the image. It will not create the container. 
- `docker inspect <container-name>` to get more details about the container. It returns all the details in JSON format such as state, mounts, network settings etc. 
- `docker logs <container-name>` to look at the logs of your container. 
- A container only lives as long as the process inside the container is alive.


### Docker environment variables
- In applications we often write something than we want to configure every time we run the application. Generally the good practice is to set that as the env variable. 
- We will use `-e` with docker run to pass the environment variable. 

#### How do you find the environment variable of a container that is already running?
- use `docker inspect <container-name>` and under the config section you will find the list of env variables.


### Docker Images
- You need to create an image if the image you are looking for is not available on the dockerhub. If you decided to containerize the application that you're developing then also you need to create docker image. 
- Think if you want to develop an application from the starting then what you will do.
```
1. OS-ubuntu // configure OS
2. update apt repo
3. install dependencies using apt
4. Install python dependencies using pip
5. Copy source code to /opt folder 
6. Run webserver using flask command
```

- Same instruction we will write in terms of Dockerfile
```
FROM Ubuntu

RUN apt-get update
RUN apt-get install python

RUN pip install flask
RUN pip install flask-mysql
COPY . /opt/source-code
ENTRYPOINT FLASK_APP=/opt/source-code/app.py flask run
```
- Once we have written the Dockerfile we will build the image with docker build. 
- `docker build -t <container-name>:<tag-name> .`
- To push your container image to docker hub use the command `docker push <container-name>:<tag-name>`

### Dockerfile
- Dockerfile is a text file written in a specific file that docker can understand. 
- It consist of two things
  - Instructions
  - Arguments

- Everything written in the uppercase on the left most side is an instruction.
- Docker builds the image in layered architecture. 
- If you wish to see the layer then you can `run docker history <image-name>`
- Rebuilding of images do not start from the absolute beginning. 
- If you update something in your application then only the layer that is added needs to be rebuild not the whole application. 
