# what are kubernetes services?

say, you have pods running nginx in a flat, cluster wide, address space. In theory, you could talk to these pods directly, but what happens when a node dies? the pods die with it, and the deployment will create new ones, with different ips. this is the problem a services solves.

kubernetes pods are mortal. they are born and when they die, they are not resurrected. if you use a deployment to run your app, it can create and destroy pods dynamically. each pod gets its own ip address, however in a deployment, the set of pods running in one moment in time could be different from the set of pods running that application a moment later.

This leads to a problem: how do the frontends find out and keep track of which IP address to connect to, so that the frontend can use the backend part of the workload?

A kubernetes service is an abstraction which defines a logical set of pods running somewhere in your cluster, that all provide the same functionality. when created, each service is assigned a unique ip address (also called cluster IP). this address is tied to the lifespan of the service, and will not change while the service is alive. pods can be configured to talk to the service, and know that communication to the service will be automatically load-balanced out to some pod that is a memeber of the service.

# deploying a kubernetes service

> kubectl apply -f nginx-svc.yaml

this specification will create a service which targets tcp port 80 on any pod with the run:my-nginx label, and expose it on an abstracted service port (targetPort: is the port the container accepts traffic on, port: is the abstracted service port, which can be any port other pods use to access the service). view service api object to see the list of supported fields in service definition. check your service.

> kubectl get svc my-nginx

as mentioned previously, a service is backed by group of pods. these pods are exposed through endponts. the service's selector will be evaluated continuously and the results will be posted to an endpoints object also named my-nginx. when a pod dies, it is automatically removed from the endpoints, and new pods matching the service's selector will automatically get added to the endpoints.

> kubectl describe svc my-nginx

you should be able to curl the nginx service on:from any node in your cluster. service ip completely virtual, it never hits the wire

# accessing the service

kubernetes supports 2 primary modes of finding a service - environment variables and dns

# environment variables

when a pod runs on a node, the kubelet adds a set of environment variables for each active service.

# dns

kubernetes offers a dns cluster addon service that automatically assigns dns names to other services

> kubectl get services kube-dns --namespace=kube-system

# cluster ip

the cluster up is the default service type. kubernetes will assign an internal ip address to your service. this ip address is reachable only from inside the cluster. you can optionally - set this ip in the service definition file. think of the case when you have a dns record that you don't want to change and you want the name to resolve to the same ip address.  
however, you cannot just add any ip address. it must be within the service-cluster-ip-range

# headless service in kubernetes

the default behavior of kubernetes is to assign an internal ip address to the service. through this ip address, the service will proxy and load balance the requests to the pods behind. if we explicitly set this ip address (cluster ip) to none, this is like telling kubernetes i don't need load balancing or proxying, just connect me to the first available pod.

# nodePort

this is one of the service types that are used when you want to enable external connectivity to your service. if you're having four nginx pods, the node port service type is going to use the ip address of any node in the cluster combined with a specific port to route traffic to those pods.

you can use the ip address of any node, the service will receive the request and route it to one of the pods.

# load balancer

this service type works when you are using a cloud provider to host your kubernetes cluster. when you choose loadbalancer as the service type, the cluster will contact the cloud provider and create a load balancer. traffic arriving at this load balancer will be forwarded to the backend pods.
