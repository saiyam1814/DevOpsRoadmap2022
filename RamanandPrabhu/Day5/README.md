#Day5 Its Containers today
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
  
	
 # To be contd