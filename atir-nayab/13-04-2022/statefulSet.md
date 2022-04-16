# statefulSets

statefulSet is the workload api object used to manage stateful applications.

manages the deployment and scaling of a set of pods, and provides guarantees about the ordering and uniquness of these pods.

like a deployment, a statefulSet manages pods that are based on an identical container spec. Unlike a deployment a statefulSet maintains a sticky identity for each of their pods. These pods are created from the same spec, but are not interchangeable: each has a persistent identifier that it maintains across any rescheduling.

if you want to use storage volumes to provide persistence for your workload, you can use a statefulSet as part of the solution. Although individual pods in a statefulSet are susceptible to failure, the persistent pod identifiers make it easier to match existing volumes to the new pods that replace any that have failed.

## using statefulSets

statefulSets are valuable for applications that require one or more of the following.

- stable, unique network identifiers
- stable, persistent storage
- ordered, graceful deployment and scaling
- ordered, automated rolling updates.
  in the above stable is synonymous with persistence across pod (re)scheduling. if an application doesn't require any stable identifiers or ordered deployment, deletion, or scaling, you should deploy your application using a workload object that provides a set of stateless replicas.
