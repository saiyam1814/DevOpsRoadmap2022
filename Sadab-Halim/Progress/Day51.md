# Kubernetes

## Kubernetes Pod
- Pod is the atomic unit on the kubernetes cluster. 
- First question that is moving around my head is that why do we need pods if we have containers?
- This is basically done to solve port allocation issue. With containers we can use a specific host port only once. Since a Pod works like a standalone virtual host having a unique IP the same port can be used multiple times without encountering conflicts.
- The pod manages all your container inside it. In other words, it is an abstraction over your containers. 
- Every pod have it's own IP address
- Containers within a pod communicate via localhost.
  
### Creating a Pod 
- There are two ways for creation of pods in kubernetes
1. Kubectl way
```
kubectl run first-pod --image=nginx
```
- You should see a output with message first-pod/Pod created. 
- This method is not suitable if you want to create large number of pods of the same kind. You need to run this command many times. 
- With this method we cannot version control our pods
- To delete the pod we will use `kubectl delete pod <pod-name>`

2. Using kubernetes manifests
- In this case we will write a YAML based declarative file and then we will tell kubernetes to launch the pod for us. 
```
apiVersion: v1
kind: Pod
metadata:
   name: first-pod
   namespace: default
   labels:
     demo: pod
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
      - containerPort: 80
```

- In the above YAML syntax apiVersion means the version of the kubernetes API you're using to create this object
- Kind is used for mentioning what type of kubernetes resource we are creating. In this case it is a Pod. 
- Metadata is data that helps uniquely identify the object, including a name string, UID, and optional namespace. 
- spec is the field where you define your object in more detail. Like in case of pod you will define the container that will run on the pod. The image of the container and ports. 

#### Kubectl commands for pods
- kubectl get <resource> - list resources
- kubectl describe <resource> - show detailed information about a resource
- kubectl logs <resource> - print the logs from a container in a pod
- kubectl exec <resource> - execute a command on a container in a pod

- We can use the commands for all the resources that we will create in the kubernetes cluster. It can be depolyment, services, replicasets, pods etc. 

