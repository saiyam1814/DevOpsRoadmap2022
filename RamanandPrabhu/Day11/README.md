#Day11 Contd Containers
* Fundamentals of Containers, Kubernetes.
  * SW applications are typically deployed as a single set of libraries and configuration files to a runtime env. They are traditionally deployed to an env that provides some services like DB or HTTP. This can be on VM or a physical host.
  *	The drawback to the above is that the update to OS might update the OS libraries used by programming langs and that might affect the running applications with incompatible updates.
  *	Moreover, with multiple applications running on same host and sharing libraries, the update of libraries for one application can affect another application.
  *	Thus any update to the running env will require through test of the application.
  *	Depending on the complexity of an application, the regression verification might not be an easy task and might require a major project. Furthermore, any update normally requires a full application stop. Normally, this implies an environment with high-availability features enabled to minimize the impact of any downtime, and increases the complexity of the deployment process. The maintenance might become cumbersome, and any deployment or update might become a complex process.
  *	Alternativly, a Sys Admin can work with containers, which are isolated partition inside a single OS system. Containers provide the same services as VM, such as security, storage, nw isolation, while requiring fewer HW resources and being quicker to launch and terminate.
  *	They also isolate the libraries and the runtime environment (such as CPU and storage) used by an application to minimize the impact of any OS update to the host OS.
  *	The use of containers helps not only with the efficiency, elasticity, and reusability of the hosted applications, but also with portability of the platform and applications. There are many container providers available, such as rkt(Rocket), Drawbridge, and LXC, but one of the major providers is Docker.
  *	Advantages of Containers
	-	Low hw footprint: Uses OS internal features to create an isolated environment where resources are managed using OS facilities such as namespaces and cgroups. This approach minimizes the amount of CPU and memory overhead compared to a virtual machine hypervisor. Running an application in a VM is a way to create isolation from the running environment, but it requires a heavy layer of services to support the same low hardware footprint isolation provided by containers.
	-	Environment Isolation: Works in a closed environment where changes made to the host OS or other applications do not affect the container. Because the libraries needed by a container are self-contained, the application can run without disruption.
	- Quick Deployment: Containers can be quickly deployed as it does not require full OS install or restart. A container only requires a restart without stopping any services on the host OS.	
	- Multiple environment deployment: In a traditional deployment scenario using a single host, any environment differences might potentially break the application. Using containers, however, the differences and incompatibilities are mitigated because the same container image is used.
	- Reusability: The same container can be reused by multiple applications without the need to set up a full OS. A database container can be used to create a set of tables for a software application, and it can be quickly destroyed and recreated without the need to run a set of housekeeping tasks. Additionally, the same database container can be used by the production environment to deploy an application.
	Often, a software application with all its dependent services (databases, messaging, filesystems) are made to run in a single container. However, container characteristics and agility requirements might make this approach challenging or ill-advised. In these instances, a multi-container deployment may be more suitable. Additionally, be aware that some application actions may not be suited for a containerized environment. For example, applications accessing low-level hardware information, such as memory, file-systems and devices may fail due to container constraints.
	
	Finally, containers boost the microservices development approach because they provide a lightweight and reliable environment to create and run services that can be deployed to a production or development environment without the complexity of a multiple machine environment.
	
	Docker provides tooling and a platform to manage the lifecycle of your containers:

	Develop your application and its supporting components using containers.
	The container becomes the unit for distributing and testing your application.
	When you’re ready, deploy your application into your production environment, as a container or an orchestrated service. This works the same whether your production environment is a local data center, a cloud provider, or a hybrid of the two.
	
  * Overview of the Docker Architecture
	- Docker uses client-server architecture. Client will run the command line tool known as 'docker' for executing the commands on the server.
	- Containers are segregated user space env for running our applications. With Docker, you can manage your infrastructure in the same ways you manage your applications. By taking advantage of Docker’s methodologies for shipping, testing, and deploying code quickly, you can significantly reduce the delay between writing code and running it in production.
	- Images are templates or definition for containers.
	- Images are stored in registeries. These registeries can be private and local on our local machines or public.
	 
	
 # To be contd