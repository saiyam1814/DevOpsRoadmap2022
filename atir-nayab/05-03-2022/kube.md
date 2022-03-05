# what is containers?

IRL:

> containers are a standardized way for you to package items together into one shipment.

shipping containers:

- standard packaging
  - islation and efficiency
- portable
- separation of concerns

IVL:

> containers are a standardized way for you to package your application, its configuration and dependencies together into a single logical object

virtual containers:

- standard packaging
  - isolation & efficiency
- portable
- separation of concerns

![application](../asset/application.png)

![container vs virtual machine](../asset/containerVsVM.png)

linux kernel features:

- namespaces - namespace allow os limit what process can see such processes, file system etc.
- control groups (cgroups) - limit how much process can use like cpu, memory etc.

To computers:  
A contaienr is a process extracted from tarballs anchored to namespaces and controlled by cgroups.

![docker visual representation](../asset/dockerVisual.png)

# concepts behind kubernetes

- declarative vs imperative
- control loops
- reconciliation

### declarative vs imperative

declarative - describes desired state  
imperative - tell about action that needs to be done

### control loop

![control loop](../asset/controlLoop.png)

### reconciliation

- reconciliation is simply making the actual state look like the desired state via the control loop
- occurs in numerous places throughout kubernetes
  - `kubectl apply -f file.yaml`
