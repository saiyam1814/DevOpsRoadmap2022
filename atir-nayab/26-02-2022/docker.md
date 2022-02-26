# Chapter 2. Container Orchestration

with container images, we confine the application code, its runtime, and all of its dependencies in a pre-defined format. and, with container runtime like runC, containerd or cri-o we can use those pre-packaged images, to create one or more contianers. all of these runtimes are good at running containers on a single host. but in practice, we would like to have a fault-tolerant and scalable solution, which can be achieved by creating a single controller/management unit, after connecting multiple nodes togther, this controller/management unit is generally referred to as a container orchestrator.

## what are contianers

containers are an application-centric method to deliver high-performing, scalable applications on any infrastructure of your choice. containers are best suited to deliver microservices by providing portable, isolated virtual environments for applications to run without interference from other running applications.

![container deployment](https://courses.edx.org/assets/courseware/v1/1256618e247da221e7c3cc4bab9af3e3/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/container_deployment.png)

microservices are lightweight applications written in various modern programming languages, with specific dependencies, libraries and environmental requirements. to ensure that an application has everything it needs to run successfully it is packaged together with its dependencies.

containers encapsulate microservices and their dependencies but do not run them directly. containers run container images.

a container image bundles the application along with its runtime, libraries and dependencies, and it represents the source of a container deployed to offer an isolate executable evironment for the application. containers can be deployed from a specific image on many platform, such as workstations, virtual machines, public cloud etc.

## what is container orchestration?

in development (dev) environment, running containers on a single host for development and testing of applications may be option. however, when migrating to quality assurance (qa) and production (prod) environments, that is no longer a viable option because the applications and services need to meet specific requirements.

- fault tolerence
- on demand scalability
- optimal resource usage
- auto discovery to automatically discover and communicate with each other
- accessibility from the outside world
- seamless updates/rollback without any downtime.

containers orchestrators are tools which group together to form clusters where containers deployment and management is automated at scale while meeting the requirements mentioned above.

## container orchestrators

with enterprises containerizing their applications and moving them to the cloud, there is a growing demand for container orchestration solutions. while there are many soclutions available

- amazon ecs
- azure container instances
- azure service fabric
- kubernetes
- marathon
- nomad
- docker swarm

## why use container orchestrators?

although we can manually maintain a couple of containers or write scripts to manage the lifecycle of dozens of containers, orchestrators make things much easier for operators especially when it comes to managing hundreds and thousands of containers running on global infrastructure.

most container orchestrators can:

- group hosts together while creating a cluster.
- schedule containers to run on hosts in the cluster based resources availability
- enable containers in a cluster to communicate with each other regardless of the host they are deployed to in the cluster.
- bind containers and storage resurces
- group sets of similar containers and bind them to load-balancing contructs to simplify access to containerized applications by creating a level of abstraction between the containers and the user
- manage and optimize resource usage.
- allow for implementation of policies to secure access to application running inside containers.

## where to deploy container orchestrators?

most container orchestrators can be deployed on the infrastructure of our choice - on bare metal, virtual machines, on-premises, on public and hybrid cloud.

# Chapter 3. Kubernetes

## what is kubernetes

_kubernetes is an open-source system for automating deployment, scaling and management of containerized applications._

kubernetes is highly inspired by the google borg system, a container and workload orchestrator for its global operations for more than a decade.

new kubernetes versions are released in 3 month cycles.

## from borg to kubernetes

_googles borg system is a cluster manager that runs hundreds of thousands of jobs, from many thousands of different applications, across a number of clusters each with up to tens of thousands of machines._

some of the features/objects of kubernetes that can be traced back to borg or to lessons learned from it are

- api servers
- pods
- ip-per-pod
- services
- labels

## kubernetes features 1

kubernetes offers a very rich set of features for container orchestration. some of its fully supported features are:

- automatic bin packing - kubernetes automatically schedules containers based on resource needs and containts, to maximize utilization without sacrificing availability.
- self-healing - kubernetes automatically replaces and reschedules containers from failed nodes. it kills and restarts containers unresponsive to health checks, based on existing rules/policy. it also prevents traffic from being routed to unresponsive containers.
- horizontal scaling - with kubernetes applications are scaled manually or automatically based on cpu or custom metrics utilization.
- service discovery and load balancing - containers receive their own IP addresses from kubernetes, while it assigns a single domain name system (dns) name to a set of containers to aid in load-balancing requests across the containers of the set.

## kubernetes features 2

- automated rollouts and rollbacks - kubernetes seamlessly rolls out and rolls back application updates and configuration changes, contantly monitoring the application's health to prevent any downtime.
- secret and configuration management - kubernetes manages sensitive data and configuration details for an application separately from the container image, in order to avoid a re-build of the respective image. secrets consist of sensitive/confidential information passed to the application without revealing the sensitive content to the stack configuration, like on github.
- storage orchestration - kubernetes automatically mounts software-defined storage (sds) solutions to containers from local storage, external cloud providers, distributed storage, or network storage systems.
- batch execution - kubernetes batch execution, long-running jobs, and replaces failed containers.

## why use kubernetes?

in addition to its fully-supported features, kubernetes is also portable and extensible. it can be deployed in many environments such as local or remote virtual machines, bare metal, or in public/private/hybrid/multi-cloud setups. it supports and it is supported by many 3rd party open source tools which enhance kubernetes capabilities and provide a feature-rich experience to its users.

kubernetes architecture is modular and pluggable. not only that it orchestrates decoupled microservices type application, but also it architecture follows decoupled microservices patterns.

## kubernetes users

with just a few years since its debut, many enterprises of various sizes run their workloads using kubernetes. it is a solution for workload management in banking, education, finance and investments, gaming, information technology, media and streaming and many others

## cloud native computing foundation (cncf)

the cloud native computing foundation (cncf) is one of the projects hosted by the linux foundation. cncf aims to accelarate the adoption of containers, microservices and cloud native applications

## cncf and kubernetes

for kubernetes, the cloud native computing foundation

- provides a neutral home for the kubernetes trademark and enforces proper usage
- provides license scanning of core and vendor code
- offers leagal guidance on patent and copyright issues
- creates open source learning
- manages a software conformance working group
- acitvely markets kubernetes
- supports ad hoc activities
- sponsors conferences and meetup events.

# Chapter 4. Kubernetes Architecture

## kubernetes architecture

at a very high level, kubernetes has the following main components:

- one or more master nodes, part of the control plane
- one or more worker nodes.

![components of kubernetes architecture](https://courses.edx.org/assets/courseware/v1/51120ad23b216a6946e3c4ebef2106bf/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/arch-1.19-components-of-kubernetes.svg)

## master node overview

the master node provides a running environment for the control plance responsible for managing the state of a kubernetes cluster, and it is the brain behind all operations inside the cluster. the control plane components are agents with very distinct roles in the cluster's management. in order to communicate with the kubernetes cluster, users send requests to control place via a command line interface (cli) tool, a web user interface (web ui) dashboard, or application programming interface (api).

it is important to keep the control plane running at all costs. losing the control plane may introduce downtime, causing service disruption to clients, with possible loss of business. to ensure the control plane fault tolerance, master node replicas can be added to cluster, configured in high-availabilituy (ha). while only on of the master nodes is dedicated to actively manage the cluster, the control plane components stay in sync across the master node replicas. this type of configuration adds resiliency to the cluster's control plane, should the active master node fail.

to persist the kubernetes cluster's state, all cluster configuration data is saved to etcd. etcd is a distributed key-value store which only holds cluster state related data, no client workload data. etcd may be configured on the master node, or its dedicated host to help reduce the chances of data store loss by decoupling it from the other control plane agents.

## master node components

a master node runs the following control plane components:

- api server
- scheduler
- control managers
- data store

in addition, the master node runs

- container runtime
- node agent
- proxy.
