# Started Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

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
