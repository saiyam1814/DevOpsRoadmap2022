# Continued Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

## AWS Global Network
The AWS Global Network represents the interconnections between AWS Global Infrastructure. <br/>
Commonly referred to as the _"Backbone of AWS"_

**Edge Locations** can acts as on and off ramps to the AWS Global Network
- **AWS Global Accelerator and AWS S3 Transfer Acceleration** <br/>
  Uses Edge Locations as an on-ramp to quickly reach AWS resources in other regions by traversing the fast AWS Global Network. 
- **Amazon CloudFront** (CDN) <br/>
  Uses Edge Location as an off-ramp, to provide at the Edge storage and compute near the end user.
  
## Point of Presence (PoP)
**Points of Presence (PoP)** is an intermediate location between an AWS Region and the end user, and this location could be a datacenter or collection of hardware.

For AWS a Point of Presnece is data center owned by AWS or a trusted partner that is utlized by AWS Services related for content delivery or expediated upload.

PoP resources are:
- Edge Locations
- Regional Edge Caches

![image](https://user-images.githubusercontent.com/74575612/152555406-b2321d8a-34f2-4ee2-b0de-9f9ec724831e.png)

**Edge Locations** are datacenters that hold cached (copy) on the most popular files (e.g. web pages, images and videos) so that the delivery of distance to the end users are reduce.

**Regional Edge Locations** are datacenters that hold much larger caches of less-popular files to reduce a full round trip and also to reduce the cost of transfer fees.

![image](https://user-images.githubusercontent.com/74575612/152556022-09ad8fe7-fec6-40eb-b228-73d0ec436ec0.png)

The following AWS Services use PoPs for content delivery or expediated upload.

**Amazon CloudFront** is a Content Delivery Network (CDN) that:
- You point your website to CloudFront so that it will route requests to nearest Edge Location cache
- Allows you to choose an origin (such as web-server or storage) that will be source of cached
- Caches the contents of what origin would returned to various Edge Location around the world.

**Amazon S3 Transfer Accelecration:** allows you to generate a special URL that can be used by end user to upload files to a nearby Edge Locationn. <br/>
Once a file is uploaded to an Edge Location, it can move much fster within the AWS Network to reach S3.

**AWS Global Accelerator** can find the optimal path from the end user to your web-servers. <br/>
Global Accelerator are deployed within Edge Locations so you send user traffuc to ab Edge Location instead if directly to you web-application.

## AWS Direct Connect
AWS Direct Connect is a private/dedicated connection between your datacenter, office, co-location and AWS.

![image](https://user-images.githubusercontent.com/74575612/152557524-d89c8b44-8e33-4b99-b594-7dc3a93117b3.png)

**Direct Connect** has two very-fast network connection options:
1. Lower Bandwidth 50MBps-500MBps
2. Higher Bandwith 1GBps or 10 GBps

- Helps reduce network costs and increase bandwidth throughput. (Great for high traffic networks)
- Provides a more consistend network experience than a typical internet based connection. (Reliable and secure)

**A co-location** is a data center where equipment, space, and bandwidth are available for rental to retail customers.

