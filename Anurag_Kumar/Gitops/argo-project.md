### Reconcillation loop
- It basically means how frequently the gitops agent (argocd) will sync the cluster state and the git repo. 
- By default it is 3 minutes in argocd. 
- To change the default reconcillation time edit the `argocd-cm config` file
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: argocd-cm
  namespace: argocd
  labels:
    app.kubernetes.io/name: argocd-cm
    app.kubernetes.io/part-of: argocd
data:
  timeout.reconciliation: 240s
```
- If you want to set your reconcillation time to 1 minute then set the value of `timeout.reconcillation: 60s`
- Now the agent will sync the changes after 1 minute

- If you're using webhooks with argocd then whenever you commit something to you gitops repo then resources in the cluster will be created. No need to wait for 3 minutes in this case. 

### Status
- Basically two types of status are there
1. Synced (Git state = desired state)
2. Out of Sync (Git state != desired state)

### Health Status
1. Healthy
2. Progessing
3. Suspended 
4. Missing
5. Degraded 
6. Unknown

- Health status depends on the underlying kubernetes resources
- For built in kubernetes resources 
  - Deployments, Replicasets, stateful sets and daemon sets are considered healthy if observed generation = Desired generation
  - Number of updated replicas equals the number of desired replicas

- For custom resouces we will define the health in lua scripts.

### Sync Strategies
1. Manual or automatic sync
    - Automatic - On finding the new version of the application, argoCD will update/create the changes in cluster. Even when you are deploying the application for the first time no need to sync it after creating the application. 
    - Manual - Argo will detect the new version of application but will not apply it to the cluster

2. Auto pruning of resources (Only applicable for automatic sync)
    - Enabled - If something is deleted in git repo then argoCD will delete that from the cluster as well
    - Disabled - Will not delete anything from the cluster

3. Self healing of clusters ( only applicable for automatic sync)
    - Enabled - If something we will try to add directly to the cluster then argoCD will discard it. This is a great advantage as it always makes your environments bulletproof against ad-hoc changes via kubectl. 

- *Usually it's a good practice to add everything through the git repo only. Do not try to add anything to the cluster directly.*
