resource "aws_s3_bucket" "b" {
    bucket = "2002my-site-api-v2-tf-bucket"
    acl = "public-read"

    cors_rule {
        allowed_methods = ["GET"]
        allowed_origins = ["*"]
        allowed_headers = ["*"]
        max_age_seconds = 3600
    }
}

resource "aws_s3_bucket_policy" "b" {
  bucket = aws_s3_bucket.b.id

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "IPAllow",
        "Effect" : "Allow",
        "Principal" : "*",
        "Action" : "s3:GetObject",
        "Resource" : "arn:aws:s3:::${aws_s3_bucket.b.bucket}/*"
      }
    ]
  })
}
