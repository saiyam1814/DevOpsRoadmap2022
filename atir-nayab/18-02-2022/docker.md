# docker inpect

give datailed information about our container.

`docker inspect <container name>`

# network logs

`docker logs <container name> -f`

> we also tell which service to run first by telling which service depend on which

# docker file best practices

**areas of improvement**

- (incremental) build time
- image size
- maintainability
- security
- consistency/repeatability

## let's improve this dockerfile

```
FROM debian
COPY . /app
RUN apt-get update
RUN apt-get -y install ....
RUN apt-get -y install ....
CMD []
```

## incremental build time

### order matter for caching

```
FROM debian
RUN apt-get update
RUN apt-get -y install ....
RUN apt-get -y install ....
COPY . /app
CMD []
```

_order from least to most frequently changing content._

## more specific COPY to limit cache busts

```
FROM debian
RUN apt-get update
RUN apt-get -y install ....
RUN apt-get -y install ....
COPY target/app.js /app
CMD ["java", "-jar", "/app/app.jar"]
```

_only copy what's needed. avoid copy if possible_

## line buddies: apt-get update & install

```
FROM debian
RUN apt-get update\
 && apt-get -y install ....\
 && apt-get -y install ....
COPY target/app.js /app
CMD ["java", "-jar", "/app/app.jar"]
```

_prvents using an outdated package cache_

## reduce image size

faster deploys, smaller attack surface

### remove unnecessary dependencies

```
FROM debian
RUN apt-get update\
 && apt-get -y install .... --no-install-recommends\
 && apt-get -y install ....
COPY target/app.js /app
CMD ["java", "-jar", "/app/app.jar"]
```

## remove package manager cache

```
FROM debian
RUN apt-get update\
 && apt-get -y install .... --no-install-recommends\
 && apt-get -y install ....
 && rm -rf /var/lib/apt/lists/*
COPY target/app.js /app
CMD ["java", "-jar", "/app/app.jar"]
```

## maintainability

### use official images where possible

- reduce time spent on maintenance (frequently updated with fixes)
- reduce size (shared layers between images)
- pre-configured for container use
- built by smart people

```
FROM openjdk
COPY target/app.jar /app
CMD ["java", "-jar", "/app/app.jar"]
```

### use more spesific tags

_the latest tag is a rolling tag, be specific to prevent unexpected changes in your base image_

### look for minimal flavors

just using a different base image reduced the image size by 540MB

### build from source in a consistent environment

make the dockerfile your blueprint:

- it describes the build environment
- correct versions of build tools installed
- prevent inconsistencies between environments
- there may be system dependencies.
- the source of truth is the source code, not the build aritfact.

# different image flavors (DRY/global arg)

`ARG flavor=alpine`

we can declare variable to use every where before first from.

## from linear dockerfile stages

- all stages are executed in sequence
- without buildkit, unneeded stages are unnecessarily executed but discarded

## ...to multi-stage graphs with buildkit

- buildkit traverses from bottom (stage name from --target) to top
- unneeded stages are not even considered

## multi stage: build concurrently

```
FROM builder-base AS builder
COPY --from=builder-someClib /out /
COPY --from=builder-someCPPlib /out /
```

## secrets

`RUN --mount=type=secret,id=aws, target=/root/.aws/credentials, required ./fetch-assets-from-s3.sh`

# docker security essentials

## what is docker

- docker allows you to build and deploy applications and services in the form of containers
- it is a platform as a service (PaaS) offering that utilizes the host os kernal as opposed to hypervisors like virtual box.
- the containers contain the dependencies and libraries that the application or service needs to run, therefore eliminating the need for installing dependencies manually.
- docker containers are much more effiecient than vm's as the utilize the host kernal.
