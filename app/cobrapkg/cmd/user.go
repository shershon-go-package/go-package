/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/10/03 15:24
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// userCmd 父命令
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户操作",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("用户列表：", list)
	},
}

// 子命令(添加用户)
var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加用户；user add --name=?",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("添加用户:", name)
	},
}

// 子命令(删除用户)
var userDelCmd = &cobra.Command{
	Use:   "del",
	Short: "删除用户；user del --name=?",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("删除用户:", name)
	},
}

var (
	name string
	list []string
)

func init() {
	// 添加子命令到父命令
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userDelCmd)
	// 添加到根命令
	rootCmd.AddCommand(userCmd)
	// 父命令接收持久标志
	userCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "用户名")
	// 父命令接收本地标志
	userCmd.Flags().StringSliceVarP(&list, "list", "l", []string{}, "用户列表")
}
