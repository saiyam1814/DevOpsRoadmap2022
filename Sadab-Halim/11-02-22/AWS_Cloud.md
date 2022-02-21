# Continued Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)

## Cloud Development Kit
AWS CDK allows you to use your favourite programming language to write Infrastructure as Code (IaC) <br/>
TypeScript, NodeJS, Python, Java, ASP.NET

- CDK is powered by CloudFormation (it generates out CloudFormation templates)
- CDK has a large library of reusable cloud components called CDK Construct
- CDK comes with it own CLI
- CDK Pipelines to quickly setup CI/CD pipelines for CDK projects
- CDK has a testing framework for Unit and Integration Testing

AWS SDK looks similar, but the key difference is CDK ensures Idempotent of your Infrastructure.

## AWS Toolkit for VSCode
AWS Toolkit is an open-source plugin for VSCode to create, debug, deploy AWS resources

1. **AWS Explorer**
  Explore a wide range of AWS resources to your linked AWS account.
2. **AWS CDK Explorer**
  Allows you to explore your stacks defined by CDK
3. **Amazon Elastic Container Service**
  Provides IntelliSense for ECS task-definition files
4. **Serverless Applications**
  Create, debug and deploy serverless applications via SAM and CFN

## Access Keys
Access Keys is a key and secret required to have programmatic access to AWS resources when interacting with the AWS API outside of the AWS Mangement Console

An Access Key is commonly referred to as AWS Credentials

A user must be granted access to use Access Keys
  
Access keys are to be stored in ~/.aws/credentials and follow a TOML file format
