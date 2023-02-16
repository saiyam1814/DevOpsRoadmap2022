# Amazon Web Services
<p align="center">
 <img src="https://user-images.githubusercontent.com/45159366/114321716-048b1c00-9ad1-11eb-828d-a5a5f2a3c726.png">
  <br />
</p>

## Resources
* [AWS Whitepapers](https://d0.awsstatic.com/whitepapers/aws-overview.pdf)
* [AWS Training Material](https://www.aws.training/LearningLibrary?query=&filters=classification%3A67&from=0&size=15&sort=_score)

## What is Cloud Computing?

Cloud computing is the on-demand delivery of compute power, database, storage, applications and other IT resources through a cloud service platform via the internet with pay-as-you-go pricing.

Cloud computing provides a simple way to access servers, storage, databases and a broad set of application services over the internet. A cloud service platform such as Amazon Web Services owns and maintains the network-connected hardware required for these application services while you provision and use what you need via a web application.

## Advantages of Cloud Computing
* **Trade capital expenses for variable expense** - Instead of having to invest heavily in data centers and servers before you know how you're going to use them, you can pay only when you consume computing resources and pay only for how much you consume. 
* **Benefit from massive economies of scale** - By using cloud computing you can achieve a lower variable cost than you can get on your own. Because usage from hundreds of thousands of customers is aggregated in the cloud, providers such as AWS can achieve higher economies of scale which translates into lower pay as-you-go prices.
* **Stop guessing capacity** - Eliminate guessing on your infrastructure capacity needs. When you make a capacity decision prior to deploying an application, you often end up sitting on expensive idle resources or dealing with limited capacity. With cloud computing these problems go away. You can access as much or as little capacity as you need and scale up and down as required with only a few minute's notice.
* **Increase speed and agility** - In a cloud computing environment, new IT resources are only a click away, which means that your reduce the time to make those resources available to your developers from weeks to just minutes. This results in a dramatic increase in agility for the organization since the cost and time it takes to experiment and develop is significantly lower.
* **Stop spending money running and maintaining data centers** - Focus on projects that differentiate your business not the infrastructure. Cloud computing lets you focus on your own customers, rather than on the heavy lifting of racking, stacking and powering servers.
* **Go global in minutes** - Easily deploy your application in multiple regions around the world with just a few clicks. This means you can provide lower latency and a better experience for your customers at minimal cost.

## Types of Cloud Computing
### Infrastructure as a Service(IaaS)
Infrastructure as a Service (IaaS) contains the basic building blocks for cloud IT and typically provides access to networking features computers(virtual or on dedicated hardware) and data storage space. IaaS provides you with the highest level of flexibility and management control over your IT resources and is most similar to existing IT resources that many IT departments and developers are familiar with today.

### Platform as a Service
Platform as a Service (PaaS) removes the need for organization to manage the underlying infrastructure(usually hardware and operating system) and allows you to focus on the deployment and management of your applications. This helps you be more efficient as you don't need to worry about resources procurement, capacity, planning, software maintenance, patching or any other undifferentiated heavy lifting involved in running your application.

### Software as a Service
Software as a Service (SaaS) provides you with a completed product that is run and managed by the service provider. In most cases, people referring to Software ad a Service are referring to end-user applications. With SaaS offering you do not have to think about how you will use that particular piece of software. A common example of SaaS application is web-based email which you can use to send and receive email without having to manage feature additions to the email product or maintain the servers and operating systems that the email program is running on.

## Cloud Computing Deployment Models
### Cloud
A cloud-based application is fully deployed in the cloud and all parts of the application run in the cloud. Applications in the cloud have either been created in the cloud or have been migrated from any existing infrastructure to take advantage of the benefits of cloud computing. Cloud-based applications can be built on low-level infrastructure pieces or can use higher level services that provide abstraction from the management, architecting and scaling requirements of core infrastructure.

### Hybrid
A hybrid deployment is a way to connect infrastructure and applications between cloud-based resources and existing resources that are not located in the cloud. The most common method of hybrid deployment is between the cloud and existing on-premises infrastructure to extend and grow an organization's infrastructure into the cloud while connecting cloud resources to the internal system. 

### On-premises
The deployment of resources on-premises, using virtualization and resource management tools is sometimes called the 'Private cloud'. On-premises deployment doesn't provide many of the benefits of cloud computing but is sometimes sought for its ability to provide dedicated resources. In most cases, this deployment model is the same as legacy IT infrastructure while using application management and virtualization technologies to try and increase resource utilization.

## Global Infrastructure

The AWS cloud infrastructure is built around AWS regions and Availability Zones. An AWS Region is a physical location in the world where we have multiple Availability Zones(AZ). AZs consist of one or more discrete data centers, each with redundant power, networking and connectivity, housed in separate facilities. These AZs offer you the ability to operate production applications and databases that are more highly available, fault-tolerant and scalable than would be possible from a single data center. The AWS cloud operates in 80 AZs within 25 geographic Regions around the world with announced plans for more AZs and Regions.
Each Amazon Region is designed to be completely isolated from the other Amazon Regions. This achieves the greatest possible fault tolerance and stability. Each Availability Zone is isolated, but the AZs in a Region are connected through low-latency links. AWS provides you with the flexibility to place instances and store data within multiple geographic regions as well as across multiple AZs within each AWS Region. Each AZ is designed as an independent failure zone. This means that Availability Zones are physically separated within a typical metropolitan region and are located in risk flood plains. In addition to discrete uninterruptible power supply and onsite backup generation facilities, data centers located in different AZ are designed to be supplied by independent substations to reduce the risk of an event on the power impacting more than one AZ. Availability Zones are all redundantly connected to multiple tire-1 transit providers.

## Security and Compliance
### Security
Cloud security at AWS is the highest priority. As an AWS customer, you will benefit from a data center and network architecture built to meet the requirements of the most security-sensitive organizations. Security in the cloud is much like security in your on-premise data centers - only without the cost of maintaining facilities and hardware. In the cloud, you don't have to manage physical servers or storage devices. Instead, you use software-based security tools to monitor and protect the flow of information into and out of your cloud resources.

An advantage of AWS Cloud is that it allows you to scale and innovate while maintaining a secure environment and paying only for the services you use. This means that you can have the security you need at a lower cost than in on-premises environment. 

As an AWS customer, you inherit all the best practices of AWS policies, architecture and operational processes built to satisfy the requirements of our most security-sensitive customers. Get the flexibility and agility you need in security controls.

The AWS Cloud enables a shared responsibility model. While AWS manages security **of** the cloud, you are responsible for security **in** the cloud. This means that you retain control of the security you choose to implement to protect your own content, platform, applications, systems and networks no differently than you would in an on-site data center.

AWS provides you with guidance and expertise through online resources, personnel and partners. AWS provided you with advisories for current issues plus you have the opportunity to work with AWS when you encounter security issues.

#### Benefits of AWS Security
* **Keep your data safe**: The AWS Infrastructure puts strong safeguards in place to help protect your privacy. All data is stored in highly secure AWS data centers.
* **Meet Compliance Requirements**: AWS manages dozens of compliance programs in its infrastructure. This means that segments of your compliance have already been completed.
* **Save Money**: Cut costs by using AWS data centers. Maintain the highest standard of security without having to manage your own facility.
* **Scale Quickly**: Security scales with your AWS Cloud usage. No matter the size of your business, the AWS Infrastructure is designed to keep your data safe.

### Compliance
AWS Cloud Compliance enables you to understand the robust controls in place at AWS to maintain security and data protection in the cloud. As systems are built on top of AWS Cloud Infrastructure, compliance responsibilities will be shared. By tying together governance-focused, audit-friendly service features with applicable compliance or audit standards, AWS Compliance enablers build on traditional programs. This helps customers to establish and operate in an AWS security control environment.

The IT infrastructure that AWS provides to its customers is designed and managed in alignment with best security practices and a variety of IT security standards.

# AWS Web Services Cloud

### Topics
* AWS Management Console
* AWS Comand Line Interface
* Software Development Kit
* Analytics
* Application Integration
* Cloud Financial Management
* Compute Services
* Contact Center
* Containers
* Database
* Developer Tools
* Management and Governance
* Migration and Transfer
* Networking and Content Delivery
* Security, Identity and Compliance
* Storage

