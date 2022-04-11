package cmd

import (
	"Moon_Trace/eng"
	"github.com/spf13/cobra"
)

var engArgs eng.Args
var engCmd = &cobra.Command{
	Use:   "eng",
	Short: "start eng",
	Run: func(cmd *cobra.Command, args []string) {
		eng.Execute(&engArgs)
	},
}

func init() {
	rootCmd.AddCommand(engCmd)
	engCmd.Flags().StringVar(&engArgs.ConfigPath, "config_path", "./conf/engConf.conf", "引擎配置文件")
	engCmd.Flags().StringVar(&engArgs.Addr, "addr", "127.0.0.1:60001", "grpc监听地址")
	engCmd.Flags().StringVar(&engArgs.CertPemPath, "pem", "./conf/cert/server.pem", "证书pem")
	engCmd.Flags().StringVar(&engArgs.CertKeyPath, "key", "./conf/cert/server.key", "证书key")
}
