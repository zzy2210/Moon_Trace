package main

import (
	"Moon_Trace/subdomain"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main(){
	logo := `
 __  __                     _____                   
|  \/  | ___   ___  _ __   |_   _| __ __ _  ___ ___ 
| |\/| |/ _ \ / _ \| '_ \    | || '__/ _  |/ __/ _ \
| |  | | (_) | (_) | | | |   | || | | (_| | (_|  __/
|_|  |_|\___/ \___/|_| |_|___|_||_|  \__,_|\___\___|	
                        |_____|
                                                     ————————author:y1nhui
`
color.Cyan(logo)

app := cli.NewApp()
app.Name ="Moon_Trace"
app.Version = "1.0.0"
app.Usage = "A easy tool framework "

var target string = "nil"

app.Flags = []cli.Flag{
	cli.StringFlag{ // 目标url的flag，参数绑定到target
		Name:        "u",
		Usage:       "target url",
		Value:       "nil",
		Destination:&target,
	},
	cli.BoolFlag{ // 参数判定，启用子域查询
		Name:        "sub",
		Usage:       "to find subdimain",
	},
}

app.Action = func(c *cli.Context) {

	if c.Bool("sub"){  //调用子域查询
		subdomain.FindSubdomain(target)
	}
}


err :=app.Run(os.Args)
if err!= nil {
	log.Fatal(err)
}



}
