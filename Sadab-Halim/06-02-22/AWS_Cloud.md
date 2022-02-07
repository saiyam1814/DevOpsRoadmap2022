# Continued Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

## AWS Ground Station
**AWS Ground Station** is a fully managed service that lets you control satellite communications, process data, and scale your operations without having to worry about buulding or managing your own ground station infrastructure.

Use Cases for Ground Station: <br/>
- weather forecasting
- surface imaging
- communications
- video broadcasts

To use Ground Station: <br/>
- You schedule a COntact (select satellite, start and end time, and the ground location).
- Use the AWS Ground Station EC2 AMI to launch EC2 instances that will uplink and downlink data during the contact or receive downlinked data in an Amazon S3 bucket.

Use Case Example: <br/>
A company reaches an agreement with a Satellite Imagery Provider to take satellite photos of a specifi region. They use AWS Ground Station to communicate that company's Satellite and download the S3 image data.

## AWS Outposts
**AWS Outposts** is a fully managed service that offers the same AWS infrastrcuture, AWS services, APIs and tools to virtually any datacenter, co-location space, or on-premises facility for a truly consistent hybrid experience.

AWS Outposts is rack of servers running AWS Infrastructure on your physical location.

**What is a Server Rack?** <br/>
A frame designed to hold and organize IT equipment

**Rack Heights** <br/>
U stands for "rack units" or "U spaces" with is equal to 1.75 inches. <br/>
The industry standard rack size is 48U (7 Foot Rack).

### AWS Outposts comes in 3 form factors: 42U, 1U and 2U
42U:This is a full rack of servers provided by AWS <br/>

![image](https://user-images.githubusercontent.com/74575612/152841200-9922b71d-ebff-4897-a27d-95a0aa2a3dd5.png)

AWS delivers it to your preferred physical site fully assembled and ready to be rolled into final position. <br/>
It is installed by AWS and the rack needs to be simply plugges into power and network.

1U and 2U are servers that you can place into your existing racks:
1U <br/>
Suitable for 19-inch wide 24-inch deep cabinets <br/>
AWS Graviton2 (up to 64 vCPUs) <br/>
128 GiB memory <br/>
4 TB of local NVMe storage

2U <br/>
Suitable for 19-inch wide <br/>
36-inch deep cabinets <br/>
Intel processor (up to 128 vCPUs) <br/>
256 GiB memory <br/>
8 TB of local NVMe storage


