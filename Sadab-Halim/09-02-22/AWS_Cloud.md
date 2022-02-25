# Continued Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

## AWS Application Programming Interface (API)
**What is an Application Programming Interface (API)?** <br/>
An API is software that allows two applications/services to talk to each other. <br/>
The most common type of API is via HTTP/S requests.

AWS API is an HTTP API and you can interact by sending HTTPS requests, using an application interacting with APIs like **Postman.**

Each AWS Service has its own **Service Endpoint** which you send requests
```
GET/HTTP/1.1
host: **monitoring.us-east-1.amazonaws.com**
x-amz-target: GraniteServiceVersion20100801.GetMetricData
x-amz-date: 20180112T092034Z
Authorization: **AWS4-HMAC-SHA256 Credential=REDACTEDREDACTED/20180411/...**
Content-Type: application/json
Content-Encoding: amz-1.0
Content-Length: 45
Connection: keep-alive
```

To authorize you will need genereated signed requests <br/>
You make a separate request with your AWS credentials and get back a token.

You need to also provide an **ACTION** and accompanying **parameters** as the payload

Rarely do users directly send HTTP requests directly to the AWS API. <br/>
Its much easier to interact with the API via a variety of Developer Tools

![image](https://user-images.githubusercontent.com/74575612/153446210-7e1d9a58-d3e5-426a-91d9-b1a0ecaa81d7.png)

## AWS Management Console
The AWS Management Console is a web-based unified console <br/>
**Build, manage, and monitor everything** from simple web apps to complex cloud deployments.

Point and Click to manually launch and configure AWS resources with limited programming knowledge.

This is known as **"ClipOps"** since you can perform all your system operations via clicks.

## AWS Management Console -- Service Console
AWS Service each have their own customized console. <br/>
You can access these consoles by **searching** the service name.

![image](https://user-images.githubusercontent.com/74575612/153447258-373ccfe5-f8a2-4c24-b47f-df160656f759.png)

The EC2 Console
![image](https://user-images.githubusercontent.com/74575612/153447341-3d36fdbb-1223-4daa-bcd7-c500418f8f2b.png)

Some AWS Services Console will act as an umbrella console containing many AWS Services: eg. <br/>
- VPC Console
- EC2 Console
- Systems Manager Console
- SageMaker Console
- CloudWatch Console

### AWS Account ID
Every AWS Account has a unique Account. <br/>
The **Account ID** can be easily found by dropping down the current user in the Global Navigation.

![image](https://user-images.githubusercontent.com/74575612/153448439-f8531e78-193e-4a3a-b48a-c359718a57e3.png)

The AWS Account ID id composed of 12 digits eg:
- 123456789012
- 121212121212
- 498241098510

![image](https://user-images.githubusercontent.com/74575612/153448506-b2632d62-2022-4bad-b870-bdf915b6645d.png)

The AWS Account ID is used: 
- when logging in with a non-root user account.
- Cross-account roles
- Support cases

## AWS Tools for PowerShell
**What is PowerShell?** <br/>
**PowerShell** is a task automation and configuration management framework. <br/>
A command-line shell and a scripting language.

Unlike most shells, which accept and return text, PowerShell is built on top of the .NET Common Language Runtime (CLR), and accepts and returns .NET objects.

**AWS Tools for PowerShell** lets you interact with the AWS API via PowerShell Cmdlets

Cmdlet is a special type of command in PowerShell in the form of capitalized verb-and-noun e.g. _New-S3Bucket_

## Amazon Resoure Name (ARNs)
**Amazon Resource Names (ARNs)** uniquely identify AWS resources. <br/>
ARNs are required to specify a resource unambigoudly across all of AWS

The ARN has the following **format variations:**
- arn: _partition:service:region:account-id:resource-id_
- arm: _partition:service:region:account-id:resource-type/resource-id_
- arn: _partition:service:region:account-id:resource-type:resource-id_

Partition
- aws - AWS Regions
- aws-cn - China Regions
- aws-us-gov - AWS GovCloud (US) Regions

Service - Identifies the service
- EC2
- S3
- IAM

Region - which AWS resource
- us-east-1
- ca-central-1

Account ID
- 121212121212
- 123456789012

Resource ID <br/>
Could be a number name or path:
- user/Bob
- instance/i-1234567890abcdef0

In the AWS Management COnsole its common to be able to copy the ARN to your clipboard.

## Path in ARNs
Resource ARNs can include a path <br/>
Paths can include a wildcard character, namely an asterisk (*)

**IAM Policy ARN Path** <br/>
arn:aws:iam::123456789012:user/Development/product_1234/*

**S3 ARN Path** <br/>
arn:aws:s3:::my_corportate_bucket/Development/*

