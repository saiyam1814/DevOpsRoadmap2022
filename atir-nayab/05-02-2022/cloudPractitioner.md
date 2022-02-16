## Infrastructure as Code (IaC)

you write a configuration script to automate creating, updating, or destroying cloud infrastructure.

- IaC is a blueprint of your infrastructure
- IaC allows you to easily share, version or inventory your cloud infrastructure as code.

**AWS has two offerings for writing infrastructure as Code.**

- aws cloud formation (CFN) - cfn is a declarative IaC tool.
  - what you see is what you get explicit
  - more verbose, but zero chance of mis-configuration
  - uses scripting language eg JSON, YAML, XML
- aws cloud development kit (CDK) - cdk is an imperative IaC tool.
  - You say what you want, and the rest is filled in implicit
  - less verbose, you could end up with misconfiguration
  - does more than declarative
  - uses programming languages eg. python, ruby, javascript

### CloudFromatioN (CFN)

aws cloudformation allows you to write infrastructure as code. (IaC) as either a JSON or YAML file.

CloudFormation is simple but it can lead to large files or is limited in some regard to creating dynamic or repeatable infrastructure compared to cdk.

cloudformation can be easier to DevOps engineer who do not have a background in web programming languages.

Since CDK generates out cloudformation its still important to be able to read and understand cloudformation in order to debug IaC stacks.

## aws toolkit

aws toolkit is an open source plugin for vscode to create, debug deploy AWS resources

## access keys

access keys is a key secret required to have programmatic access to aws resources when interacting with the aws api outside of the aws management credentials.

- never share your access keys
- never commit access keys to a codebase

## access keys

access keys are to be stored in ~/.aws/credentials and follow a TOML file format

## aws documentation

aws documentation is a large collection of technical documentation on how to use aws services.

## shared responsibility model

the shared responsibility model is a cloud security framework that define the security obligations of the customer versa the cloud service provider (CSP)

## aws shared responsibility model

- aws
  - software
    - compute
    - storage
    - database
    - networking
  - hardware
    - regions
    - availability zones
    - edge locations
    - physical security
- customer

  - configuration of managed services or third party software
    - platforms
    - applications
    - identity and access management (IAM)
  - configuration of virtual infrastructure and systems
    - operating system
    - network
    - firewall
  - security configuration of data
    - client side data encryption
    - server side encryption
    - networking traffic protection
    - customer data

- IN - customers are responsible for security in the cloud
  - data
  - configuration
- OF - aws is responsible for security of the cloud

  - hardware operation of managed services.
  - operation of managed services
  - global infrastructure

## computing services

elastic compute cloud (EC2) allows you to launch Virtual machines (VM)

what is viratual machine?
A virtual machine (vm) is an emulation of a physical computer using software. Server virtualization allows you to easily create, copy, resize or migrate your server.
Multiple VMs can run on the same physical server so you can share the cost with other customers. Imagine if your server or computer was an executable file on your computer.

_when we launch a virtual machine we call it an instance_

### virtaul machines

an emulated of physical computer using software

- amazon lighshell - it is the manged virtual server service. It is the friendly version of EC2 virtaul machine
- containers virtaulizing an operating system to run multiple workloads on a single OS instance. Containers are generally used in micro-service architecture.
- elastic container registry - it is repository for container images in order to launch a container you need an image. An image just means a saved copy.
- ECS fargate - it is a serverless orchestration container service. It is the same as ECS expect you pay-on demand per running container
- Elastic kubernetes service - is a fully managed kubernetes service. Kubernetes is an open-source orchestration software that was created by google and is generally the standard for managing microservices.
- aws lambda - it is a serverless function service. you can run code without provisioning or managing or managing servers.

## High Performance Computing Services

The Nitro system a combination of dedicated hardware and lightweight hypervisor enabling faster innovation and enhanced security. All new ec2 instance types use the nitro system

- Nitro cards - specialize cards for VPC, EBS and instance storage and controller card
- Nitro security chips - integrated into motherboard. Protects hardware resources
- Nitro hypervisor - lightweight hypervisor memory and cpu allocation bare metal like performance.

### Bare-metal instance

you can launch ec2 instance that have no hypervisor so you can run workloads directly on the hardware for maximum performance and control. The M5 and R% ec2 instances run are bare metal.

### bottlerocket -

It is a linux based open source operation system that is purpose-built by AWS for running containers on VIrtaul Machines or bare metal hosts.

### what is high performance computing (HPC)?

A cluster of hundreds of thousands of servers with fast connections between each of them with the purpose of boosting computing capacity.

### aws parallelCluster

It is an aws supported open source cluster management tool that makes it easy for you to deploy and manage high performance computing (HPC) clusters on AWS.
