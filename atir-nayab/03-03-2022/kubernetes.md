# chapter 12. kubernetes volumen management

all data is stored inside a contianer is deleted if the container crashes. however, the kubelet will restart it with a clean slate, which means that it will not have any of the old data.

to overcome this problem, kubernetes uses volumes, storage abstractions that allow various storage technologies to be used by kubernetes and offered to containers in pods as storage media. a volume is essentially a mount point on the container's file system backed by a storage medium. the storage medium, content and access mode are determined by the volume type.

![shared volume in pod](https://courses.edx.org/assets/courseware/v1/f3ec9ad18ca681bc581c166aa3c5cba2/asset-v1:LinuxFoundationX+LFS158x+3T2020+type@asset+block/pod.svg)

in kubernetes, a volume is linked to a pod and can be shared among the containers of that pod. although the volume has the same life span as pod, meaning that it is deleted together with the pod, the volumes outlives the container of the pod - this allows data to be preserved across container restarts.

## volume types

a directory which is mounted inside a pod is backed by the underlying volume type. a volume type decides the properties of the directory.

- emptyDir - an empty volume is created for the pod as soon as it is scheduled on the worker node. the volume's life is tightly coupled with the pod. if the pod is terminated, the content of emptyDir is deleted forever.
- hostPath - with the hostPath volume type, we can share a directory between the host and the pod. if the pod is terminated, the content of the volume is still availble on the host.
- gcePersistentDisk - with the gcePersistentDisk volume type, we can mount a goolge compute engine (gce) persistent dist into a pod.
- awsElasticBlockStore - with the awsElasticBlockStore volume type, we can mount an aws ebs volume into a pod.
- azureDisk - with azureDisk we can mount a microsoft azure data disk into a pod
- azurefile - with azurefile we can mount a microsoft azure file volume into a pod.
- cephfs - with cephfs, an existing cephFs volume can be mounted into a pod. when a pod terminates, the volume is unmounted and the content of the volume are preserved.
- nfs - with mfs, we can mount as nfs share into a pod.
- iscsi - with iscsi, we can mount an iscsi share into a pod.
- secret - with the secret volume type, we can pass sensitive information, such as passwords, to pods.
- configMap - with configMap objects, we can provide configuration data, or shell commands and agruments into a pod.
- persistentVolumeClaim - we can attach a persistentVolume to a pod using a persistentVolumeClaim.
