# deployment 101

ReplicaSet have one major drawback: once you select the pods that are managed by a replicaSet, you cannot change their pod templates.

For example, if you are using a replicaSet to deploy four pods with nodejs running and you want to change the nodejs to a newer version, you need to delete the replicaSet and recreate it.  
Restarting the pods causes downtime till the images are available and the pods are running again.

A deployment resource uses a replicaSet to manage the pods. However it handles updating them in a controlled way.
