# chapter 9. Authentication, Authorization, Admission Control

## introduction

every api request reaching the api server has to go through several control stages before being accepted by the server and acted upon. in this chapter, we will learn about authentication, authorization, and adminssion control stages of the kubernetes api access control.

## authentication, authorization, and admission control - overview

to access and manage kubernetes resources or objects in the cluster, we need to access a specific api endpoint on the server. each access request goes through the following access control stages.

- authentication - logs in a user
- authorization - authorizes the api submitted by the authenticated user.
- admission control - software modules that validate and/or modify user requests based.

## authentication I

kubernetes does not have an object called user, nor does it store usernames or other related details in its object store. however, even without that, kubernetes can use usernames for the authentication phase of the api access control, and to request logging as well.

kubernetes supports two kinds of users

- normal users - they are managed outside of the kubernetes cluster via independent services like user/client certificates, a file listing usernames/passwords, google accounts etc.
- service accounts - service accounts allow in-cluster processes to communicate with the api server to perform various operations. Most of the service accounts are created automatically via the api server, but they can also be created manually. the service accounts are tied to a particular namespace and mount the respective credentials to communicate with the api server as secrets.

if properly configured, kubernetes can also support anonymous requersts, along with requests form normal users and service accounts. user impersonation is also supported allowing a user to act as another user, a helpful feature for administrators when troublshooting authorization policies.

## authentication II

- x509 client cerificates - to enable ceritificate authentication, we need to reference a file contianing one or more certificate authorities by passing the --clinet-ca-file=SOMEFILE option to the api server. the certificate authorities mentioned in the file would validate the client certificates presented by users to the api server.
- static token file - we can pass file containing pre-defined bearer tokens with the --token-auth-file=SOMEFILE opiton to the api server. currently, these tokens would last indefinitely, and they cannot be changed without restarting the api server.
- bootstrap tokens - tokens used for bootstrapping new kubernetes clusters.
- service account tokens - automatically enabled authenticators the use signed bearer tokens to verify requests. these tokens get attached to pods using the serviceAccount ccontroller, which allows in cluster processes to talk to the api server.
- openid connect tokens - openid connect helps us connect with oauth2 providers, such as azure active directory, salesforce, and google, to offload the authentication to external services.
- webhook token authentication - with webhook based authentication, verification of bearer tokens can be offloaded to a remote service.
- authenticating proxy - allows for the programming of additional authentication logic

we can enable multiple authenticators, and the first module to successfully authenticate the request short - circuits the evaluation. to ensure successful user authentication, we should enable at least two methods: the service account tokens authenticator and one of the user authenticators.

## authorization I

after a successful authentication, users can send the api requests to perform different operations. here, these api requests gets authorized by kubernetes using various authorization modules, that allow or deny requests.

## authorization II

authorization modes (part 1)

- node - node authorization is a special-purpose authorization mode which specifically authorizes api requests made by kubelets. it authorizes the kubelet's read operations for the services, endpoints, or nodes, and writes operations for nodes, pods, and events.
- attribute-based access control (ABAC) - with the abac authorizer, kubernetes grants access to api requests, which combine policies with attributes. in the following example, user student can only read pods in the namespace lfs158

```
{
  "apiVersion": "abac.authorization.kubernetes.io/v1beta1",
  "kind": "Policy",
  "spec": {
    "user": "student",
    "namespace": "lfs158",
    "resource": "pods",
    "readonly": true
  }
}
```

- webhook - in webhook mode, kubernetes can request authorization decisions to be made by third party services, which would return true for successful authorization, and false for failure. in order to enable the webhook authorizer, we need to start the api server with the --authorization-webhook-config-file=SOME_FILENAME option, where SOME_FILENAME is the configuration of the remote authorization service.

### authorization III

- role-based access control (rbac) - in general, with rbac we regualte the access to resouces based on the roles of individual users. in kubernetes, multiple roles can be attached to subjects like users, service accounts, etc. while creating roles, we restrict resource access by specific operations, such as create, get, update, patch etc. these operations are reffered to as verbs.

role  
a role grants access to resources within a specific namespace

clusterRole  
a clusterRole grants the same permissions as role does, but its scope is cluster-wide.

roleBinding  
it allows us to bind users to the same namespace as a role. we could also refer a clusterRole in roleBinding, which would grant permissions to namespace resources defined in the clusterRole within the roleBindings namespace.

clusterRoleBinding  
it allows us to grant access to resources at a cluster level and to all namespaces.

## admission control

admission controllers are used to specify granular access control policies, which include allowing privileged containers, checking on resource quota, etc. we force these policies using different admission controllers, like resourceQuota etc. we force these policies using different admission controllers, like resourceQuota, defaultStorageClass, alwaysPullImages, etc. they come into effect only after api requests are authenticated and authorized.

kubernetes admission control can also be implemeted though custom plugins, for a dynamic admission control method. these plugins are developed as extensions and run as admission webhooks.

# chapter 10 - services

## connecting users or applicaitons to pods

to access the application, a user or another application need to connect to the pods, as pods are ephemeral in nature, resources like IP addesses allocated to them cannot be static. pods could be terminated abruptly or be resheduled based on existing requirements.

![user connected pods via their ip addresses](https://courses.edx.org/assets/courseware/v1/804a02a9297c12e78e820ba6693f5c3c/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/service-1.png)

unexpectedly, one of the pods the user/client is connected to is terminated, and a new pod is created by the controller. the new pod will have a new ip address, that will not be immediately known by the user/client of the earlier pod.

![a new pod is created after an old one terminated unexpectedly](https://courses.edx.org/assets/courseware/v1/3f0cf38178b8639549e3b78b7c634ba3/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/service-2.png)

to overcome this situation, kubernetes provides a higher level abstraction called service, which logically groups pods and defines a policy to access them. this grouping is achieved via labels and selectors.

## services

labels and selectors use a key/value pair format. in the following graphical representation, app is the key, frontend and db are label values for different pods.

![grouping of pods using labels and selectors](https://courses.edx.org/assets/courseware/v1/957221cb68bbacd773891e2a8be97d59/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/service1.png)

## kube-proxy

each cluster node runs a daemon called kube-proxy, that matches the api server on the master node for the addition, updates, and removal of services and endpoints. kube-proxy is responsible for implementing the service configuration on behalf of an administrator or developer, in order to enable traffic routing to an exposed application running in pods.

![kube-proxy, services and endpoints](https://courses.edx.org/assets/courseware/v1/f6184f33a4c81a2c59eb9c28bf79c3ae/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/kubeproxy.png)

## service discovery

as services are primary mode of communication between containerized applications managed by kubernetes

- environment variables - as soon as the pod starts on any worker node, the kubelet daemon running on that node adds a set of environment variables in the pod for all active services.
- dns - kubernetes has an add-on for dns, which creates a dns record for each service and its format is my-svc.my-namespace.svc.cluster.local. sevices within the same namespace find other services just by their names.

## service type

while defining serice, we can also choose its access scope. we can decide whether the service:

- is only accessible within the cluster
- is accessible from within the cluster and the external world
- maps to an entity which resides either inside or outside the cluster.

acess scope is decided by service type property, defined when creating the service.

### service type: cluster ip and nodeport

cluster ip is the default service type. a service receives a virtual ip address, known as its clusterip. this virtual ip address is used for communicating with the service and is accessible only from within the cluster.

![nodeport](https://courses.edx.org/assets/courseware/v1/c9ddb793c9e82594d751d9abcf412356/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/NodePort.png)

the nodeport servicetype is useful when we want to make our services accessible form the external world. the end-user connects to any worker node on the specified high-port, which proxies the request internally to the the clusterIP of the service, then the request is forwarded to the applications running inside the cluster.

### service tyep: load balancer

with the load balancer service type:

- nodeport and clusterIP are automatically created, and the external load balancer will route to them
- the service is exposed at a static port on each worker node
- the service is epxosed externally using the undrelying cloud provider's laod balancer feature.

![load balancer](https://courses.edx.org/assets/courseware/v1/aabed46d4f2f9c203dc16c5b9221cac4/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/LoadBalancer.png)

### service type: external ip

a service can be mapped to an external ip address if it can route to one or more of the worker nodes. traffic that is ingressed into the cluster with the externalIp on the service port, gets routed to one of the service endpoints. this type of service requires an external cloud provider such as google cloud provider

![external ip](https://courses.edx.org/assets/courseware/v1/0d4109849e01e60a5f67f5d715fb4ea7/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/ExternalIP.png)

### service type: external name

external name is a special service type, that has no selctors and does not define any endpoints. when accessed within the cluster, it returns a cname record of an externally configured service.
