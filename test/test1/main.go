/**
 * @Author: DollarKiller
 * @Description: fileDir upload test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:42 2019-10-09
 */
package main

import (
	"fmt"
	simple_aws_tools "github.com/dollarkillerx/simple-aws-tools"
	"log"
	"os"
	"time"
)

func main() {
	tools := simple_aws_tools.AwsTools{}
	tools.Init(simple_aws_tools.WithAccessKey("AKIAVH4VX2QL4I3BY7UV"),simple_aws_tools.WithRegion("eu-central-1"),simple_aws_tools.WithSecretKey("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"))

	upload := tools.InitUpload("pre-europe")

	// 遍历目录
	filelist := simple_aws_tools.GetFileList("img") // 返货目录下所有文件 的[]string
	for _,item := range filelist {
		file, e := os.Open(item)
		if e != nil {
			continue
		}
		defer file.Close()
		ko:
		e = upload.UploadFile(file, "fr-kok/"+item)
		if e != nil {
			log.Println(e)
			time.Sleep(time.Second * 3)
			goto ko
		}else {
			fmt.Println("上传完毕： "+item)
		}
	}
	log.Println("上传完毕")

}

