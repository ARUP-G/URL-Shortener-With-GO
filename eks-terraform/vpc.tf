provider "aws" {
  region = var.aws_region
}

# Retrive the list of azs in the region, used in vpc creation
data "aws_availability_zones" "available" {}
  # filter {
  #   name   = "zone-name"
  #   values = ["us-west-1a", "us-west-1b"] 
  # }


locals {
  cluster_name = "ard-eks-${random_string.suffix.result}"
#   cluster_name = var.cluster_name
}
# Generates a random string resource
resource "random_string" "suffix" {
  length = 8
  special = false 
}
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "5.13"

  name = "ard-eks-vpc"
  cidr = var.vpc_cidr
  azs = data.aws_availability_zones.available.names
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
  public_subnets = ["10.0.4.0/24" , "10.0.5.0/24"]

  # Enables a NAT Gateway for the VPC, 
  # allowing private subnets to access the internet.  
  enable_nat_gateway = true
  # Configures a single NAT Gateway instead of multiple (to save costs).
  single_nat_gateway = true
  enable_dns_hostnames = true
  enable_dns_support = true
  
  tags ={
  "kubernetes.io/cluster/${local.cluster_name}" = "shared"
  }
  public_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    # Marks the public subnets for use by AWS Elastic Load Balancers (ELBs).
    "kubernetes.io/role/elb" = "1"
  }
  private_subnet_tags ={
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    # Marks private subnets for use by internal Elastic Load Balancers (ELBs).
    "kubernets.io/role/internal-elb" = "1"
  }

}

