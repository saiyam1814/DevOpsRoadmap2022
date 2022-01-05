provider "aws" {
    region = "ap-south-1"
    profile = "test"
}

resource "aws_instance" "web" {
    ami = "ami-1233hfh"
    instance_type = "t2.micro"
    subnet_id = "sub-123uhfg"
    tags {
        name = "my-first-tf-instance"
    }
}
