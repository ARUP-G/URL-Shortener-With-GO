variable "Kubernetes_version" {
  default = 1.28
  description = "Kubernetes version"
}
variable "vpc_cidr" {
  default = "10.0.0.0/16"
  description = "VPC CIDR"
}
variable "aws_region" {
  default = "us-west-1"
  description = "aws region"
}
# For custom cluster name
# variable "cluster_name" {
#   default = "ard-cluster-default"
#   description = "Name of the cluster given in terminal"
#   type = string
# }