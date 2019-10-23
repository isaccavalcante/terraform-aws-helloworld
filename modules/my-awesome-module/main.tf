provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "my_instance" {
  ami                    = "ami-0a313d6098716f372" # ubuntu 18.04
  instance_type          = "t2.small"

  tags = {
    Name = "My Instance"
    MyCostTag = "Free Tier"
  }
}