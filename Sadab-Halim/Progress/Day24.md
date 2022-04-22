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

Resources:
- [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)