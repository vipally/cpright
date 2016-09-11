package cpright

import (
	"github.com/vipally/cmdline"
)

var (
	copyRight = `CopyRight @Ally Dale(vipally@gmail.com) 2016
    Author  : Ally Dale(vipally@gmail.com)
    Blog    : http://blog.csdn.net/vipally
    Site    : https://github.com/vipally
`
)

func init() { //set copyright for vipally
	cmdline.CopyRight(copyRight)
}

//get the CopyRight declaration for vipally
func CopyRight() string {
	return copyRight
}
