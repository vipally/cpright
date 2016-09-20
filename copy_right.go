//    CopyRight @Ally Dale 2016
//    Author  : Ally Dale(vipally@gmail.com)
//    Blog    : http://blog.csdn.net/vipally
//    Site    : https://github.com/vipally

//package cpright declaration CopyRight of vipally
package cpright

import (
	"github.com/vipally/cmdline"
)

//<version>
//<buildtime>
//<thiscmd>
var (
	copyRight = `CopyRight @Ally Dale 2016
    Author  : Ally Dale(vipally@gmail.com)
    Blog    : http://blog.csdn.net/vipally
    Site    : https://github.com/vipally
	BuildAt : <buildtime>
	Version : <version>
`
	version = "1.0.0"
)

func init() { //set copyright for vipally
	cmdline.CopyRight(copyRight)
}

//get the CopyRight declaration of vipally
func CopyRight() string {
	return copyRight
}

func Version() string {
	return version
}
