/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动http服务,使用方法: app server --port=?",
	Run: func(cmd *cobra.Command, args []string) {
		if port == "" {
			fmt.Println("port不能为空!")
			os.Exit(-1)
		}
		engine := gin.Default()
		_ = engine.Run(":" + port)
	},
}

var port string

func init() {
	// 添加到根命令
	rootCmd.AddCommand(serverCmd)
	// 接收参数
	serverCmd.Flags().StringVar(&port, "port", "", "端口号")
}
