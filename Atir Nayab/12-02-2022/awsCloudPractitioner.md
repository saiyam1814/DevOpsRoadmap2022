# aws free tier

## aws free tier

the aws free tier enables you to begin using certain services without having to worry about incurring costs for the specified period.

three types of offers are available

- always free
- 12 months free
- trials

### always free

these offers do not expire and are available to all aws customers

### 12 months free

these offers are free for 12 months following your initial sign-up date to aws.

### trials

short-term trial offers start from the date you activate a particular service. the length of each trial might vary by number of days or the amount of usage in the service.

# aws pricing concepts

## how aws pricing works

aws offers a range of cloud computing services with pay as you go pricing

### pay for what you use.

for each service, you pay for exactly the amount of resource that you actually use, without requiring long-term contracts or complex licensing

### pay less when you reserve

some services offer reservation options that provide a significant discount compared to on-demand instance pricing.

### pay less with volume based discounts when you use more

some services offer tiered pricing, so the per-unit cost is incrementally lower with increased usage.

## aws pricing calculator

the aws pricing calculator lets you explore aws services and create an estimate for the cost of your use cases on aws. you can organize your aws estimates by groups that you define. a group can reflect how your company is organized, such as providing estimates by cost center.

![calculator](../asset/calender.jpg)

## aws pricing examples

### aws lambda

for aws lambda, you are charged based on the number of requests for your functions and time that it take for them to run.

### amazon ec2

with amazon ec2 you pay for only the compute time that you use while your instances are running.

### amazon s3

- storage - you pay for only the storage that you use.
- requests and data retrievals - you pay for requests made to your amazon s3 objects and buckets.
- data transfer - there is no cost to transfer data between different amazon s3 buckets or from amazon s3 to other services within the same aws region.
- management and replication - you pay for the storage management features that you have enabled on your account amazon s3 buckets.
- management and replication - you pay for the storage management features that you have enabled on your account's amazon s3 buckets. these features include amazon s3 inventory, analytics, and object tagging.

# billing dash board

use the aws billing & cost management dash board to pay your aws bill, monitor your usage, and analyze and control your costs.

- compare your current month to date balance with the previous month, and get a forecast of the next month based on current usage.
- view month-to-date spend on service
- view free tier usage by service
- access cost ecplorer and create budgets
- purchase and manage savings plan.
- publish aws cost and usage reports.

# consolidated billing

## consolidated billing

the consolidated billing feature of aws organization enables you to receive a single bill for all aws accounts in your organization. by consolidating you can easily track the combined costs of all the linked accounts in your organization.

# aws budgets

## aws budgets

in aws budgets you can create budgets to plan your service usage, service costs, and instance reservations.

the information in aws budget updated three times a day. this helps you to accurately determine how close your usage is to your budgeted amounts or to the aws free tier limits.

# aws cost explorer

## aws cost explorer

aws cost explorer is tool that enables you to visualize, understand, and manage you aws costs and usage over time.

![cost explorer](../asset/cost.png)

# aws support plans

## aws support

aws offers four different support plans to help you troubleshoot issues, lower costs, and efficiently use aws services.

you can choose from the following support plan to meet your company's needs:

- basic
- developer
- business
- enterprise

### basic support

basic support is free for all aws customers. it includes access to whitepapers, documentation and support communities. with basic support you can also contact aws for billing questions and service limit increases.

with basic support, you have access to a limit selection of aws trusted advisor checks.

### developer support

- best practice guidance
- client side diagnostic tools
- building block architecture support, which consists of guidance for how to use aws offerings, features and services together.

### business support

- use case guidance to identify aws offerings, features and services that can best support your specific needs
- all aws trusted advisor checks
- limited support for third party software, such as common operating systems and application stack components

### enterprise support

- application architecture guidance, which is a consultative relationship to support your company's specific use cases and applications
- infrastructure event management: a short term engagement with aws support that helps your company gain a better understanding of your use cases.
- a technical account manager

### technical account manager (TAM)

if your company has an enterprise support plan, the tam is your primary point of contact at aws.

# aws marketplace

## aws marketplace

aws marketplace is digital catalog that includes thousands of software listing from independent software vendors. you can use aws marketplace to fing, test, and buy software that runs on aws.

## aws marketplace categories

- business application
- data & analytics
- devops
- infrastructure software
- internet of things (iot)
- machine learning
- migration
- security

# aws cloud adoption framework (aws caf)

## six core perspective of the cloud adoption framework

at the highest level, the aws cloud adoption framework (aws caf) organizes guidance into six areas of focus, called perspective. each perspective addresses distict responsibilities. the planning process helps the right people across the organization prepare for the changes ahead.

### business perspective

the business perspective ensures that it aligns with business needs and that it investment link to key business results.

use the business perspective to create a strong business case for cloud adoption and prioritize cloud adoption initiatives. ensure that your business strategies and goals aligns with your it strategies and goals.

common business perspective

- business managers
- finance managers
- budget owners
- strategy stakeholders

### people perspective

the people perspective supports development of an organization-wide change management strategy for successful cloud adoption

use the people perspective to evaluate organizational structure and roles, new skill and process requirements, and identify gaps. this helps prioritize training, staffing and organizational changes

commong roles in the people perspective include

- human resources
- staffing
- people managers

### gevernance perspective

the governament perspective focuses on the skills and processes to aling it strategy with business strategy. this ensure that you maximize the business value and minimize risks.

common riles in the governance perspective include

- chief information officer (cio)
- program managers
- enterprise architects
- business analysts
- portfolio managers

### platform perspective

the platform perspective includes principles and patterns for implementing new solutions on the cloud, and migrating on-premise workloads to the cloud.

common roles in the platform perspective include

- chief technology officer(cto)
- it managers
- solutions architects

### security perspective

the security perspective ensures that the organization meets security objectives for visibility, auditability, control and agility

common roles in the security perspective include

- chief information security officer (ciso)
- it security managers
- it security analysts

### operations perspective

the operation perspective helps you to enable, run, use operate and recovre it workloads to the level agreed upon with your business stakeholders

common roles in the operations perspective include

- it operations managers
- it support managers

# migration strategies

## 6 strategies for migration

when migrating applications to the cloud, six of the most common migration strategies that you can implement are:

- rehosting
- replatforming
- refactoring/re-architecting
- repurchasing
- retaining
- retiring

### rehosting

rehosting also known as lift-and-shift involves moving applications without changes.

in the scenario of a large legacy migrations in which the company is looking to implement its migration and scale quickly to meet a business case. the majority of applications are rehosted.

### replatforming

replatforming also known as lift tinker and shift, involves making a few cloud optimizations to realize a tangible benefit. Optimization is achieved without changing the core architecture of the application.

### refactoring/re-architecting

refactoring (also known as re-architecting) invlolves reimagining how an application is architected and developed by using cloud-native features. refactoring is driven by a strong business need to add features, scale or performance that would otherwise be difficult to achieve in the applications existing environment.

### repurchasing

repurchasing involves from a traditional license to a software as a service model.

### retaining

retaining consists of keeping applications that are critical for the business in the source environment. this might include applications that require major refactoring before they can be migrated, or, work that can be postponed until a later time.

### retiring

retiring is the process of removing applications that are no longer needed.

# aws snow family

## aws snow family members

the aws snow family is a collection of physical devices that help to physically transport up to exabytes of data into and out of aws.

aws snow family is composed of aws snowcone, aws snowball and aws snowmobile.

### aws snowcone

aws snowcone is a small, reugged and secure edge computing and data transfer device.

### aws snowball

aws snowball offers two types of devices

- snowball edge storage optimized devices are will suited for large scale data migrations and recurring transfer workflows, in addition to local computing with higher capacity needs.

  - storage - 80 tb of hdd for block volumes and amazon s3 compatible object storage and 1 tb ssd for block volumes.
  - compute: 40 vcpus and 80 gib of memory to support amazon ec2 sbe1 instances

- snowball edge compute optimized provides powerful computing resources for use cases such as machine learning, full motion video analysis, analytics, and local computing stacks,
  - storage: 42 tb usable hdd capacity for amazon s3 compatible object storage or amazon ebs compatible block volume and 7.6tb of usable nvme ssd
  - compute: 52 vcpus 208 gib of memory and optional nvidia tesla v100 gpu.

# innovation with aws

## innovate with aws sevices

when examining how to use aws services, it is important to focus on the desired outcomes, you are properly equiped to drive innovation in the cloud if you can clearly articulate the following conditions:

- the current state
- the desired state
- the problems you are trying to solve

### serverless applications

with aws serverless refers to applications that don't require you to provision. maintain or administer server. you don't need to worry about fault tolerance or availability. aws handles these capabilities for you.

### aritificial intelligence

aws affers a variety of services powered by artifical intelligence (ai)

### machine learning

traditional machine learning (ml) development is complex, expensive, time consuming and error prone. aws offers amazon sagemaker to remove the difficult work from the process and empower you to build train and deploy ml models quickly

# the aws well architected framework

## the aws well architected framework

the aws well architected framework helps you understand how to design and operate reliable, secure, effiecient and cost effective system in the aws cloud. it provides a way for you to consistently measure against best practices and design principles and identify areas improvement.

### operational excellence

operational excellence is the ability to run and monitor systems to deliver business value and to continuously improve supporting processes and procedures.

design principles for operational excellence in the cloud include performing operations as code, annotating documentation, anticipating failure, and frequently making small reversible changes.

### security

the security pillar is the ability to protect information, systems and assets while delivering business value through risk assessments and mitigation strategies.

when considering the security of your architecture, apply these best practices

- automate security best practices when possible
- apply security at all layers
- protect data in transit and at rest.

### reliability

reliability is the ability of a system to do the following

- recover from infrastructure or service desruptions
- dynamically acquire computing resources to meet demand
- mitigate disruptions such as misconfigurations or transient network issues.

### performance efficiency

performance efficiency is the ability to use computing resources efficiently to meet system requirements and to maintain that efficiency as demand changes and technologies evolve.

### cost optimization

cost optimizaiton is the ability to run system to deliver business value at the lowest price point

# benefits of the aws cloud

## advantages of cloud computing

### trade upfront expennse for variable expense

upfront expenses include data centers, physical servers, and other resources that you would need to invest in before using computing resources.

### benefit from massive economies of scale.

by using cloud computing, you can achieve a lower variable cost than you can get on your own

### stop guessing capacity

with cloud computing, you don't have to predict how much infrastructure capacity you will need before deploying an application.

### increase speed and agility

the flexibility of cloud computing makes it easier for you to develop and deploy applications

### stop spending money running and maintaining data centers

cloud computing in data centers often requires you to spend more money and time managing infrastructure and servers.

### gl global in minutes

the aws cloud global foorprint enables you to quickly deploy applications to customers around the world, while providing them with low latency
