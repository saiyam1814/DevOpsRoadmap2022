# workloads

a workloads is an application running on kubernetes. whether your workload is a single component or serveral that twork together, on kubernetes you run it inside a set of pods. In kubernetes a pod represents a set of running containers on your cluster.

kubernetes pods have a defined lifecycle. for example once a pod is running in your cluster then a critical fault on the node where that pod is running means that all the pods on that node fail. kubernetes treats that level level of failure as final: you would need to create a new pod to recover even if the node later becomes healthy

however to make life considerably easier, you don't need to manage each pod directly.

kubernetes provides several built in workload resources

- deployment and replicaset (replacing the legacy resource replciation controller). deployment is a good fit for managing a stateless application workload on your cluster, where any pod in the deployment is interchangeable and can be replaced if needed.
- statefulset - lets you run one or more related pods that do track state somehow.
- daemonset defines pods that provide node-local facilities. these might be fundamental to the operation of your cluster, such as a networking helper tool, or be part of an add-on
- job and cronjob defines tasks that run to completion and then stop. jobs represent one-off tasks, whereas cronJobs recur according to a schedule.

# services, load balancing and networking

## the kubernetes network model

every pod gets its own ip address. this means you do not need to explicitly create links between pods and you almost never need to deal with mapping container ports to host ports.
this creates a clean, backward-compatible model where pods can be treated much like vms or physical hosts from the perspective of ports allocation, naming, service discovery, load balancing application configuration, and migration.

kubernetes networking addresses four concerns:

- containers within a pod use networking to communicate via loopback.
- cluster networking provides communication between different pods.
- the service resource lets you expose an application running in pods to be reachable from outside your cluster.
- you can also use services to publish services only for consumption inside your cluster.

## service

an abstract way to expose an application running on a set of pods as a network service.

### define a service

a service in kubernetes is a REST object, similar to a pod. like all of the REST objects, you can POST a service definition to the api server to create a new service. the name of a service object must a valid rfc label name

for example, suppose you have a set of pods where each listens on tcp port 9376 and contains a label app=MyApp

```
apiVersion: v1
kind: service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```

this specification creates a new service object named "my-service" which targets tcp port 9376 on any pod with the app=MyApp label.

kubernetes assigns this service an IP address (sometimes called the cluster ip) which is used by the service proxies

the controller for the service selector continuously scans for pods that match its selector, and then posts any updates to an endpoint object also names my-service

> a service can map any incoming port to a targetPort. By default and for convenience, targetPort is set to the same value as the port field.

### services without selectors

services most commonly abstract access to kubernetes pods, but they can also abstract other kids of backends.

- you want to have an external database cluster in production, but in your test environment you use your own databases.
- you want to point your service to a service in a different namespace or on another cluster.
- you are migrating a workload to kubernetes. while evaluating the approach, you run only a portion of your backends in kubernetes.

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```

because this service has no selector, the corresponding endpoints object is not created automatically. you can manually map the service to the network address and port where it's running, by adding endpoints object manually

```
apiVersion: v1
kind: Endpoints
metadata:
  name: my-service
subnets:
  - addresses:
    - ip: 192.0.2.42
  ports:
    - port: 9376
```

# ingress

an api object that manages external access to the services in a cluster, typically http. ingress may provide load balancing, ssl termination and name-based virtual hosting.

Ingress exposes http and https routes from outside the cluster to services within the cluster. traffic routing is controlled by rules defined on the ingress resource.
