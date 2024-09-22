resource "aws_security_group" "all_worker_mgmt" {
  name_prefix = "all_worker_management"
  vpc_id = module.vpc.vpc_id
}
resource "aws_security_group_rule" "all_worker_mgmt_ingress" {
  description = "Allow inbounf=d traffic from eks"
  from_port = 0
  protocol = "-1" # allow all protocols.
  to_port = 0
  security_group_id = aws_security_group.all_worker_mgmt.id
  type = "ingress"
  cidr_blocks = [ 
    "10.0.0.0/8", # private IP ranges used in AWS VPC.
    "172.16.0.0/12", # Another common private IP range.
    "192.168.0.0/16", # The last private IP range defined in RFC 1918.
     ]
}
resource "aws_security_group_rule" "all_worker_mgmt_egress" {
  description       = "allow outbound traffic to anywhere"
  from_port         = 0
  protocol          = "-1"
  security_group_id = aws_security_group.all_worker_mgmt.id
  to_port           = 0
  type              = "egress"
  cidr_blocks       = ["0.0.0.0/0"]
}
