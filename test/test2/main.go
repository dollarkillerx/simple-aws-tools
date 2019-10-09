/**
 * @Author: DollarKiller
 * @Description: 并发版上传
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:03 2019-10-09
 */
package main

import (
	"fmt"
	simple_aws_tools "github.com/dollarkillerx/simple-aws-tools"
	"log"
	"os"
	"sync"
	"time"
)

var (
	chan1 chan string
	end chan string
	end2 chan string
	tou string
	num int
)

func main() {
	chan1 = make(chan string,100000)
	end = make(chan string,0)
	end2 = make(chan string,0)

	fmt.Print("存储文件名称:")
	fmt.Scan(&tou)
	fmt.Print("协程数:")
	fmt.Scan(&num)


	go grenerurl(end2)
	go upload(end)

	<-end2
	<- end
	log.Println("全部上传完毕")
}

func grenerurl(end chan string) {
	fileList := simple_aws_tools.GetFileList("img") // 返货目录下所有文件 的[]string
	for _,path := range fileList {
		// 下发任务
		chan1 <- path
	}
	close(chan1)
	end<-"end"
}

func upload(end chan string) {
	tools := simple_aws_tools.AwsTools{}
	tools.Init(simple_aws_tools.WithAccessKey("AKIAVH4VX2QL4I3BY7UV"),simple_aws_tools.WithRegion("eu-central-1"),simple_aws_tools.WithSecretKey("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"))

	upload := tools.InitUpload("pre-europe")

	wg := sync.WaitGroup{}
	ch := make(chan int,num)

	bb:
	for {
		select {
		case data,ok:= <- chan1:
			if ok {
				ch<-1
				wg.Add(1)
				go func(url string) {
					defer func() {
						wg.Done()
						<-ch
					}()
					file, e := os.Open(url)
					if e != nil {
						return
					}
					defer file.Close()
					ko:
					e = upload.UploadFile(file, tou + "/"+url)
					if e != nil {
						log.Println(e)
						time.Sleep(time.Second * 10)
						goto ko
					}else {
						fmt.Println("上传完毕： "+url)
					}
				}(data)
			}else {
				wg.Wait()
				end <- "end"
				break bb
			}
		}
	}

}