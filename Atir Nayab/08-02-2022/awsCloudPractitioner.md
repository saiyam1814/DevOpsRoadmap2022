# Directing traffic with elastic load balancing

## elastic load balancing

elastic load balancing is the aws service that automatically distributes incoming application traffic across multiple resources, such as amazon ec2 instances.

A load balancer acts as a single point of contact for all incoming web traffic to your auto scaling group. This means that as you add or remove amazon ec2 instances in response to the amount of incoming traffic, these requests route to the load balancer first. Then the requests spread across multiple resources that will handle them, For eg, if you have multiple amazon ec2 instances, elastic load balancing distributes the workload across the multiple instances so that no single instance has to carry the bulk of it.

Although elastic load balancing and amazon ec2 auto scaling are separate services, they work together to help ensure that applications running in amazon ec2 can provide high performance and availability

![low demand period](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/Mhs-brNLrz0idxgI_gxyNgUGrmnijDl62.png)

### low-demand period

suppose that a few customers have come to the coffee shop and are ready to place their orders.

if only a few registers are open, this matches the demand of customers who need service.

![high demand period](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/3OO8icoc4ZvQE-x1_q_fQ0lM9jrHOGE2v.png)

### high demand period

throughout the day, as the number of customers increases, the coffee shop opens more registers to accomodate them, in the diagram the auto scaling group represents this.

# Messaging and Queueing

## monolithic applications and microservices

![monolithic application](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/E0FJ5FMD_3svDFPS_sqlUJljjEBTJxbym.png)

applications are mode of multiple components. The components communicate with each ohter to transmit data, fulfill requrests, and keep the application running.

suppose that you have an application with tightly coupled components. These components might include databases, servers, the user interface, business logic so on. This type of architecture can be considered a monolithic application

In this approach to application architecture, if a single component fails, other components fail, and possibly the entire application fails.

## microservices

![microservices](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/e-tZJ4wfD5gtJNCA_oZ_ymliKJjdl2_No.png)

In a microservices approach, application components are loosely coupled, In this case, if a single component fails, the other components continue to work because they are communicating with each other, The loose coupling prevents the entire application from failing

## amazon simple notification service (amazon sns)

amazon sns is a publish/subscribe service. Using amazon sns topics, a publisher publishes messages to subscribers. This is similar to the coffee shop; the cashier provides coffee orders to the barista who makes the drinks.

## amazon simple queue service (amazon sqs)

using amazon sqs you can send, store and receive messages between software components, without losing messages or requiring other services to be available. In amazon SQS an application sends messages into a queue, A user or service retrieves a message from the queue, processes it and then deletes it from the queue.

# additional compute services

## serverless computing

Earlier in this module, you learned about amazon ec2, a service that lets you run virtaul servers in the cloud. If you have applications that you want to run in amazon ec2

1. provision instances (virtual servers)
1. upload your code
1. continue to manage the instances while your application is running.

![serverless compute](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/FOkwI-hxTA07WaKF_AKDPOn1ZYEHZsxk4.png)
The term 'serverless' means that your code runs on servers, but you do not need to provision or manage these servers. With serverless computing, you can focus more on innovating new products and features instead of maintaining servers

another benefit of serverless computing is the flexibility to scale serverless applications automatically. Serverless computing can adjust the applications capacity by modifying the units of consumptions, such as throughput and memory.

An aws service for serverless computing is aws lambda.

## aws lambda

aws lambda is a service that lets you run code without needing to provision or manage servers.

while using aws lambda, you pay only for the compute time that you consume. Charges apply only when your code is running. You can also run code for virtually any type of application or backend service, all with zero administration.

![aws lambda working](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/mJqf2HpZ33b0nHmJ_ssdAdxlZRBGG7c57.png)

1. upload your code to lambda
1. set your code to trigger from an event source, such as aws services, mobile applications, or HTTP endpoints.
1. lambda runs your code only when triggered.
1. you pay only for the compute time that you use. In the previous example of resizing images, you would pay only for the compute time that you use when uploading new images. Uploading images triggers lambda to run code for the image resizing function.

**in aws, you can also build and run containerized applications**

## containers

containers provide you with a standard way to package your application's code and dependencies into a single object. You can also containers for processors and workflows in which there are essential requirements for security, reliability and scalability

container orchestration services help you to deploy manage, and scale your containerized applications. Next you will learn about two services that provide container orchestration. Amazon elastic container service and amazon elastic kubernetes service.

## amazon elastic container service (amaozn ecs)

amazon ecs is highly scalable, high-performance container management system that enables you to run and scale containerized application on aws.

amazon ecs supports docker containers. docker is a software platform that enables you to build, test, and deploy applications quickly. AWS supports the use of open-source docker community edition and subscription-based docker enterprise edition.

## amazon elastic kubernetes service (amazon eks)

amazon eks is a fully managed service that you can use to run kubernetes on aws.

kubernetes is open-source software that enables you to deploy and manage containerized applications at scale.

## aws fargate

aws fargate is a serverless compute engine for containers. It works with both amazon ecs and amazon eks.

when using aws fargate, you do not need to provision or manage servers. AWS fargate manages your server infrastructure for you. you can focus more on innovating and developing your applications, and you pay only for the resources that are required to run your containers.

# aws global infrastructure

## selecting a region

when determining the right region for your services, data, and applications, consider the following four business factors.

### compliance with data governance and legal requirements

depending on your company and location you might need to run you data of specific areas. for eg if your company requires all of its data to reside within the boundaries of the UK. you should choose the london region

### proximity to your customers

selecting a region that is close to your customers will help you to get content to them fasters.

### availble services within a region

sometimes the closest region might not have all the features that you want offer customers. AWS is frequently innovating by creating new services and expanding on features within existing services.

### pricing

suppose that you are considering running applications in both the united states and brazil. the way brazil tax structure is set up, it might cost 50% more to run the same workload out of the region compared to the other region.

## availability zones

![availability zones](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/old4pJtvN7HvddhL_zBL1ijNGsMMu3-4j.png)

an availability zone is a single data center or a group of data centers within a region. availability zones are located tens of miles apart from each other. this is close enough to have low latency between availability zones. however if a disaster occurs in one part of the region, they are distant enough to reduce the chance that multiple availability zones are affected.

# edge locations

![edge locations](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644343200/53fqugduJvar8WnFEsRrRQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/X3WxHCXa0QwTmj_6_OvAnh0SdGr1YBIds.png)

an edge location is a site that amazon cloudfront uses to store cached copies of your content closer to your customers for faster delivery.

amazon uses cdn as services as cloudfront

# how to provision aws resources

## ways to interact with aws services

### aws management console

the aws management console is a web-based interface for accessing and managing aws services. you can quickly access recently used services and search for other services by name, keyword, or acronym. the console include wizards and automated workflows that can simplify the process fo completing tasks.

### aws command line interface

to save time when making api requrests you can use the aws command line interface (aws cli). aws cli enables you to control multiple aws services directly from the command line within one tool.

by using aws cli you can automate the actions that your service and applications perform through scripts.

### software development kits

another option for accessing and managing aws services is the software development kits (SDKs). sdks make it easier for you to use aws services through an api designed for your programming language or platform. sdks enable you to use aws services with your existing applications or create entirely new applications that will run on aws.

## aws elastic beanstalk

with aws elastic beanstalk, you provide code and configuration settings, and elastic beanstalk deploys the resources necessary to perform the following tasks.

- adjust capacity
- load balancing
- automatic scaling
- application health monitoring

## aws cloudformation

with aws cloudformation you can treat your infrastructure as code. this means that you can build an evironment by writing lines of code instead of using the aws management console to invdividually provision resources.

aws cloudformation provisions your resources is a safe, repeatable manner, enabling you to frequently build your infrastructure and applications without having to perform manual actions. it determines the right operations to perform when managing your stack and rolls back changes automatically if it detects errors.

# connectivity to aws

## amazon virtual private cloud (amazon vpc)

a networking service that you can use to establish boundaries around your aws resources is amazon virtual private cloud (amazon vpc)

amazon vpc enables you to provision an isolated section of the aws cloud. in this isolated section, you can launch resources in a virtual network that you define. within a virtual private cloud (vpc) you can organize your resources into subnets. a subnets is a section of a vpc that can contain resources such as amazon ec2 instances.

## internet gateway

to allow public traffic from the internet of access your vpc, you attach an internet gateway to the vpc.

![internet gateway](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/Q_HnMl_BAEsDZGxf_NEblbQjD0vn0-pPU.png)

an internet gateway is a connection between a vpc and the internet. you can think of an internet gateway as being similar to a doorway that customers use to enter the coffee shop. wihtout internet gateway no one can access the resources within your vpc.

## virtual private gateway

to access private resources in a vpc, you can use a virtual private gateway.

the virtual private gateway is the component that allows protected internet traffic to enter into the vpc.

![virtual private gateway](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/tthacSS-FyYNWwE3_s8U3lQzEONXm1FMX.png)

a virtaul private gateway enables you to establish a virtual private network (VPN) connection between your vpc and a private network, such as on-premises data center or internal corporate network. a virtual private gateway allows traffic into the vpc only if it is coming from an approved network.

## aws direct connect

aws direct connect is a service that enables you to establish a dedicated private connection between your data center and vpc.

![aws direct connect](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/p53HDtoqu2euSy0Y_YdzRvczPABE_j-yV.png)

the private connection that aws direct connect provides helps you to reduce network costs and increase the amount of bandwidth that can travel through your network.

## subnets

![subnets](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/-HQN-9SqkWc5e-yv_pi4TpfvEYaalWuIu.png)

a subnet is a section of a vpc in which you can group resources based on security or operational needs. subnets can be public or private.

### public subnets

public subnets contain resources that need to be accessible by the public, such as an online store's website.

### private subnets

private subnets contains resources that should be accessible only through your private network, such as a database that contains customer's personal information and order histories.

in a vpc subnets can communicate with each ohter.

## network traffic in a vpc

when a customer requests data from an application hosted in the aws cloud, this request is sent as a packet. a packet is a unit of data sent over the internet or a network.

it enters into a vpc through an internet gateway. before a packet can enter into a subnet or exit from a subnet it checks for permissions. these permissions indicate who sent the packet and how the packet is trying to communicate with the resources in a subnet.

the vpc component that checks packet permissions for subnets is a network access control lists (acl)

## network access control lists (acls)

a network access control lists (acl) is a virtual firewall that controls inbound and outbound traffic at the subnet level.

![acls](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/Iw1uUP_PAhIJktew_e5z0Nx1wJ-H36Lsd.png)

each aws account includes a default network acl. when configuring your vpc, you can use your account'd default network acl or create custom network acls.

by default your account's network acl allows all inbound and outbound traffic, but you can modify it by adding your own rules. for custom acls, all inbound and outbound traffic is denied until you add rules to specify which traffic to allow. additionally all network acls have an explicit deny rule.

## stateless packet filtering

network acls perform stateless packet filtering. they remember nothing and check packets that cross the subnet border each way; inbound and outbound

when a packet response for that request comes back to the subnet, the network acl does not remember your previous request. the network acl checks the packet response against its list of rules to determine whether to allow or deny.

![stateless packet filtering](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/-IqvewqIwJScqVre_y7vKKj12R8wZy8gK.png)

## securit groups

a security group is a virtaul firewall that controls inbound and outbound traffic for an amazon ec2 instance.

by default a security group denies all inbound traffic and allows all outbound traffic. you can add custom rules to configure which traffic to allow or deny.

if you have multiple amazon ec2 instances within a subnet, you can associate them with the same security group or use different security groups for each instance.

## stateful packet filtering

security groups perform stateful packet filtering. they remember previous decisions made for incoming packets.

when a packet response for that request returns to the instance, the security group remembers your previous request. the security group allows the response to proceed, regardless of inbound security group rules.

![stateful packet filtering](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/nf1lZAA-I6EQxVy1_unCx4hVGemR5ybFS.png)

both acls and security group enable you to configure custom rules for the traffic in you vpc.

![acls](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/QkcDe-SJB4lQAuyB_ha8um-1InZb0jryB.png)

# global networking

## domain name system (dns)

you can think of dns as being the phone book of the internet. dns resolution is the process of thranslating a domain name to an ip address.

![dns](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/Pd32jad3VWk7EjGe_AFBIDcTkYnYj1OOb.png)

1. when you enter the domain name into your browser, this request is sent to a customer dns resolver
1. the customer dns resolver asks the company dns server for the ip address that corresponds to any company website.
1. the company dns server responds by providing the ip address for anycompany websiete 192....

## amazon route 53

amazon route 53 is a dns web service. it gives developers and businesses a reliable way to route end users to internet applications hosted in aws.

amazon route 53 connects user requests to infrastructure running in aws. it can route users to infrastructure outside of aws.

another feature of route 53 is the ability to manage the dns records for domain names. you can register new domain names direclty in route 53. you can also transfer dns records for existing doamin names managed by other domain registrars.

![route 53 and cloudfront](https://assets.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1644354000/Yjb_a9p405n9vOf5WDwEzQ/tincan/31d9c0cca79c54bdceaf3e938fd424e97c98c7e8/assets/mR1nvYoC4OSUVg9a_WE71CA369xcdceJ2.png)

1. a customer requests data from the application by going to any company's website
1. amazon route 53 uses dns resolution to identify anycompany.com corresponding ip address 192.0.... this information is sent back to the customer
1. customer request is sent to the nearest edge location through amazon cloudfront
1. amazon cloudfront connects to the application laod balancer which sends the incoming packet to an amazon ec2 instance.
