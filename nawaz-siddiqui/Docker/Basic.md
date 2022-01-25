# Docker

Before understanding Docker lets first understand what is a container?

As the name suggests, a container is a package but in terms of software. It is a way to package applications with all the necessary dependencies and configurations needed to run the software. So it's a package that is portable and can easily be shared. With these features, containers make the development and deployment cycle more efficient. Now you might be thinking that It's great to package and ship applications and their dependencies, but where does a container live? It will consume a lot of storage in case of multiple applications so where should it be stored?

Container lives in container repository, some company has their own private repository. For example, we are going to use Docker for this purpose, and docker has a public repo called docker hub to store our application.

### How do containers improve the development process?

Before the use of containers, a software development process is like this:

You have a javascript application that depends on Python, Redis, SQL, and other software. Now each time you need to run the javascript app you must have the dependent software installed in your local machine. For a single time it is okay, but think of it like when you have ten dependencies and you need to install all of them in your local machine and every developer of the team has to do that. But not every installation cycle is smooth. Some errors might happen. These are the combination of problem containers try to address and solve.

With containers in action, you don't need to install any dependency on your local machine, and chances of error and configuration time are highly reduced making the development cycle smooth and efficient. Some advantages of docker based containers are:

- Own isolated environment
- Packaged with all needed configuration
- One command to install the app
- Can run the same app with 2 different version

## What is a Container?

Now you understand containers in a very simple way. Let’s understand it with a more technical approach. A container contains:

- Multiple Layers of images (mostly Linux based due to its low size)
- Application image

### Docker Image

- It is the actual package
- It can be moved around
- It doesn't run and is used for the purpose of portability

### Docker Container

- Actually starts the application
- container environment is created
- It can be run on the machine
- It is a running environment for images.

## Docker Installation

Refer to this link and follow the instruction for your machine:

[https://docs.docker.com/engine/install/](https://docs.docker.com/engine/install/)

