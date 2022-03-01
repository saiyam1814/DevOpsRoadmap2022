# chapter 7 - accessing minikube

## introduction

we can use a variety of external clients or custom scripts to access our cluster for administration purpose. we will explore kubectl as a cli tool to access the minikube kubernetes cluster, the kubernetes dashboard as a web-based user interface to interact with the cluster, and the curl command with proper credentials to access the cluster via apis

## accessing minikube

- cli tools and scripts
- web-based user interface (web-ui) from a web browser
- apis from cli or programmatically

these methods are applicable to all kubernetes clusters.

### accessing minikube: command line interface (cli)

kubectl is the kubernetes command line interface (cli) client to manage cluster resources and applications. it is very flexible and easy to integrate with other systems, therefore it can be used standalone, or part of scripts and automation tools. once all required credentials and cluster access points have been configured for kubectl, it can be used remotely from anywhere to access a cluster.

### accessing minikube: web-based user interface (web ui)

the kubernetes dashboard provides a web-based user interface (web ui) to interact with a kubernetes cluster to manage resources and containerized applications. in one of the later chapters, we will be using it to deploy a containerized application.

![kubernetes dashboard user interface](https://courses.edx.org/assets/courseware/v1/b03d0b289dc978667ed80413ec78dd4a/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/ui-dashboard.png)

### accessing minikube: apis

the main component of the kubernetes control plane is the api server, responsible for exposing the kubernetes apis. the apis allow operators and users to directly interact with the cluster. using both cli tools and the dashboard ui, we can access teh api server running on the master node to perform various operations to modify cluster's state. the api server is accessible through its endpoints by agents and users possessing the required credentials.

![http api directory tree kubernetes](https://courses.edx.org/assets/courseware/v1/7ebcb514203a3af89dc1599625779c1f/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/api-server-space_.jpg)

**HTTP api directory tree of kubernetes can be divided into three independent group types**

- core group (/api/v1) - this group includes objects such as pods, services, nodes, namespace, configMaps, secrets etc.
- named group - this group includes objects in /apis/$NAME/$VERSION format. these different api versions imply different levels of stability and support.
- system-wide - this group consists of system-wide api endpoints, like /healthz, /logs, /metrics, /ui etc.

## kubectl

kubectl allows us to manage local kubernetes clusters like the minikube cluster, or remote clusters deployed in the cloud. it is generally installed before installing and starting minikube, but it can also be installed after the cluster boostrapping step. once installed, kubectl receives its configuration automatically for minikube kubernetes cluster access. however in different kubernetes cluster steups, we may need to manually configure the cluster access points and certificates required by kubectl to securely access the cluster.

## kubectl configuration file

to access the kubernetes cluster, the kubectl client needs the master node endpoint and appropriate credentials to be to securely interact with the api server running on the master node. while starting minikube, the startup process creates, by default, a configuration file, config, inside the `.kube` directory, which resides in the user's home directory. the configuration file has all the connection details required by kubectl. by default the kubectl binary parses this file to fine the master node's connection endpoint, along wiht endpoint.

although for the kubernetes cluster installed by minikube the ~/.kube/config file gets created automatically, this is not the case for kubernetes clusters installed by other tools. in other cases, the config file has to be created manually and sometimes re-configured to suit various networking and client/server setups.

## kubernetes dashboard with kubectl proxy

issuing the kubectl proxy command kubectl authenticates with the api server on the master node and makes the dashboard available on a slightly different url than the one earlier, this time through the default proxy port 8001.

it locks the terminal for as long as the proxy is running, unless we run it in the background (with kubectl proxy &)

## apis with kube's proxy

when kubectl proxy is running, we can send requests to the api over the localhost on the default proxy port 8001 (from another terminal, since the proxy locks the first terminal when running in foregroung)

## api with authentication

when not using the kubectl proxy, we need to authenticate to the api server when sending api requests, we can authenticate by providing a bearer token when issuing a curl, or by providing a set of keys and ceritificates.
