# Started Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

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

## Regional Services and Global Services

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
