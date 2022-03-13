# CNCF certifcations

- CKAD - certified kubernetes application developer for helping developers build, deploy and manage applications on kubernetes

- CKA - certified kubernetes administrator for administrators, for administrator responsible for managing kubernetes

- CKS - certified kubernetes security specialist, focuses on securing container-based applications on kubernetes during build, deployment and runtime.

# **Revision**

# overview

## what is kubernetes?

kubernetes is a portable, extensible, open-source platform for managing containerized workloads and services, that facilitates both declarative configuration and automation. It has a large, rapidly growing ecosystem.

![going back in time](https://d33wubrfki0l68.cloudfront.net/26a177ede4d7b032362289c6fccd448fc4a91174/eb693/images/docs/container_evolution.svg)

kubernetes provides you with:

- service discovery and load balancing - kubernetes can expose a container using the dns name or using their own ip address. if traffic to a container is high, kubernetes is able to load balance and distribute the network traffic so that the deployment is stable.
- storage orchestration - kubernetes allows you to automatically mount a storage system of your choice, such as local storages, public cloud providers, and more.
- automated rollouts and rollbacks - you can describe the desired state for your deployed containers using kubernetes, and it can change the actual state to the desired state at a controlled rate. for example you can automate kubernetes to create new containers for your deployemnt, remove existing containers and adopt all their resources to the new container.
- automatic bin packing - you provide kubernetes with a cluster of nodes that it can use to run containerized tasks. you tell kubernetes how much cpu and memory (ram) each container needs. kubernetes can fit containers onto your nodes to make the best use of your resources.
- self-healing - kubernestes restarts containers that fail, replaces containers, kills containers that don't respond to your user-defined health check, and doesn't advertise them to clients until they are ready to serve.
- secret and configuration management - kubernetes lets you store and manage sensitive information, such as passwords, oauth tokens and ssh keys. you can deploy and update secrests and application configuration without rebuilding your container images,a nd without exposing secrets in your stack configuration.

## kubernetes components

a kubernetes cluster consists of a set of worker machines called nodes, that run containerized applications. every cluster has at least one worker node.

the worker nodes host the pods that are the components of the application workload. the control plane manages the worker nodes and the pods in the cluster. in production environments, the control plane usually runs accross multiple computers and a cluster usually runs multiple nodes, providing fault-tolerance and high availability.

![the components of a kubernetes cluster](https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg)

### control plane components

the control plane's components make global decisions about the cluster, as well as detecting and responding to cluster events (for example scheduling), as well as detecting and responsing to cluster events

control plane components can be run on any machine in the cluster. however, for simplicity, set up scripts typically start all control plane components on the same machine, and do not run user containers on this machine.

### kube-apiserver

the api server is a component of the kubernetes control plane that exposes the kubernetes api. the api server is the front end for the kubernetes control plane.

### kubernetes api

the kubernestes api lets you query and manipulate the state of objects in kubernetes. the core of kubernetes control plane is the api server and the http api that it exposes. users hte different parts of your cluster, and external components all communicate with one another through the api server.

the main implementation of kubernetes api server is kube-apiserver. kube-apiserver is designed to scale horizontally - that is, it scales by deploying more instances. you can run serveral instances of kube apiserver and balance traffic between those instances.

### etcd

consistent and highly available key value store used as kubernetes backing store for all cluster data.

if your kubernetes cluster uses etcd as its abcking store, make sure you have a backup plan for those data.

### kube-scheduler

control plane component that watches for newly created pods with no assigned node, and selects a node for them to run on.

factors taken into account for scheduling decisions include: invidual and collective resource requriements, hardware/software/policy constraints, affinity and anti-affinity specifications,data locality, inter workload interference, and deadlines.

### kube-controller-manger

control plane component that runs controller processes.

logically, each controller is a separate process, but to reduce complexity, they are all compiled into a single binary and run in a single process.

some types of these controllers are:

- node controller: responsible for noticing and responding when nodes go down.
- job controller: watches for job objects that represent one-off tasks, then creates pods to run those tasks to competion.
- endpoints controller: populates the endpoints object
- service accoutn & token controllers: create default accounts and api access tokens for new namespaces.

### cloud-controller-manager

a kubernetes control plane component that embeds cloud specific control logic. the cloud controller manger lets you link your cluster into your cloud provider's api, and separates out the components that interact with that cloud platform from components that only interact with your cluster.

the following controllers can have cloud provider dependencies:

- node controller: for checking the cloud provider to determine if a node has been deleted in the cloud after it stops responding
- route controller: for setting up routes in the underlying cloud infrastructure
- service controller: for creating,updating and deleting cloud provider load balancers

### Node components

node components run on every node, maintaining running pods and providing the kubernetes runtime evironment.

#### kubelet

an agent that runs on each node in the cluster. it makes sure that containers are running in a pod.

the kubelet takes a set podSpecs that are provided through various mechanisms and ensures that the containers desccribed in those podSpecs are running and healthy. the kubelet doesn't manage containers which were not created by kubernets

#### kube-proxy

kube-proxy is a network proxy that runs on each node in your cluster, implementing part of the kubernetes service concept.

kube-proxy maintains network rules on nodes. these network rules allow network communication to your pods from network sessions inside or outside of your cluster.

kube-proxy uses the operating system packet filtering layer it there is one and it's available. ohterwise kube-proxy forwards the traffic itself.

### container runtime

the container runtime is the software that is responsible for running containers.

kubernetes supports container runtimes such as containerd, CRI-O and any other implementation of the kubernetes CRI

### addons

addons use kubernetes resources to implemeant cluster features. because these are providing cluster level features, namespaced resources for addons belong within the kube-system namespace.

### dns

while the other addons are not strictly requried, all kubernetes clusters should have cluster dns, as many examples rely on it.

cluster dns is a dns server, in addition to the other dns server in your environment which servers dns records for kubernetes services.

### web ui (dashboard)

dashboard is a general purpose, web-based ui for kubernetes clusters.

### contianer resource monitoring

container resource monitoring records generic time series metrics about containers in a central database and provides a ui for browsing that data.

### cluster-level logging

container resource monitoring records generic time-series metrics about containers in a central database, and provides a ui for browsing that data.
