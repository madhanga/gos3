package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)

	bucket := "mf-core-files/DEV/bank-reverse-files/05082022/"
	item := "05082022AxisReverseFile.xls"

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create("./abc.xls")
	if err != nil {
		log.Fatalf("Unable to open file %q, %v", item, err)
	}

	defer file.Close()

	buf := aws.NewWriteAtBuffer([]byte{})
	_, err = downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})

	if err != nil {
		log.Fatal(err)
	}

	strData := string(buf.Bytes())
	scanner := bufio.NewScanner(strings.NewReader(strData))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
