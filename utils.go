/**
 * @Author: DollarKiller
 * @Description: 工具库
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:53 2019-10-09
 */
package simple_aws_tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileList(path string) []string {
	data := make([]string,0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {return nil}
		data = append(data,path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return data
}
