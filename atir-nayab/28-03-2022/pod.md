# deploying your first nginx pod

## what are k8s pods?

- kubernetes pods are the foundational unit for all higher kubernetes objects.
- a pod hosts one or more containers
- it can be created using either a command or a yml/json file
- use kubectl to create pods, view the running ones, modifying their configuration or terminate them. kubernetes will attempt to restart a failing pod by default.
- if the pod fails to start indefinitely, we can use the kubectl describe command to know what went wrong

## why does kubernetes use a pod as the smallest deployable unit, and not a single container?

there are good reasons to add a layer of abstraction represented by the pod.  
To manage a container, kubernetes needs additional information, such as a restart polucy, which defines what to do with a container when it terminates, or a liveness probe, which defines an action to detect if a process in a container is still alive from the applications perspective, such as web server responding to http requests.  
instead of overloading the existing "thing" with additional properties, kubernetes architects have decided to use a new entity, the pod, that logically contains (wraps) one or more containers that should be managed as a single entity.

## why does kubernetes allow more than one container in a pod?

containers in a pod run on a logical host; they use the same network namespace (in other words, the same IP address and port space), and the same IPC namespace. They can also use shared volumes. These properties make it possible to these containers to efficiently communicate, ensuring data locality. also pods enable you to manage several tightly coupled application containers as a single unit.

## use cases for multi-container pods

the primary purpose of a multi-container pod is to support co-located, co-managed helper processes fo ra primary application. There are some general patterns for using helper processes in pods.  
sidecar container help the main container. some examples include log or data change watchers, monitoring adapters and so on.
proxies, bridges and adapters connect the main container with the external world.

## shared volumes in a kubernetes pod

in kubernetes, you can use a kubernetes volume as a simple and efficient way to share data between containers in a pod. for most cases, it is sufficient to use a directory on the host that is shared with all containers within a pod.  
kubernetes volumes enables data to survive container restarts, but these volumes have the same lifetime as the pod. That means that the volume (and the data it holds) exists exactly as long as that pod exists. if that pod is deleted for any reason, even if an identical replacement is created, the shared volume is also destroyed and create a new.

multi container pod example

```yml
apiVersion: v1
kind: Pod
metadata:
  name: mc1
spec:
  volumes:
    - name: html
      emptyDir: {}
  containers:
    - name: 1st
      image: nginx
      volumeMounts:
        - name: html
          mountPath: /usr/share/nginx/html
    - name: 2nd
      image: debian
      volumeMounts:
        - name: html
          mountPath: /html
      command: ["/bin/sh", "-c"]
      args:
        - while true; do
          date >> /html/index.html;
          sleep 1;
          done
```
