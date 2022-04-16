# service

an abstract way to expose an application running on a set of pods as a network service. With kubernetes you don't you don't need to modify your application to use an unfamiliar service discovery mechanism. Kubernetes gives pods their own IP address and a single dns name for a set of pods, and can load-balance across them.

## motivation

pods are created and destroyed to matches the desired state. each pod get its own ip address. We can't keep track of the ip addresses if pods are created and destroyed. That's where services come in.

## service resource

A service is an abstraction which defines logical sets of pods and policy which to access them. the set of pods targeted by a service is usually determined by a selector.
