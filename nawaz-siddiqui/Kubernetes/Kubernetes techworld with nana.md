# Kubernetes

Kubernetes is an open source container orchestration tool, Developed by Google. In other words, Kubernetes will help you manage containerized applications in different deployment environment like physical, virtual and hybrid.

### Features Kubernetes offer:

High availability or no downtime.

Scalability or high performance

Disaster Recovery

### Pod

Smallest unit of K8s.

Abstraction over container.

Usually one application per pod.

Each pod get its own IP address, offering a virtual network inside Kubernetes.

New IP address on recreation.

### Service

Permanent IP address

The lifecycle of Pod and Service are not connected.

It is also a load balancer.

### Ingress

Kubernetes Ingress is **an API object that provides routing rules to manage external users' access to the services in a Kubernetes cluster**, typically via HTTPS/HTTP. With Ingress, you can easily set up rules for routing traffic without creating a bunch of Load Balancers or exposing each service on the node.

### ConfigMap:

Contains configuration data like URLs of databases or some other services you use

It's an external configuration of your application

### Secret:

Its just like a config map but used to store secret data.

Stores the content in base64 encoded format

The built-in security mechanism is not enabled by default.

### Volumes

If a pod does or restarts the data is not accessible anymore because Kubernetes doesn't manage data persistence. The function of volume is to attach physical storage on the hard drive to your pod, it can be local or remote/cloud storage.

### Deployment

Once you create an application in Kubernetes, its time to make it accessible with no downtime. Deployment helps with this issue:

Its a blueprint for your app.

You can create deployments whenever in need to scale up and vice versa

Its an abstraction of Pods

Deployment is for Stateless apps.

### StatefulSet

For stateful apps(which contains data)

StatefulSet is for Stateful apps or databases

Avoids data inconsistency

Its a bit hard to deploy database applications using StatefulSet. A common practice is to host database applications outside of the Kubernetes cluster.

## Kubernetes Architecture

Each node has multiple Pods on it.

Two types of nodes, worker and master node. Worker node does the actual work.

The Kubernetes cluster is made up of multiple nodes.

A node consists of Kubelet, Kubeproxy, and Container Runtime to function properly.

**Container Runtime**

A container runtime, also known as container engine, is **a software component that can run containers on a host operating system**.

**Kubelet**

The kubelet is the primary "node agent" that runs on each node. It can **register the node with the apiserver using** one of: the hostname; a flag to override the hostname; or specific logic for a cloud provider. Kubelet interacts with both - the container and node. It helps start the pod with a container inside.

**Kubeproxy**

kube-proxy is **a network proxy that runs on each node in your cluster**, implementing part of the Kubernetes Service concept. kube-proxy maintains network rules on nodes. These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.

**Worker Node**

A worker node consist of container runtime, kubelet and kubeproxy.

**Master Node**

Managing processes are done by Master Node. Below are some of the component which makes full-fledged master nodes run master processes.

### Master Processes

**API Server**

cluster gateway

act as a gatekeeper for authentication

**Scheduler**

Intellegently schedules and distributed piods in different worker nodes.

**Controller manager**

It detects the cluster change states and tries to recover the state ASAP.

After detecting the cluster change, it request the scheduler for the recovery state.

**etcd**

It is a key-value pair.

It is a cluster brain.

Cluster changes get stored in the key-value store.

Application data is not stored in etcd.

**Process of adding new nodes**

Get a bare server

Install all the worker/master node processes.

Add it to the cluster.

### Minikube

Creates a virtual box in your laptop

Node runs in that virtual box

It is a 1 node K8s cluster.

**Kubectl**

It's a command-line tool for K8s cluster

Interacts with the API server

**Layers of Abstraction**

Deployment manages a replica set

Replicaset manages all replica of a pod

A pod is an abstraction of a container

Everything below the deployment is managed by Kubernetes

### K8s configuration files

Each configuration files has three parts:

Metadata means what you want to create.

Specification means specific attributes to the kind.

Status is automatically generated and added by Kubernetes via the help of etcd.

### Connecting deployment to pods

Pods get label through the template blueprint.

The label is matched by selector.