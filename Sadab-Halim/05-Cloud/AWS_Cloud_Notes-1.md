## What is Certified Cloud Practitioner?
The Certified Cloud Practitioner is the entry-level AWS certification teaching cloud fundamentals e.g. Cloud Concepts, Cloud Architecture, Cloud Deployment Models

## AWS Certification Roadmap

![image](https://user-images.githubusercontent.com/74575612/151797071-0aed42d6-cc98-4da1-bf84-33d9052e3dc4.png)

## Cloud Concepts

What is Cloud Computing
Cloud Computing is the practice of using a network of remote servers hosted on the internet to store, manage, and process data, rather than a local server or a personal computer.

Evolution of Cloud Hosting
- Dedicated Server <br/>
  One physical machine dedicated to a single business.
  Runs a single web-app/site.
  
  Key Points: Very Expensive, High Maintenance, High Security
  
  ![image](https://user-images.githubusercontent.com/74575612/151810923-e325c752-41ae-4981-a4e2-cb8a8aa374a0.png)

- Virtual Private Server (VPS) <br/>
  One physical machine dedicates to a single business. <br/>
  The physical machine is virtualized into sub-machines. <br/>
  Runs multiple web-apps/sites.
  
  Key Points: Better Utilization and Isolation of Resources
  
  ![image](https://user-images.githubusercontent.com/74575612/151810993-881cafe7-dd93-44db-b65c-cbaf41b8e23d.png)

- Shared Hosting
  One physical machine, shared by hundred of business <br/>
  Relies on most tenants under utilizing their resources.
  
  Key Points: Very Cheap, Limited functionality, Poor isolation
  
  ![image](https://user-images.githubusercontent.com/74575612/151811058-3a1a5365-21a0-4b94-852f-51f5c2763d3b.png)

- Cloud Hosting
  Multiple physcial machines that act as one system. <br/>
  The system is abstracted into multiple cloud services <br/>
  
  Key Points: Flexible, Scalable, Secure, Cost Effective, High Configurability
  
  ![image](https://user-images.githubusercontent.com/74575612/151811146-b2b6ae7d-a5cb-4cbb-8d08-7962d788f872.png)

## What is AWS?
AWS is a collection of cloud services that can be used together under a single unified API to build a different types of workloads.

Cloud Service Providers can be initialized as CPSs

Timeline of AWS Services
- Simple Suque Service (SQS) was the first AWS service launched for public use in 2004
- Simple Storage Service (S3) was launched in March of 2006
- Elastic Compute Cloud (EC2) was launched in August of 2006
In November 2010. Amazon's retail sites had migrated to AWS

## What is a Cloud Service Provider?
A Cloud Service Provider (CSP) is a company which 
- provides multiple CLoud Services, e.g. tens to hundreds of services
- those Cloud Services can be chained together to create cloud architectures
- those Cloud Services are accessible via Single Unified API e.g. AWS API
- those Cloud Services utilized metered billing based on usage e.g., per second, per hour
- thouse Cloud Services have rich mointoring built in eg. AWS CloudTrail
- thouse Cloud Services have an Infrastructure as a Service (IaaS) offering
- those Cloud Services offers automation via Infrastructure as Code (IaC)
  
  ![image](https://user-images.githubusercontent.com/74575612/151818626-8d8fb7bd-fae3-48f3-8f09-42855f6055ae.png)

**NOTE:** If a company offers multiple cloud services under a single UI but do not meet most of or all of those requirements, it would be referred to as a Cloud Platform e.g. Twilio, HashiCorp, Databricks

**Landscape of CSPs**
- Tier-1 (Top Tier): Early to market, wide offering, strong synergies between services, well recognized in the industry
  1. Amazon Web Services (AWS)
  2. Microsoft Azure 
  3. Google Cloud Platform (GCP)
  4. Alibaba Cloud
  
- Tier-2 (Mid Tier): Backed by well-known tech companies, slow to innovate and turned to specialization.
  1. IBM Cloud
  2. Oracle Cloud
  3. Rackspace (OpenStack)
  
- Tier-3 (Light Tier): Virtual Private Servers (VPS) turned to offer core Iaas offering. Simple, cost-effective. 
  1. Vultr
  2. Digital Ocean
  3. Linode

## Common Cloud Services
A Cloud Service Provider can have hundreds of cloud services that are grouped into various types of services.

The four most common types (four core) of cloud services for Infrastructure as a Service (IaaS) would be:
1. **Compute:** Imagine having a virtual computer that can run application, programs and code.
2. **Networking:** Imagine having virtual network defining internet connection or network isolation between services or outbound to the internet
3. **Storage:** Imagine having a virtual hard-drive that can store files.
4. **Database:** Imagine a virtual database for storing reporting data or a database for general purpose web-application.

AWS Technology Overview
Cloud Service Provider (CSPs) that are Infrastructure as a Service (IaaS) will always have 4 core cloud service offerings:

![image](https://user-images.githubusercontent.com/74575612/151825535-7c23843a-b716-4e11-87d1-66e4c8484bd7.png)

## Evolution of Computing

- **Dedicated:**
  - A physical server wholly utilized by a single customer
  - You have to guess your capacity
  - You'll overpay for an underutilized server
  - You can't vertical scale, you need a manual migration
  - Replacing a server is very difficult
  - You are limited by your Host Operating System
  - Multiple apps can result in conflicts in resource sharing
  - You have a gurantee of security, privacy, and full utility of underlying resources.

  ![image](https://user-images.githubusercontent.com/74575612/152021935-6ca47f9f-99be-48bf-ad73-df56f4f4423d.png)

- **VMs:**
  - You can run multiple Virtual Machines on one machine.
  - Hypervisor is the software layer that lets you run the VMs
  - A physical server shared by multiple customers
  - You are to pay for a fraction of the server
  - You'll overpay for an underutilized Virtual Machine
  - You are limited by your Guest Operating System
  - Multiple apps on a single Virtual Machine can result in conflicts in resource sharing.
  - Easy to export or import images for migration
  - Easy to Vertical or Horizontally scale

  ![image](https://user-images.githubusercontent.com/74575612/152022670-e81f4147-0d31-4a3a-a0eb-808e054216fe.png)

- **Containers:**
  - Virtual Machine running multiple containers
  - Docker Deamon is the name of the software layer that lets you run multiple containers
  - You can maximixe the utilize of the available capacity which is more cost-effective.
  - Your containers share the same underlying OS so containers are more efficient than multiple VMs
  - Multiple apps can run side by side without being limited to the same OS requirements and will not cause conflicts during resource sharing.
  
  ![image](https://user-images.githubusercontent.com/74575612/152023297-85d34130-f51c-44d9-96b8-a5c4f5fbe3f5.png)

- **Functions:**
  - Are managed VMs running managed containers.
  - Known as Serverless Compute
  - You upload a piece of code, choose the amount of memory and duration.
  - Only responsible for code and data, nothing else
  - Very cost-effective, only pay for the time code is running. VMs only run when there is code to be executed
  - Cold Starts is a side-effect of this setup

  ![image](https://user-images.githubusercontent.com/74575612/152023812-625a67fd-2c1c-4814-9ee3-ee3fcf6d748a.png)

## Types of Cloud Computing
- **SaaS** Software as a Service (For Customers) <br/>
  A product that is run and managed by the service provider. <br/>
  
  _Don't worry about how the service is maintained._ <br/>
  _It just works and remains available_
  
  Example: salesforce, Mail, Office 365
  
- **PaaS** Platform as a Service (For Developers) <br/>
  Focus on the deployment and management of your apps. <br/>
  
  _Don't worry about, provisioning, configuring or understanding the hardware or OS. <br/>
  
  Example: heroku, Elastic, Google App Engine
  
- **IaaS** Infrastructure as a Service (For Admins) <br/>
  The basic building blocks for cloud IT. Provides access to networking features, computers and data storage space. <br/>
  
  _Don't worry about IT staff, data centers and hardware_.
  
  Example: Microsoft Azure, AWS, Oracle
  
## Cloud Computing Deployment Models

- **Public Cloud:** <br/>
  Everything (the workload of project) is built on the CSP. <br/>
  Also known as: Cloud-Native or Cloud First
  
  ![image](https://user-images.githubusercontent.com/74575612/152025707-3711d9a9-e4a6-4f88-91ae-4d4bb5c763cd.png)

- **Private Cloud** <br/>
  Everything built on company's datacenters, also known as On-Premise <br/>
  The cloud could be OpenStack
  
  ![image](https://user-images.githubusercontent.com/74575612/152025821-bd54f8a5-a8fd-481b-8be7-91a9fdaf0f49.png)

- **Hybrid** <br/>
  Using both On-Premise and A Cloud Service Provider
  
  ![image](https://user-images.githubusercontent.com/74575612/152026675-294ada21-b8a7-4b6f-b970-21b53ff40f4c.png)

- **Cross-Cloud** <br/>
  Using Multiple Cloud Providers
  Aka multi-cloud, "hybrid-cloud"
  
  ![image](https://user-images.githubusercontent.com/74575612/152027292-443976fb-faa4-48b1-af26-2a152fd5efc6.png)

## Deployment Model Use Cases:
| Cloud | Hybrid | On-Premise |
| ----- | ------ | ---------- |
| Fully utilizing cloud computing | Using both Cloud and On-Premise | Deploying resources on-premises, using virualization and resource management tools, is sometimes called "private cloud". |
| Companies that are starting out today, or are small enough to make the leap from a VPS to a CSP. | Organizations that started with their own datacenter, can't fully move to cloud due to effort of migration or security compliance | Organizations that cannot run on cloud due to strict regulatory compliance or the sheer size of their organization |
| Examples: Startups, SaaS offerings, New projects and companies | Examples: Banks, FinTech, Investment Management, Large Professional Service Providers Legacy On-Premise | Examples: Public Sector, e.g. Government, Super Sensitive Data e.g. Hospitals, Large Enterprise with heavy regulation e.g. Insurance Companies |

## Innovation Waves
**Kondratiev waves** (aka Innovation Waves or K-Waves) are hypothesized cycle-like phenomena in the global world economy. <br/>
The phenomenon is closely connected with technology life cycles.

![image](https://user-images.githubusercontent.com/74575612/152029465-9fcfc0b4-9d54-4808-bfe3-0fcd74dae5a6.png)

## Burning Platform
**Burning Platform** is a term used when a company abandons old technology for new technology with the uncertainty of success and can be motivated by fear that the organisation future survival hinges on its _digital transformation._ <br/>

## Evolution of Computing Power?
**What is Computing Power?** <br/>
The throughput measured at which a computer can complete a computational task.

- **General Computing:** Xeon CPU Processor
- **GPU Computing:** 50x faster than traditional CPUs
- **Quantum Computing:** 
  - D-Wave 2000Q
  - Rigetti 16Q Aspen-4
  - IonQ linear ion trap
  - 100 Million times faster

**AWS Service Offering:**
- Elastic Compute Cloud EC2
- AWS Infrastructure (Ef1)
- AWS Bracket

## The Benefits of Cloud
- Agility
  - Increase speed and agility
- Pay-as-you go pricing
  - Trade capital expense for variable expense
- Economy of scale
  - Benefit from massive economies of scale
- Global Reach
  - Go global in minutes
- Security
- Readability
  - Stop spending money on running and maintaining data centers.
- High Availability
- Scalability
  - Benefit from massive economies of scale
- Elasticity

## Six Advantages of Cloud
1. **Trade capital expense for variable expense** <br/>
    Pay-On-Demand: _Instead of paying for upfront costs of data centers and servers._
  
2. **Benefit from massive economies of scale** <br/>
    Sharing the cost with other customers to get unbeatable savings.
  
4. **Stop guessing capacity** <br/>
    Scale up or down to meet the current need. Launch and destroy services whenever we need.
  
6. **Increase speed and agility** <br/>
    Launch resources within a few clicks in minutes
    
8. **Stop spending money on running and maintaining data centers** <br/>
    Focus on your own customers, developing and configuring your applications.
    
10. **Go global in minutes** <br/>
    Deploy your app in multiple regions around the world with a few clicks.

# AWS Cloud

## Seven Advantages of Cloud
- **Cost-Effective** <br/>
  Pay for what you consume, no up-front cost, Pay-as-you-go
  
- **Global** <br/>
  Launch workloads anywhere in the world, just choose a region
  
- **Secure** <br/>
  Cloud provider takes care of physical security. Cloud services can be secure by default.
  
- **Reliable** <br/>
  Data backup, disaster recovery, data replication, and fault tolerance
  
- **Scalable** <br/>
  Increase or decrease resources and services based on demand
  
- **Elastic** <br/>
  Automate scaling during spikes and drop in demand
  
- **Current** <br/>
  The underlying hardware and managed software is patched, upgraded and replaced by the cloud provider without any interruption.
  
## AWS Global Infrastructure
**What is AWS Global Infrastructure?**
The AWS Global Infrastructure is globally distributed hardware and datacenters that are physically networked together to act os one large resource for the end customer.

Regions are geographically distinct locations consisting of one or more Availability Zones.

Every region is physically isolated from and independento of every other resion in terms of location, power, water supply.

![image](https://user-images.githubusercontent.com/74575612/152508632-47e576bc-bab7-461d-baa9-d5848f6acb15.png)

![image](https://user-images.githubusercontent.com/74575612/152508728-ee8a89bc-93d7-4cbf-bdf3-46ff2d7e6d97.png)


Each region generally has three availability zones:
- Some new useers are limited to two e.g. US-West
New services almost always become available first in US-EAST <br/>
Not all AWS Services are available in all regions <br/>
All your billing information appears in US-EAST (North Virgina) <br/>
The cost of AWS services vary per region <br/>

When you choose a region there are **four factors** you need to consider:
1. What Regulatory Compliance does this region meet?
2. What is the cost of AWS services in this region?
3. What AWS services are available in this region?
4. What is the distance or latency to my end-users?

### Regional Services and Global Services

**Regional Services**
AWS scopes their AWS Management Console on a selected Region.
This will determine where an AWS service will be launched and what will be seen within an AWS Service's console.

![image](https://user-images.githubusercontent.com/74575612/152510278-55b2e709-ed14-4bf0-8407-65f9cc6c9846.png)

**Global Services**
Some AWS Services operate across multiple region and will be fixed to "Global" <br/>
E.g. Amazon S3, CloudFront, Route53, IAM

![image](https://user-images.githubusercontent.com/74575612/152510177-b94080ce-94a4-4e1e-84a0-d134dcdddc80.png)

For these global services at the time of creation:
- There is no concept of region --> e.g. IAM User
- A single region must be explicitly chosen --> e.g. S3 Bucket
- A group or regions chosen --> CloudFront Distribution

**Availability Zones**
An Availability Zone (AZ) is physical location made up of one or more datacenter.

A datacenter is a secured building that contains hundreds of thousands of computers.

A region will generally contain 3 Availability Zones

Datacenters within a region will be isolate from each other. But, they will be close enough to provide low-latency (<10 ms). <br/>

It s common practice to run workloads in at least 3 AZs to ensure services remain available in case one or two datacenters fail. (High Availability)

AZs ae represented by a Region Code, followed by a letter identifier e.g. us-east-1a

A Subnet is associated with an Availability Zone.

You never choose the AZ when launching resources. <br/>

![image](https://user-images.githubusercontent.com/74575612/152550148-93aadb21-719a-4b61-8220-9c7883070dc6.png)

You choose the Subnet which is associated to the AZ.

Example of an architectural diagram, representing two AZs. the Subnets associated with those AZs, and EC2 instances (Virtuaal Machines) launched in those subnets.

![image](https://user-images.githubusercontent.com/74575612/152550437-dc4421b6-491d-47f5-9a05-e26d13e3c21e.png)

A region has multiple Availability Zones

An Availability Zone is made up of one or more datacenters.

All AZs in an AWS Region are interconnected with high-bandwidth, low-latency networking, over fully rebundant, dedicated metro fiber providing high throughput, low-latency networking between.

All traffuc between AZs is encrypted

AZs are within 100 km (60 miles) of each other.

![image](https://user-images.githubusercontent.com/74575612/152551257-59c26ba2-e130-4963-bc72-cb94d6614c1e.png)

### Fault Tolerance
**What is fault domain?**
A fault domain is a section of a network that is vulnerable to damage if a critical device or system fails. <br/>
The purpose of a fault domain is that if a failure occurs, It will not cascade outside that domain, limiting the damage possible.

![image](https://user-images.githubusercontent.com/74575612/152552073-f6c998d7-9187-40d9-a57e-a99ab3eb3936.png)

We can have fault domains nested inside fault domains.

**What is a fault level?**
A fault level is a collection of fault domains.

The scope of a fault domain could be:
- specific servers in a rack
- an entire rack in a datacenter
- an entire room in a datacenter
- the entire data center building

Its upto the Cloud Service Provider (CSPs) to define the boundaries of domain.

An AWS Region would be a Fault Level
![image](https://user-images.githubusercontent.com/74575612/152552544-b6da2d5a-c69d-4b72-b3c3-6d2359ac6f54.png)

A Availability Zone would be a Fault Domain
![image](https://user-images.githubusercontent.com/74575612/152552596-9b2b7ff8-4dd0-4e1a-ac0b-cc0d7fbd3f81.png)

Each Amazon Region is desinged to be completely isolated from the other Amazon Regions.
- This achieves the greatest possible fault tolerance and stability.

Each Availability Zone is isolated, but the Availability Zones in a region are connected through low-latency links.

Each Availability Zone is desined as an independent failure zone.

**Failure Zone:**
- Availability Zones are physically separated within a typical metropolitan region and are located in lower risk flood plants.
- Discrete uninterruptible power supply (UPS) and onsite backup generation facilities
- Data centers located in different Availability Zones are designed to be supplied by independent substations to reduce the risk of an event on the power grid impacting more than one Availabilty Zone.
- Availabiilty Zones are all redundantly connected to multiple tier-1 transit providers.

If an application is partitione across AZs, companies are better isoated and protected from issues such as power outage, lightining strike, tornadoes, earthquake and more.

# AWS Cloud

## AWS Global Network
The AWS Global Network represents the interconnections between AWS Global Infrastructure. <br/>
Commonly referred to as the _"Backbone of AWS"_

**Edge Locations** can acts as on and off ramps to the AWS Global Network
- **AWS Global Accelerator and AWS S3 Transfer Acceleration** <br/>
  Uses Edge Locations as an on-ramp to quickly reach AWS resources in other regions by traversing the fast AWS Global Network. 
- **Amazon CloudFront** (CDN) <br/>
  Uses Edge Location as an off-ramp, to provide at the Edge storage and compute near the end user.
  
## Point of Presence (PoP)
**Points of Presence (PoP)** is an intermediate location between an AWS Region and the end user, and this location could be a datacenter or collection of hardware.

For AWS a Point of Presnece is data center owned by AWS or a trusted partner that is utlized by AWS Services related for content delivery or expediated upload.

PoP resources are:
- Edge Locations
- Regional Edge Caches

![image](https://user-images.githubusercontent.com/74575612/152555406-b2321d8a-34f2-4ee2-b0de-9f9ec724831e.png)

**Edge Locations** are datacenters that hold cached (copy) on the most popular files (e.g. web pages, images and videos) so that the delivery of distance to the end users are reduce.

**Regional Edge Locations** are datacenters that hold much larger caches of less-popular files to reduce a full round trip and also to reduce the cost of transfer fees.

![image](https://user-images.githubusercontent.com/74575612/152556022-09ad8fe7-fec6-40eb-b228-73d0ec436ec0.png)

The following AWS Services use PoPs for content delivery or expediated upload.

**Amazon CloudFront** is a Content Delivery Network (CDN) that:
- You point your website to CloudFront so that it will route requests to nearest Edge Location cache
- Allows you to choose an origin (such as web-server or storage) that will be source of cached
- Caches the contents of what origin would returned to various Edge Location around the world.

**Amazon S3 Transfer Accelecration:** allows you to generate a special URL that can be used by end user to upload files to a nearby Edge Locationn. <br/>
Once a file is uploaded to an Edge Location, it can move much fster within the AWS Network to reach S3.

**AWS Global Accelerator** can find the optimal path from the end user to your web-servers. <br/>
Global Accelerator are deployed within Edge Locations so you send user traffuc to ab Edge Location instead if directly to you web-application.

## AWS Direct Connect
AWS Direct Connect is a private/dedicated connection between your datacenter, office, co-location and AWS.

![image](https://user-images.githubusercontent.com/74575612/152557524-d89c8b44-8e33-4b99-b594-7dc3a93117b3.png)

**Direct Connect** has two very-fast network connection options:
1. Lower Bandwidth 50MBps-500MBps
2. Higher Bandwith 1GBps or 10 GBps

- Helps reduce network costs and increase bandwidth throughput. (Great for high traffic networks)
- Provides a more consistend network experience than a typical internet based connection. (Reliable and secure)

**A co-location** is a data center where equipment, space, and bandwidth are available for rental to retail customers.

**Direct Connect Locations** are trusted partnered datacenters that you can establish a dedicated high speed, low-latency connectin from you on-premise to AWS.

You would use the AWS Direct Connect service to order and establisj a connection.

### Local Zones
**Local Zones** are datacenters located very close to a densely populated area to provide single-digit millisecond low latency performance (e.g. 7ms) for that area.

- Los Angeles, California was the first Local Zone to be deployed.
  - It is a logical extension of the US-West Region
  - The identifier looks like the following: **us-west-w-lax-1a**
- Only specific AWS Services have been made available
  - EC2 Instance Types (T3, C5, R5, R5d,  I3en, G4)
  - EBS (io1 and gp2)
  - Amazon FSx
  - Application Load Balancer
  - Amazon VPC

The purpose of Local Zone is the support highlt-demaind application sensitive to latencies:
- Media & Entertainment
- Electronic Design Automation
- Ad-Tech
- Machine Learning

### Wavelength Zones
AWS Wavelength Zones allows for edge-computing on 5G Networks. <br/>
Applications will have ulta-low latency being as close as possible to the users.

AWS has partnered with various Telecom companies to utlilize their 5G networks.

You create a Subnet tied to a Wavelength Zone and then you can launch VMs to the edge of the targeted 5G networks.

### Data Residency
**What is Data Residency?** <br/>
The physcial or geogeaphic location of where an organization or cloud resources reside.

**What is Compliance Boundaries?** <br/>
A regulatory compliance (legal requirement) byy a government or organisation that describes where data and cloud resources are allowed to reside.

**What is Data Sovereignty?** <br/>
Data Sovereignty is the jurisdictioncal cotrol or legal authorit that can be asserted over data because it's physical location is within jurisdictional boundaries.

For workloads that need to meet compliance boundaries strictly definind the data residency of data and cloud resources in AWS you can use:

**AWS Outposts** is physcial rack of servers that you can put in your data center. Your data will reside whenevrr the Outpost physcially resides.

**AWS Config** is a policy as Code service. You can create rules to continuous check AWS resources configuration. If they deviate from your expectations you are alerted or WS Config can in some cases auto-rmediate.

**IAM Policies** cam be written explicitly deny access to specific AWS regions. A **Service Control Policy (SCP)** are permissions applied organisations wide.
