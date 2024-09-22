output "cluster_ip" {
    description = "EKS cluster ip"
    value = module.eks.cluster_id
}
output "cluster_endpoint" {
  description = "Endpoint for EKS control plane"
  value = module.eks.cluster_endpoint
}
output "cluster_security_group_id" {
 description = "Secut=rity group ids attached to the cluster control plane." 
 value = module.eks.cluster_security_group_id
}
output "region" {
  description = "AWS region"
  value = var.aws_region
}
output "oidc_provider_arn" {
  value = module.eks.oidc_provider_arn
}
