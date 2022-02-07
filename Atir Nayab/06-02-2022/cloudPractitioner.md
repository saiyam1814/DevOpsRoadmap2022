## edge and hybrid computing

what is edge computing?\
when you push your computing workoads outside of your network to run close to the destination location. eg pushing computing to run on phones, IoT devices, or external servers not within your cloud network

what is hybrid computing?\
when you're able to run workloads on both your on-premise datacenter and aws virtual private cloud (VPC)

## Types of storage service

### elastic block store (EBS) - block

- data is split into evenly split blocks
- directly accessed by the operation system
- supports only a single write volume.

### AWS elastic file storage (EFS)- file

- File is stored with data and metadata
- multiple connections via network share
- supports multiple reads, writing locks the file.
- when you need a file-share where multiple users or vms need to access the same drive.

### amazon simple storage service(s3) - object

- object is stored with data, metadata and unique ID
- scales with limited no file limit or storage limit
- supports multiple reads and writes (no locks)

## Intro to S3

what is Object Storage (Object-based Storage)?
data storage architecture that manages data objects, as opposed to other storage architectures:

- file systems which manages data as a files and fire hierarchy.
- block storage which manages data as blocks within sectors and tracks.

_S3 provides you with unlimited storage._

- S3 Object
  - Objects contain your data, They are like files.
  - Object may consist of:
    - key: this the name of the object
    - value the data itself made up of a sequence of bytes
    - version ID when versioning enabled, the version of object
    - metadata additional information attached to the object
- S3 Bucket
  - Buckets hold objects. Buckets can also have folders which in turn hold objects
  - S3 is a universal namespace so bucket names must be unique

## S3 storage classes

AWS offers a range of S3 storage classes that trade retrievel time, accessibility and durability for Cheaper Storage

- S3 standard (default)
  - fast! 99.99% availability, 11 9's durability, Relplication across at least three AZs.
- S3 intelligent tiering
  - uses ML to analyze object usage and determine the appropriate storage class.
  - Data is moved to the most cost-effective access tier, without any performance impact or added overhead.
- S3 standard-IA (Infrequent Access)
  - Still Fast! cheaper if you access files less than once a month.
  - Additional retrieval fee is applied. 50% less than standard (reduced availability)
- S3 One-zone-IA
  - Still Fast! Objects only exist in one AZ. Availability is (99.5%), but cheaper than standard IA by 20% less (reduce durability) data could get destroyed. A retrievel fee is applied.
- S3 Galcier
  - For long-term cold storage. Retrieval of data take minutes to hours but the off is very cheap storage
- S3 galcier deep archieve
  - The lowest cost storage class, data retrievel time is 12 hours.

## aws snow family

AWS snow family are storage and compute devices used to physically move data in or out the cloud. When moving data over the internet or private connection it too slow, difficult or costly

- snowcone
  - comes in two sizes
    - 8tb of storage
    - 14tb of storage
- snowball edge
  - comes generally in two type:
    - storage optimized
      - 80tb
    - compute optimized
      - 39.5 tb
- snowmobile
  - 100 pb of storage

## storage services

- simple storage service (S3) - is a serverless storage object service. You can upload very large files and an unlimited amount
  of files. You pay for what you store. You don't worry about the underlying file-system, or upgrading the disk size.
- s3 glacier is a cold storage service. It design as a low cost storage solution for archiving and long-term backup. It uses previous generation HDD drives to get that low cost. Its highly secure and durable.
- Elastic Block store(EBS) is a persistent block storage service. It is a virtual drive in the cloud you attach to EC2 instances. You can coose different kinds of hard dirves: SSD, IOPS, throughput HHD, Cold HHD
- Elastic file storage (EFS) is a cloud native NFS file system service. File storage you can mount to multiple EC2 instances at the same time. When you need to share files between multiple servers.
- storage gateway - is a hybrid cloud storage service that extends your on-premise storage to cloud.

## what is database?

A database is a data-store that stores semi-structured and structured data.

A database is more complex data stores because it requires using formal design and modeling techniques.

Database can be generally categorized as either:

- Relational databases
  - structured data that strongly represents tabular data (tables, rows and columns)
    - row oriented or columar-oriented
- Non-relational databases
  - semi-structured that may or may not distantly resemble tabular data

## what is data warehouse

A relational datastore designed for analytic workloads, which is generally column-oriented data store

Data warehouses generally perform aggregation

- aggregation is grouping data
- data warehouse are optimized around columns since they need to quickly aggregate column data

## Cloud-Native Networking Service

- vpc - It is a logically isolated section of the AWS cloud where you can launch AWS resources
- route tables - it determines where network traffic from your subnets are directed
- internet gateway - Enable access to the internet
- region - The geographical location of your network
- az - the data center of your aws resources
- subnets - it is a logical partition of an IP network into multiple, smaller network segments
- security groups - acts a firewall at the instance level
- nacls - acts as a firewalls at the subnet level.

## Enterprise/hybrid networking

- DirectConnect dedicated gigabit connection from on-premise data-center to AWS (a bery fast connection)
- AWS Virtual Private Network (VPN) a secure connection between on-premise, remote offices, mobile employees
- PrivateLinks - (VPC Interface Endpoints) keeps traffic within the AWS network and not traverse the internet to keep traffic secure.

## Virtual Private Cloud (VPC) and subnets

- Virtual Private Cloud (VPC) is a logically isolated section of the AWS Network where you launch you AWS resources. You choose a range of IPs using CIDR Range.
- subnets a logical partition of an IP network into multiple smaller network segments. You are breaking up your IP range for VPC into smaller networks.
- a public subnet is one that can reach the internet
- a private subnet is one that cannot react the internet.

## Security Groups vs NACLs

- Network Access Control Lists (NACLs) - acts as a virtual firewall at the subnet level. You create allow and deny rules. eg Block a specific IP address known for abuse
- Security Groups - Acts as a virtual at the instance level implicitly denies all traffic. You create only allow rules. eg You cannot block a single IP address.

## EC2

elastic compute cloud ec2 is a highly highly configurable virtual server. EC2 is resizable compute capacity. It takes minute to launch new instances. Anyhting and everything on AWS uses EC2 instance underneath.

1. choose os
1. choose instance type
1. add storage
1. configure instance

## EC2 Instance Families

what are instance families?\
Instance families are different combinations of CPU, memory, storage and networking capacity.

instance families allow you to choose the appropriate combinations of capacity to meet your application's unique requirements.

Different instance families are different because of the varying hardware used to give them their unique properties.

- general purpose
  - T2 A1 T3 T4g M4 M5a M5n M6zn M6g M6i Mac
  - balance of compute, memory and networking resources
  - use-cases web servers and code repositories
- Compute Optimized
  - CS C4 Cba C5n C6g C6gn
  - ideal for compute bound application that benefit from high performance processor
  - use-cases scientific modelling, dedicated gaming servers and ad server engines
- Memory Optimized
  - R4 R5 RSa RSb X1 X1E high Memory z1d
  - fast performance for workloads that process large data sets in memory.
  - use-cases in memory caches, in-memory database, real time big data analytics
- Accelerated Optimized
  - P2 P3 P4 G3 G4ad G4dn F1 Inf1 VT1
  - hardware accelerator, or co-processors
  - use-cases machine learning, computational finance analysis, speec recognition
- Storage Optimized
  - I3 I3en D2 D3 D3en H1
  - high sequential read and write access to very large data sets on local storage
  - use-cases NoSQL, in-memoryor transactional database, data warehousing

## EC2 Instance Types

An instance type is a particular instance size and instance family:

A common pattern for instance sizes:

- nano
- micro
- small
- medium
- large
- xlarge
- 2xlarge
- 4xlarge
- 8xlarge

## EC2 - Dedicated Host

Dedicated Hosts are single-tenant EC2 instance designed to let you Bring-Your-Own-License based on machine characteristics

|                                             | Dedicated Instance  |                           Dedicated Hosts                            |
| :-----------------------------------------: | :-----------------: | :------------------------------------------------------------------: |
|                  Isolation                  | instance isolation  |                      physical server isolation                       |
|                   Biling                    | per instance biling |                           per host biling                            |
|   Visibility of Physical characteristics    |   no visibilities   |                       sockets, cores, host ID                        |
|    Affinity between a host and instance     |     no affinity     | consistency deploy to the same instances to the same physical server |
|         Targeted instance placement         |     no control      |    additional control over instance placement on physical server     |
|        Automatic Instance placement         |         yes         |                                 yes                                  |
| Automatic capacity usign allocation request |         no          |                                 yes                                  |

## EC2 tenancy

ec2 has three levels of tenancy

- dedicated host
  - your server lives here and you have control of the physical attributes
- dedicated instance
  - your server always lives here
- default
  - you instance live here until reboot

## EC2 pricing models

there are 5 different ways to pay for EC2 (virtual machines)

- on demand (least commitment)
  - low cost and flexible
  - only pay per hour or the \*second
  - short term, spiky, unpredictable workloads
  - cannot be interrupted
  - For the first time apps
- spot upto 90% (biggest savings)
  - request spare computing capacity
  - flexible start and end times
  - can handle interruptions (server randomly stopping and starting)
  - for non-critical background jobs
- reserved upto 75% off (best long term)
  - steady state or predictable usage
  - commit to ec2 over a 1 or 3 year term
  - can resell unused reserved instances
- dedicated (Most expensive)
  - dedicated servers
  - can be on demand or reserved or spot
  - when you need a guarentee of isolate hardware (enterprise requirements)

## AWS Savings Plan

savings plan offer you the similar discounts as reserved instances (RI) but simplifies the purchasing process

There are 3 types of savings plans:

- Compute savings plans
- EC2 instance savings plans
- sageMaker savings plan
