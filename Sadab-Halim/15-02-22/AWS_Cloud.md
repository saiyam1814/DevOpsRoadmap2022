# Directing traffic with elastic load balancing

## Elastic Load Balancing

elastic load balancing is the aws service that automatically distributes incoming application traffic across multiple resources, such as amazon ec2 instances.

A load balancer acts as a single point of contact for all incoming web traffic to your auto scaling group. This means that as you add or remove amazon ec2 instances in response to the amount of incoming traffic, these requests route to the load balancer first. Then the requests spread across multiple resources that will handle them, For eg, if you have multiple amazon ec2 instances, elastic load balancing distributes the workload across the multiple instances so that no single instance has to carry the bulk of it.

Although elastic load balancing and amazon ec2 auto scaling are separate services, they work together to help ensure that applications running in amazon ec2 can provide high performance and availability

### Low-Demand Period

suppose that a few customers have come to the coffee shop and are ready to place their orders.

if only a few registers are open, this matches the demand of customers who need service.

### High Demand Period

throughout the day, as the number of customers increases, the coffee shop opens more registers to accomodate them, in the diagram the auto scaling group represents this.

# Messaging and Queueing

## Monolithic applications and microservices

applications are mode of multiple components. The components communicate with each ohter to transmit data, fulfill requrests, and keep the application running.

suppose that you have an application with tightly coupled components. These components might include databases, servers, the user interface, business logic so on. This type of architecture can be considered a monolithic application

In this approach to application architecture, if a single component fails, other components fail, and possibly the entire application fails.

## Microservices

In a microservices approach, application components are loosely coupled, In this case, if a single component fails, the other components continue to work because they are communicating with each other, The loose coupling prevents the entire application from failing

## amazon simple notification service (amazon sns)

amazon sns is a publish/subscribe service. Using amazon sns topics, a publisher publishes messages to subscribers. This is similar to the coffee shop; the cashier provides coffee orders to the barista who makes the drinks.

## Amazon Simple Queue Service (Amazon SQS)
Using amazon sqs you can send, store and receive messages between software components, without losing messages or requiring other services to be available. In amazon SQS an application sends messages into a queue, A user or service retrieves a message from the queue, processes it and then deletes it from the queue.

# Additional Compute Services

## Serverless Computing

Earlier in this module, you learned about amazon ec2, a service that lets you run virtaul servers in the cloud. If you have applications that you want to run in amazon ec2

1. provision instances (virtual servers)
1. upload your code
1. continue to manage the instances while your application is running.

The term 'serverless' means that your code runs on servers, but you do not need to provision or manage these servers. With serverless computing, you can focus more on innovating new products and features instead of maintaining servers

another benefit of serverless computing is the flexibility to scale serverless applications automatically. Serverless computing can adjust the applications capacity by modifying the units of consumptions, such as throughput and memory.

An aws service for serverless computing is aws lambda.

## AWS Lambda
AWS Lambda is a service that lets you run code without needing to provision or manage servers.

while using aws lambda, you pay only for the compute time that you consume. Charges apply only when your code is running. You can also run code for virtually any type of application or backend service, all with zero administration.

1. upload your code to lambda
1. set your code to trigger from an event source, such as aws services, mobile applications, or HTTP endpoints.
1. lambda runs your code only when triggered.
1. you pay only for the compute time that you use. In the previous example of resizing images, you would pay only for the compute time that you use when uploading new images. Uploading images triggers lambda to run code for the image resizing function.

**In AWS, you can also build and run containerized applications**

## Containers

containers provide you with a standard way to package your application's code and dependencies into a single object. You can also containers for processors and workflows in which there are essential requirements for security, reliability and scalability

container orchestration services help you to deploy manage, and scale your containerized applications. Next you will learn about two services that provide container orchestration. Amazon elastic container service and amazon elastic kubernetes service.

## Amazon Elastic Container Service (Amaozn ECS)
Amazon ECS is highly scalable, high-performance container management system that enables you to run and scale containerized application on aws.

Amazon ECS supports docker containers. docker is a software platform that enables you to build, test, and deploy applications quickly. AWS supports the use of open-source docker community edition and subscription-based docker enterprise edition.

## Amazon Elastic Kubernetes Service (Amazon EKS)
Amazon EKS is a fully managed service that you can use to run kubernetes on aws.

kubernetes is open-source software that enables you to deploy and manage containerized applications at scale.

## AWS Fargate
AWS Fargate is a serverless compute engine for containers. It works with both amazon ecs and amazon eks.

When using aws fargate, you do not need to provision or manage servers. AWS fargate manages your server infrastructure for you. you can focus more on innovating and developing your applications, and you pay only for the resources that are required to run your containers.

# AWS Global Infrastructure

## Selecting a region

when determining the right region for your services, data, and applications, consider the following four business factors.

### Compliance with data governance and legal requirements

depending on your company and location you might need to run you data of specific areas. for eg if your company requires all of its data to reside within the boundaries of the UK. you should choose the london region

### Proximity to your customers

selecting a region that is close to your customers will help you to get content to them fasters.

### Availble Services within a region

sometimes the closest region might not have all the features that you want offer customers. AWS is frequently innovating by creating new services and expanding on features within existing services.

### Pricing

suppose that you are considering running applications in both the united states and brazil. the way brazil tax structure is set up, it might cost 50% more to run the same workload out of the region compared to the other region.

## Availability Zones
an availability zone is a single data center or a group of data centers within a region. availability zones are located tens of miles apart from each other. this is close enough to have low latency between availability zones. however if a disaster occurs in one part of the region, they are distant enough to reduce the chance that multiple availability zones are affected.

# Edge Locations
An Edge location is a site that amazon cloudfront uses to store cached copies of your content closer to your customers for faster delivery.

Amazon uses cdn as services as cloudfront

# How to provision aws resources

## Ways to interact with aws services

### AWS Management Console
The aws management console is a web-based interface for accessing and managing aws services. you can quickly access recently used services and search for other services by name, keyword, or acronym. the console include wizards and automated workflows that can simplify the process fo completing tasks.

### AWS Command Line Interface
To save time when making api requrests you can use the aws command line interface (aws cli). aws cli enables you to control multiple aws services directly from the command line within one tool.

by using aws cli you can automate the actions that your service and applications perform through scripts.

### Software Development Kits
Another option for accessing and managing aws services is the software development kits (SDKs). sdks make it easier for you to use aws services through an api designed for your programming language or platform. sdks enable you to use aws services with your existing applications or create entirely new applications that will run on aws.

## AWS Elastic Beanstalk
With aws elastic beanstalk, you provide code and configuration settings, and elastic beanstalk deploys the resources necessary to perform the following tasks.

- Adjust capacity
- Load balancing
- Automatic scaling
- Application health monitoring

## AWS CloudFormation
With aws cloudformation you can treat your infrastructure as code. this means that you can build an evironment by writing lines of code instead of using the aws management console to invdividually provision resources.

AWS CloudFormation provisions your resources is a safe, repeatable manner, enabling you to frequently build your infrastructure and applications without having to perform manual actions. it determines the right operations to perform when managing your stack and rolls back changes automatically if it detects errors.

# Connectivity to AWS

## Amazon Virtual Private Cloud (Amazon VPC)
A networking service that you can use to establish boundaries around your aws resources is amazon virtual private cloud (amazon vpc)

Amazon VPC enables you to provision an isolated section of the aws cloud. in this isolated section, you can launch resources in a virtual network that you define. within a virtual private cloud (vpc) you can organize your resources into subnets. a subnets is a section of a vpc that can contain resources such as amazon ec2 instances.

## Internet Gateway
To allow public traffic from the internet of access your vpc, you attach an internet gateway to the vpc.

An internet gateway is a connection between a vpc and the internet. you can think of an internet gateway as being similar to a doorway that customers use to enter the coffee shop. wihtout internet gateway no one can access the resources within your vpc.

## Virtual Private Gateway
To access private resources in a vpc, you can use a virtual private gateway.

The virtual private gateway is the component that allows protected internet traffic to enter into the vpc.

A virtaul private gateway enables you to establish a virtual private network (VPN) connection between your vpc and a private network, such as on-premises data center or internal corporate network. a virtual private gateway allows traffic into the vpc only if it is coming from an approved network.

## AWS Direct Connect
AWS Direct Connect is a service that enables you to establish a dedicated private connection between your data center and vpc.

The private connection that aws direct connect provides helps you to reduce network costs and increase the amount of bandwidth that can travel through your network.

## Subnets
A subnet is a section of a vpc in which you can group resources based on security or operational needs. subnets can be public or private.

### Public Subnets
Public Subnets contain resources that need to be accessible by the public, such as an online store's website.

### Private Subnets
Private Subnets contains resources that should be accessible only through your private network, such as a database that contains customer's personal information and order histories.

In a VPC subnets can communicate with each ohter.

## Network Traffic in a VPC
When a customer requests data from an application hosted in the aws cloud, this request is sent as a packet. a packet is a unit of data sent over the internet or a network.

It enters into a vpc through an internet gateway. before a packet can enter into a subnet or exit from a subnet it checks for permissions. these permissions indicate who sent the packet and how the packet is trying to communicate with the resources in a subnet.

The VPC component that checks packet permissions for subnets is a network access control lists (acl)

## Network Access Control Lists (ACLS)
A network access control lists (acl) is a virtual firewall that controls inbound and outbound traffic at the subnet level.

Each aws account includes a default network acl. when configuring your vpc, you can use your account'd default network acl or create custom network acls.

By default your account's network acl allows all inbound and outbound traffic, but you can modify it by adding your own rules. for custom acls, all inbound and outbound traffic is denied until you add rules to specify which traffic to allow. additionally all network acls have an explicit deny rule.

## Stateless Packet Filtering
Network acls perform stateless packet filtering. they remember nothing and check packets that cross the subnet border each way; inbound and outbound

when a packet response for that request comes back to the subnet, the network acl does not remember your previous request. the network acl checks the packet response against its list of rules to determine whether to allow or deny.

## Securit Groups
A security group is a virtaul firewall that controls inbound and outbound traffic for an amazon ec2 instance.

By default a security group denies all inbound traffic and allows all outbound traffic. you can add custom rules to configure which traffic to allow or deny.

If you have multiple amazon ec2 instances within a subnet, you can associate them with the same security group or use different security groups for each instance.

## Stateful Packet Filtering
Security groups perform stateful packet filtering. they remember previous decisions made for incoming packets.

When a packet response for that request returns to the instance, the security group remembers your previous request. the security group allows the response to proceed, regardless of inbound security group rules.

Both acls and security group enable you to configure custom rules for the traffic in you vpc.

# Global Networking

## Domain Name System (DNS)
You can think of dns as being the phone book of the internet. dns resolution is the process of thranslating a domain name to an ip address.

1. when you enter the domain name into your browser, this request is sent to a customer dns resolver
1. the customer dns resolver asks the company dns server for the ip address that corresponds to any company website.
1. the company dns server responds by providing the ip address for anycompany websiete 192....

## Amazon Route 53
Amazon route 53 is a dns web service. it gives developers and businesses a reliable way to route end users to internet applications hosted in aws.

amazon route 53 connects user requests to infrastructure running in aws. it can route users to infrastructure outside of aws.

another feature of route 53 is the ability to manage the dns records for domain names. you can register new domain names direclty in route 53. you can also transfer dns records for existing doamin names managed by other domain registrars.

1. a customer requests data from the application by going to any company's website
1. amazon route 53 uses dns resolution to identify anycompany.com corresponding ip address 192.0.... this information is sent back to the customer
1. customer request is sent to the nearest edge location through amazon cloudfront
1. amazon cloudfront connects to the application laod balancer which sends the incoming packet to an amazon ec2 instance.

