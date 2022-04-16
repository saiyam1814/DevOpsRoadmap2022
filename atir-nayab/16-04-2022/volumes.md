# volumes

one problem is the loss of data when container crashes. the kubelet restarts the container but with a clean state. a second problem occurs when sharing files between containers running together in a pod.

docker volume is a directory on disk or in another container. docker provides volume drivers, but the functionality is somewhat limited.

kubernetes supports many types of volumes. a pod can use any number of volume types sumultaneously. ephemral volume types have a lifetime of a pod, but persistent volumes exist beyond the lifetime of a pod. when a pod ceases to exist, kubernetes destroys ephmeral volumes; however kubernetes does not destroy persistent volumes. for any kind of volume in a given pod data is preserved across container restarts.

# persitent volumes

the PersistentVOlume subsystem provides an API for users and administrators that abstracts details of how storage is provided from how it is consumed.

- PV - a PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by am administrator or dynamically provisioned using storage classes. is is a resource in the cluster just like a node is a cluster resource. pvs are volume plugins like volumes,
- PVC - a PersistentVolumeClaim (PVC) is a request for storage by a user. it is similar to a pod. pods consume node resources and pvcs consume pv resources. pods can request specific levels of resources (cpu and memory). claims can request specific size and access modes (e.g., they can be mounted ReadWriteOnce, ReadOnlyMany or ReadWriteMany).

while presistentVolumeClaims allow a user to consume abstract storage resources. it is common taht users need persistentVolumes with varying properties, such as performance, for different problems.
