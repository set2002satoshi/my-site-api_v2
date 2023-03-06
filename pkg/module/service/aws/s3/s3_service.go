package s3

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/rs/xid"
)

func UploadImage(file *multipart.FileHeader) (string, string, error) {
	return "http://iamge.iamge", "img_key", nil
}

func DeleteImage(URL string) error {
	return nil
}

func DeleteUserImage(objKey string) error {
	bucket, sess, err := initS3()
	if err != nil {
		return nil
	}
	svc := s3.New(sess)
	_, err = svc.DeleteObject(
		&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(objKey),
		},
	)
	if err != nil {
		return fmt.Errorf("not delete")
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objKey),
	})
	if err != nil {
		return err
	}
	return nil

}

func UploadUserImage(useType, nickname, email string, file multipart.File) (string, string, error) {
	bucket, sess, err := initS3()
	if err != nil {
		return "", "", fmt.Errorf("error init: %v", err)
	}

	uploader := s3manager.NewUploader(sess)
	objectKey, err := objectKeyGeneration(useType, nickname, email)
	if err != nil {
		return "", "", err
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", "", fmt.Errorf("error upload: %v", err)
	}
	fmt.Println(objectKey)
	return objectKey, fmt.Sprintf("https://2002my-site-api-v2-tf-bucket.s3.amazonaws.com/%s", objectKey), nil

}

func objectKeyGeneration(useType, nickname, email string) (string, error) {
	key := fmt.Sprintf("%s-%s", nickname, email)
	guid := xid.New()
	return fmt.Sprintf("%s/%s/%s", useType, key, guid), nil
}

func initS3() (string, *session.Session, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", nil, err
	}
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(os.Getenv("AWS_REG")),
	})
	if err != nil {
		return "", nil, fmt.Errorf("error creating session: %v", err)
	}
	bucket := os.Getenv("AWS_BUCKET")
	return bucket, sess, err
}
