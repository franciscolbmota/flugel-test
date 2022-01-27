provider "aws" {
	region = "us-east-2"
}

#key definitions
resource "aws_key_pair" "flugel-key" {
	key-name = "Flugel"
	public_key = "ssh - rsa $PUBLIC_KEY$"
}

#security resource creation
resource "aws_security_group" "Flugel" {
	ingress {
		from_port	= 22
		to_port		= 22
		protocol	= "tcp"
		cidr_blocks	= ["0.0.0.0/0"] #change for the specific network address
	}

#s3 resource creation
resource "aws_s3_bucket" "Flugel" {
    bucket = "bucket-flugel"
    acl = "private"
    versioning {
            enabled = true
    }
    owners = ["InfraTeam"]
	tags {
        Name = "Flugel"
    }
	}	
}
#ec2 resource creation
resource "aws_instance" "app_server" {
	ami = "ami-830c94e3"
	instance_type = "t2.micro"

	owners = ["InfraTeam"]

	tags = {
		Name = "Flugel"
  }
}