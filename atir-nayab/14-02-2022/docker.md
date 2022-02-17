# docker on aws

we'are going to use aws beanstalk to get our application up and running in a few clicks.

## docker push

the first thing that we need to do before we deploy our app to aws is to publish our image on a registry which can be accessed by aws.

if this is the first image your are pushing an image, the client will ask you to login.

## Beanstalk

aws elastic beanstalk (eb) is a paas (plaform as a service) offered by aws.

## multi-container environments

pull docker elastic search

`docker pull docker.elastic.co/elasticsearch/elasticsearch:6.3.2`

and then run it in development mode by specifying ports and setting an evironment variable that configures the elasticsearch cluster to run as a single-node

`docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2`

once the container is started, we can see the logs by running `docker container logs` with the container name (or ID) to inspect logs

`docker container ls`

`docker container logs es`

sending a request to the elasticsearch container

we start off with the ubuntu LTS base image and use the package manager apt-get to install the dependencies namely -python and node. the yqq flag is used to supress output and assumes 'yes' to all prompt.

we then use the add command to copy our application into a new volume in the container. `/opt/flask-app`. this is where our code will reside.

## docker network

`docker network ls`

when docker is installed, it creates three networks automatically.

the bridge network is the network in which containers are run by default.

bridge network is shared by every container by default, this method is not secure. we need to isolate our network

create our own network

`docker network create foodtrucks-net`

the `network create` command creates a new bridge network, which is what we need at the moment. in terms of docker a bridge network uses a software bridge which allows containers connected tot the same bridge network to communicate, while providing isolation from containers which are not connected ot that bridge network. the docker bridge driver automatically installs rules in the host machines so that containers on different bridge networks cannot communicate directly with each other.

now that we have a network, we can launch our containers inside this network using the --net flag.

`docker container stop es`

`docker containre rm es`

`docker run -d --name es --net foodtrucks-net -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2`

on user defined networks like foodtrucks-net containers can not only communicate by address, but can also resolve a container name to an IP address. this capability is called automatic service discovery.
