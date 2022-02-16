# Certified cloud practitioner

## what is aws certified cloud practitioner?
The certified cloud practitioner is the entry level certification
- The cloud fundamentals e.g. cloud concepts, cloud architecture, cloud deployment models
- A close look at the aws core services
- a quick look at the vast amount of AWS services
- identify, security and governance of the cloud
- billing, pricing and support of aws services

*The course code is `CLF-C01` but its commongly referred to as the CCP*

Various roles
- cloud pratitioner
  - SysOps administration
  - developer
  - solutions architect
    - DevOps
    - solutions Architect
      - Security
      - Advanced networking
      - database
      - machine learning
      - data analysis


## The exam of questions in 4 doamins
- 26% domain1: cloud concepts
- 25% Domain2: Security and compliance
- 33% domain3: technology
- 16% domain4: billing and pricing

## What is cloud computing?
the practice of using a network of remote servers hosted on the internet to store, manage, and process data, rather than a local server or a personal computer.

on-premise
- you own the servers
- you hire the IT people
- You pay or rent the real-estate
- you take all the risk

cloud providers
- someone else owns the servers
- someone else hires the IT people
- someone else pays or rents the real-estate
- You are responsible for your configuring cloud services and code, someone else takes care of the rest.

## Evolution of cloud Hosting
- Dedicated server
  - one physical machine dedicated to single business.
  - Runs a single web-app/site
  - very expensive, High maintenance, *High Security
- virtual private server (VPS)
  - one physical machine dedicated to a single business. The physical machine is virtualized into sub-machine
  - runs multiple web-apps/sites
  - better utilization and isolation of resources.
- shared hosting
  - one pysical machine, shared by hundred of businesses relies on most tenants under utilizing their resources.
  - very cheap, limited functionality, poor isolation
- cloud hosting
  - multiple physical machines that act as one system. The system is abstracted into multiple cloud services.
  - flexible, scalable, secure cost-effective, high configurability

## What is amazon
Amazon has expanded beyond just an online e-commerce store into:
- cloud computing
- digital streaming
  - amazon prime video
  - amazon prime music
- grocery stores
- artifical intelligence
- low orbit satellites

## what is aws?
amazon calls their cloud provider service Amazon Web Services commongly referred to just aws    

## what is cloud service provider (CSP)?
A cloud service provider (CSP) is a company which
- provides multiple cloud services e.g. tens to hundreds of services
- those cloud services can be chained togther to create cloud architecture
- those cloud services are accessible via single unified API eg. AWS API
- those cloud services utilized metered billing based on usage e.g. per second, per hour
- those cloud services have rich monitoring built in eg AWS cloudTrail
- those cloud services have an infrastructure as a service (IaaS) offering
- Those cloud services offers automation via infrastructure as code (IaC)

*If a company offers multiple cloud services under a single UI but do not meet most of or all of these requriements, it would be reffered to as a cloud platform eg Twilio*

## Landscape of CSPs
- Tier 1 (Top Tier) - early to market, wide offering, storing synergies between services, well recognized in the industry.
  - aws
  - azure
  - google cloud platform
  - alibaba cloud
- Tier 2 (middle tier) - backed by well known tech companies, slow to innovate and turned to specialization
  - IBM cloud
  - oracle cloud
  - rackspace
- Tier 3 (ligh tier) - virtual private servers (VPS) turned to offer core Iaas offering, Simple and cose effective.
   - vultr
   - digital ocean
   - linode

## Gartner Magic quadrant for cloud
Magic Quadrant (MQ) is a series of market research reports published by IT consulting firm Gartner that rely on proprietary qualitive data analysis methods to demonstrate market trends, such as direction maturity and participants.

## common cloud services
A cloud service provider can have hundred of cloud services that are grouped into various types of services. The four most common types of cloud services (the 4 core) for Infrastructure as a service(IaaS) would be:
- compute
Imagine having a virtual computer that can run application, programs and code.
- networking
Imagine having virtual network defining internet connections of network isolations between services or outbound to the internet.
- storage
Imagine having a virtual hard-drive that can store files
- databases
imagine a virtual database for storing reporting data or a database for general purpose web-application

*The term cloud computing can be used to refer to all categories, even though it has compute in the name*

## Technology overview
- compute - ec2
- network - vpc private cloud network
- storage - ebs virtual  hard drives
- database - RDS sql database
- analytics
- application integration
- ar & vr
- aws cost management
- blockchain
- business applications
- containers
- customer engagement
- developer tools
- end user computing
- game tech
- internet of things
- machine learning
- management & gevernance
- media services
- migration & transfer
- mobile
- quantum
- robotics
- satellites
- security, identity & compliance

## evolution of computing
dedicated -> Vms -> containers -> functions

- dedicated 
  - a physical server wholly utilized by a single customer
  - you have to guess your capacity
  - you'll overplay for an underutilized server
  - you can't veritcal scale, you need a manual migration
  - replacing a server is very difficult
- VMs
  - you can run multiple virtual machines on one machine
  - hypervisor is the software layer that lets you run the vms
  - a physical server shared by multiple customers
  - you are to pay for a fraction of the server
  - you'll overpay for an underutilized virtual machine
  - you are limited by the guest operating system
  - easy to export or import images for migration
  - easy to verical or horizontal scale
- containers
  - virtual machine running multiple containers
  - docker daemon is the name of the software layer that lets you run multiple containers.
  - you can maximize the utilize tha available capacity which is more cost effective
  - your container share the same underlying OS so contaiers are more effeicient than multiple VMs
  - Multiple app can run side by side without being limited to the same OS requirements and will not cause conflicts during resource sharing.
- functions
  - are managed VMs running managed containers
  - known as serverless compute
  - you upload a piece of code, choose the amount of memory and duration
  - only reponsible for code and data, nothing else
  - very cost effective, only pay for the time code is running, VMs only run when there is code to be executed
  - cold starts is a side-effect of this setup.

## Types of cloud computing
- SaaS (Software as a Service) for customer
  - A product that is run and managed by the service provider eg gmail, office 365
- PaaS (Platform as a Service) for developers
  - focus on deployment and management of your apps. Don't worry about, provisioning, configuring or understanding the hardware or OS.eg Heroku
  - Iaas (Infrastructure as a Service) for admins
    - the basic building blocks for cloud IT. Providers access to network features, computer and data storage space.

## Cloud Computing Deployment Models
- public cloud
Everything (the workload or project) is  built on the csp also known as: Cloud-Native or Cloud First

- private cloud
Everything built on company's data centers also know as On-Premise. The cloud can be OpenStack

- hybrid
Using both on-premise and cloud service provider

- cross cloud
means using using cloud providers at the same time

|cloud|hybrid|on-premise|
|:-:|:-:|:-:|
|fully utilizing cloud computing|using both cloud and on-premise|deloying resources on prmises, using virtualization and resource management tools, is something called private cloud|
| startup, SaaS offerings, new projects and companies|organization that started with their own data center, can't fully move to cloud due to effort of migration or security compliance|organization that cannot run on cloud due strict regulatory compliance or the sheer size of their organization|