module "eks" {
  source = "terraform-aws-modules/eks/aws"
  version = "20.24.1"
  cluster_name = local.cluster_name
  cluster_version = var.Kubernetes_version
  subnet_ids = module.vpc.private_subnets
  enable_irsa = true # IAM Roles for Service Accounts (IRSA)
  tags = {
    cluster = "demo"
  }
  vpc_id = module.vpc.vgw_id

  eks_managed_node_group_defaults = {
    ami_type = "AL2_x86_64"
    instance_types         = ["t3.medium"]
    vpc_security_group_ids = [aws_security_group.all_worker_mgmt.id]
  }
  eks_managed_node_groups = {
    node_group = {
        min_size     = 2
        max_size     = 2
        desired_size = 2
    }
  }
}