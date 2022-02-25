# Continued Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

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

