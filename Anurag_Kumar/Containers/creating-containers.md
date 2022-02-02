- docker build -t image-name .

- FROM - this specifies the parent image of your dockerfile. 


- LABEL - it is used for providing the key-value pair to the image. You can add as many key-value pairs as you want. 
```
LABEL key1="value 1" key2="value 2" 
```

- ADD - it is a way to copy files from your local filesystem to the filesystem of the images you are going to build. We can also used remote URL and local compressed archive file. 
```
ADD <source> <destination>
```

- ENV - it is used to set environment variable during the run and build stages of the image. 
```
ENV <key> <value>
ENV <key>=<value>
```

- COPY - It is used to copy the local files into the docker files. It is different from because this only works with the local files.

- EXPOSE - It defines on which port container services are listening 
```
EXPOSE <port>
EXPOSE <port>/tcp
EXPOSE <port>/udp
```
- RUN - It executes the instructions given during the build process inside the image filesystem environment. 

```
RUN command 
```

- VOLUME - it is used to creating mount points in the within the container image filesystem. It can mount to your local system or between the containers. 

```
VOLUME ["/data/vol1", "/data/vol2"]
VOLUME /data/vol1 /data/vol2
```

- [credits](https://www.suse.com/c/rancher_blog/how-to-build-and-run-your-own-container-images/)
