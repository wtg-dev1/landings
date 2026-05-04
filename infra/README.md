# Infra

Terraform stack for deploying the WTG landings site to AWS.

## Architecture

```
GitHub Actions  ──(OIDC role)──►  S3 (private)
                                   ▲
                                   │ OAC
Route53 ──► ACM (us-east-1) ──► CloudFront ──► viewers
```

No Go runtime in production. CI runs `go run ./cmd/export`, syncs the resulting `dist/` to S3, then invalidates CloudFront.

## Files

| File | Purpose |
|---|---|
| `versions.tf` | Terraform + AWS provider versions, S3 backend block (commented until bootstrap is run) |
| `main.tf` | Default + `us_east_1` aliased provider, common tags, derived bucket name |
| `s3.tf` | Site bucket: versioning, SSE, public-access-block, OAC-only bucket policy, lifecycle |
| `cloudfront.tf` | OAC, cache + response-header policies, distribution with `/brooklyn/*` HTML behavior |
| `dns.tf` | ACM cert (us-east-1), DNS validation records, Route53 alias for apex + each SAN |
| `github_oidc.tf` | OIDC provider, deploy role pinned to `repo:<org>/<repo>:ref:refs/heads/main` |
| `variables.tf` | Inputs |
| `outputs.tf` | Bucket name, distribution ID + domain, deploy role ARN, cert ARN |
| `bootstrap/` | One-time stack creating the state bucket + lock table (local state) |

## First-time setup

1. **Bootstrap state backend** (once per AWS account):
   ```sh
   cd infra/bootstrap
   terraform init
   terraform apply
   ```
   Note the `state_bucket` and `lock_table` outputs. Commit `terraform.tfstate` (it's small and changes rarely).

2. **Wire the backend** — uncomment the `backend "s3"` block in `versions.tf` and fill in `bucket` / `dynamodb_table` from step 1.

3. **Init the main stack**:
   ```sh
   cd infra
   terraform init
   ```

4. **Apply** with required vars (consider a `terraform.tfvars` file you don't commit, or pass `-var` flags):
   ```sh
   terraform apply \
     -var domain_name=lp.williamsburgtherapygroup.com \
     -var route53_zone_id=Z0123456789ABC \
     -var github_repo=your-org/landings
   ```

5. **Wire GitHub Actions** — set these as **repo variables** (not secrets; none are sensitive) under Settings → Secrets and variables → Actions → Variables:
   - `AWS_DEPLOY_ROLE_ARN` — from `github_actions_role_arn` output
   - `S3_BUCKET` — from `site_bucket_name` output
   - `CLOUDFRONT_DISTRIBUTION_ID` — from `cloudfront_distribution_id` output

6. **First deploy** — push to `master` (or run the workflow via `workflow_dispatch`). The workflow renders, syncs, and invalidates.

## Required variables

| Variable | Notes |
|---|---|
| `domain_name` | Apex CloudFront serves from |
| `route53_zone_id` | Hosted zone must already exist |
| `github_repo` | `<org>/<repo>` — pinned to `master` branch in the trust policy |

Optional: `subject_alternative_names` (e.g. `["www.lp.williamsburgtherapygroup.com"]`), `aws_region` (default `us-east-1`), `environment` (default `prod`), `price_class` (default `PriceClass_100`), `bucket_name_override`.

## Smoke tests

- Pre-DNS: `curl -I https://<cloudfront_domain_name>/brooklyn/therapy/` → 200, `cache-control: public, max-age=300, s-maxage=86400`
- Bucket is private: `curl -I https://<bucket>.s3.amazonaws.com/brooklyn/therapy/index.html` → 403
- Post-DNS: same `/brooklyn/*` URLs against `domain_name`

## Notes

- ACM cert lives in us-east-1 (CloudFront requirement) via the `aws.us_east_1` aliased provider — do not change.
- CloudFront's default-root-object only handles the apex; subdirectory URLs require an explicit `index.html` per route, which `cmd/export` writes.
- The 404 page is served from `/404.html`; `cmd/export` writes a minimal one.
- Static assets are not fingerprinted yet, so static cache TTL is held to 1 day. Bump after fingerprinting is added.
- Do **not** widen the OIDC `:sub` claim past `repo:<org>/<repo>:ref:refs/heads/main` — that would let any branch / any PR assume the deploy role.
