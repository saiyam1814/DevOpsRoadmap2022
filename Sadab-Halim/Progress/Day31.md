# AWS Cloud

## AWS Documentation
AWS Documentation is a large collection of technical documentation on how to use aws services.

## Shared Responsibility Model
The Shared Responsibility Model is a **cloud security framework** that define the security obligations of the customer versa the cloud service provider (CSP) e.g. AWS

Each CSP has their own variant of the Shared Responsibility Model but they are all generally the same.

![image](https://user-images.githubusercontent.com/74575612/154994859-ccb39c6e-2b85-44fe-b51f-3b1ff3aae33d.png)

### NOTE:
The type of cloud deployment model and/or the scope of clourd service category can result in specialized Shared Responsibility Model.

## AWS Shared Responsibility Model

![image](https://user-images.githubusercontent.com/74575612/154995296-401afa00-a555-4627-9161-83f458b9f18f.png)

Customers are responsible for Security in the Cloud <br/>
AWS is responsible for Security of the Cloud

![image](https://user-images.githubusercontent.com/74575612/154995416-af79bb7b-aac9-474c-be6d-97a79e80f307.png)

## Types of Cloud Computing Responsibility

![image](https://user-images.githubusercontent.com/74575612/154997102-7e79ad6d-719f-461b-a0a3-510e3d07fb77.png)

## Shared Responsibility Model - Compute

![image](https://user-images.githubusercontent.com/74575612/154997261-501927fa-81fe-49c6-96b7-2ccfaa477638.png)

![image](https://user-images.githubusercontent.com/74575612/154997371-7d80b97c-f90c-4b26-afa2-bb0f2888a98b.png)

## Shared Responsibility Model Alternative
The **Shared Responsibility Model** is a simple visualization that helps determine what the customer is responsible for and what the CSP is responsible for related to AWS.

The customer is responsible for the data and the configuration of access controls that resides in AWS.

The customer is responsible for the configuration of cloud services and granting access to users via permissions.

CSP is generally responsible for the underlying Infrastructure.

**Responsibility in the cloud** <br/>
If you can configure or store it then you (the customer) are responsible for it.

**Responsibility of the cloud** <br/>
If you cannot configure it then the CSP is not responsible for it.

![image](https://user-images.githubusercontent.com/74575612/154998085-d99997de-6923-4af1-ad3b-c943aa34e9b1.png)

## Shared Responsibility Model Architecture
**Serverless** / Functions <br/>
No more servers, just worry about data and code

**Microservices** / Containers
Mix and match languages, better utilization of resources

**Traditional** / VMs
Global workforce is most familiar with this kind of architecture and lots of documentation, frameworks and support

![image](https://user-images.githubusercontent.com/74575612/154998438-4444e304-dabc-4684-9c97-6291e7c9aa78.png)


## Computing Services
**Elastic Compute Cloud (EC2)** allows you to launch Virtual machines (VM)

## Virtual Machines
A Virtual Machine (VM) is an emulation of a physical computer using software. Server virtualization allows you to easily create, copy, resize or migrate your server.
Multiple VMs can run on the same physical server so you can share the cost with other customers. <br/>

_Imagine if your server or computer was an executable file on your computer._

_When we launch a virtual machine we call it an instance_

EC2 is highly configurable server where you can choose **AMI** that affects options such as:
- The amount of CPUs
- The amount of Memory (RAM)
- The amount of Network Bandwidth
- The Operation System (OS) eg. Windows 10, Ubuntum Amazon Linux 2
- Attach multiple virtual hard-drives for storage eg. Elastic Block Store (EBS)

**AMI (Amazon Machine Image)** is a predefined configuration for a Virtual Machine

EC2 is considered the backbone of AWS because the majority of AWS services are using EC2 as their underlying servers. eg. S3, RDS, DynamoDB, Lambdas

## VMs, Containers and Serverless

**Virtual Machines:** an emulation of physical computer using software
- **Amazon Lightshell** it is the manged virtual server service. It is the friendly version of EC2 virtaul machine

_When you need to launch a Linux or Windows server but don't have much AWS knowledge eg. Launch a Wordpress_
 
**Containers:** virtaulizing an Operating System (OS) to run multiple workloads on a single OS instance. <br/>
Containers are generally used in micro-service architecture.
- **Elastic Container Service (ECS)** is a container orchestration service that support Docker containers. <br/>
  Launches a cluster of server(s) on EC2 instances with Docker installed. <br/>
  
  _When you need Docker as a Servicem or you need to run containers._
  
- **Elastic Container Registry (ECR)** is repository for container images in order to launch a container you need an image. An image just means a saved copy.
- **ECS Fargate** it is a serverless orchestration container service. It is the same as ECS expect you pay-on demand per running container. <br/>
  AWS manages the underlying server, so you don't have to scale or upgrade the EC2 server.
- **Elastic Kubernetes Service** is a fully managed kubernetes service. Kubernetes is an open-source orchestration software that was created by google and is generally the standard for managing microservices. <br/>

 _When you need to run Kubernetes as a service_
 
**Serverless:** When the underlying servers are managed by AWS. You don't worry or configure servers.
- **AWS Lambda** it is a serverless function service. you can run code without provisioning or managing or managing servers.

## High Performance Computing Services

### The Nitro System
**The Nitro System** a combination of dedicated hardware and lightweight hypervisor enabling faster innovation and enhanced security. All new ec2 instance types use the nitro system

- Nitro cards - specialize cards for VPC, EBS and instance storage and controller card
- Nitro Security Chips - integrated into motherboard. Protects hardware resources
- Nitro Hypervisor - lightweight hypervisor memory and cpu allocation bare metal like performance.

### Barel-Metal Instance
You can launch ec2 instance that have no hypervisor so you can run workloads directly on the hardware for maximum performance and control. The M5 and R% ec2 instances run are bare metal.

- **Bottlerocket** It is a linux based open source operation system that is purpose-built by AWS for running containers on VIrtaul Machines or bare metal hosts.

### High Performance Computing (HPC)
A cluster of hundreds of thousands of servers with fast connections between each of them with the purpose of boosting computing capacity.

- **AWS ParallelCluster** It is an aws supported open source cluster management tool that makes it easy for you to deploy and manage high performance computing (HPC) clusters on AWS.

Resources:
- [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)
