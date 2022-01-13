You must have played a game on your computer. Think about the different consituents of the game. A game consist of graphics, sounds, actions, characters, weapons, levels and many more such things. 
Now think of the installation process of the game. What if I told you that in order to play the game you need to install each and everything seperately. It will of course be a very tedious job. Instead of that we get a executable .exe file which contains all this things. This install all the necessary requirements to play the game in one go. 

HELM is similar to that. A cloud native application consists of many kubernetes objects.

In kubernetes we describe everything we do with yaml files. These YAML files represent objects such as deployments, pods, services, config maps, secrets and more. 
Kubernetes works on declrative concepts. It always monitors the desired state with the actual state. 
Managing this YAML files can be little cumbersome.


With the help of HELM we can install, upgrade, rollback and uninstall application in one go. We don't have to micromanage all the kubernetes object by ourselves. HELM can do that for us.

### three big concepts

1. Charts - Charts are the packages that contains all the resource definitions necessary to run application, tool or service inside a kubernetes cluster. Think of it like a gaming application which contains all the necessary components which are required to run the game(graphics, levels, characters, challenges etc.)

2. Repository - the place where charts can be collected and shared


3. Release - release is an running instance of a kubernetes cluster. Every time you install a new chart it gets created as a seperate release. So, if you want two databases running in your cluster, you can install MySQL chart twice.

## helm search

- Basically search in helm works in two ways :
    - Searching on the hub (helm search hub meshery)
    - Searching in the local client (helm search repo )
    
- Searching in the local client
    - for searching in the local client you have to add the repo first to your local client. For adding the repo to your local helm client use the command
        
        `helm repo add bitnami https://charts.bitnami.com/bitnami` 
        
        `helm repo add brigade https://brigadecore.github.io/charts`
        
        helm search repo bitnami will give you all the charts maintained by bitnami.
        
- Searching on the HUB
    - It will search the charts from [Artifact Hub](https://artifacthub.io)
    
