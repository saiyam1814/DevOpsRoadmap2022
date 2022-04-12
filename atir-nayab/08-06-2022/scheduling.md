# what is kubernetes scheduling?

- the kubernetes scheduler is a core component of kubernetes: after a user or a controller creates a pod, the kubernetes scheduler, monitoring the object store for unassigned pods, will assign the pod to a node. Then, the kubelet, monitoring the object store for assingned pods, will execute the pod.

# what is the scheduler for?

![scheduler](../asset/scheduler.png)

the kubernetes scheduler is in charge of scheduling pods onto nodes. basically it works like this:

1. you create a pod
1. he scheduler notices that the new pod you created doesn't have a node assigned to it.
1. the scheduler assigns a node tothe pod

it's not responsible for actually running the pod - that's the kubelet's job. so it basically just needs to make sure every pod has a node assigned to it.

kubernetes in general has this idea of a controller. a controller's job is to:

- look at the state of the system
- notice ways in which the actual state does not match the desired state (like 'this pod needs to be assinged a node')
- repeat.

the scheduler is a kind of controller. there are lots of different controllers and they all have different jobs and operate independently.

# what is node affinity?

- in simple words this allows you to tell kubernetes to schedule pods only to specific subsets of nodes.
- the initial node affinity mechanism in early versions of kubernetes was the node selector field in the pod specification. the node had too include all the labels specified in that field to be eligible to become the target for the pod.

# node affinity

- node affinity is conceptually similar to nodeSelector - it allows you to constrain which nodes your pod is eligible to be scheduled on, based on labels on the node.
- there are currenlty two types of node affinity.
  1. requiredDuringSchedulingIgnoredDuringExecution (Required during scheduling, ignored during execution; we are also known as 'hard' requirements)
  1. preferredDuringSchedulingIgnoredDuringExecution (Preferred during scheduling, ignored during execution; we are also known as 'soft' requirements)

# anti-node affinity?

- some scenarios require that you don't use one or more nodes except for particular pods. think of the nodes that host your monitoring application
- those nodes shouldn't have many resources due to the nature of their role. thus, if ohter pods than those which have the monitoring app are scheduled to those nodes, they hurt monitoring and also degrades the application they are hosting.
- in such a case, you need to use anti-affinity to keep pods away form a set of nodes.

# what is node taints and tolerations?

- this kubernetes feature allows users to mark a node (taint the node) so that no pods can be scheduled to it, unless a pod explicitly tolerates the taint.
- when you taint a node, it is automatically excluded from pod scheduling. when the schedule runs the predicate tests on a tainted node, thay'll fail unless the pod has toleration for that node.
- let assume new member joins the development team, writes a deployment for her application, but forgets to exclude the monitoring node from the target nodes? kubernetes administrators need a way to repel pods from nodes without having to modify every pod definition.

# tolerations

- toleration is a way of ignoring a taint during scheduling. Tolerations aren't applied to nodes, but rather the pods. if we apply a toleration to teh podSpec, we could 'tolarate' the slow disks on that node and still use it.

- an important thing to notice, though, is that tolerations may enable a tainted node to accept a pod but it does not guarantee that this pod runs on that specific node.
- in other words, the tainted node will be considered as one of the candidates for running our pod. however, if another node has higher priority score, it will be chosen instead. for situations like this, you need to combine the toleration with nodeSelector or node affinity parameters.
