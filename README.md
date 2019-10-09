# simple-aws-tools
simple-aws 简单AWS  工具盒


## Aws S3

### Aws S3术语
- Region: 存储数据所在的地理区域
- Endpoint: 存储服务入口,Web服务入口点的URL
- Bucket: 存储桶S3中用于存储对象的容器
- Object: 对象是S3中存储的基本实体,由对象数据和原数据组成
- Key: 键是存储桶中对象的唯一标识符,桶内的每个对象都只能有一个key

### 安装
``` 
go get github.com/dollarkillerx/simple-aws-tools
```

### 上传
``` 
	tools := simple_aws_tools.AwsTools{}
	tools.Init(simple_aws_tools.WithAccessKey("xxxxxx"),simple_aws_tools.WithRegion("xxxx"),simple_aws_tools.WithSecretKey("xxxxxxx"))

	upload := tools.InitUpload("pre-europe")

	// 遍历目录
	filelist := simple_aws_tools.GetFileList("test") // 返货目录下所有文件 的[]string
	for _,item := range filelist {
		file, e := os.Open(item)
		if e != nil {
			continue
		}
		defer file.Close()
		upload.UploadFile(file,item)
	}
```