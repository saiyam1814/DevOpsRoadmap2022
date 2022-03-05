# Started Learning [Kubernetes from Civo Academy](https://www.civo.com/account/academy)

## Basic Linux Commands
-	`ls`			lists information about the files
-	`mkdir`			creating a new directory
-	`cd`			changing the directory
-	`touch`			creating an empty file
-	`vi`			text editor
-	`cat`			reading the files sequentially
-	`cp`			copy command
-	`mv`			move command
-	`chmod`		changing the permissions
-	`kill`			killing the process
-	`env`			getting the list of environment variables
-	`export`		exporting the variables
-	`curl`			transfer data from or to a server
-	`tar`			archive utility
-	`du`			disk usage command
-	`ping`			network connectivity
-	`netstat`		network statistics to monitor network connections
-	`nslookup`		name server lookup
-	`ps`			process states
-	`tail`			prints lasts N lines of a file
-	`systemctl`		to control services
-	`journalctl`		view systems logs
-	`top`			memory and cpu utilisation view

## Monolithic VS Microservices

### Monolithic Application
If all the functionalities of a project exists in a single codebase, then that application is known as monolithic application.

_Advantages of Monolithic Application:_
-	Simple to develop relative to microservices where skilled developers are required in order to identify and develop the services.
-	Easier to deploy as only a single jar/war file is deployed.
-	The problems of network latency and security are relatively less in comparison to microservices architecture.

_Disadvantages of Monolithic Application:_
-	It becomes too large in size with time and hence difficult to manage.
-	We need to redeploy the whole application even for a small change
-	As the size of the application increase, its start-up and deployment time also increases.

### Microservices Application
It is an architectural development style in which the application is made up of smaller services communicating with each other directly using lightweight protocols like HTTP. 
Microservices are the small services that work together.

_Advantages of Microservices Application:_
-	It is easy to manage as it is relatively smaller in size
-	Microservices are self-contained and hence, deployed independently. Their start-up and deployment time are relatively less.
-	Each microservice can use different technology based on the business requirements.

_Disadvantages of Microservices Application:_
-	Being a disrupted system, it is much more complex that monolithic applications. Its complexity increases with the increase in number of microservices.
-	Independent deployment of microservices is complicated.
-	Microservices are costly in terms of network usage as they need to interact with each other and all these remote calls results into network latency.

![image](https://user-images.githubusercontent.com/74575612/156892886-8c689d0f-cde3-46b7-92df-981096789a62.png)

## Introduction to Containers
A **container** is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another.

![image](https://user-images.githubusercontent.com/74575612/156892907-644d7944-8074-4133-8e38-350593a7c1fa.png)

_Learnt from the modules of Before Kubernetes and Kubernetes Introduction_
_Rest of the notes will be updates soon._
