###############################################################################
# Bootstrap stack — runs ONCE per AWS account, with LOCAL state.
#
# Creates the S3 bucket + DynamoDB table the main stack uses for its own
# remote state. Chicken-and-egg: you can't put the state for "create the
# state bucket" in the state bucket itself.
#
# Run order:
#   cd src/infra/bootstrap
#   terraform init
#   terraform apply -var aws_region=us-east-1
#
# Then commit `terraform.tfstate` (or store it durably). It changes
# rarely — only when these two resources change.
###############################################################################

terraform {
  required_version = ">= 1.6.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  type        = string
  default     = "us-east-1"
  description = "Region for the state bucket + lock table. Pick once and never change."
}

variable "state_bucket_name" {
  type        = string
  default     = null
  description = "Override the auto-generated bucket name. Leave null to derive from account ID."
}

data "aws_caller_identity" "current" {}

locals {
  bucket_name = coalesce(var.state_bucket_name, "wtg-landings-tfstate-${data.aws_caller_identity.current.account_id}")
  tags = {
    Project   = "wtg-landings"
    Component = "tfstate"
    ManagedBy = "terraform"
  }
}

resource "aws_s3_bucket" "tfstate" {
  bucket = local.bucket_name
  tags   = local.tags
}

resource "aws_s3_bucket_versioning" "tfstate" {
  bucket = aws_s3_bucket.tfstate.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "tfstate" {
  bucket = aws_s3_bucket.tfstate.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "tfstate" {
  bucket                  = aws_s3_bucket.tfstate.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_dynamodb_table" "tflock" {
  name         = "wtg-landings-tflock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = local.tags
}

output "state_bucket" {
  value       = aws_s3_bucket.tfstate.id
  description = "Plug into infra/versions.tf backend \"s3\" block as `bucket`."
}

output "lock_table" {
  value       = aws_dynamodb_table.tflock.name
  description = "Plug into infra/versions.tf backend \"s3\" block as `dynamodb_table`."
}

output "region" {
  value       = var.aws_region
  description = "Plug into infra/versions.tf backend \"s3\" block as `region`."
}
