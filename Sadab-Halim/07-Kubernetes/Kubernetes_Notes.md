## Kubernetes Namespaces

- Namespaces as a virtual cluster inside your kubernetes cluster
- In a production environment it is very easy for a team to accidently overwrite or disrupt another team without even realizing it.
- The ideal approach is to create multiple namespaces and segment them into sizeable chunks.
- Think of a school, In a big school we have different classrooms and different departments for different operations. Same here we can create different namespaces for different purposes. Say if we want to split our project into different teams then we can do that.
- Namespaces are a way to divide cluster resources between multiple users (via resource quota)
- When you start kubernetes it starts with 4 default namespaces
    - Default
    - kube-public
    - kube-system
    - kube-node-lease
- to view the namespaces use the command `kubectl get namespace`

## Creating a Namespace

- kubectl create namespace demo will create you a namespace with the name demo.
- You can also write a YAML file for creating a namespace.
- There are two ways to explicitly tell the kubernetes which namespace you want for your resources.
  - Using the --namespace flag
  - mentioning it in the yaml configuration
- We usually don't prefer writing in the yaml declaration because everytime the pod will be created in the same namespace. 


```
apiVersion: v1
kind: Namespace
metadata:
    name: demo
    labels:
        name: demo
```
- After creating this YAML file we can run `kubectl apply -f demo.yaml`
- to list all the namespaces run the command `kubectl get namespace`

```
apiVersion: v1
kind: Pod
metadata: mypod
labels:
    name: mypod
spec:
    containers:
        - name: mypod
          image: nginx
```
`kubectl apply -f mypod.yaml` this will create the resources in the default namespace. 
- To create your file in the demo namespace you need to apply it with the namespace flag `kubectl apply -f mypod.yaml --namespace=demo`




## Needs for namespaces

- When you are working on a kubernetes cluster with default namespace and you and your colleague have deployed a resource with same name then one will overwrite the other.
- When we need to split our team into sub teams for a project then we can create a different namespace for each of them.
  - Team Alpha
  - Team Beta
  - Team Gamma



## kubens
- When you run the kubens command, you should see all the namespaces with the active namespace highlighted. To switch you active namespace to your demo namespace simply run `kubens demo`

## Kubernetes Architecture

The Kubernetes Architecture is divided into two components
  - Master Node (Also called control plane)
  - Worker Nodes 
  
### Control Plane 
- All the operations that run on the worker nodes are governed by control plane.
- A user interacts with control plane via three routes :
  - CLI
  - Web UI
  - API

- Components of Control Plane 
1.  *API Server* 
    - The API server intercepts RESTful calls from users, operators and external agents then validates and process them
    - The API Server is the only master plane component to talk to the etcd data store, both to read from and to save Kubernetes cluster state information - acting as a middle interface for any other control plane agent inquiring about the cluster's state.

2. *Controller Manager* 
    - It keeps track of what's happening in the cluster like if there is something that needs to be repaired or something. It continuously compares the desired state (provided in configuration file) with the current state (obtained from ETCD data store via the API server)
3. *Scheduler* 
    - The role of scheduler is to assign new workload objects such as pods to nodes. It receives the requirementes from the API server about what needs to be configured.The scheduler makes decision based on the current state of the kubernetes cluster. It obtains the cluster information from ETCD datastore via the API server. 
5. *ETCD* 
    - It holds the current status of the kubernetes cluster at any point of time.

*NOTE*: In production environment we have at least two master nodes running because it is the vital part of the kubernetes architecture and if this stops then the whole application will go down. So, at any point of time we should have a backup of our master node.

### Worker Node 
- It is that part of the cluster on which your containerized application is running. 

- A worker node has following components :
  - container runtime
  - kubelet
  - proxy
  - Addons for DNS, Dashboard user interface, cluster level monitoring and logging

1. *Container Runtime*
    - The smallest unit that we can deploy in kubernetes is a pod and kubernetes don't interact directly with containers. Therefore we need a container runtime for our containerized application to run. Kubernetes support runtime like containerd, cri-o, rkt etc. 
2. *Kubelet*
    - Kubelet acts like an agent and is used for communication between nodes and the API server.
    - It also interacts with container runtime on the node to run the containerized microservices associated with the pod.
3. *Kube-proxy*
    - It is the network agent which runs on each node responsible for dynamic updates and maintenance of all networking rules on the node.

4. Addons
    - `DNS` - DNS server to assign DNS records to kubernetes objects and resources
    - `Dashboard` - A web based interface for cluster management
    - `Monitoring` - collects cluster level container metrics and saves them to a central data store
    - `Logging` - collects cluster-level container logs and saves them to a central log store for analysis.


