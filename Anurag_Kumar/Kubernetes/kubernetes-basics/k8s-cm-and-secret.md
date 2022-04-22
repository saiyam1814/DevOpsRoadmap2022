- `kubectl config set-context --current --namespace=kube-system` to change the active namespace

### Environment Varibales

- Direct way of exposing environment varibales in your pod manifest is 
```yaml
env:
  - name: APP_COLOR
    value: yellow
```

## Exposing environment variables using config maps and secrets. 
- Everything is similar, the only difference is that instead of defining the `value` directly we define the `valueFrom` and they we define the reference to config map or secret. 

```yaml
env:
  - name: MY_USERNAME
    valueFrom:
      configMapKeyRef:
        name: app-config
        key: USERNAME
```

```yaml
env:
  - name: MY_SECRET
    valueFrom:
      secretMapKeyRef:
        name: app-config
        key: SECRET
```

### ConfigMaps
- When you have lots of configuration data then it becomes difficult to manage all these values as environment variables. What we can do in this case is that we can take this data out of the pod manifest store it centrally as a config map. 
- Configmaps are used to pass key-value data in kubernetes. 
- When a pod is created inject the configmap into the pod so the key-value pairs are available as environment variables for the application hosted inside the container in the pod. 
- It's a two step process
  - Create a configmap
  - Inject the configmap into the pod 

##### Imperative way to create configmap
- `kubectl create configmap <configmap-name> --from-literal=<key>=<value>`
- If you want to pass more than one value in the config then pass `--from-literal` multiple times
```
kubectl create configmap \
    app-config --from-literal=APP_COLOR=blue \
               --from-literal=APP_MOD=prod
```

- Another way to create a config map is by using a file which contains key-value pairs and use the `--from-file` flag to pass the file while executing the command. 
```
kubectl create configmap \ 
    app-config --from-file=app-config.properties
```

#### Declarative way 
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config 
data: 
  APP_COLOR: blue
  APP_MOD: prod
```
- Create config map using the command `kubectl apply -f configmap.yaml`

#### Injecting that into a Pod 
- We will pass our configmap under the spec section of pod and we will do it via `envFrom`

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-simple-pod
spec:
  containers:
  - name: simplewebapp
    image: busybox
    envFrom:
    - configMapRef:
        name: app-config
```
- Please note that if we are passing only one key then we will use `configMapKeyRef` and if we are passing a existing config map then we will use `configMapRef`. 

### Secret 
- Again it's very very similar to what we havae studied in config map

##### Imperative way to create a secret 
- `kubectl create secret generic <secret-name> --from-literal=<key>=<value>`
- `kubectl create secret generic app-secret --from-literal=DB_HOST=mysql`
- Here also if you wish to pass multiple secret then type `--from-literal` flag multiple times. 
- `--from-file` flag is used when you have multiple secrets and it's a tedious task to get it done. 
- How will you encode what you're trying to pass while creating a secret in via Declarative approach.
- We will use the command `echo -n 'mysql' | base64 ` 
- How to decode the secret present
- We will use the command `echo -n 'abdfgc=23 | base64 --decode `
- `envFrom` property is a list. 