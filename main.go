package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/shakinm/xlsReader/xls"
)

func main() {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)

	bucket := "mf-core-files/DEV/bank-reverse-files/05082022/"
	//item := "05082022AxisReverseFile.xls"
	//item := "05082022SBIReverseFile_286538.xls"
	item := "286538_04-08-2022_04-08-2022_0408110650 (1).xls"

	downloader := s3manager.NewDownloader(sess)
	buf := aws.NewWriteAtBuffer([]byte{})

	_, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})

	if err != nil {
		log.Fatal(err)
	}

	//strData := string(buf.Bytes())
	//fmt.Println(strData)

	reader := bytes.NewReader(buf.Bytes())
	workbook, err := xls.OpenReader(reader)
	if err != nil {
		log.Panic(err.Error())
	}
	sheet, err := workbook.GetSheet(0)
	if err != nil {
		log.Panic(err.Error())
	}

	for i := 0; i <= sheet.GetNumberRows(); i++ {
		if row, err := sheet.GetRow(i); err == nil {
			if cell, err := row.GetCol(1); err == nil {

				// Значение ячейки, тип строка
				// Cell value, string type
				fmt.Println(cell.GetString())
			}
		}
	}

	/*file, err := excelize.OpenReader(strings.NewReader(strData), excelize.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()*/

}
