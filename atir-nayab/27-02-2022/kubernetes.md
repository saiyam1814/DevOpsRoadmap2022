## master node components

a master node runs the following control plane components:

- api server
- scheduler
- controller manager
- data store

in addition, the master node runs:

- container runtime
- node agent
- proxy

### master node components: api server

all the administrative tasks are coordinated by the kub-apiserver, a central control plane component running on the master node. the api server intercepts restful calls from users, operators and external agents, then validates and processes them. during processing the api server reads the kubernetes cluster's current state from the etcd data store, and after a call's execution, the resulting state of the kubernetes cluster is saved in the distributed key-value data store for persistence. the api server is the ony master plane component to talk to the etcd data store, both to read from and to save kubernetes cluster state information.

### master node components: scheduler

the role of the kube-scheduler is to assign new workload objects, such as pods, to nodes. during the scheduling process, decisions are made based on current kubernetes cluster state and new object's requirements. the scheduler obtains from the etcd data store, via the api server, resource usage data for each worker node in the cluster. the scheduler also receives from the api server the new object's requirements which are part of its configuration data.

the scheduler is highly configurable and customizable through scheduling policies, plugins, and profiles. additional custom schedulers are also supported, then the object's configuration data should include the name of the custom scheduler expected to make the scheduling descision for that particular object; if no such data is included, the default scheduler is selected instead.

### master node components: controller managers

the controller managers are control plane components on the master node running controller to regulate the state of the kubernetes cluster. controllers are watch-loops continuously running and comparing the cluster's desired state (provided by object's configuration data) with its current state (obtained from etcd data store via the api server). in case of a mismatch corrective action is taken in the cluster until its current state matches the desired state.

the kube-controller-manager runs controllers responsible to act when nodes become unavailable, to ensure pod counts are as expected, to create endpoints, service accounts, and api access tokens.

the cloud-controller-manager runs controller responsible to interact with the underlying infrastructure of a cloud provider by a cloud service, and to manage load balancing and routing.

### master node components: data store

etcd is a strongly consitent, distributed key-value data-store used to persist a kubernetes clusters state. new data is written to the data store only by appending to it, data is never replaced in the data store. obsolete data compacted periodically to minimize the size of the data store.

out of all the control plane components, only the api server is able to communicate with the etcd data store.

etcd's cli management tool - etcdctl. provides backup snapshot and restore capabilities which come in handy especially for a single etcd kubernetes cluster.

some kubernetes cluster bootstrapping tools, such as kubeadm, by default, provision stacked etcd master nodes, where the data store runs alongside and shares resources with other control plane components on the same master node.

![stacked etcd topoligy](https://courses.edx.org/assets/courseware/v1/7aed6c75efbc84567cbf1b7d61ec150a/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/kubeadm-ha-topology-stacked-etcd.svg)

**stacked stcd topology**  
for data store isolation from the control plane components, the bootstrapping process can be cofigured for an external etcd topology, where the data store is provisioned on a dedicated separate host, thus reducing teh chances of an etcd failure.

![external etcd topology](https://courses.edx.org/assets/courseware/v1/9abe43636a43eecfed5dfd3478ddae45/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/kubeadm-ha-topology-external-etcd.svg)

etcd is written in the go programming language. in kubernetes, besides storing the cluster state, etcd is also used to store configuration details such as subnets, configMaps, secrets etc.

## worker node overview

a worker node provides a running environment for client applications. through containerized microservices, these applications are encapsulated in pods, controlled by the cluster control plane agents running on the master node. pods are scheduled on worker nodes, where they find required compute, memory and storage resources to run, and networking to talk to each other and the outside world. a pod is the smallest scheduling unit in kubernetes. it is a logical collection of one or more containers scheduled together, and the collection can be started, stopped, or rescheduled as a single unit of work.

also, in a multi-worker kubernetes cluster, the network traffic client users and the containerized applications deployed in pods is handled directly by the worker nodes, and is not routed through the master node.

## worker ndoe components

a worker node has the following components:

- contaienr runtime
- node agent - kubelet
- proxy - kube-proxy
- addons for dns, dashboard user interface, cluster-level monitoring and logging.

### worker node components: container runtime

although kubernetes is described as a container orchestration engine, it does not have the capability to directly handle containers. in order to manage a containers lifecycle, kubernetes requires a container runtime on the node where a pod and its container are to be scheduled. kubernetes support many container runtime.

- docker - although a container platform which uses containerd as a container runtime, it is the most popular container runtime used with kubernetes
- cri-o a lightweight container runtime for kubernetes, it also supports docker image registries
- containerd - a simple and portable container runtime providing robustness
- frakti - a hypervisor - based container runtime for kubernetes.

### worker node components: node agent-kubelet

the kubelet is an agent running on each node and communicates with the control plane components from the master node. it reaceives pod definitions, primarily from the api server, and interacts with the container runtime on the node to run containers associated with the pod. it also monitors the health and resources of pods running containers.

the kubelet connects to container runtime though a plugin interface - the container runtime interface (cri). the cri consists of protocol buffers, grpc api, libraries and additional specifications and tools that are currently under development. in order to connect to interchangeable container runtimes, kubelet uses a shim application which provides a clear abstraction layer between kubelet and the container runtime.

![container runtime environment](https://courses.edx.org/assets/courseware/v1/ab209f7c32ceb17ed43dcf6b66056cea/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/CRI.png)

the kubelet acting as grpc client connects to the cri shim acting as grpc server to perform container and image operations. the cri implements two services: imageService and RunimeService. the ImageService is responsible for all the image-related operations, while the RuntimeService is responsible for all the pod and container-related operations.

container runtimes used to be hard-coded into kubelet, but since the cri was introduced, kubernetes has become more flexible to use different container runtimes without the need to recompile. any container runtime that implements the cri can be used by kubernetes to manage pods, containers, and container images.

### worker node components: kubelet - cri shims

shims are cri implementations, or interfaces, specific to each container runtime supported by kubernetes. below we present some examples of cri shims.

- dockershim - with dockershim, contaienrs are created using docker installed on the worker nodes. internally docker uses containerd to create and manage containers.
  ![dockershim](https://courses.edx.org/assets/courseware/v1/aa11f8d767939eb27a989d12423e5ae6/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/dockershim.png)

- cri-containerd - with cri-containerd, we can directly use containerd to create and manage containers:
  ![cri-containerd](https://courses.edx.org/assets/courseware/v1/4d76490e58857edcf3a9c335f46fdcb9/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/cri-containerd.png)

- cri-o enables the use of any open container intiaitve (oci) compatible runtime with keubernetes. at the time this course was created, cri-o supported runc and clear contaienrs as container runtimes. however, in principle, any oci-complaint runtime can be plugged-in.

- frakti enables cri implementation through hardware virtualization, aimed to achieve a higher level of security and isolation than the traditional os level contaienrs based on cgroups and namespaces. the frakti cri shim is aimed at enabling kubelet to interact with kata containers.
  ![farkti](https://courses.edx.org/assets/courseware/v1/9e71fcc2a9b585d0a3513db3ed2a1a45/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/frakti.png)

### worker node components: proxy-kube-proxy

the kube-proxy is the network agent which runs on each node responsible for dynamic updates and maintenance of all networking rules on the node. it abstracts the details of pods networking and forwards connection requests to pods.

the kube-proxy is responsible for tcp, udp and sctp stream forwarding or round robin forwarding across a set of pod backends, and it implements forwarding rules defined by users through service api objects.

### worker node components: addons

addons are cluster features and functionality not yet available in kubernetes, therefore implemented through 3rd-party pds and services

- dns -cluster dns is a dns server required to assign dns records to kubernetes objects and resources
- dashboard - a general purposed web-based user interface for cluster management
- monitoring - collects cluster-level container metrics and saves them to a central data store.
- loggin - collects cluster level container logs and saves them to a central log store for analysis.

## networking challenges

decoupled microservices based applications rely heavily on networking in order to mimic the tight-coupling once available in the monolithic era. networking, in general, is not the easiest to understand and implement. kubernetes is no exception

- as a containerized microservices orchestrator it needs to address a few distinct networking challenges
- container-to-container communication inside pods
- pod to service communication within the same namespaces
- external-to-service communication for clients to access applications in a cluster.

all the networking challenges must be addressed before deploying a kubernetes cluster.

## container-to-container communication inside pods

making use of the underlying host operating system's kernel virtaulization features, a container runtime creates an isolated network space for each container it starts. on linux this isolated network space is referred to as a network namespace. a network namespace can be shared across containers, or with the host operating system.

when a pod is started, a special pause container is initialized by the container runtime for the sole purpose to create a network namespace for the pod. all additional containers, created through user requests, running inside the pod will share the pause container's network namespace so that they can all talk to each other via localhost

## pod-to-pod communication across nodes

in a kubernetes cluster pods are scheduled on nodes in a nearly unpredictable fashion. regardless of their host node, pods are expected to be able to communicate with all other pods in a cluster, all this without the implmentation of network address translation (nat), this is a fundamental requirement of any networking implentation in kubernetes

![cni core plugins](https://courses.edx.org/assets/courseware/v1/e7f4f4c4b79ffd11fb57659d8558943b/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/Container_Network_Interface_CNI.png)

## pod-to-external world communication

a successfully deployed containerized application running in pods inside a kubernetes cluster may require accessibility from the outside workd. kubernetes enables external accessibility through services, complex encapsulations of network routing rule definitions stored in iptables on cluster nodes and implented by kube-proxy agents. by exposing services to the external world with the aid of kube-proxy, applications become accessible from outside the cluster over a virtual ip address.

# chapter 5: installing kubernetes

## kubernetes configuration

- all in one single node installation
- single-master and multi-worker istallation
- single master with single-node etcd, and multi-worker installation
- multi-master and multi-worker installation
- multi-master with multi-node etcd, and multi-worker installation

## infrastructure for kubernetes installation

- should we set up kubernetes on bare metal, public cloud, private or hybrid cloud?
- which underlying os should we use? should we choose a linux distribution such as RHEL, cnetos, coreos or windows?
- which networking solution should we use?

## localhost installation

- minikub - single-node local kubernetes cluster, recommended for a learning environment deployed on a single host
- kind - multi-node kubernetes cluster deployed in docker containers acting as kubernetes nodes, recommended for learning environment
- docker desktop - including a local kubernetes cluster for docker users.
- microK8s - local and cloud kubernetes cluster, from canonical
- k3s - lightweight kubernetes cluster for local, cloud, edge iot deployments, from rancher

## on-premise installation

- on-premise vms - kubernetes can be installed on vms created via vagrant, vmware vSphere, kvm, or another configuration management tool in conjunction with a hypervisor software. there are different tools to automate the installation such as ansible and kudeadm
- on-premise bare metal - kubernetes can be installed on on-premise bare metal, on top of different operating systems, like rhel, coreos, centos fedora ubuntu etc

## cloud installation

- hosted solutions
- turnkey cloud solutions
- trunkey on-premise solutions

## kubernetes installation tools/resources

- kubeadm
- kubespray
- kops
- kube-aws

# chapter 6: minikube - a local kubernetes

it is the easiest and most popular method to run all in one kubernetes cluster in a virtual machine (vm) locally on our workstations. minikube is the tool responsible for the installation of kubernetes components, cluster boostrapping, and cluster tear-down when no longer needed.

## requirements for running minikube

- VT-x/AMD-v virtualization must be enabled on the local workstation in bios, and/or verify if wirtualization is supported by your workstation's os.
- kubectl is a binary used to access and manage any kubernetes cluster. it is installed through minikube and accessed through minikube command, or it can be installed separately.
- type-2 hypervisor or container runtime
- internet connection at least on first minikube run - to download packages, dependencies, updates and pull images needed to initialize the minikube kubernetes cluster.
