# init containers

specialized containers that run before app containers in a pod. init containers can contain utilities or setup scripts not present in an app image.

a pod can have multiple containers running apps within it, but it can also have one or more inti containers, which are run before the app containers are started.

init containers are exaclty like regular containers except -

- init containers always run to completion
- each init container must complete successfully before the next one starts.

if a pod's init container fails, the kubelet repeatedly restarts that init container until it succeeds. however if the pod has `restartPolicy` of Never, and an init container fails during startup of that pod, kubernetes treats the orverall pod as failed.

# differences from regular containers

init containers support all the fields and features of app containers, including resource limits, volumes, and security settings.

also inti containers do not support `lifecycle`, `livenessProbe`, `readinessProbe`, or `startupProbe` because they must run to completion before the pod can be ready.
