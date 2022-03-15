# the kubernetes api

the core fo kubernetes control plane is the api server. the api server exposes an http api that lets end users, different parts of the cluster, and external components communicate with one another.

the kubernetes api lets you query and manipulate the state of api objects in kubernetes

# understanding kubernetes objects

kubernetes objects are persistent entitites in the kubernetes sysntem. kubernetes uses these entities to repreent the state of your cluster. specifically, they can describe

- what containerized applications are running (and on which nodes)
- the resources available to those applications
- the policies around how those applications behave, such as restart policies, upgrades and fault tolerance

a kubernetes object is a record of intent --once you create the object, the kubernets system will constantly work to ensure that object exists. by creating an object, you're effectively telling the kubernetes system what you want your cluster's workload to look like; this is your cluster's desired state.

to work with kubernetes objects --whether to create, modify or delete them you'll need to use the kubernetes api. when you use the kubectl command line interface, for example, the cli makes the necessary kubernetes api calls for you. you can also use the kubernetes api directly in your own programs using one of the client libraries.

## describing a kuberntes object

when you create an object in kubernets, you must provide the object spec that describes its desired state, as well as some basic information about the object (such as name) when you use the kuberntes api to create the object (weither directly or via kubectl), that api request must include that information as json in the requrest body. most often, you provide the information to kubectl in a yml file. kubectl converts the information to json when making the pai request.

# Nodes

kubernetes runs your workload by placing containers into pods to run on nodes. a node may be a virtual or physical machine, depending on the cluster. each node is managed by the control plane and containers the services necessary to run pods.

the components on a node incluse the kubelet, a container runtime and the kube-proxy.

## management

there are two main ways to have nodes added to the api server

- the kubelet on a node self-registers to the control plane
- you manually add a node object

## capacity and allocatable

describes the resources available on the node: cpu, memory and the maximum number of pods taht can be scheduled onto the node.

## info

describes general information about the node, such as kernel version, kubernetes version,container runtime details, and which operating system the node uses. the kubelet gathers this information from the node and publishes it into the kubernets api.

## heartbeats

heartbeats sent by kubernetes nodes, help your cluster determine the availability of each node, and to take action when failures are detected

for nodes there are two forms of hearbeats

- updated to the status of a node
- lease objects within the kube-node-lease namespace. each node has an associated lease object

## node controller

the node controller is a kubernetes control plane component that manages various aspects of nodes.

the node controller has multiple roles in a node's life. the first isassigning a cidr block to the node when it is registered

the second is keeping the node controller internal list of nodes up to date with the cloud providers list of available machines. when running a cloud environment and whenever a node is unhealthly, the node controller asks the cloud provider if the vm for that node is still available. if not, the node controller deletes the node from its list of nodes.

the third is monitoring the nodes health. the node controller is responsible for

- in the case that a node becomes unreachable, updating the nodeReady confition of within the node's status
- if a node remains unreachable: triggering api-initiated eviction for all of the pods on the unreachable node. by default, the node controller wais 5 minutes between marking the nodes as conditionUnkown and submutting the first evicition request

## rate limits on eviction

in most cases, the node controller limits the eviciton rate to --node-eviction-rate (default 0.1) per second, meaning it won't evict pods from more than 1 node per 10 seconds.

## resouce capacity tracking

node objects track information about the node's resource capacity: for example, the amount of memory available and the number of cpus. nodes that self register report their capacity during registration. it you manually add a node, then you need to set the node's capacity information when you add it.

the kubernetes scheduler ensures that there are enough resources for all the pods on a node.

## graceful node shutdown

the kubenelet attempts to detect node system shutdown and terminates pods running on the node.

kubelet ensures that pods follow the normal pod termination process during the node shutdown.

# control plane-node communation

## node to control plane

kubernetes has a hub-and-spoke api pattern. all api usage from nodes (or the pods they run) terminates at the apiserver. none of the other control plane components are designed to expose remote services. the apiserver is configured to listen for remote connections on a secure https port with one or more forms of client authentication enabled.

## control plane to node

there are two primary communication paths from the control plane (apiserver) to the nodes. the first is from th eapiserver to the kubelet process which runs on each node in the cluster. the second is from the apiserver to any node, pod, or service through the apiserver's proxy functionality

## apiserver to kubelet

the connections from the apiserver to the kubelet are used for

- fetching logs for pods.
- attaching (through kubectl) to running pods
- providing the kubelet's port forwarding functionality

these connections terminate at the kubelet's https endpoint. by default the apiserver does not verify kubelet's serving certificate, which makes the connection subject to man-in-the-middle attacks and unsafe to run over untrusted and/or public networks

to verify this connection, use the --kubelet-certificate-authority flag to provide the apiserver with a root certificate bundle to use to verify the kubelet's serving certificate.

## apiserver to nodes, pods, and services

the connections from the apiserver to a node, pod or service default to plain http connections and are therefore neither authenticated nor encrypted. they can be run over a secure http connection by prefixing `https:` to the node,pod or service name in the api url

## ssh tunnels

kubernetes supports ssh tunnels to protext the control plane to nodes communication paths. in this configuration, the apiserver initiates an ssh tunnel to each node in the cluster and passes all traffic destined for a kubelet, node, pod or service through the tunnel.

# controllers

a control loop is a non-terminating loop that regulates teh state of a system

## controller pattern

a controller pattern tracks at least one kubernetes resource type. these objects have a spec field that represents the desired state. the controller(s) for that resource are responsible for making the current state come closer to that desired state.

## control via api server

the job controller is an example of a kubernetes built in controller. built in controllers manage state by interacting with the cluster api server.

job is a kubernetes resource that runs a pod, or perhaps serveral pods to carry out a task and then stop.

## direct control

by contrast with job, some controllers need to make changes to things outside of your cluster.

for example if you use a control loop to make sure there are enough nodes in your cluster, then that controller needs something outside the current cluster to set up new nodes when needed.

# cloud controller manager

cloud infrastructure technologies let you run kubernetes on public, private, and hybrid clouds. kubernetes believes in automated, api-driven infrastructure without tight coupling between components.

the cloud-controller-manger is a kubernetes control plane component that embeds cloud-specific control logic. the cloud controller manager lets you link your cluster into your cloud provider's api.

# container runtime interface (cri)

the cri is a plugin interface which enables the kubelet to use a wide variety of container runtimes, without having a need to recompile the cluster components.

you need a working container runtime on each node in your cluster, so that the kubelet can launch pods and their containers.

the cri is the main protocol for the communication between the kubelet and container runtime.

the kubernetes container runtime interface cri defines the main grpc protocol for the communication between the cluster components kubelet and container runtime.

# garbage collection

garbage collection is a collective term for the various mechanisms kubernetes uses to clean up cluster resources. this allows the clean up of resources like the following.

## cascaing deletion

kubernetes checks for and deletes objects that no longer have owner references, like the pods left behind when you delete a replicaSet.

# containers

each container that you run is repeatable, the standardization form having dependencies included means that you get the same behavior wherever you run it.

container decouple applications from underlying host infrastructure. this makes deployment easier in different cloud or os environments.

## container images

a container image is a ready-to-run software package, containing everything needed to run an application; the code and any runtime it requires, application and system libraries, and default values for any essential settings.

by design a container is immutable; you cannot change the code of a container that is already running. if you have a containerized application and want to make changes, you need to build a new image that includes the change, then recreate the container to start from the updated image.

## container runtime

the container runtime is the software that is responsible for running containers.

kubernetes supports cotainer runtimes such as containerd, cri-o and any other implementation of the cri.
