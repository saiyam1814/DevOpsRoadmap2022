# liveness, readiness and startup probes

the kubelet uses liveness probes to know when to restart a container. for example, liveness probes could catch a deadlock, where an application is running, but unable to make progress.

the kubelet uses readiness probes to know when a container is ready to start accepting traffic. A pod is considered ready when all of its containers are ready. one use of this signal is to control which pods are used as backends for services. when a pod is not ready, it is removed from service load balancers.

# liveness

many applications running for long periods of time eventually transition to broken states, and connot recover except by being restarted. kubernetes provides liveness probes to detect and remedy such situations.

# readiness probes

sometimes, applications are temporarily unable to server traffic. for example, an application might need to load large data or configure files during startup, or depend on external services after startup. in such cases, you don't want to kill the application, but you don't want to send it requrest either. kubernetes provides readiness probes to detect and mitifate these situations.

> readiness probes runs on the container during its whole lifecycle.

> liveness probes do not wait for readiness probes to succeed. it you want to wait before executing a liveness probe you should use initialDelaySeconds or a startupProbe.

# protect slow starting containers with startup

sometimes you have to deal with legacy applications that might require an additional startup time on their first initialization, in such cases, it can be tricky to set up liveness probe parameters without compromiaing the fast response to deadlocks that motivated such a probe.
