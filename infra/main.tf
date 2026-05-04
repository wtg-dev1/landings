###############################################################################
# WTG landings — AWS infra
#
# Shape:
#   Route53 → ACM (us-east-1) → CloudFront (OAC) → S3 bucket of pre-rendered HTML
#
# The Go server is not deployed; pages are rendered to static files by
# `go run ./cmd/export` in CI and synced to S3. CloudFront serves them.
###############################################################################

provider "aws" {
  region = var.aws_region
}

# CloudFront only accepts ACM certs from us-east-1, regardless of where the
# rest of the stack lives. This aliased provider exists for that single
# resource path (cert + validation records).
provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"
}

locals {
  common_tags = {
    Project     = "wtg-landings"
    Environment = var.environment
    ManagedBy   = "terraform"
  }

  bucket_name = coalesce(var.bucket_name_override, "wtg-landings-${var.environment}-${data.aws_caller_identity.current.account_id}")
}

data "aws_caller_identity" "current" {}
