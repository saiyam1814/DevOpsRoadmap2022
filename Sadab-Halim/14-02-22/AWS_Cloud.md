# Continued Learning [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

## AWS Service Offerings

- Compute
- Storage
- Network Security
- Blockchain
- Machine Learning
- Artificial Intelligence
- Robot Development
- Video Production
- Orbital Satellites

## What is a Client-Server Model?
In computing a client can be web browser or desktop application that a person interacts with to make requests to computer servers. A server can serivces such as Amazon Elastic compute Cloud (Amazon EC2), a type of virtual server.

## Deployment models for cloud computing?

When selecting a cloud strategy, a company must consider factors such as required cloud application components, preferred resource management tools, and any legacy IT infrastructure requirements.

The three cloud computing deployment models are cloud-based, on-premises and hybrid

- Cloud based Deployment
  - run all parts of the application in the cloud
  - migrate existing applications to the cloud
  - design and build new applications in the cloud
  - In a cloud-baseddeployment model you can migrate existing applications to the cloud, or you can design and build new applications to the cloud, or you can design and build new applications in the cloud
- On-Premises Deployment
  - deploy resources by using virtualization and resource management tools.
  - iscrease resource uitlization by using application management and virtualization technologies.
  - on-premises deployment is also known as a private cloud deployment. In this model, respurces are deployed on premises by using virtualization and resource management tools.
  - for example you might have applications that run on technology that is fully kept in your on-premises data center. Though this model is much like legacy IT infrastructure, its incorporation of application management and virtualization technologies helps to increase resource utilization.
- hybrid deployment
  - connect cloud based resources to an-premises infrastructure
  - integrate cloud-based resources with legacy IT applications.
  - In a hybrid deployment, cloud based resources are connected to on-premises infrastructure. for eg, you have legacy applications that are better maintained on premises, or government regulations require your business to keep certain records on premises.

## Benefits of cloud computing

consider why a company might choose to take a particular cloud computing approach when addressing business needs.

### trade upfront expense for variable expense

upfront expense refers to data centers, physical servers, and other resources that you would need to invest in before using them. Variable expense means you only pay for computing resources you consume instead heavily in data centers and servers before you know you're going to use them.

### Stop spending money to run and maintain data centers

computing in data centers often reuqires you to spend more money and time managing infrastructure

a benefit of cloud computing is the ability to focus less on these tasks and more on your application and customers.

### Stop guessing capacity

With cloud computing, you don't have to predict how much infrastructure capacity you will need before deploying an application.

### Benefit from massive economies of scale

By using cloud computing, you can achieve a lower variable cose that you can get on you own.

Because usage from hundreds of thousands of customers can aggregate in the cloud, can achieve higher economies of scale. The economu of scale transalate into lower pay-as-you-go prices.

### Increase speed and agility

The flexibility of cloud computing makes it easier for you to develop and deploy applications.

The flexibility provides you with more time to expirement and innovate.

### Go global in minutes

The global footprint of the AWS cloud enables you to deploy applications to customers around the world quickly, while providing them with low latency.

### Cloud Computing models

- Infrastructure as a Service (IaaS)
  - infrastructure as a service, sometimes abbreviated as IaaS, contains the basic building blocks for cloud IT and typically provide access to networking features.
- Platform as a service (Paas)
  - Platforms as a service remove the need for organization to manage underlying infrastructure and allow you to focus on the deployment and management of your applicaiton.
- Software as a Service(Saas)
  - Software as a service provides you with a completed product that is run and managed by the service provider.

## Amazon Elastic Compute Cloud (Amazon EC2)

amazon ec2 provides secure, resizable compute capacity in the cloud as amazon ec2 instances.

- you can provision and launch EC2 instance within minutes
- you can stop using it when you have finished running a workload
- you pay only for the compute time you use when an instance is running, not when it is stopped or terminated
- you can save costs by paying only for server capacity that you need or want.

## How amazon ec2 works

1. launch
1. connect
1. use

## amazon ec2 isntance types

amazon ec2 instance types are optimized for different tasks. When selecting an instance type, consider the specific needs of your workloads and applications. This might include requirements for compute, memeory, or storage capabilities

### General purpose instance

general purpose instances provide a balance of compute, memory, and networking resources. You can use them for a variety of workloads,

- application servers
- gaming servers
- backend servers for enterprise applications
- small and medium databases.
  suppose that you have an application in which the resource needs for compute, memory, and networking are roughly equivalent. You might consider running it on a general purpose instance because the application does not require optimization in any single resource area.

### Compute optimized instances

compute optimized instances are ideal for compute-bound applications that benefit from high-performance processors.

The difference is compute optimized applications are ideal for high-performance web-servers, compute-intensive applciations servers, and dedicated gaming servers. You can also use compute optimized instances for batch processing workloads that require processing many transactions in a single group.

### Memory optimized instances

memory optimized instances are designed to deliver fast performance for workloads that process large datasets in memory. In computing, memory is a temporary storage area. It holds all the data and instructions that a central processing unit.

Memoru optimized instances enable you to run workloads with high memory needs and receive great performance.

### accelerated computing instances

accelerated computing instances use hardware accelerators, or coprocessors to perform some functions more efficiently than is possible in software running on CPUs.

In computing, a hardware accelerator is a component that can expedite data processing. Accelerated computing instance are ideal for workloads such as graphics applications, game streaming, and application streaming

### storage optimized instances

storage optimized instances are designed for workloads that require high, sequential read and write access to large data sets on local storage.

In computing the term input/output operations per seconds (IOPS) is a metric that measures the performance of a storage device. It indicates how many different input or output operations a device can perform in one second.

## amazon ec2 pricing

with amazon ec2 you pay only for the compute time that you use. Amazon ec2 offers a variety of pricing options for different use cases.

### On-demand

on-demand instance are ideal for short-term, irregular workloads that cannot be interrupted. No upfront costs or minimum contracts apply. The instance run continuously until you stop them, and you pay for only the compute time you use.

sample use cases for on-demand instances include developing and testing applications and running applications that have unpredictabel usage patterns. On-demand instances are not recommended for workloads that last a year or longer because these workloads can experience greater cost savings using reserved instances.

### amazon ec2 savings plan

aws offers savings plan for several compute services, including amazon ec2. Amazon ec2 savings plan enable you to reduce your compute costs by commiting a consistent amount of compute usage for a 1-year or 3-year term. This term commitment results in savings of upto 66% over 0n-demand costs.

Any usage beyond the commitment is charged at regular on-demand rates.

### reserved instances

reserved instance are a billing discount applied to the use of on-demand instances in your account. You can purchase standard reserved and convertible reserved instances for a 1-year or 3-year term.

at the end of a reserved instance term, you can continue using the amazon ec2 instance without iterruption. However, you are charged on-demand rated until you do one of the following

- terminate the instance
- purchase a new reserved instance that matches the instance attributes.

### spot instances

spot instances are ideal for workloads with flexible start and end times, or that can withstand interruptions. Spot instances use unused amazon ec2 computing capacity and offer you cost savings at up to 90% off of on-demand prices.

Suppose that you have a background processing job that can start and stop as needed.

After you have launched a spot instance, if capacity is no longer available or demand for spot instances increases, your instance may be interrupted. This might mot pose any issues for your background processing job.

### dedicated hosts

dedicated hosts are physical servers with amazon ec2 instance capacity that is fully dedicated to your use.

You can use your existing per-socket, per-core, or per-VM software licenses to help maintain license complaince. You can purchase on-demand dedicated hosts and dedicated reservations. Of all the amazon ec2 options that were covered, dedicated hosts are the most expensive.

## Scalability

scalability involves begining with only the resources as you need and designing your architecture to automatically respond to changing demand by scaling out or in. As a result you pay only the resources you use.

If you want the scaling process to happen automatically, which aws service would you use? The aws service that provides this functionality for amazon ec2 instance is amazon ec2 auto scaling.

## amazon ec2 auto scaling

amazon ec2 auto scaling enables you to automatically add or remove amazon ec2 instances in response to changing application demand. By automatically scaling your instances in and out as needed.

withing amazon ec2 auto scaling, you can use two approaches

- dynamic scaling responds to changing demand
- predictive scaling automatically schedules the right number of amazon ec2 instance based on predictive demand.

_to scale faster, you can use dynamic scaling and predictive scaling together._

## example amazon ec2 auto scaling

suppose that you are preparing to launch an application on amazon ec2 instances. When configuring the size of your auto scaling gorup, you might set the minimum number of amazon ec2 instance at one. This means that at all times, there must be at least one amazon ec2 instance running.

when you create an auto scaling group, you can set the minimum capacity is the number of amazon ec2 instances that launch immediately after you have created the auto scaling group.

_if you don't specify the desired number of amazon ec2 instance in an auto scaling group, the desired capacity defaults to your minimum capacity._

Next you can set the desired capacity at two amazon ec2 instances even though your application nees a minimum of a single amazon ec2 instance to run.

The third configuration that you can set in an auto scaling group is the maximum capacity.

## Zero trust model

The zero trust model operates on the principal of "trust no one, verify everything"

Malicious actors being able to by-pass conventional access controls demonstrates traditional security measures are no longer sufficient.

In zero trust model Identity becomes the primary security perimeter.

What is the primary security perimeter?
The primary or new security perimeter defines the first line of defense and its security controls that protect a company's cloud resources and assets.

## Zero Trust on aws

Identity security controls you can implement on AWS to meet the zero trus model

- IAM policies
- Permission boundaries
- service control policies (organization-wide policies)
- IAM policy conditions
  - aws:SourceIp - Restrict on IP address
  - AWS:RequestedRegion - Restrict on Region
  - aws:MultifactorAuthPresent - restrict if MFA is turned off
  - aws: currentTime - restrict access based on time of day

## zero trust on AWS with third parties

aws does technically implement a zero trust model but does not allow for onteligent identity security controls.

for eg azure active directory has Real tiem and calculated risk detection based more data points than aws

## Directory service

what is directory service?\
A directory service maps the names of network resouces to their network addresses.

A directory service is shared information infrastructure for locating, managing, administering and organizing resources.

A directory service is a critical component of a network operating system. A directory service is a server which provides a directory service.

## Identity providers (IdPs)

Identity Provider (IdP) a system entity that creates, maintains and manages identity information for principals and also provides authentication services to applications within a federation or distributed network. A trusted provider of your identity that lets you use authenticate to access other services. Identity providers could be facebook, amazon, google, twitter, github, lindkIn

