# instance stores and amazon elastic block store (amazon ebs)

## instance stores

an instance store provides temporary block-level storage for an amazon ec2 instance.an instance store is disk storage that is physically attached to the host computer for an ec2 instance, and therefore has the same lifespan as the instance. when the instance is terminated you lose any data in the instance store.

amazon ec2 instances are virtual servers. if you start an instance from a stopped state, the instance might start on another host, where the previously used instance store volume does not exist. therefore aws reccommends instance stores for use cases that involve temporary data that you do not need in the long term.

amazon elastic block store (ebs) is a service that provides block-level storage volumes that you can use with amazon ec2 instances. if you stop or terminate an amazon ec2 instance, all the data on the attached ebs volume remains available.

to create an ebs volume, you define the configuration (such as volume and type) and provision it. after you create an ebs volume, it can attach to an amazon ec2 instance.

because ebs volumes are for data that needs to persis, it's important to back up the data. you can take incremental backups of ebs volumes by creating amazon ebs snapshots.

## amazon ebs snapshots

an ebs snapshot is an incremental backup. this means that the first backup taken of a volume copies all the data. for subsequent backups, only the blocks of data that have changed since the most recent snapshots are saved,

incremental backup are different from full backups, in which all the data in a storage volume copies each time a backup occurs. the full backup includes data that has not changed since the most recent backup.

# amazon simple storage service (amazon s3)

in object storage, each object consists of data, metadata, and a key

the data might be an image, video, text document or any other type of file, metadata contains information about what the data is, how it is used the object size, and so on. an object's key is its unique identifier.

_when a file in object storage is modified the entire object is updated_

## amazon simple storage service (s3)

amazon simple storage service (amazon s3) is a service that provides object level storage. amazon s3 stores data as object in buckets.

you can upload any type of file to amazon s3, such as images, videos text files and so on.

when you upload a file to amazon s3 you set permissions to control visibility and access to it. you can also use the amazon s3 versioning feature to track changes to your objects over time.

## amazon s3 storage classes

with amazon s3 you pay only for what you use. you choose from a range of storage classes to select a fit for your business and cost needs. when selecting an amazon s3 storage consider these two factors:

- how often you plan to retrieve your data
- how available you need your data to be.

### s3 standard

- designed for frequently accesses data
- stores data in a minimum of three availability zones

s3 standard provides high availability for objects. this makes it a good choice for a wide range of use cases, such as websites content distribution and data analytics. s3 standard has a higher cost than other storage classes intended for infrequently accessed data and archival storage.

### s3 standard-infrequent access (s3 standard-ia)

- ideal for infrequent accesses data
- similar to se standard but has a lower storage price and higher retrivel price.

s3 standard-ia ideal for data infrequently accesses but requires high availability when needed. both s3 standard and s3 standard ia store data in a minimum of three availability zones. s3 standard-ia provides the same level of availability as s3 standard but with a lower storage price and a higher retrievel price.

### s3 one zone-infrequent access (s3 one zone ia)

- stores data in a single availability zone
- has a lower storage price than s3 standard-ia

compared to s3 standard and s3 standard-ia. which store data in a minimum of three availability zones, s3 one availability zone-ia stores data in a single availability zone. this makes it a good storage class to consider if the following conditions apply

- you want to save costs on storage
- you can easily reproduce you data in the event of an availability zone failure.

### s3 intelligent tiering

- ideal for data unknown or changing access patterns
- requires a small monthly monitoring and automation fee per object

in the s3 intelligent tiering storage class, amazon s3 monitors object's access patterns. if you haven't accessed an object for 30 consecutive days. amazon s3 automatically moves it to the infrequent access tier, s3 standard-ia. if you access an object in the infrequent access tier. amazon s3 automatically moves it to the frequent access tier, s3 standard

### glacier

- low cost storage designed for data archiving
- able to retrieve objects within a few minutes to hours

s3 glacier is a low cost storage class that is ideal for data archiving. for eg you might use this storage class to store archieved customer records or older photos and video files.

### s3 glacier deep archive

- lowest cost object storage class ideal for archiving
- able to retrieve objects within 12 hours

when deciding between amazon s3 glacier and amazon s3 glacier deep archive, consider how quickly you need to retrieve archived objects. you can retrieve objects stored in the s3 glacier storage class within a few minutes to a few hours. by comparison you can retrieve objects stored in the s3 glacier deep archives storage class within 12 hours.

# amazon elastic file system (amazon efs)

## file storage

in file storage multiple clients can access data that is stored in shared file folders. in this approach a storage server uses block storage with a local file system to organize files. clients access data through file paths

compared to block storage and object storage file storage is ideal for use cases in which a large number of services and resources need to access same data at the same time.

amazon elastic file system (amazon efs) is a scalable file system used with aws cloud services and on-premises resources. as you add and remove files, amazon efs grows and shrinks automatically. it can scale on demand to perabytes without disrupting applications.

## comparing amazon ebs and amazon efs

### ebs

an amazon ebs volume stores data in a single availability zone

to attach an amazon ec2 instances to an ebs volume, both the amazon ec2 instance and the ebs volume must reside within the same availability zone.

### efs

amazon efs is a regional service. it stores data in and across multiple availability zones.

the duplicate storage enables you to access data concurrently form all the availability zones in the region where a file system is located. additionally on-premises servers can access amazon efs using aws direct connect.

# amazon relational database service (amazon rds)

## relational database

in a relational database, data is stored in a way that relates it to other pieces of data.

relational database use structured query language (sql) to store and query data. this approach allows data to be stored in an easily understandable, consistent, and scalable way.

## amazon relational database service

amazon relational database service (amazon rds) is a service that enables you to run relational databases in aws cloud.

amazon rds is a managed service that automates tasks such as hardware provisioning. database setup, patching and backups. with these capabilities, you can spend less time completing administrative tasks and more time using data to innovate your applications. you can integrate amazon rds with other services to fulfill your business and operational needs such as using aws lambda to query your database from a serverless application.

amazon rds provides a number of different security options. many amazon rds database engines offer encryption at rest (protecting data while it is stored) and encryption in transit (protecting data while it is being sent and received)

## amazon rds database engines

amazon rds is available on six database engines, which optimize for memory, performance or input/output (i/o). supported database engine include

- amazon aurora
- postgresql
- mysql
- mariadb
- oracle database
- microsoft sql server

## amazon aurora

amazon aurora is an enterprise-class relational database. it is compatible with mysql and postgresql relational databases. it is up to five times faster than standard mysql databases and upto three times faster than standard postgresql databases.

amazon aurora helps to reduce your database costs by reducing unnecessary input/output (i/o) operations, while ensuring that your database resource remain reliable and available

consider amazon aurora if your workloads require high availability. it replicates six copies of your data across three availability zones and continuously backs up your data to amazon s3.

# amazon dynamoDb

## nonrelational databases

in a nonrelational database, you create tables. a tables is a place where you can store and query data.

nonrelational databases are sometimes referred to as nosql databases becuase they use structure other than rows and columns to organize data. one type of structural approach for nonrelational databases is key-value pairs. with a key-value pairs, data is organized into items (keys) and items have attributes. you can think of attributes as being different features of your data.

## amazon dynamodb

amazon dynamodb is a key value database service. it delivers single digit milisecond performance at any scale.

dynamodb is a serverless, which means that you do not have to provision patch or manage servers

you also do not have to install maintain or operate software

as the size of your database shrinks or grows, dynamodb automatically scales to adjust for changes in capacity while maintaining consistent performance.

this makes it a suitable choice for use cases that require high performance while scaling.

# amazon redshift

## amazon redshift

amazon redshift is a data warehousing service that you can use for big data analytics. it offers the ability to collect data from many sources and helps you to understand relationships and trends across your data

# aws database migration service

## aws database migration serivce (aws dms)

aws database migration service (aws dms) enables you to migrate relational databases, nonrelational databases, and ohter types of data stores.

with aws dms you move data between a source database and a target database. the source and target database can be of the same type or different types. during the migration your source database remains operational reducing downtime for any applications that rely on the database

## other use cases for aws dms

enabling developers to test applications against production data without affecting production users.

cobining serveral databases into a single database

sending ongoing copies of your data to other target sources instead of doing a on-time migration

# additional database services

## amazon documentdb

amazon documentdb is a document database service that supports mongodb workloads.

## amazon neptune

amazon neptune is a graph database service

you can use amazon neptune to build and run applications that work with highly coonected datasets, such as recommendation engines, fraud detection, and knowledge graphs.

## amazon quantum ledger database (amazon qldb)

amazon quantum ledger database (amazon qldb) is a ledger database service.

you can use amazon qldb to review a complete history of all the changes that have been made to your application data.

## amazon managed blockchain

amazon managed blockchain is a service that you can use to create and manage blockchain networks with open-source frameworks

blockchain is a distributed ledger system that lets multiple parites run transactions and share data without a central authority.

## amazon elasticache

amazon elasticache is a service that adds caching layers on top of your databases to help improve the read times of common requests.

it supports two types of data stores; redis and memcached

## amazon dynamodb accelator

amazon dynamodb accelerator (dax) is an in-memory cache for dynamodb. it helps improve response time for single digit miliseconds to microseconds.

# shared responsibility model

## the aws shared resposibility model

the shared responsibility model divides into customer responsibility (commonly referred to as "security" in the cloud) and AWS responsibilities (commonly referred to as security of the cloud)

![shared responsibility](../asset/sharedResp.png)

you can think of this model as being similar to the division of responsibilities between a homeowner and a homebuilder.

### customers: security in the cloud

customers are responsible for the security of everything that they create and put in the aws cloud

the security steps that you take will depend on factors such as the services that you use, the complexity of your systems and your company's specific operational and security needs. steps include selecting configuring and patching the operating systems that will run on amazon ec2 instances, configuring security groups and managing user accounts.

### aws: security of the cloud

aws is responsible for security of the cloud

aws operates manages and controls the components at all layers of infrastructure. this includes areas such as the host operating system, the virtualization layer, and even the physical security of the data centers from which services operate.

aws is responsible for protecting the global infrastructure that runs all of the services offered in the aws cloud. this infrastructure includes aws regions, availability zones and edge locations

aws manages the security of the cloud specifically the physical infrastructure that hosts your resources which include

- physical security of data centers
- hardware and software infrastructure
- network infrastructure
- virtualization infrastructure
