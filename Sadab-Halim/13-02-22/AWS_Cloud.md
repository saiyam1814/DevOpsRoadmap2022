# Continued Learning [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

## Edge and Hybrid Computing

**What is Edge Computing?** <br/>
When you push your computing workoads outside of your network to run close to the destination location. eg pushing computing to run on phones, IoT devices, or external servers not within your cloud network

**What is Hybrid Computing?** <br/>
When you're able to run workloads on both your on-premise datacenter and aws virtual private cloud (VPC)

**AWS Outposts** is physical rack of servers that you can put in your datacenter. AWS Outposts allows you to use AWS API and Services such as EC2 right in your datacenter.

**AWS Wavelength** allows you to build and lanuch your application in a telecom datacenter. By doing this your applications which have ultra-low latency since they will be pushed over the 5G network and be closest as possible to the end user.

**VMWare Cloud on AWS** allows you to manage on-premise virtual machines using VMWare as EC2 instances. The datacenter must being using VMWare for Virtualizations.

**AWS Local Zones** are edge datacenters located outside of an AWS region so you can use AWS closer to end destination. <br/>

_When you need faster computing, storage and databases in populated areas that are outside of an AWS region._

## Cost and Capacity Management Computing Services

**Cost Management** How do we save money?
**Capacity Management** How do we meet the demand of traffic and usages though adding or upgrading servers?

**EC2 Spot Instances, Reserved Instanced and Savings Plan** <br/>
Ways to save on computing, by paying up in full or partially, by commiting to a yearly contracts or by being flexible about availability and interruption to computing service.

**AWS Batch** plans, schedules, and executes your batch computing workloads across the full range of AWS compute services, can utilize Spot instance to save money.

**AWS Compute Optimizer** suggests how to reduce costs and improve performance by using machine learning to analyze your previous usage history.

**EC2 Autoscaling Groups (ASGs)** <br/>
Automatically adds or remove EC2 servers to meet the current demand of traffic. Will save your money and meet capacity since you only run the amount of servers you need.

**Elastic Load Balancer (ELB)** <br/>
Distributes traffic to multiple instance, can re-route traffic from unhealthy instace to healthy instances, can route traffic to EC2 instances running in different Availability Zones.

**AWS Elastic Beanstalk (EB)** is for easily deploying web-application developers having to worry about seeting up and underlying the AWS Services. Similar to Heroku


## Types of Storage Service

### Elastic Block Store (EBS) - Block
- Data is split into evenly split blocks
- Directly accessed by the operation system
- Supports only a single write volume.

When you need a virtual ahrd drive attached to a VM

![image](https://user-images.githubusercontent.com/74575612/155006439-fc6bb7bd-02d7-4a8b-af04-71ac52200752.png)

### AWS Elastic File Storage (EFS) - File
- File is stored with data and metadata
- Multiple connections via network share
- Supports multiple reads, writing locks the file.

When you need a file-share where multiple users or vms need to access the same drive.

![image](https://user-images.githubusercontent.com/74575612/155006584-c490ac37-c545-4183-98ed-cb4dfcdaeb37.png)

### Amazon Simple Storage Service (S3) - Object
- Object is stored with data, metadata and unique ID
- Scales with limited no file limit or storage limit
- Supports multiple reads and writes (no locks)

When you just want to upload files and not have to worry about underlying infrastructure. Not intended for high IOPs

![image](https://user-images.githubusercontent.com/74575612/155006819-d6e454e9-d12e-4180-bd65-d160fcd8f386.png)

## Introduction to S3

**What is Object Storage (Object-based Storage)?**
Data storage architecture that manages data objects, as opposed to other storage architectures:
- **file systems** which manages data as a files and fire hierarchy.
- **block storage** which manages data as blocks within sectors and tracks.

_S3 provides you with unlimited storage._

**S3 Object**
- Objects contain your data, They are like files.
- Object may consist of:
  - **Key** this the name of the object
  - **Value** the data itself made up of a sequence of bytes
  - **Version ID** when versioning enabled, the version of object
  - **Metadata** additional information attached to the object
  
**S3 Bucket**
  - Buckets hold objects. Buckets can also have folders which in turn hold objects
  - S3 is a universal namespace so bucket names must be unique

_You can store an individual object from 0 bytes to 5 Terabytes in size._

## S3 Storage Classes

AWS offers a range of S3 storage classes that trade retrievel time, accessibility and durability for Cheaper Storage

- **S3 Standard (default)**
  - Fast! 99.99% availability, 11 9's durability, Relplication across at least three AZs.
- **S3 intelligent Tiering**
  - Uses ML to analyze object usage and determine the appropriate storage class.
  - Data is moved to the most cost-effective access tier, without any performance impact or added overhead.
- **S3 standard-IA (Infrequent Access)**
  - Still Fast! cheaper if you access files less than once a month.
  - Additional retrieval fee is applied. 50% less than standard (reduced availability)
- **S3 One-Zone-IA**
  - Still Fast! Objects only exist in one AZ. Availability is (99.5%), but cheaper than standard IA by 20% less (reduce durability) data could get destroyed. A retrievel fee is applied.
- **S3 Glacier**
  - For long-term cold storage. Retrieval of data take minutes to hours but the off is very cheap storage
- **S3 Glacier Deep Archive**
  - The lowest cost storage class, data retrievel time is 12 hours.

## AWS Snow Family

AWS snow family are storage and compute devices used to physically move data in or out the cloud. When moving data over the internet or private connection it too slow, difficult or costly

- **Snowcone**
  - Comes in two sizes
    - 8tb of storage
    - 14tb of storage
    
- **Snowball edge**
  - Comes generally in two type:
    - Storage optimized
      - 80tb
    - Compute optimized
      - 39.5 tb
      
- **Snowmobile**
  - 100 pb of storage

## Storage Services

- **Simple Storage Service (S3)** is a serverless storage object service. You can upload very large files and an unlimited amount
  of files. You pay for what you store. You don't worry about the underlying file-system, or upgrading the disk size.
  
- **S3 Glacier** is a cold storage service. It design as a low cost storage solution for archiving and long-term backup. It uses previous generation HDD drives to get that low cost. Its highly secure and durable.

- **Elastic Block Store(EBS)** is a persistent block storage service. It is a virtual drive in the cloud you attach to EC2 instances. You can coose different kinds of hard dirves: SSD, IOPS, throughput HHD, Cold HHD

- **Elastic File Storage (EFS)** is a cloud native NFS file system service. File storage you can mount to multiple EC2 instances at the same time. When you need to share files between multiple servers.

- **Storage Gateway** is a hybrid cloud storage service that extends your on-premise storage to cloud.

**File Gateway** extends your local storage to AWS S3 <br/>
**Volume Gateway** caches your local drives to S3 so you have a countinous backuo of local files in the cloud <br/>
**Tape Gateway** stores files onto virtual tapes for backing up your files on very cost effective long term storage.

**AWS Snow Family** are storage devices used to physically migrate large amounts of data to the cloud.
- **Snowball** and **Snowball Edge** are briefcase size data storage devices. **50-80 Terabytes**
- **Snowmobile** is a cargo container filled with racks of storage and compute that is transported via semi-trailer tractor truck to transfer up to **100PB** of data per trailer.
- **Snowcone** is a very small version of Snowball that can transfer **8TB** of data.

**AWS Backup** a fully managed backup service that makes it easy to centralize and automate the backup of data across multiple AWS services eg. EC2, EBS, RDS, DynamoDB, EFS, Storage Gateway. You can create backup plans

**CloudEndure Disaster Recovery** continuously replicates your machines into a low-cost staging area in your target AWS account and prefereed Region enabling fast and reliable recovery in case of IT data center failures.

**Amazon FSx** is a feature rich and highly-performant file system. That can be used for Windows (SMB) or Linux (Lustre).

- **Amazon FSx for Windows File Server** uses the SMB protocol and allows you to mount FSx to Windows servers

- **Amazon FSx for Lustre** uses Linux's Lustre file system and allows you to mount FSx to Linux servers.


## What is Database?
A database is a data-store that stores semi-structured and structured data.

A database is more complex data stores because it requires using formal design and modeling techniques.

Database can be generally categorized as either:
- Relational databases
  - Structured data that strongly represents tabular data (tables, rows and columns)
    - Row oriented or columar-oriented
- Non-relational databases
  - Semi-structured that may or may not distantly resemble tabular data

Databases have a rich set of functionality:
- specified language tp query (retrieve data)
- specialized modelling strategies to optimize retrieval for different use cases
- more fine tune control over the transformation or the data into useful data structures or reports

![image](https://user-images.githubusercontent.com/74575612/155084186-04a4662b-b894-4431-a618-e48aa030d814.png)

## What is Data Warehouse?
A relational datastore designed for analytic workloads, which is generally column-oriented data store

Data warehouses generally perform aggregation
- aggregation is grouping data
- data warehouse are optimized around columns since they need to quickly aggregate column data

Data warehoueses are generally designed be HOT
- Hot means they can return queries very fast even though they have vast amounts of data

Data warehouses are infrequently accessed meaning they aren't intented for real-time reporting but maybe once or twice a day or once a week to generate business and user reports.

A data warehouse needs to consume data from a relational databases on a regular basis.

![image](https://user-images.githubusercontent.com/74575612/155084773-f9ef5a95-c6e6-42de-8ddb-7b447260f1b8.png)

## What is a Key/Value store?
A **key-value database** is a type of non-relational database (NoSQL) that uses a simple key-value method to store data.

Key values stores are dumb and fast. <br/>
They generally lack features like:
- Relationships
- Indexes
- Aggregation

![image](https://user-images.githubusercontent.com/74575612/155085224-123c7505-9e57-4737-ac07-8d2ac1b1334a.png)

A simple key/value store will interpret this data resembling a dictionary 
![image](https://user-images.githubusercontent.com/74575612/155085325-7526d385-c665-4691-8ed0-30a5b550f94d.png)

A key/value store can resemble tabular data, it dos not have to have the consistent columns per row.

Due to their simple design they can scale beyond a relational database. <br/>
A key/value stores a unique key alongside a value

![image](https://user-images.githubusercontent.com/74575612/155085634-b593a200-82e8-446d-8cd3-a80deff88951.png)

## What is a Document Store?
A document store is a NoSQL database that stores documents as its primary data structure. <br/>
A document could be a XML but more commonly is JSON or JSON-Like <br/>
Document stores are sub-class of Key/Value stores

The components of a document store compared to Relational Database <br/>
![image](https://user-images.githubusercontent.com/74575612/155086220-780d09ad-3a46-45a1-b34d-a398ab59ecb3.png)

## NoSQL Database Service
**DynamoDB** is a serverless NoSQL key/value and document database. It is designed to scale to billions of records with guranteed consistent data return in at least a second. You don't have to worry about managing shards!

- **DynamoDB** is AWS's _flagship database service_ meaning whenever we think of a databse service that just scales, is cost effective and very fast we should think DynamoDB

_When we want a massively scalable database_

**DocumentDB** is a NoSQL document database that is "MongoDB compatible" <br/>
MongoDB is very popular NoSQL among developers. There were open-source licensing issues around using open-source MongoDB, so AWS got around it by just building their own MongoDB database.

**Amazon Keyspace** is a fully managed Apache Cassandra database. Cassandra is an open-source NoSQL key/value database similar to DynamoDB in that is columnar store database but has some additional functionality. <br/> 

_When you want to use Apache Cassandra_

## Relational Database Service

**Relational Database Service (RDS)** is a relational database service that supports multiple SQL engines. Relational is synonymous with SQL and Online Transactional Processing (OLTP). <br/>

Relational database are the most commonly used type of database among tech companies and start-up.

**RDS Supports the following SQL Engines:**
- **MySQL** - The most popular open-source SQL database that was purchased and now owned by Oracle
- **MariaDB** - When Oracle bought MySQL, MariaDB made a fork of MySQL was made under a different open-source license.
- **Postgres (PSQL)** - Most popular open-source SQL database amond developers. Has rich-features over MySQL but at added complexity.
- **Oracle** - Oracle's proprietary SQL database. Well used by Enterprice companies. You have to buy a license to use it.
- Microsoft SQL Server - Microsoft's proprietary SQL database. You have to buy a license to use it.
- Aurora - Fully managed database

**Aurora** is a fully managed database of either MySQL (5x faster) and PSQL (3x faster) database. <br/>
_When you want a highly available, durable, scalable and secure relational database for Postgres or MySQL_

**Aurora Serverless** is the serverless on-demand version of Aurora. <br/>
_When you want "most" of the benefits of Aurora but can trade to have cold-starts or you don't have lots of traffic demand_

**RDS on VMWare** allows you to deploy RDS supported engines to an an-premise data-center. The datacenter must be using VMWare for server virtualization. <br/>
_When you want databases managed by RDS on your own datacenter_

**Redshift** is a petabyte-size data warehouse. Data warehouses are for Onlince Analytical Processing (OLAP) <br/>
Data warehouses can be expensive because they are keeping data "hot". Meaning that we can run a very complex query and a large amount of data and get that data back very fast.

**ElastiCache** is a managed database of the in-memory and caching open-source databases Redis or Memcached. <br/>
_When you need to improve the performance of application by adding a chaching layer in-front of web-server or database._

**Neptune** is a managed graph database. Data is represented as interconnected nodes. <br/>
_When you need to understand the conncection between data eg. Mapping Fraud Rings or Social Media relationships._

**Amazon Timestreams** is a fully managed time series databse. <br/>
When you need to measure how things change over time.

**Amazon Quantum Ledge Database** is a fully managed ledger database that provides transaparent, immutable and cryptographically variable transaction logs. <br/>
_When you need to record history of financial activities that can be trusted._

**Database Migration Service (DMS)** is database migration service. You can migrate from:
- on-premise database to AWS
- from two database in different or same AWS accounts using different SQL engines.
- from SQL to NoSQL database


## Cloud-Native Networking Service

![image](https://user-images.githubusercontent.com/74575612/155519425-687601a2-d290-4a55-8b8b-76ea95ea464c.png)

- **VPC** - It is a logically isolated section of the AWS cloud where you can launch AWS resources
- **Route Tables** - it determines where network traffic from your subnets are directed
- **Internet Gateway** - Enable access to the internet
- **Region** - The geographical location of your network
- **AZ** - the data center of your aws resources
- **Subnets** - it is a logical partition of an IP network into multiple, smaller network segments
- **Security Groups** - acts a firewall at the instance level
- **NACLs** - acts as a firewalls at the subnet level.

## EnterpriseHhybrid networking

![image](https://user-images.githubusercontent.com/74575612/155519553-bf2ade6c-cea7-472b-abb0-ce13862cb35a.png)

- **DirectConnect** dedicated gigabit connection from on-premise data-center to AWS (a bery fast connection)
- **AWS Virtual Private Network (VPN)** a secure connection between on-premise, remote offices, mobile employees
- **PrivateLinks** - (VPC Interface Endpoints) keeps traffic within the AWS network and not traverse the internet to keep traffic secure.

## Virtual Private Cloud (VPC) and Subnets

**Virtual Private Cloud (VPC)** is a logically isolated section of the AWS Network where you launch you AWS resources. You choose a range of IPs using CIDR Range.

CIDR Range of 10.0.0.0/16 = 65,536 IP Addresses

![image](https://user-images.githubusercontent.com/74575612/155519932-4a27f118-b6ab-45ba-80ae-aa12bef7d31d.png)

**Subnets** a logical partition of an IP network into multiple smaller network segments. You are breaking up your IP range for VPC into smaller networks.
- a public subnet is one that can reach the internet
- a private subnet is one that cannot react the internet.

## Security Groups vs NACLs

- **Network Access Control Lists (NACLs)** - acts as a virtual firewall at the subnet level. You create allow and deny rules. eg Block a specific IP address known for abuse
- **Security Groups** - Acts as a virtual at the instance level implicitly denies all traffic. You create only allow rules. eg You cannot block a single IP address.

![image](https://user-images.githubusercontent.com/74575612/155520366-9d4c4bb6-c609-4405-ad5d-ddf6c3bd2f1a.png)

## Introduction to EC2

**Elastic Compute Cloud EC2** is a highly highly configurable virtual server. EC2 is resizable compute capacity. It takes minute to launch new instances. Anyhting and everything on AWS uses EC2 instance underneath.

- Choose os
- Choose instance type
- Add storage
- Configure instance

## EC2 Instance Families

**What are Instance Families?** <br/>
Instance families are different combinations of CPU, memory, storage and networking capacity.

Instance families allow you to choose the appropriate combinations of capacity to meet your application's unique requirements.

Different instance families are different because of the varying hardware used to give them their unique properties.

- General purpose
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

EC2 has three levels of tenancy

- Dedicated Host
  - your server lives here and you have control of the physical attributes
- Dedicated Instance
  - your server always lives here
- Default
  - you instance live here until reboot

## EC2 pricing models

there are 5 different ways to pay for EC2 (virtual machines)

- On Demand (Least Commitment)
  - low cost and flexible
  - only pay per hour or the \*second
  - short term, spiky, unpredictable workloads
  - cannot be interrupted
  - For the first time apps
- Spot upto 90% (Biggest Savings)
  - request spare computing capacity
  - flexible start and end times
  - can handle interruptions (server randomly stopping and starting)
  - for non-critical background jobs
- Reserved upto 75% off (Best Long-Term)
  - steady state or predictable usage
  - commit to ec2 over a 1 or 3 year term
  - can resell unused reserved instances
- Dedicated (Most Expensive)
  - dedicated servers
  - can be on demand or reserved or spot
  - when you need a guarentee of isolate hardware (enterprise requirements)

## AWS Savings Plan

savings plan offer you the similar discounts as reserved instances (RI) but simplifies the purchasing process

There are 3 types of savings plans:

- Compute savings plans
- EC2 instance savings plans
- sageMaker savings plan
