## Networking Challanges in Kubernetes
- In the monolithic architecture if we needed a service then it was just a in memory function call which was easy and quick.
- In the microservices architecture we are building containerized services for our application. 
- These services will communicate over a network means we need to figure out a lot of things like latency issue, authorisation issue, protocols, security etc. 
- In kubernetes we will see the following networking :
  - Container-to-container communication inside Pods
  - Pod-to-Pod communication on the same node and across cluster nodes
  - Pod-to-Service communication within the same namespace and across cluster namespaces
  - External-to-Service communication for clients to access applications in a cluster.