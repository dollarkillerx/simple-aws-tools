/**
 * @Author: DollarKiller
 * @Description: s3 test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:51 2019-10-08
 */
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func main() {
	accessKey := "AKIAVH4VX2QL4I3BY7UV"
	secretKey := "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"
	endPoint := "s3.amazonaws.com" //endpoint设置，不要动
	region := "eu-central-1"

	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endPoint),
		Region:           aws.String(region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})

	if err != nil {
		panic(err)
	}

	//svc := s3.New(sess)
	//result, err := svc.ListBuckets(nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, b := range result.Buckets {
	//	fmt.Printf("* %s created on %s\n",
	//		aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	//}
	//
	//for _, b := range result.Buckets {
	//	fmt.Printf("%s\n", aws.StringValue(b.Name))
	//}

	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("pre-europe"),
		Key:    aws.String("fr-kok/img/info/test.html"),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}

	log.Println("上传成功")


	//fp, err := os.Open("README.md")
	//if err != nil {
	//	panic(err)
	//}
	//defer fp.Close()
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	//defer cancel()
	//
	//server := s3.New(sess)
	//_, err = server.PutObjectWithContext(ctx, &s3.PutObjectInput{
	//	Bucket: aws.String("pre-europe"),
	//	Key:    aws.String("fr-kok/img/info/test.html"),
	//	Body:   fp,
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
}
