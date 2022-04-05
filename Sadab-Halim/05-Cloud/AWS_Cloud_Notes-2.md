# AWS Cloud

## AWS for Government
**What is Public Sector?** <br/>
Public sectors include public goods and governmental services such as:

- military
- public education
- law enforcement
- health care
- infrastructure
- the government itself
- public transit

AWS can be utilized by public sector or organizations developing cloud workloakds for the public sector.

AWS achieves this by meeting regulatory compliance programs along with the specific governance and security controls.

AWS has special regions for US regulation called GovCloud

### GovCloud (US)
**Federal Risk and Authorization Management Program (FedRAMP)** A US government-wide program that provides a standardized approach to security assessment, authorization, and continuous monitoring for cloud products and services.

**What is GovCloud?**
A Cloud Service Provider (CSP) generally will offer an isolated region to run FedRAMP workloads.

AWS GovCloud Regions allows customers to host sensiticve Controleed Unclassified Information and other tyoes of regulated workloads.
- GovCloud Regions are only operated by employess who are U.S. citizends, U.S. soil.
- They are only accessible to U.S. entities and root accound holders who pass a screening process

Customers can architect secure cloud solutions that comply with:
- FedRAMP High baseline
- DOJ's Criminal Justice Information Systems (CJIS) Security Policy
- U.S. International Traffic in Arms Regulations (ITAR)
- Export Administration Regulation (EAR)
- Department of Defense (DoD) Cloud Computing Sercurity Requirements Guide

### Sustainability
AWS Cloud's Sustainability goals are composed of three parts:
1. **Renewable Energy** <br/>
  AWS is working towards having their AWS Global Infrastructur powered by 100% renewableenergy by 2025.
  
  AWS purchases and retires environmental attributes to cover the non-renewable energy for AWS Globaal Infrastructure:
  - Renewable Energy Credits (RECs)
  - Gurantee of Orgin (GOs)
  
2. **Cloud Efficiency** <br/>
  AWS's Infrastructure is 3.6 times more energy thsn thr median of U.S. enterprise data centers surveyed.
  
3. **Water Stewardship** <br/>
  Direct evaporate technology to cool our data center. <br/>
  Use of non-portable water for cooling purposes (recycled water) <br/>
  On-site water treatment allows us to remoce scale-forming minerals and reuse water for more cycle. <br/>
Water efficient metrics determings and monitor optimal water use for each AWS region.

# AWS Cloud

## AWS Ground Station
**AWS Ground Station** is a fully managed service that lets you control satellite communications, process data, and scale your operations without having to worry about buulding or managing your own ground station infrastructure.

Use Cases for Ground Station: <br/>
- weather forecasting
- surface imaging
- communications
- video broadcasts

To use Ground Station: <br/>
- You schedule a COntact (select satellite, start and end time, and the ground location).
- Use the AWS Ground Station EC2 AMI to launch EC2 instances that will uplink and downlink data during the contact or receive downlinked data in an Amazon S3 bucket.

Use Case Example: <br/>
A company reaches an agreement with a Satellite Imagery Provider to take satellite photos of a specifi region. They use AWS Ground Station to communicate that company's Satellite and download the S3 image data.

## AWS Outposts
**AWS Outposts** is a fully managed service that offers the same AWS infrastrcuture, AWS services, APIs and tools to virtually any datacenter, co-location space, or on-premises facility for a truly consistent hybrid experience.

AWS Outposts is rack of servers running AWS Infrastructure on your physical location.

**What is a Server Rack?** <br/>
A frame designed to hold and organize IT equipment

**Rack Heights** <br/>
U stands for "rack units" or "U spaces" with is equal to 1.75 inches. <br/>
The industry standard rack size is 48U (7 Foot Rack).

### AWS Outposts comes in 3 form factors: 42U, 1U and 2U
42U:This is a full rack of servers provided by AWS <br/>

![image](https://user-images.githubusercontent.com/74575612/152841200-9922b71d-ebff-4897-a27d-95a0aa2a3dd5.png)

AWS delivers it to your preferred physical site fully assembled and ready to be rolled into final position. <br/>
It is installed by AWS and the rack needs to be simply plugges into power and network.

1U and 2U are servers that you can place into your existing racks:
1U <br/>
Suitable for 19-inch wide 24-inch deep cabinets <br/>
AWS Graviton2 (up to 64 vCPUs) <br/>
128 GiB memory <br/>
4 TB of local NVMe storage

2U <br/>
Suitable for 19-inch wide <br/>
36-inch deep cabinets <br/>
Intel processor (up to 128 vCPUs) <br/>
256 GiB memory <br/>
8 TB of local NVMe storage

# AWS Cloud

## Cloud Architecture Technologies

**What is Solutions Archtecht?** <br/>
A role in a technical organization that architects a technical solution using multiple systems via researching, documentation, experimentation

**What is a Cloud Architect?** <br/>
A solutions architect that is focused solely on architecting technical solutions using cloud services.

A cloud architect need to understand the following terms and factor them into their designed and architecture based on the business requirements.
- **Availability:** Your ability to ensure a service remains available eg. Highly Available (HA)
- **Scalability:** Your ability to grow rapidly or unimpeded
- **Elasticity:** Your ability to shrink and grow to meet the demand
- **Fault Tolerance:** Your ability to prevent a failure
- **Disaster Recovery:** Your ability to recover from a failure eg. Highly Durable (DR)

A Solutions Archtect needs to always consider the following business factors: <br/>
- Security: How secure is this solution?
- Cost: How much is this going to cost?

## High Availability
Your ability for your service to remain available by ensuring there is no single point of failure and/or ensure a certain level of performance.

**Elastic Load Balancer:** <br/>
A load balancer allows you to evenly distribute traffic to multiple servers in one or more datacenter. If a datacenter or server becomes unavailable (unhealthy) the load balancer will route the traffic to only availabe datacentes with servers.

![image](https://user-images.githubusercontent.com/74575612/152843787-85fd24d2-cc5f-4002-8308-e8a2c206c790.png)

Running your workload across multiple **Availability Zones** ensures that if 1 or 2 AZs become unavailable your service / applications remain available.

## High Scalability
Your ability to increase your capacity based on the increasing demand of traffic, memory and computing power.

![image](https://user-images.githubusercontent.com/74575612/152844132-4d16104f-0e37-4959-9ef1-ac65fbc67ce7.png)

## High Elasticity
Your ability to automatically increase or decrease your capacity based on the current demand of traffic, memory and computing power.

**Auto Scaling Groups (ASG)** is an AWS feature that will automatically add or remore servers based on scaling rules you define based on metrics.

**Horizontal Scaling** <br/>
Scaling Out - Add more servers of the same size <br/>
Scaling In - Removing underutilzied servers of the same data <br/>

Vertical Scaling is generally hard for traditional architecture so you'll usually only see horizontal scaling described with Elasticity.

## Highly Fault Tolerant
Your ability for you service to ensure there is no single point of failure. Preventing the chance of failure.

**Fail-overs** is when you have a plan to shift traffic to a redundant system in case the primary system fails.

![image](https://user-images.githubusercontent.com/74575612/152844931-bf72133f-b122-4ebf-b1ea-b10b726033ad.png)

A common example is having a copy (secondary) of your databas where all ongoing changes are synced. The secondary system is not in-use until a fail over occurs and it becomes the primary databse.

**RDS Multi-AZ** is when you run a duplicate standby database in another Availability Zone in case your primary database fails.

## High Durability
Your ability to recover from a disaster and to prevent the loss of data Solutions that recover from a disaster is known as **Disaster Recover (DR)**
- Do you need a backup?
- How fast can you restore that backup?
- Does your backup still work?
- How do you ensure current live data is not corrupt?

**CloudEndure Disaster Recovery** continiously replicates your machines into a low-cost staging area in your target AWS account and preferred Region enabling fast and reliable recover in case of IT data center failutes.

## Business Continuity Plan (BCP)
A **business continuity plan** (BCP) is a document that outlines how a business will continue operating during an unplanned disruption in services.

**Recovery Point Objective (RPO)** the maximum acceptable amount of data loss after an unplanned data-loss incident, expressed as an amount of time.

**Recovery Time Objective (RTO)** the maximum amount of downtime your business can tolerate without incurring a significant finacial loss

![image](https://user-images.githubusercontent.com/74575612/152846335-df70e4b9-f190-40eb-87f6-fc821c4e66e4.png)

## Disaster Recovery Options
There are multiple options for recovery that trade cist vs time to recover

![image](https://user-images.githubusercontent.com/74575612/152846602-ee3f23f3-8616-4a1a-8a28-7dc7165d0068.png)

## RTO
**Recovery Time Objective (RTO)** is the maximum acceptable delay between the interruption of service and restoration of service. This objective determines what is considered an acceptable time window when service in unavailable and is defined by the organization.

![image](https://user-images.githubusercontent.com/74575612/152846929-220d8ec1-22bf-4dd1-9a3f-9342d6696ebc.png)

## RPO
**Recovery Point Objective (RPO)** is the maximum acceptable amount of time since the last data recovery point. The objective determines what is considered an acceptable loss of data between the last recovery and the interruption of service and is defined by the organisation.

![image](https://user-images.githubusercontent.com/74575612/152847187-1cde98ca-3d35-4275-9e90-e02667b8b114.png)

# AWS Cloud

## AWS Application Programming Interface (API)
**What is an Application Programming Interface (API)?** <br/>
An API is software that allows two applications/services to talk to each other. <br/>
The most common type of API is via HTTP/S requests.

AWS API is an HTTP API and you can interact by sending HTTPS requests, using an application interacting with APIs like **Postman.**

Each AWS Service has its own **Service Endpoint** which you send requests
```
GET/HTTP/1.1
host: **monitoring.us-east-1.amazonaws.com**
x-amz-target: GraniteServiceVersion20100801.GetMetricData
x-amz-date: 20180112T092034Z
Authorization: **AWS4-HMAC-SHA256 Credential=REDACTEDREDACTED/20180411/...**
Content-Type: application/json
Content-Encoding: amz-1.0
Content-Length: 45
Connection: keep-alive
```

To authorize you will need genereated signed requests <br/>
You make a separate request with your AWS credentials and get back a token.

You need to also provide an **ACTION** and accompanying **parameters** as the payload

Rarely do users directly send HTTP requests directly to the AWS API. <br/>
Its much easier to interact with the API via a variety of Developer Tools

![image](https://user-images.githubusercontent.com/74575612/153446210-7e1d9a58-d3e5-426a-91d9-b1a0ecaa81d7.png)

## AWS Management Console
The AWS Management Console is a web-based unified console <br/>
**Build, manage, and monitor everything** from simple web apps to complex cloud deployments.

Point and Click to manually launch and configure AWS resources with limited programming knowledge.

This is known as **"ClipOps"** since you can perform all your system operations via clicks.

## AWS Management Console -- Service Console
AWS Service each have their own customized console. <br/>
You can access these consoles by **searching** the service name.

![image](https://user-images.githubusercontent.com/74575612/153447258-373ccfe5-f8a2-4c24-b47f-df160656f759.png)

The EC2 Console
![image](https://user-images.githubusercontent.com/74575612/153447341-3d36fdbb-1223-4daa-bcd7-c500418f8f2b.png)

Some AWS Services Console will act as an umbrella console containing many AWS Services: eg. <br/>
- VPC Console
- EC2 Console
- Systems Manager Console
- SageMaker Console
- CloudWatch Console

### AWS Account ID
Every AWS Account has a unique Account. <br/>
The **Account ID** can be easily found by dropping down the current user in the Global Navigation.

![image](https://user-images.githubusercontent.com/74575612/153448439-f8531e78-193e-4a3a-b48a-c359718a57e3.png)

The AWS Account ID id composed of 12 digits eg:
- 123456789012
- 121212121212
- 498241098510

![image](https://user-images.githubusercontent.com/74575612/153448506-b2632d62-2022-4bad-b870-bdf915b6645d.png)

The AWS Account ID is used: 
- when logging in with a non-root user account.
- Cross-account roles
- Support cases

## AWS Tools for PowerShell
**What is PowerShell?** <br/>
**PowerShell** is a task automation and configuration management framework. <br/>
A command-line shell and a scripting language.

Unlike most shells, which accept and return text, PowerShell is built on top of the .NET Common Language Runtime (CLR), and accepts and returns .NET objects.

**AWS Tools for PowerShell** lets you interact with the AWS API via PowerShell Cmdlets

Cmdlet is a special type of command in PowerShell in the form of capitalized verb-and-noun e.g. _New-S3Bucket_

## Amazon Resoure Name (ARNs)
**Amazon Resource Names (ARNs)** uniquely identify AWS resources. <br/>
ARNs are required to specify a resource unambigoudly across all of AWS

The ARN has the following **format variations:**
- arn: _partition:service:region:account-id:resource-id_
- arm: _partition:service:region:account-id:resource-type/resource-id_
- arn: _partition:service:region:account-id:resource-type:resource-id_

Partition
- aws - AWS Regions
- aws-cn - China Regions
- aws-us-gov - AWS GovCloud (US) Regions

Service - Identifies the service
- EC2
- S3
- IAM

Region - which AWS resource
- us-east-1
- ca-central-1

Account ID
- 121212121212
- 123456789012

Resource ID <br/>
Could be a number name or path:
- user/Bob
- instance/i-1234567890abcdef0

In the AWS Management COnsole its common to be able to copy the ARN to your clipboard.

## Path in ARNs
Resource ARNs can include a path <br/>
Paths can include a wildcard character, namely an asterisk (*)

**IAM Policy ARN Path** <br/>
arn:aws:iam::123456789012:user/Development/product_1234/*

**S3 ARN Path** <br/>
arn:aws:s3:::my_corportate_bucket/Development/*

#  AWS Cloud

## AWS Command Line Interface (CLI)

**What is a CLI?** <br/>
A Command Line Interface (CLI) processes commands to a computer program in the form of lines of text. <br/>
Operating system implement a command-line interface in a shell.

**What is a Terminal?** <br/>
A terminal is a text only interface (input/output environment)

**What is a Console** <br/>
A Console is a physical computer to physically input information into a terminal.

**What is a Shell?** <br/>
A shell is the command line program that users interact with to input commands. <br/>
Popular shell programs:
- Bash
- Zsh
- PowerShell

AWS Command Line Interface (CLI) allows users to programtically interact with the AWS API via entering single or multi-line commands into a shell or terminal.

![image](https://user-images.githubusercontent.com/74575612/153470547-e3467b80-523a-434a-a363-2b3d6e816892.png)

The AWS CLI is a Python executable program.
- Python is required to install AWS CLI <br/>
The AWS CLI can be installed on Windows, Mac or Linux/Unix. <br/>
The name of the CLI program is aws

## AWS Software Development Kit (SDK)
A Software Development Kit (SDK) is a collection of software development tools in one installable package.

You can use the AWS SDK to programtically create, modify, delete or interact with AWS resources.

AWS SDK is offered in various programming languages:
- Java
- Python
- Node.js
- Ruby
- Go
- .NET
- PHP
- JavaScript
- C++

## AWS CloudShell
AWS CloudShell is a browser-based shell built into the AWS Management Console. <br/>
AWS CloudShell is scoped per region, Same credentials as logged in user. Free Service

**Preinstalled Tools** <br/>
AWS CLI, Python, Node.js, git, make, pip, sudo, tar, tmux, vim, wget, and zip and more

**Storage included** <br/>
1 GB of storage free per AWS region

**Saved files and settings** <br/>
Files saved in you home directory are available in future sessions for the same AWS region

**Shell Environments** <br/>
Seamlessly switch between
- Bash
- Powershell
- Zsh

## Infrastructure as Code (IaC)
You write a configuration script to automate creating, updating or destroying cloud infrastructure.
- IaC is a blueprint of your infrastructure.
- IaC allows you to easily share version or inventory your cloud infrastructure.

AWS has two offerings for writing infrastrucutre as a code.

**AWS CloudFormation (CFN)**
CFN is a Declarative IaC tool

**Declarative**
- What you see is what you get. **_Explicit_**
- More verbose, but zero chance of mis-configuration
- Uses scripting languages eg. JSON, YAML, XML

**AWS Cloud Development Kit (CDK)**
CDK is an Imperative IaC tool.

**Imperative**
- You say what you want, and the rest is filled in. **_Implicit_**
- Less verbose you could end up with misconfiguration
- ```
- Does more than Declarative
- Uses programming languages eg. Python, Ruby & JavaScript.

## CloudFormation
AWS CloudFormation allows you to write Infrastructure as Code (IaC) as either a JSON or YAML file.

CloudFormation is simple but it can lead to large files or is limited in some regard to creating dynamic or repeatable infrastructure compared to CDK.

CloudFormation can be easier for DevOps Engineers who do not have a background in web programming languages.

Since CDK generates out CloudFormation its sill important to be able to read and understand CloudFormation in order to debug Iac stacks.

![image](https://user-images.githubusercontent.com/74575612/153473889-1c5df7f1-88e1-4699-93ea-81085c71113b.png)

# AWS Cloud

## Cloud Development Kit
AWS CDK allows you to use your favourite programming language to write Infrastructure as Code (IaC) <br/>
TypeScript, NodeJS, Python, Java, ASP.NET

- CDK is powered by CloudFormation (it generates out CloudFormation templates)
- CDK has a large library of reusable cloud components called CDK Construct
- CDK comes with it own CLI
- CDK Pipelines to quickly setup CI/CD pipelines for CDK projects
- CDK has a testing framework for Unit and Integration Testing

AWS SDK looks similar, but the key difference is CDK ensures Idempotent of your Infrastructure.

## AWS Toolkit for VSCode
AWS Toolkit is an open-source plugin for VSCode to create, debug, deploy AWS resources

1. **AWS Explorer**
  Explore a wide range of AWS resources to your linked AWS account.
2. **AWS CDK Explorer**
  Allows you to explore your stacks defined by CDK
3. **Amazon Elastic Container Service**
  Provides IntelliSense for ECS task-definition files
4. **Serverless Applications**
  Create, debug and deploy serverless applications via SAM and CFN

## Access Keys
Access Keys is a key and secret required to have programmatic access to AWS resources when interacting with the AWS API outside of the AWS Mangement Console

An Access Key is commonly referred to as AWS Credentials

A user must be granted access to use Access Keys
  
Access keys are to be stored in `~/.aws/credentials` and follow a TOML file format

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

