# docker-compose

all docker configuration are stored in yml file so with a single file we can control/configure multiple containers

to create container from docker compose `docker-compose up -d` to run in detach mode

to stop container and delete all volume mounts `docker-compose down -v`

docker-compose is dump if container was down and did some changes to file it will build image again, we have to force build the image.

adding multiple docker compose file based on production evironment
`docker-compose -f <file1> -f <file2> up -d`, here oder of the files matters, file at begining loads first.
