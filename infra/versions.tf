terraform {
  required_version = ">= 1.6.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  # Populated after running infra/bootstrap once. See infra/README.md.
  backend "s3" {
    bucket         = "wtg-landings-tfstate-455999870532"
    key            = "main/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "wtg-landings-tflock"
    encrypt        = true
  }
}
