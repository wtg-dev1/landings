resource "aws_cloudfront_origin_access_control" "site" {
  name                              = "${local.bucket_name}-oac"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

# Short TTL for HTML — gives content updates a 5-min floor before a manual
# CloudFront invalidation lands. Browsers cache for 5 min; edge holds 1 day.
resource "aws_cloudfront_cache_policy" "html_short" {
  name        = "wtg-landings-html-short-${var.environment}"
  default_ttl = 300
  max_ttl     = 86400
  min_ttl     = 0

  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config { cookie_behavior = "none" }
    headers_config { header_behavior = "none" }
    query_strings_config { query_string_behavior = "none" }
    enable_accept_encoding_brotli = true
    enable_accept_encoding_gzip   = true
  }
}

# 1-day TTL for static. Assets are not fingerprinted, so a longer TTL would
# trap stale CSS until a manual invalidation. Revisit when filenames are
# fingerprinted.
resource "aws_cloudfront_cache_policy" "static_long" {
  name        = "wtg-landings-static-long-${var.environment}"
  default_ttl = 86400
  max_ttl     = 86400
  min_ttl     = 0

  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config { cookie_behavior = "none" }
    headers_config { header_behavior = "none" }
    query_strings_config { query_string_behavior = "none" }
    enable_accept_encoding_brotli = true
    enable_accept_encoding_gzip   = true
  }
}

resource "aws_cloudfront_response_headers_policy" "html" {
  name = "wtg-landings-html-headers-${var.environment}"

  custom_headers_config {
    items {
      header   = "Cache-Control"
      value    = "public, max-age=300, s-maxage=86400"
      override = true
    }
  }

  security_headers_config {
    strict_transport_security {
      access_control_max_age_sec = 31536000
      include_subdomains         = true
      preload                    = true
      override                   = true
    }
    content_type_options {
      override = true
    }
    referrer_policy {
      referrer_policy = "strict-origin-when-cross-origin"
      override        = true
    }
  }
}

resource "aws_cloudfront_response_headers_policy" "static" {
  name = "wtg-landings-static-headers-${var.environment}"

  custom_headers_config {
    items {
      header   = "Cache-Control"
      value    = "public, max-age=86400"
      override = true
    }
  }
}

resource "aws_cloudfront_distribution" "site" {
  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"
  price_class         = var.price_class
  aliases             = concat([var.domain_name], var.subject_alternative_names)
  comment             = "wtg-landings ${var.environment}"
  tags                = local.common_tags

  origin {
    domain_name              = aws_s3_bucket.site.bucket_regional_domain_name
    origin_id                = "s3-site"
    origin_access_control_id = aws_cloudfront_origin_access_control.site.id
  }

  # Default behavior = static assets. Most request volume by URL count is
  # /static/*, so cache them long by default and override for HTML paths.
  default_cache_behavior {
    target_origin_id           = "s3-site"
    viewer_protocol_policy     = "redirect-to-https"
    allowed_methods            = ["GET", "HEAD"]
    cached_methods             = ["GET", "HEAD"]
    compress                   = true
    cache_policy_id            = aws_cloudfront_cache_policy.static_long.id
    response_headers_policy_id = aws_cloudfront_response_headers_policy.static.id
  }

  # /<city>/* → HTML. Short TTL + HTML headers/security policy.
  ordered_cache_behavior {
    path_pattern               = "/brooklyn/*"
    target_origin_id           = "s3-site"
    viewer_protocol_policy     = "redirect-to-https"
    allowed_methods            = ["GET", "HEAD"]
    cached_methods             = ["GET", "HEAD"]
    compress                   = true
    cache_policy_id            = aws_cloudfront_cache_policy.html_short.id
    response_headers_policy_id = aws_cloudfront_response_headers_policy.html.id
  }

  # S3 returns 403 (not 404) for missing keys when public access is blocked.
  custom_error_response {
    error_code         = 403
    response_code      = 404
    response_page_path = "/404.html"
  }
  custom_error_response {
    error_code         = 404
    response_code      = 404
    response_page_path = "/404.html"
  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate_validation.site.certificate_arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2021"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
}
