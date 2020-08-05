package Global

import "github.com/urfave/cli"

// 仅用于存储变量方便调用
var Moon *cli.App




func init(){
	Moon = cli.NewApp()
	Moon.Name ="Moon_Trace"
	Moon.Version = "1.0.0"
	Moon.Usage = "A easy tool framework "
}