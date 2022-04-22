# AWS Cloud

## Cloud Architecture Technologies

**What is Solutions Archtecht?** <br/>
A role in a technical organization that architects a technical solution using multiple systems via researching, documentation, experimentation

**What is a Cloud Architect?** <br/>
A solutions architect that is focused solely on architecting technical solutions using cloud services.

A cloud architect need to understand the following terms and factor them into their designed and architecture based on the business requirements.
- **Availability:** Your ability to ensure a service remains available eg. Highly Available (HA)
- **Scalability:** Your ability to grow rapidly or unimpeded
- **Elasticity:** Your ability to shrink and grow to meet the demand
- **Fault Tolerance:** Your ability to prevent a failure
- **Disaster Recovery:** Your ability to recover from a failure eg. Highly Durable (DR)

A Solutions Archtect needs to always consider the following business factors: <br/>
- Security: How secure is this solution?
- Cost: How much is this going to cost?

## High Availability
Your ability for your service to remain available by ensuring there is no single point of failure and/or ensure a certain level of performance.

**Elastic Load Balancer:** <br/>
A load balancer allows you to evenly distribute traffic to multiple servers in one or more datacenter. If a datacenter or server becomes unavailable (unhealthy) the load balancer will route the traffic to only availabe datacentes with servers.

![image](https://user-images.githubusercontent.com/74575612/152843787-85fd24d2-cc5f-4002-8308-e8a2c206c790.png)

Running your workload across multiple **Availability Zones** ensures that if 1 or 2 AZs become unavailable your service / applications remain available.

## High Scalability
Your ability to increase your capacity based on the increasing demand of traffic, memory and computing power.

![image](https://user-images.githubusercontent.com/74575612/152844132-4d16104f-0e37-4959-9ef1-ac65fbc67ce7.png)

## High Elasticity
Your ability to automatically increase or decrease your capacity based on the current demand of traffic, memory and computing power.

**Auto Scaling Groups (ASG)** is an AWS feature that will automatically add or remore servers based on scaling rules you define based on metrics.

**Horizontal Scaling** <br/>
Scaling Out - Add more servers of the same size <br/>
Scaling In - Removing underutilzied servers of the same data <br/>

Vertical Scaling is generally hard for traditional architecture so you'll usually only see horizontal scaling described with Elasticity.

## Highly Fault Tolerant
Your ability for you service to ensure there is no single point of failure. Preventing the chance of failure.

**Fail-overs** is when you have a plan to shift traffic to a redundant system in case the primary system fails.

![image](https://user-images.githubusercontent.com/74575612/152844931-bf72133f-b122-4ebf-b1ea-b10b726033ad.png)

A common example is having a copy (secondary) of your databas where all ongoing changes are synced. The secondary system is not in-use until a fail over occurs and it becomes the primary databse.

**RDS Multi-AZ** is when you run a duplicate standby database in another Availability Zone in case your primary database fails.

## High Durability
Your ability to recover from a disaster and to prevent the loss of data Solutions that recover from a disaster is known as **Disaster Recover (DR)**
- Do you need a backup?
- How fast can you restore that backup?
- Does your backup still work?
- How do you ensure current live data is not corrupt?

**CloudEndure Disaster Recovery** continiously replicates your machines into a low-cost staging area in your target AWS account and preferred Region enabling fast and reliable recover in case of IT data center failutes.

## Business Continuity Plan (BCP)
A **business continuity plan** (BCP) is a document that outlines how a business will continue operating during an unplanned disruption in services.

**Recovery Point Objective (RPO)** the maximum acceptable amount of data loss after an unplanned data-loss incident, expressed as an amount of time.

**Recovery Time Objective (RTO)** the maximum amount of downtime your business can tolerate without incurring a significant finacial loss

![image](https://user-images.githubusercontent.com/74575612/152846335-df70e4b9-f190-40eb-87f6-fc821c4e66e4.png)

## Disaster Recovery Options
There are multiple options for recovery that trade cist vs time to recover

![image](https://user-images.githubusercontent.com/74575612/152846602-ee3f23f3-8616-4a1a-8a28-7dc7165d0068.png)

## RTO
**Recovery Time Objective (RTO)** is the maximum acceptable delay between the interruption of service and restoration of service. This objective determines what is considered an acceptable time window when service in unavailable and is defined by the organization.

![image](https://user-images.githubusercontent.com/74575612/152846929-220d8ec1-22bf-4dd1-9a3f-9342d6696ebc.png)

## RPO
**Recovery Point Objective (RPO)** is the maximum acceptable amount of time since the last data recovery point. The objective determines what is considered an acceptable loss of data between the last recovery and the interruption of service and is defined by the organisation.

![image](https://user-images.githubusercontent.com/74575612/152847187-1cde98ca-3d35-4275-9e90-e02667b8b114.png)

Resources:
- [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)
