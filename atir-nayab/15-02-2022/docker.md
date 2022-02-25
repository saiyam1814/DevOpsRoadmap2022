## docker compose

there are a bunch of other open-source tools which play very nicely with docker. a few of them are:-

- docker machine - create docker hosts on your computer, on cloud providers, and inside your own data center.
- docker compose - a tool for defining and running multi-container docker applications.
- docker swarm - a native cluster solution for docker
- kubernetes - kubernetes is an open system for automating deployment, scaling and management of containerized applications.

compose is a tool that is used for defining and running multi-container docker apps in an easy way. it provides a configuration file called `docker-compose.yml` that can be used to bring up an application and the suite of services it depends on with just one command. compose works in all environments; production, staging development, testing, as wellas ci workflows, although compose is ideal for for development and testing environments.

breakdown docker compose file

at the parent level, we define the names of our services es and web. the image parameter is always required, and for each service that we want docker to run, we can add additional parameters. for es, we just refer to the elasticsearch image available on elastic reqistry,

other parameters such as command and ports provide more information about the container. the volumes paramter specifies a mount point in our web container where the code will reside. this is purely optional and is useful if you need access to logs, etc. we'll later see how this can be useful during development

> note: you must be inside the directory with the docker-compose.yml file in order to execute most compose commands.

before we start, we need to make sure the ports and names are free. so if you have the flask and es containers running lets turn them off.
