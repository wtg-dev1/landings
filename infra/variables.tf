variable "aws_region" {
  type        = string
  default     = "us-east-1"
  description = "AWS region for the S3 bucket and most resources. ACM stays in us-east-1 regardless via an aliased provider."
}

variable "environment" {
  type        = string
  default     = "prod"
  description = "Environment name — used for tags and the default bucket name."
}

variable "domain_name" {
  type        = string
  description = "Apex domain CloudFront serves from (e.g. lp.williamsburgtherapygroup.com)."
}

variable "subject_alternative_names" {
  type        = list(string)
  default     = []
  description = "Additional hostnames on the cert + CloudFront aliases (e.g. [\"www.lp.williamsburgtherapygroup.com\"])."
}

variable "route53_zone_id" {
  type        = string
  description = "Hosted zone ID for the apex domain. Must already exist."
}

variable "github_repo" {
  type        = string
  description = "GitHub repo allowed to assume the deploy role, in '<org>/<repo>' form."
}

variable "price_class" {
  type        = string
  default     = "PriceClass_100"
  description = "CloudFront price class. PriceClass_100 = NA + EU edges only."
}

variable "bucket_name_override" {
  type        = string
  default     = null
  description = "Override the auto-generated bucket name. Set this once and don't change it — the bucket is created at apply time."
}
