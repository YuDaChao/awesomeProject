/**
* @Author: YuDC
* @Date: 2019-06-25 23:42
* @Description:
 */
package main

import (
	"fmt"
	// 导入包
	"awesomeProject/helloword/demo_01/model"
)

func main() {

	// 包名.标示符(和文件没有关系)
	fmt.Println("访问model中的变量名 = ", model.UserName)
}
