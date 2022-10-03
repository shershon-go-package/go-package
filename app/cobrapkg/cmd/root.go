package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// 接受参数
	rootCmd.PersistentFlags().String("version", "", "版本")
}

// rootCmd represents the base command when called without any subcommands
// 根命令
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "命令行的简要描述....",
	Long: `学习使用Cobra,开发cli项目,
- app: 指的是编译后的文件名。`,
	//// 根命令执行方法，如果有要
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("args:",args)
	//},
}

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
