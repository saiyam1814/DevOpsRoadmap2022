# Started Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

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
