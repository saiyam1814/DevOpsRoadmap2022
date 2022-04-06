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

Resources:
- [Docker for Beginners: Full Course](https://www.youtube.com/watch?v=zJ6WbK9zFpI)