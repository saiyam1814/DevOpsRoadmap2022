# chapter 8. kubernetes building blocks

## kubernetes building blocks

### kubernetes object model

kubernetes has a very rich object model, representing different persistent entities in the kubernetes cluster. those entities describe

- what containerized applications we are running
- the nodes where the containerized applications are deployed
- application resource consumption
- policies attached to applications, like restart/upgrade policies, fault tolerence etc.

with each object, we declare our intent, or the desired state of the object, in the spec section. the kubernetes system manages the status section for objects, where it records the actual state of the object. at any given point in time, the kubernetes control plane tries to match the object's actual state to the object's desired state.

examples of kubernetes objects are pods, replicaSets, deployments, namespace etc.

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.15.11
        ports:
        - containerPort: 80
```

the apiVersion field is the first required field, and its specifies the api endpoint on the api server which we want to connect to; it must match an existing version for the object type defined. the second required field is kind, specifying the object type - in our case it is deployment, but it can pod, replicaset, namespace, service etc. the third required field metadata, holds the object's basic information, such as name, label, namespace etc. our example shows two spec fields (spec and spec.template.spec). the fourth required field spec marks the begining of the block defining the desired state of the deployment object. the pods are created using the pod template defined in spec.template. a nested object, such as the pod being part of a deployment retains its metadata and spec and loses the apiVersion and kind - both being replaced by template. in spec.template.spec, we define the desired state of the pod. our pod creates a single container running the nginx:1.15.11 image from docker hub.

once the deployment object is created, the kubernetes system attaches the status field to the object and populate it will all necessary status fields.

### pods

a pos is the smallest and simplest kubernetes object. it is the unit of deployment in kubernetes, which represents a single instance of the application. a pod is a logical collection of one or more containers, which

- are scheduled together on the same host with the pod
- share the same network namespace, meaning that they share a single IP address originally assigned to the pod.
- have access to mount the same external storage (volumes)

![single and multi container pods](https://courses.edx.org/assets/courseware/v1/ccc5ba54a8a06ac2a87fe447bb53dcf1/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/pods-1-2-3-4.svg)

pods are emphemeral in nature, and they do not have the capability to self-heal themselves. that is the reason they are used with controllers which handle pods replication fault tolerence, self healing etc

below is an example of a pod object's configuration manifest in yaml format.

```
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
  labels:
    app: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.15.11
    ports:
    - containerPort: 80
```

the apiversion field must specify v1 for the pod object definition. the second required field is kind specifying the pod object type. the third required field metadata, holds the object's name and label. the fourth required field spec marks the begining of the block defining the desired state of the pod object - also named the podspec. our pod creates a single container running the nginx:1.15.11 image from docker hub.

### labels

labels are key-value pairs attached to kubernetes objects. labels are used to organize and select a subset of objects, based on the requirements in place. many objects can have the same label. labels do not provide uniqueness to objects. controller use labels to logically group togther decoupled objects, rather than using objects names or ids.

![labels](https://courses.edx.org/assets/courseware/v1/6669997d43534cbd2f251a57ebe0587c/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/Labels.png)

in the image above, we have used two label keys: app and env. based on our requirements, we have given different values to our four pods. the label env=dev logically selects and groups the top two pods, while the label app=frontend logically selects and groups the left two pods. we can select one of the four pods - bottm left, by selecting two labels: app=frontend AND env=qa

### label selectors

controllers use label selector to select a subset of objects. kubernetes supports two types of selectors:

- equality-based selectors - equality-based selectors allow filtering of objects based on label keys and values. matching is achieved using the =, == (equals, used interchangeably), or != (not equals) operators. for example, with env==dev or env=dev we are selecting the objects where the env label key is set to value dev.
- set-based selectors - set-based selectors allow filtering of objects based on a set of values. we can use in, notin operators for label values, and exists/does not exist operator for label keys. for example with env in (dev, qa) we are selecting objects where the env label is set to either dev or qa; with !app we select objects with no label key app.

![selectors](https://courses.edx.org/assets/courseware/v1/2f83a85bea7ffd9fd861195725d7f146/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/Selectors.png)

### replication controllers

although no longer a recommended controller, a replicationContainer ensures a specified number of replicas of a pod is running at any given time, by constantly comparing the actual state with the desired state of managed application. it there are more pods than the desired count, the replication controller randomly terminates the number of pods exceeding the desired count, and, if there are fewer pods than the desired count, then the replication controller requersts additional pods to be created unitl the actual count matches the desired count. genrally, we do not deploy a pod independently, as it would not be able to re-start itself if terminated in error becuase a pod misses the much desired self-healing feature that kubernetes ohterwise promises. the recommended method is to use some type of a controller to run and manage pods.

the default recommend controller is the deployment which configures a replicaset controller to manage pods lifecycle.

### replicaSets I

a replicaSet is in part, the next-generation replicationController, as it implements the replication and self-healing aspects of the replicationController. replicasets supports both equality and set-based selectors, whereas replicationController only support equality based selectors.

with the help of a replicaSet, we can scale the number of pods running a specific application container image. scaling can be accomplished manually or through the use of an autoscaler.

![replicaSet](https://courses.edx.org/assets/courseware/v1/bdb9c27c39d027fc22133456a2882fa6/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/Replica_Set1.png)

### replicaSets II

let's continue with same replicaSet example and assume that one of the pods is forced to unexpectedly terminate (due to insufficient resources, timeout, etc) causing the current state to no longer match the desired state.

![replicaset](https://courses.edx.org/assets/courseware/v1/010297d9a0a76859de8110273eb56ae1/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/ReplicaSet2.png)

### replicaSet III

the replicaSet detects that the current state is no longer matching the desired state and triggers a request for an additional pod to be created, thus ensuring that the current state matches the desired state

![replicaSet](https://courses.edx.org/assets/courseware/v1/454b618e38084900bf247aecec0b3679/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/ReplicaSet3.png)

### deployments I

deployment objects provide declarative updates to pods and replicaSets. the deploymentController is part of the master node's controller manager, and as a controller it also ensures that the current state always matches the desired state. it allows for seamless application updates and rollbacks through rollouts and rollbacks, and directly manages its replicaSets for application scaling.

![deployment](https://courses.edx.org/assets/courseware/v1/5b2c8cbd6bed63c68f6a7d2566615d7f/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/Deployment_Updated.png)

### namespaces

if multiple users and teams use the same kubernetes cluster we can partition the cluster into virtual sub-cluster using namespaces. the names of the resources/objects created inside namespace are unique, but not across namespace in the cluster.
