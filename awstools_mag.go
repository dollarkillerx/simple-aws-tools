/**
 * @Author: DollarKiller
 * @Description: tools mag 管理
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:27 2019-10-09
 */
package simple_aws_tools

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
)

type AwsTools struct {
	Session *session.Session
}

type AwsUpload struct {
	upload *s3manager.Uploader
	bucket string
}

// 初始化 aws upload
func (a *AwsTools) Init(opts ...Option) {
	options := Options{}
	for _,fu := range opts {
		fu(&options)
	}
	if options.EndPoint == "" {
		options.EndPoint = "s3.amazonaws.com"
	}
	// 初始化session
	a.initAws(options)
}

// 文件上传  会自动创建文件夹
func (a *AwsUpload) UploadFile(file io.Reader,fileName string) error {
	_, err := a.upload.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	return err
}

// 初始化session
func (a *AwsTools) initAws(options Options) {
	var err error
	a.Session, err = session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(options.AccessKey, options.SecretKey, ""),
		Endpoint:         aws.String(options.EndPoint),
		Region:           aws.String(options.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})
	if err != nil {
		log.Fatal(err)
	}
}

// 初始化upload
func (a *AwsTools) InitUpload(bucket string) *AwsUpload {
	return &AwsUpload{
		bucket:bucket,
		upload:s3manager.NewUploader(a.Session),
	}
}

