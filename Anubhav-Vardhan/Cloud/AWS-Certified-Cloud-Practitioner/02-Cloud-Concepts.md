# What is Cloud Computing

The practice of using a network of remote servers hosted on the internet to store,
manage and process data, rather than a local server or a personal computer.

|ON PREMISE|CLOUD PROVIDERS|
|---|---|
|You own the server|Someone else owns the server|
|You hire the IT people|Someone else hires the IT people|
|You pay or rent the real estate|Someone else pays or rent the real estate|
|You take all the risk|You are responsible for configuring cloud services and code. someoen else takes care of the rest|

# Evolution of Cloud Hosting

### Dedicated Sever

- One physical machine dedicated to a single business
- Runs a single web-app/site
- Very expensive
- High maintainance
- High Security

### Virtual Private Server (VPS)

- One physical machine dedicated to a single business
- The physical machine is virtualized into sub-machines
- Runs multiple web-apps/sites
- Better utilization and isolation of resources

### Shared Hosting

- One physical machine, shared by hundreds of businesses
- Relies on most tenants under-utilizing their resources
- Very cheap
- limited functionality
- Poor isolation

### Cloud Hosting

- Multiple physical machines that act as one system
- The system is abstracted into multiple cloud services
- Flexible
- Scalable
- Secure
- Cost-effective
- High configurability

# What is Amazon?

An American multinational computer technology corporation headquatered
in Seattle, Washington

Amazon was founded in 1994 by Jeff Bezos and the company started as an
online store for books and expanded to other products

Amazon has expanded beyond just an online store into:
- cloud computing (Amazon web services)
- digital streaming
- grocery store
- AI
- low orbit satellites

# What is AWS?

Amazon calls their cloud provoder service Amazon web Services, commonly
referred to just AWS

AWS was launched in 2006 and is the leading cloud service provider in the world

Cloud service provider can be initilized as CSPs

### Some history of services

- Simple Queue Service (SQS) was the first AWS service launched in 2004
- Simple Storage Service (S3) was launched in March of 2006
- Elastic Compute Cloud (EC2) was launched in August of 2006

- In November 2010, all of Amazon.com's retail sites had migrated to AWS
- To support industry wide training and skill standradization, AWS began offering
  certification program for computer engineers, on April, 2013
  
# What is a Cloud Service Provider?

A CSP is a company which 
- provides multiple cloud services eg. tens to hundreds of services
- those cloud services can be chained together to create cloud architecture
- those cloud services are available via single unified API eg. AWS API
- those cloud services utilize metered billing based on usage eg. per second, per hour
- those cloud services have rich monitoring built in eg. AWS CloudTrail
- those cloud services have Infrastucture as a Service (IaaS) offering
- those cloud services offer automation via Infrastucture as code (IaC)
- Example of an architecture
  ![image](https://user-images.githubusercontent.com/71918957/159246801-0a16e0f9-5351-4fdd-827f-ccb0ffd9e85d.png)

If a company offers multiple cloud services under a single UI but do not 
meet most of or all of these requirements, it would be referred to as a 
Cloud Platform eg, Twilio, HashiCorp, Databricks

# Landscape of CSPs

### Tier-1 (Top Tier)

- Early to market 
- wide offering
- strong synergies between services
- well recognized in the industry

Example:
- Amazon Web Services (AWS)
- Microsoft Azure
- Google Cloud Platform (GCP)
- Alibaba Cloud

### Tier-2 (Mid Tier)

- Backed by well known tech companies
- slow to innovate and turned to specialization

Example:
- IBM Cloud
- Oracle Cloud
- Rackspace (OpenStack)

### Tier-3 (Light Tier)

- Virtual Private Servers (VPS) turned to offer core IaaS offering.
- Simple
- Cost-effective

Example:
- Vultr
- Digital Ocean
- Linode

# Gartner Magic Quadrant for Cloud

Magic Quadrant is a series of market research reports published by 
IT consulting firm Gartner that rely on proprietary qualitative data
analysis methods to demonstrate market trends, such as direction, 
maturity and participants

![image](https://user-images.githubusercontent.com/71918957/159248771-f3a89e89-5a09-4cf3-972c-5535e80654ab.png)

# Common Cloud Services

A cloud service provider can have 100s of cloud services that are
grouped into various types of services.
The four most common types of cloud services (the 4 core) for infrastucture
as a Service (IaaS) would be:

### Compute
like a virtual computer that can run applications, programs, and code

### Networking
like a virtual network defining internet connections or network isolations
between services or outbound to the internet

### Storage
like having a virtual hard drive that can store files

### Databases
like a virtual database for storing reporting data or a database for general
purpose web application

**AWS has over 200+ cloud services**
The term "cloud computing" can be used to refer all categories, even though it has
"compute" in the name.

# AWS Technology Overview

Cloud Service Providers (CSPs) that are Infrastucture as a Service (IaaS)
will always have 4 core cloud offerings:
- Compute: **EC2** virtual machines
- Storage: **EBS** virtual hard drives
- Database: **RDS** SQL databases
- Networking (and content Delivery): **VPC** private cloud network

Other categories:
- Analytics
- Application Integration
- AR and VR
- AWS Cost Management
- Blockchain
- Business Application
- Containers
- Costumer Engagement
- Developer Tools
- End User Computing
- Game Tech
- Internet of Things
- Machine Learning
- Management and Governance
- Media Services
- Migration and Tranfer
- Mobile
- Quantum Technologies
- Robotics
- Satelites
- Security, Identity and Compliance

# AWS Services Preview

We can explore all the services offered by AWS by going to https://aws.amazon.com/ and clicking `products`

![image](https://user-images.githubusercontent.com/71918957/159497614-e7ff0b00-cbd1-4490-8e9c-b432e195b246.png)

# Evolution of Computing

### Dedicated

- A physical server wholly utilized by a single costumer
- You have to guess your capacity
- You'll overpay for an underutilized server
- You can't vertically scale, you need manual migration
- Replacing a server is very difficult
- You are limited by your host OS
- Multiple apps can result in conflicts in resource sharing
- You have a guarantee* if security, privacy and full utility of underlying resources

### VMs

- You can run multiple Virtual Machines on one machine
- Hypervisor is a software layer that lets you run VMs
- A physical server shared by multiple costumers
- You are to pay for a fraction of server
- You'll overpay for an underutilized virtual machine
- You are limited by your guest OS
- multiple apps on single VM can result in conflict in resource sharing
- Easy to export and import images for migration
- easy to vertical or horizontally scale

### Containers

- virtual machines running multiple containers
- Docker Daemon is a software that lets you run multiple containers
- 
