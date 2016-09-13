//    CopyRight @Ally Dale 2016
//    Author  : Ally Dale(vipally@gmail.com)
//    Blog    : http://blog.csdn.net/vipally
//    Site    : https://github.com/vipally

//package cpright declaration CopyRight of vipally
package cpright

import (
	//"strings"

	"github.com/vipally/cmdline"
)

var (
	copyRight = `CopyRight @Ally Dale 2016
    Author  : Ally Dale(vipally@gmail.com)
    Blog    : http://blog.csdn.net/vipally
    Site    : https://github.com/vipally
`
	//copyRightCode = "//    " + strings.Replace(copyRight, "\n", "\n//", 3)
)

func init() { //set copyright for vipally
	cmdline.CopyRight(copyRight)
}

//get the CopyRight declaration of vipally
func CopyRight() string {
	return copyRight
}

//get the CopyRight declaration of vipally for comment go code
//func CopyRightCode() string {
//	return copyRightCode
//}
