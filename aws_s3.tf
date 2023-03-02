resource "aws_s3_bucket" "bucket1" {
    bucket = "2002my-site-api-v2-tf-bucket"
    acl = "private"
}