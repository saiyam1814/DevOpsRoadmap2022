## Zero trust model

The zero trust model operates on the principal of "trust no one, verify everything"

Malicious actors being able to by-pass conventional access controls demonstrates traditional security measures are no longer sufficient.

In zero trust model Identity becomes the primary security perimeter.

What is the primary security perimeter?
The primary or new security perimeter defines the first line of defense and its security controls that protect a company's cloud resources and assets.

## Zero Trust on aws

Identity security controls you can implement on AWS to meet the zero trus model

- IAM policies
- Permission boundaries
- service control policies (organization-wide policies)
- IAM policy conditions
  - aws:SourceIp - Restrict on IP address
  - AWS:RequestedRegion - Restrict on Region
  - aws:MultifactorAuthPresent - restrict if MFA is turned off
  - aws: currentTime - restrict access based on time of day

## zero trust on AWS with third parties

aws does technically implement a zero trust model but does not allow for onteligent identity security controls.

for eg azure active directory has Real tiem and calculated risk detection based more data points than aws

## Directory service

what is directory service?\
A directory service maps the names of network resouces to their network addresses.

A directory service is shared information infrastructure for locating, managing, administering and organizing resources.

A directory service is a critical component of a network operating system. A directory service is a server which provides a directory service.

## Identity providers (IdPs)

Identity Provider (IdP) a system entity that creates, maintains and manages identity information for principals and also provides authentication services to applications within a federation or distributed network. A trusted provider of your identity that lets you use authenticate to access other services. Identity providers could be facebook, amazon, google, twitter, github, lindkIn
