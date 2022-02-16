## AWS Direct Connect

aws direct connect is a private/dedicated connection between your datacenter, office, co-location and AWS.

direct connect has two very fast network

- lower bandwidth 50MBps-500MBps
- higher bandwidth 1GBps or 10 GBps

helps reduce network costs and increase bandwidth throughout (great for high traffic networks).

Provides a more consistent network experience than a typical internet-based connection (reliable and secure)

Direct connect locations are trusted partnered datacenters that you can establish a dedicated high speed, low latency connection from your on-premise to AWS.

## Local Zones

local zones are data centers located very close to a densely populated area to provide single digit miliseconds low latency performance for that area.

## Wavelength Zones

AWS wavelength zones allows for edge computing on 5G Networks. Applications will have ultra-low latency being as close as possible to the users.

## Data Residency

what is data residency?\
The physical or geographic location of where an organization or cloud resources reside.

what is compliance boundaries?\
a regulatory compliance (legal requirement) by a government or organization that describes where data and cloud resources are allowed to reside.

what is data sovereignty?\
data sovereingnty is the jurisdictional control or legal authority that can be asserted over data because it's physical location is within jurisdictional boundaries.

for workloads that need to meet compliance boundaries strictly defining the data residency of data and cloud resources in AWS you can use:

- aws config
  aws config is a policy as code service you can create rules to continuous check AWS resources configuration. If they deviate form your expectations you are alerted or aws config can some cases auto-remediate
- IAM policies
  can be written explicitly deny access access to specific AWS Regions.
- AWS outposts
  that you can put in your data center. Your data will reside whenever the outpost physically resides.

## aws for government

what is public sector?\
public sectors include public goods and governamental services such as

- military
- law enforcement
- infrastructure
- public transit
- public education
- health care
- the governemt itself

Aws can utilized by public sector or organizations developing cloud workloads for the public sector.

aws achieves this by meeting regulatory compliance programs along with specific governance and security controls.

aws has special regions for US regulation called GovCloud

## GovCloud (US)

Federal Risk and Authorization Mangement Program (FedRAMP)\
a US government-wide program that provides a standardized approach to security assessment, authorization and continuous monitoring for cloud products and services.

What is GovCloud?\
a cloud service provider (CSP) generally will offer an isolated region to run FedRAMP workloads.

## AWS in china

aws china is the aws cloud offerings in mainland china. aws china is completely isolate intentionally from aws global to meet regulatory compliance for mainland china. aws china has its own domain at: amazonaws.cn

## Sustainability

aws cloud's sustainability goals are composed of three parts.

- renewable enery
- cloud efficiency
- water stewardship

## aws ground station

aws ground station is a fully managed service that lets you control satellite communications, process data and scale your operations without having to worry about building or managing you own ground station infrastucture.

use cases for gorund station

- weather forecasting
- surface imaging
- communications
- video broadcasts

## aws outposts

aws outposts is fully managed service that offers the same aws infrastrucure, aws services, APIs and tools to virtually any data center, co-location space or on premises facility for a truly consistent hybrid experience.

## cloud architecture terminologies

- what is a solution architect?
  a role in a technical organization that architects a technical solution using multiple systems.
- what is a cloud architect?
  a solution architect that is focused solely on architecting technical solutions using cloud services.

a cloud architect need to understand the following terms and factor them into their designed architecture based on the business requirements

- availability - ability to ensure a service remains available
- scalability - ability to grow or unimpeded
- elasticity - you ability to shrink and grow to meet the demand
- fault tolerance - your ability to prevent a failure
- disaster recovery - ability to recover from a failure

a solutions architect needs to always consider the following business factors:

- how secure is this solution
- how much is this going to cost?

### High Availability

your ability for your service to remain available by ensuring there is no single point of failure and/or ensure a certain level of performance.

running your workload across multiple availability zones ensures that if 1 or 2 AZs become unavailable your service / applications remains available

### high scalability

your ability to increase your capacity based on the increasing demand of traffic, memory and computing power.

- vertical scaling (scaling up)
  upgrade to a bigger server
- horizontal scaling (scaling out)
  add more servers on the same size

### high elasticity

your ability to automatically increase or decrease your capacity based on the current demand or traffic, memory and computing power

horizontal scaling

- scaling out - add more servers of the same size
- scaling in - removing underutilized servers of the same size

veritcal scaling is generally hard for traditional architecture so you'll usually only see horizontal scaling described with elasticity

### high fault tolerant

your ability for your service to ensure there is no single point of failure, preventing the chance of failure

fail-over is when you have a plan to shift to a redundant system in case the primary system fails.

aws provide RDS Multi-AZ

### high durability

your ability to recover form a disaster and to prevent the loss of data solutions that recover from a disaster is known as disaster recovery

CloudEndure Disaster Recovery continuously replicates your machine into a low-cost staging area

### Business continuity plan (BCP)

A business continuity plan (BCP) is a document that outline how a business will continue operating during an unplanned disruption in services

- Recovery Point Objective (RPO)
  the maximum acceptable amount of data loss after an unplanned data-loss incident

- recovery time object (RTO)
  the maximum amount to downtime your business can tolerate without incurring a significant financial loss.

## disaster recovery options

there are multiple options for recovery that trade cost vs time to recover.

![aws multi-region](https://docs.aws.amazon.com/wellarchitected/latest/reliability-pillar/images/image16.pngf)

- Backup & Restore - your backup your data and restore it to new infrastructure
  - hours
  - lower priority use cases
  - deploy resource after the event
- Pilot light - data is replicated to another region with the minimal services running
  - less stringent RTO & RPO
  - core services
  - start and scale resources after event
- warm standby - scaled down copy of your infrastructure running ready to scale up
  - business critical services
  - scale resources after event
- multi-site active/active - scaled up copy of your infrastruture in anoter region
  - zero downtime
  - nero zero loss
  - mission critical services
