# before kubernetes

## linux commands

- ls - list information about the files
- mkdir - creating a new directory
- cd - changing the directory
- touch - create an empty file
- vi - text editor
- cat - reading the files sequentially
- cp - copy command
- mv - move command
- chmod - changing file permissions
- kill - killing the process
- env - getting the list of environment variables
- export - exporting the variables
- curl - transfer data from or to a server
- tar - achive utility
- du - disk uility
- ping - network connectivity
- netstat - network statistics to monitor network connections
- nslookup - name server lookup
- ps - process status
- tail - prints last n lines of a file
- systemctl - to control services
- journalctl - view systemd logs
- top - memory and cpu utilization view.

## monolithic vs microservices

The monolithic architecture is considered to be a traditional way of building applications. A monolithic application is built as a single and indivisible unit. Usually, such a solution comprises a client-side user interface, a server side-application, and a database. It is unified and all the functions are managed and served in one place.

While a monolithic application is a single unified unit, a microservices architecture breaks it down into a collection of smaller independent units. These units carry out every application process as a separate service. So all the services have their own logic and the database as well as perform the specific functions.

# Introduction to Kubernetes

## Chapter 1. From Monolith to Microservices

### the legacy monolith

a monolith has a rather expensive taste in hardware. being a large, single piece of software which continuously grows, it has to run a single system which has to satisfy its compute, memory, storage and networking requirements. the hardware of such capacity is both complexity and extremely pricey.  
since the entire monolith application runs as a single process, the scaling of individual features of the monolith is almost impossible. it internally supports a hardcoded number of connections and operations. however scaling the entire application can achieved by manually deploying a new instance of the monolith on another server, typically behind a load balancing appliance - another pricey solution.

microservices can be deployed individually on separate servers provisioned with fewer resources - only what is required by each service and the host system itself, helping to lower compute resource expenses.  
microservices-based architecture is aligned with event-driven architecture and service-oriented architecture (soa) principles, where complex applications are composed of small independent processes which communicate with each other through apis over a network.
