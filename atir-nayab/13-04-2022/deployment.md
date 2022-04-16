# deployments

_a deployment provides declarative updated for pods and replicaSets_

you describe a desired state in a deployment, and the deployment controller changes the actual state to the desired state at a controlled rate. you can define deployments to create new replicaSets, or to remove existing deployments and adopt all their resources with new deployments.

> Do not manage replicaSets owned by a deployment.

## use case

- create a deployment to rollout a replicaSet. the replicaset creates pods in the background. check the status of the rollout to see if it succeeds or not.
- declare the new state of the pods by updating the PodTemplateSpec of the deployment. A new replicaset is created and the deployment manages moving the pods from the old replicaSet to the new one at a controlled rate. each new replicaset updates the revision of the deployment.
- rollback to an earlier deployment revision if the current state of deployment is not stable. each rollback updates the revision of the deployment.
- scale up the deployment to facilitate more load.
- pause the rollout of a deployment to apply multiple fixes to its PodTemplateSpec and then resume it to start a new rollout.
- use the status of the deployment as an indicator that a rollout has stuck.
- clean up older replicaSets that you need anymore.
