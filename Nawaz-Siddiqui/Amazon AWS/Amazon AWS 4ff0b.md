# Amazon AWS

What is a cloud provider?

A cloud service provider (CSP) is a company which:

Provides multiple cloud services

Those cloud services could be chained together to create cloud architectures

Those cloud services can be accessed via single unified API. eg AWS API

Have Infrastructure as a service (IAAS) offering

Offer automation via Infrastructure as code (IAC)

Common cloud services for Infrastructure as a Service (IAAS) would be:

Compute

Networking

Storage

Databases

Types of cloud computing:

SaaS (Software as a Service)

Its usually for customers. A product run and managed by the service provider. You don't have to worry about how the service is maintained. It works and remain available. 

For example: Gmail, Salesforce, Office 365 etc.

PaaS (Platform as a Service)

Its for developers. It focus on the deployment and management of your apps. You don't have to worry about provisioning , configuring, or understanding the hardware or OS.

For example: Heroku, elastic beanstalk,  Google app engine etc.

IaaS (Infrastructure as a Service)

Its for administrators. The basic building block for cloud IT. Provides access to networking features, computers and data storage space. You don’t have to worry about IT staff, data centers and hardware.

For example: Microsoft Azure, AWS, GCP etc.

Cloud Computing Deployment Models:

Public cloud

Private cloud

Hybrid cloud

Cross cloud

AWS Braket

Amazon Braket is a fully managed quantum computing service designed to **help speed up scientific research and software development for quantum computing.**

You can use quantum computing via AWS Braket.

Advantages of cloud

- **Trade fixed expense for variable expense** – Instead of having to invest heavily in data centers and servers before you know how you’re going to use them, you can pay only when you consume computing resources, and pay only for how much you consume.
- **Benefit from massive economies of scale** – By using cloud computing, you can achieve a lower variable cost than you can get on your own. Because usage from hundreds of thousands of customers is aggregated in the cloud, providers such as AWS can achieve higher economies of scale, which translates into lower pay as-you-go prices.
- **Stop guessing capacity** – Eliminate guessing on your infrastructure capacity needs. When you make a capacity decision prior to deploying an application, you often end up either sitting on expensive idle resources or dealing with limited capacity. With cloud computing, these problems go away. You can access as much or as little capacity as you need, and scale up and down as required with only a few minutes’ notice.
- **Increase speed and agility** – In a cloud computing environment, new IT resources are only a click away, which means that you reduce the time to make those resources available to your developers from weeks to just minutes. This results in a dramatic increase in agility for the organization, since the cost and time it takes to experiment and develop is significantly lower.
- **Stop spending money running and maintaining data centers** – Focus on projects that differentiate your business, not the infrastructure. Cloud computing lets you focus on your own customers, rather than on the heavy lifting of racking, stacking, and powering servers.
- **Go global in minutes** – Easily deploy your application in multiple regions around the world with just a few clicks. This means you can provide lower latency and a better experience for your customers at minimal cost.

Availability zones

An Availability Zone (AZ) is **one or more discrete data centers with redundant power, networking, and connectivity in an AWS Region.**

A region will generally have 3 AZ but isolated from each other, but close enough to provide low latency. All AZ are within 100km from each other.

Each availability zone is designed as an independent failure zone.

AZ’s are represented by region code, followed by a letter identifier. eg: us-east-1a

A subnet is associated with an Availability zone. You never choose the AZ when launching resources. You choose the subnet which associated to the AZ.

Fault Domain

Its a section of a network that is vulnerable to damage if a critical device or system fails.The purpose is if the failure occurs, it should not cascade outside that domain, limiting the damage possible.

You can have fault domains nested inside fault domains.

Fault level

A fault level is a collection of fault domains.

Edge location

A site that CloudFront uses to cache copies of your content for faster delivery to users at any location.

Points of Presence (POP)

Its an intermediate location between an AWS region and the end user, and this location could be datacenter or collection of hardware.

For AWS a point of presence is a data center owned by AWS or trusted partner that is utilized by AWS services related for content delivery or expedited upload.

PoP resources are:

Edge locations

Regional edge caches

AWS services using PoPs

Amazon cloudfront (CDN)

Amazon S3 transfer acceleration

AWS global accelerator

AWS direct connect

It is a private/dedicated connection between your datacenter, office, co-location and AWS.

It helps reduce network cost and increase bandwidth throughput, also a consistent network experience than a typical internet based connection.

Local zones

Local zones are datacenters located very close to a densely populated area to provide single-digit millisecond low latency performance(eg. 7ms for that area. 

Los Angeles, California was the first local zone to be deployed. Its identifier looks like: us-west-2-Iax-1a

Wavelength Zones

It allows for edge-computing on 5G networks. Applications will have ultra-low latency being as close as possible to the users.

Data residency

For workload that need to meet compliance boundaries strictly defining the data residency of data and cloud resources in AWS you can use:

AWS outposts

Its a physical rack of servers that you can put in your data center. Your data will reside whenever the outpost physically resides.

AWS config

It is policy as a code service. You create rules and when you deviate from it you get alerted or AWS config in some cases auto-remediate.

IAM policies

can be written explicitly deny access to specific AWS region. A Service Control Policy (SCP) are permissions applied organization-wide.

AWS Ground Station

AWS Ground Station is a fully managed service that lets you control satellite communications, process data, and scale your operations without having to worry about building or managing your own ground station infrastructure.

Graviton2

Graviton2 processors are **a type of server processor released by AWS in 2020**. They're a successor to the first generation of Graviton processors (simply called Graviton). They're custom built by AWS using 64-bit Arm architecture, with the goal of offering better performance across EC2 instances vs competitors.

Cloud architecture terminologies:

**High-Availability**

Ability of service to remain available all the time by ensuring no single point of failure and ensure a certain level of performance. Elastic load balancer is used to route the traffic to different available server to ensure availibility.

**High-Scalability**

Ability to increase up your capacity based on the high traffic demand by:

Vertical scaling

Horizontal scaling

**High-Elasticity**

Your ability to automatically increase or decrease your capacity based on the current demand of traffic, memory, and computing power. Can be done via **ASG (Auto scaling groups)**.

**Highly-Fault tolerance**

To ensure there is no single point of failure, preventing the chance of failure.

**Fail-overs** is when you plan to shift traffic to a redundant system in case the primary system fails.

**RDS multi-AZ** is when you run a duplicate standby database in other availability zones in case your primary database fails.

**High-durability**

Ability to recover from a disaster and to prevent the loss of data solutions that recovers from a disaster is known as disaster recovery(DR).

**CloudEndure Disaster Recovery** is an automated IT resilience solution that lets you recover your environment from unexpected infrastructure or application outages, data corruption, ransomware, or other malicious attacks. Block-level, continuous data replication enables recovery point objectives (RPOs) of seconds.

### **Business Continuity Plan (BCP)**

Its a document that outlines how a business will continue operating during unplanned disruption in service.

**Recovery point objective (RPO)**

The maximum acceptable amount of data loss after an unplanned data-loss incident, expressed as an amount of time.

**Recovery time objective (RTO)**

The maximum amount of downtime your business can tolerate without incurring a significant financial loss.

### **Disaster recovery options**

The preference is in the low to high order, first is the lowest.

**Backup and restore**

You backup your data and restore it to a new infrastructure

**Pilot light**

Data is replicated to another region with the minimal services running

**Warm Standby**

Scaled down copy of your infrastructure running ready to scale up

**Multi-site active/active**

Scaled up copy of your infrastructure in another region

**Powershell**

Powershell is a task automation and configuration management framework. A command-line shell and a scripting language. Generally used by people operating Azure or Microsoft workloads.

### **Amazon resource names (ARN)**

Uniquely identifies AWS resources. ARNs are required to specify a resource unambiguously across all of AWS.

example:

arn:partition:service:region:account-id:resource-type/resource-id

### AWS Software development kit (SDK)

An SDK is a collection of software development tools in one installable package.

You can use the AWS SDK to the programmatically create, modify, delete or interact with AWS resources. It is offered in various programming languages.

Cloud9 is the native IDE of AWS.

### AWS IAC

AWS has two offerings for IAC.

**AWS Cloudformation (CFN) :** A declarative tool supports JSON and YAML format.

**AWS Cloud Development Kit (CDK):** An imperative tool, supports languages like JAVA, Python, Nodejs, Typescript etc.

Both CFN and CDK have their own CLI.

### AWS Toolkit for VS Code

AWS toolkit is an open-source plugin for VS code to create, debug, deploy AWS resources like:

**AWS Explorer**

Explore a wide range of AWS resources to your linked AWS account.

**AWS CDK Explorer**

Allows you to explore your stacks defined by CDK.

**Amazon Elastic Container Service**

Provides IntelliSense for ECS task-definitions files.

**Serverless Applications**

Create, debug and deploy serverless applications via SAM and CFN.

### Access Keys

Access keys is a key and secret required to have programmatic access to AWS resources when interacting with the AWS API outside of the AWS management console. You can have two access keys at a given time.

Access keys are stored in ~/.aws/credentials and follow a TOML file format. You can store multiple access keys by giving the profile names.

### Shared responsibility model

It is a cloud security framework that defines the obligations of the customer versus the cloud service provider(CSP).

**Customers**

Customers are responsible for security in the cloud. i.e. data configuration.

**AWS**

AWS is responsible for the security of the cloud. i.e. Hardware, operations of managed services, and global infrastructure.