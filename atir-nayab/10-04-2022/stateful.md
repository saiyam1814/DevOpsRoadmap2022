# what is statefulset and how is it different from deployment?

a statefulset is a kubernetes controller that is used to manage and maintain one or more pods. however so do other controller like replicaSets and the more robust, deployments.

a stateless application is one that does not care which network it is using, and it does not need permanent storage. examples of stateless apps may include web servers (apache, nginx or tomcat).

on other hand, we have stateful apps. let's say you have a solr database cluster that is managed by several zookeeper instances. for such an application to function correctly, each solr instance must be aware of the zookeeper instances that are controlling it. similarly, the zookeeper instances themselves establish connections between each other to elect a master node .due to such a design, solr clusters are an example of stateful applications. other examples of stateful applications include mysql clusters, redis, kafka, mongodb and others.

given this difference, deployment is more suited to work with stateless applications. as far as a deployment is concerned, pods are interchangeable. while a stateful keeps a unique identity for each pod it manages. it uses the same identity whenever it needs to reschedule those pods.

# exposing a statefulSet

a kubernetes service acts as an abstraction layer. in a stateless application like an nginx web server, the client does not (and should not) care which pod receives a response to the request. the connection reaches the service, and it routes it to any backend pod. this is not the case in stateful apps. solr pod may need to reach the zookeeper master, not any pod. for this reason, part of the statefulset definition entails a headless service. a headless service does not contain a cluster ip. instead, it creates several endpoints that are used to produce dns records. each dns record is bound to a pod. all of this is done internally by kubernetes, kut it's good to have an idea about how it does it.

# storage class

storage classes are kubernetes objects that let the users specify which type of storage they need from the cloud provider. different storage classes represent various service quality, such as disk latency and throughput, and are selected depending on the scenario they are used for and the cloud provider's support.

# persistent volumes and persistent volume claims

persistent volumes act as an abstraction layer to save the user from going into the details of how storage is managed and provisioned by each cloud provider. by definition statefulsets are the most frequent users of persistent volumes since they need permanent storage of their pods.

a persistent volume claim is a request to use a persistent volume. if we are to use the pods and nodes analogy, then consider persistent volumes as the nodes and persistent volume claims as the pods that use the node resources.
