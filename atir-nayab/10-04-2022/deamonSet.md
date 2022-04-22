# what is daemonSet?

daemonsets are used to ensure that some or all of your k8s nodes run a copy of a pod which allows you to run a daemon on every node.

when you add a new node to the cluster a pod gets added to match the nodes. similarly, when you remove a node from your cluster, the pod is put into the trash. deleting a daemonset cleans up the pods that it previously created.

a daemonset is another controller that manages pods like deployments, replicasets, and statefulsets. it was created for one particular purpose: ensuring that the pods it manages to run on all the cluster nodes. as soon as a node joins the cluster, the daemonSet ensures that it has necessary pods running on it. when node leaves the cluster, those pods are garbage collected.

daemonsets are used in kubernetes when you need to run one or more pods on all (or a subset of) the nodes in a cluster. the typical use case for a daemonset is logging and monitoring for the hosts.

why use of daemonSets

- to run a daemon for cluster storage on each node, such as glusterd - ceph
- to run a daemon for logs collection on each node, such as ; fluentd - logstash
- to run a daemon for node monitoring on ever note, such as: prometheus node exporter - collectd - datadog agent
- as your use case gets more complex, you can deploy multiple daemonsets for one kind of daemon, using a variety of flags or memory and cpu requests for various hardware types.
