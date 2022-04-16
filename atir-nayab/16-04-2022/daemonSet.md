# DaemonSet

a daemonset ensures that all (or some) nodes run a copy of a pod. As nodes are added to the cluster, pods are added to them. as nodes are removed from the cluster, those pods are garbage collected. deleting a daemonSet will clean up the pods it created.

daemonSet use cases

- running a cluster storage daemon on every node
- running logs collection daemon on every node
- running a node monitoring daemon on every node.
