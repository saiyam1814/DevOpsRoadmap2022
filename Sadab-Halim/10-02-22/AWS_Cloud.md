# Continued Learning for [AWS Certified Cloud Practitioner](https://www.youtube.com/watch?v=SOTamWNgDKc)


## AWS Command Line Interface (CLI)

**What is a CLI?** <br/>
A Command Line Interface (CLI) processes commands to a computer program in the form of lines of text. <br/>
Operating system implement a command-line interface in a shell.

**What is a Terminal?** <br/>
A terminal is a text only interface (input/output environment)

**What is a Console** <br/>
A Console is a physical computer to physically input information into a terminal.

**What is a Shell?** <br/>
A shell is the command line program that users interact with to input commands. <br/>
Popular shell programs:
- Bash
- Zsh
- PowerShell

AWS Command Line Interface (CLI) allows users to programtically interact with the AWS API via entering single or multi-line commands into a shell or terminal.

![image](https://user-images.githubusercontent.com/74575612/153470547-e3467b80-523a-434a-a363-2b3d6e816892.png)

The AWS CLI is a Python executable program.
- Python is required to install AWS CLI <br/>
The AWS CLI can be installed on Windows, Mac or Linux/Unix. <br/>
The name of the CLI program is aws

## AWS Software Development Kit (SDK)
A Software Development Kit (SDK) is a collection of software development tools in one installable package.

You can use the AWS SDK to programtically create, modify, delete or interact with AWS resources.

AWS SDK is offered in various programming languages:
- Java
- Python
- Node.js
- Ruby
- Go
- .NET
- PHP
- JavaScript
- C++

## AWS CloudShell
AWS CloudShell is a browser-based shell built into the AWS Management Console. <br/>
AWS CloudShell is scoped per region, Same credentials as logged in user. Free Service

**Preinstalled Tools** <br/>
AWS CLI, Python, Node.js, git, make, pip, sudo, tar, tmux, vim, wget, and zip and more

**Storage included** <br/>
1 GB of storage free per AWS region

**Saved files and settings** <br/>
Files saved in you home directory are available in future sessions for the same AWS region

**Shell Environments** <br/>
Seamlessly switch between
- Bash
- PowerShekk
- Zsh

## Infrastructure as Code (IaC)
You write a configuration script to automate creating, updating or destroying cloud infrastructure.
- IaC is a blueprint of your infrastructure.
- IaC allows you to easily sharem version or inventory your cloud infrastructure.

AWS has two offerings for writing infrastrucutre as a code.

**AWS CloudFormation (CFN)**
CFN is a Declarative IaC tool

**Declarative**
- What you see is what you get. **_Explicit_**
- More verbose, but zero chance of mis-configuration
- Uses scripting languages eg. JSON, YAML, XML

**AWS Cloud Development Kit (CDK)**
CDK is an Imperative IaC tool.

**Imperative**
- You say what you want, and the rest is filled in. **_Implicit_**
- Less verbosem you could end up with misconfiguration
- Does more than Declarative
- Uses programming languages eg. Python Rubym JavaScript

## CloudFormation
AWS CloudFormation allows you to write Infrastructure as Code (IaC) as either a JSON or YAML file.

CloudFormation is simple but it can lead to large files or is limited in some regard to creating dynamic or repeatable infrastructure compared to CDK.

CloudFormation can be easier for DevOps Engineers who do not have a backgroung in web programming languages.

Since CDK generates out CloudFormation its sill important to be able to read and understand CloudFormation in order to debug Iac stacks.

![image](https://user-images.githubusercontent.com/74575612/153473889-1c5df7f1-88e1-4699-93ea-81085c71113b.png)


