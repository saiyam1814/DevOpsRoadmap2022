## Innovation Waves

Kondratiev waves (aka innovation waves or K-waves) are hypothesized cycle-like phenomena in the global world economy. The phenomenon is closely connected with technology life cycles.
Each wave irreversibly changes the society on a global scale. The latest wave is cloud Technology

## Burning Platform

Burning platform is a term used when a company abandons old technology for new technology with the uncertainty of success and can be motivated by fear that the organization future survival hinges on its digital transformation

## Digital Transformation Checklist

## Evolution of computing power

what is computing power?
The throughput measured at which a computer can complete a computational task.

- general computing
  - xeon CPU processor
- GPU computing
  - 50x faster than traditional CPUs
- Quantum Computing
  - 100 million times faster
  - D-wave 2000q
  - Regetti 16q Aspen-4

## AWS bracket

aws offers quantum computing with the through aws bracket

## The benefit of cloud

The benefits of the cloud is a summary of reasons why an organization would consider adopting or migrating to utilizing public cloud.

- Agility
  - Increase speed and agility
- Pay as you go pricing
  - Trade capital for variable exoense
- Economy of scale
  - benefit from massive economies of scale
- global reach
  - go global in minutes
- security
- reliability
  - stop spending money on running and maintaining data centers
- High availability
- scalability
  - benefit form massive economics of scale
- elasticity

## Six Advantage to cloud

1. trade capital expense for variable expense - You can pay on demand meaning there is no upfront-cost and you pay for only what you consume or pay by the hour, minutes or seconds.

1. benefit from massive economies of scale - You are sharing the cost with other customer to get ubeatable savings

1. stop guessing capacity - scale up or down to meet the current need. Launch and destroy services whenever.

1. Increase speed and agility - launch resources within a few clicks in minutes

1. Stop spending money on running and maintaining data centers - focus on your own customers

1. go global in minutes - deploy you app in multiple regions around the world with a few clicks.

## AWS global infrastructure

what is the aws global infrastructure?
The aws global infrastructure is globally distributed hardware and datacenters that are physically networked together to act as on large resource for the end customer.

## Global Infrastructure- Regions

Regions are geographically distinct locations consisting of one or more availability zones.

Every region is physically isolated from and independent if every other region in terms of location, power, water supply

Eachi region generally has three Availability zones

- some new users are limited to two eg. US-West
  New services almost always become available first in US-EAST. Not all aws services are available in all regions. All your billing information appears in US-EAST-1 (North Virginia) The cost of AWS services vary per region.

when you choose a region there are four factors you need to consider:

- what regulatory compliance does this region meet?
- what is the cost of AWS services in this region
- whaw AWS services are available in this region?
- What is the distance or latency to my end-users

## Regional vs Global Services

- Regional Services - AWS scopes their AWS Management Console on its selected region. This will determine where an AWS service will be launched and what will be seen within an AWS services console. You generally don't expilictly set the region

- global services - Some aws services operate across multiple regions and the region will be fixed to "Global"

## Availability Zones

An availability zone (AZ) is physical location made up of one or more datacenter.

A datacenter is a secured building that contains hundreds of thousands of computers.

A region will genrally contain 3 availability zones

Datacenters within a region will be isolated from each other (differnt buildings). But they will be close enough to provide low-latency (<10ms)

Its common practice to run workloads in at least 2 AZs to ensure services remain available in case one or two data centers fail.

A subnet associated with an availability zone.

You never choose the AZ when launching resources. You choose the subnet associated to the AZ.

A region has multiple Availability zone is made up of one or more datacenters

All AZs in an aws region are interconnected with high bandwidth, low-latency networking, over fully redundant, dedicated metro fiber providing high throughput, low-latency networking between.

All traffic between AZs is encrypted

AZs are within 100 km (60 miles) of each other

## Fault Tolerance

- what is a fault domain?
  A fault domain is a section of a network that is vulnerable to damage if a critical device or system fails. The purpose of a fault domain is that if a failure occurs it will not cascade outside that domain, limiting the damage possible.
- what is a fault level?
  a fault level is a collection of fault domains.
- The scope of a fault domain could be
  - specific servers in a rack
  - an entire rack in a datacenter
  - an entire room in a datacenter
  - the entire data center building

Each amazon region is designed to be completely isolated from the other amazon regions.

- This achives the greatest possible fault tolerance and stability.
  Each availability zones is isolated, but the availability zones in a region are connected through low-latency links Each availability zone is designed as an independent failure zone

**failure zone**

- Availability zones are physically separated within a typical metropolitan region and are located in lower risk flood plains.
- discrete uninterruptible power supply (UPS) and onsite backup deneration facilities
- data centers lcaoted in different availability zones are designed to be supplied by independent substations to reduce the risk of an event on the power grid impacting more than one availability zone.
- Availability zone are all redundantly connected to multiple tier-1 transit providers

## AWS global Network

The AWS global network represent the interconnection between AWS global infrastructure. Commonly referred to as the "The Backbone os AWS".

This of if as private expressway, where things can move very fast between data centers.

AWS global accelator, aws s3 transfer accelaration Uses edge locations as on on-ramp to quickly reach AWS resources in other regions by traversing the fast AWS global network.

amazon cloudFront (CDN) uses edge locations as on off-ramp, to provide at the edge storage and compute near the end user.

## Points of Presence(PoP)

Points of Presence (PoP) is an intermediate location between an AWS region and the end user, and this location could be a datacenter or collection of hardware.

For AWS a point of presence is a data center owned by aws or a trusted partner that is utilized by aws services related for content delivery or expediated upload.

### Edge locations

edge locations are datacenters that hold cached (copy) on the most popular files eg web pages, images, and videos so that the delivery of distance to the end users are reduce

### Regional Edge Locations

regional edge locations are data centers that hold much larger caches of less popular files to reduce a full round trip and also to reduce the cost of transfer fees.

### Tier 1 network

tier 1 network is a network that can reach every other network on the internet without purchasing IP transit or paying for peering.

Aws availability zones are all redundantly connected to multiple tier-1 transit providers

The following AWS services use PoPs for content delivery or expediated upload.

- amazon cloudfront is a content delivery network (CDN) service that:
  - you point your website to cloudfront so that it will route requests to nearest edge ocation cache
  - allows you to choose an origin that will source of cached
  - caches the contents of what origin would returned to various edge locations around the world.
- amazon s3 transfer accelaration allows you to generate a special URL that can be used by end users to upload files to a nearby edge location. Once a file is uplaoded to an edge location, it can move much faster within the aws network to reach s3
- aws global accelerator can find the optimal path form the end user to your web-servers. Global accelerator are deployed within edge locaitons so you send user traffic to an edge location istad of directly to your web-application.
