package cmd

import (
	"Moon_Trace/web"
	"github.com/spf13/cobra"
)

var webArgs web.Args
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "start moon trace web api",
	Run: func(cmd *cobra.Command, args []string) {
		web.Execute(&webArgs)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	// webCmd.Flags().StringVar()
	webCmd.Flags().StringVar(&webArgs.CertPemPath, "pem", "./conf/cert/server.pem", "证书pem")
	webCmd.Flags().StringVar(&webArgs.ConfPath, "config_path", "./conf/webConf.conf", "web配置文件")
}
