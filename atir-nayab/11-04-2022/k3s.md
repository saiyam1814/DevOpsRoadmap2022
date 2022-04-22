# k3s - what is it?

- fully compatible kubernetes distribution
- originally started by rancher labs (now suse)
- "3 is less than 8" - aims to be more approachable and less complex than k8s

# differences between distributions

- packaged as a single binary
- opinionated
- control plane node = worker node

## difference 1: single binary

- under 100MB
- installs everything it needs to run either as a control plane node or join a cluster
  - api server, schdeuler,controller
  - kubelet
  - contianerd

## difference 2: opinionated

- default ingress controller - traefik
- 1 container runtime - containerd
- virtual networking within the cluster - flannel
- cluster service discovery - coreDNS
- embedded load balancer - klipper-lb

**k3s throws out...**

- cloud provider specific code
  - replaced with cloud controller manager (CCM)
- third party-specific storage drivers
  - replaced with container storage interface (CSI)

## difference 3: Single node

- a control plane node can also perform work
  - memory requirements increase
- result: you can run a single-node "cluster"
  - minimum system requirements ~512MB for control plane
