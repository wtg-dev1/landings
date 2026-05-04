output "site_bucket_name" {
  value       = aws_s3_bucket.site.id
  description = "S3 bucket the GitHub Actions deploy syncs into. Set as repo variable S3_BUCKET."
}

output "cloudfront_distribution_id" {
  value       = aws_cloudfront_distribution.site.id
  description = "CloudFront distribution ID. Set as repo variable CLOUDFRONT_DISTRIBUTION_ID."
}

output "cloudfront_domain_name" {
  value       = aws_cloudfront_distribution.site.domain_name
  description = "CloudFront-assigned domain (use for pre-DNS smoke testing)."
}

output "github_actions_role_arn" {
  value       = aws_iam_role.github_actions_deploy.arn
  description = "IAM role GitHub Actions assumes via OIDC. Set as repo variable AWS_DEPLOY_ROLE_ARN."
}

output "acm_certificate_arn" {
  value       = aws_acm_certificate_validation.site.certificate_arn
  description = "Validated ACM cert ARN attached to CloudFront."
}
