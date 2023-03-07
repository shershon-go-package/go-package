/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/10/03 17:01
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"shershon1991/go-cli/app/config"
)

func init() {
	// 初始化配置信息
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./app.yaml | ./config/app.yaml)")
}

// 根命令
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "命令行的简要描述...",
	Long:  `学习使用Cobra,开发cli项目,app: 指的是编译后的文件名。`,
}

var (
	cfgFile   string
	appConfig *config.AppConfig
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// 接收指定的配置文件
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// 设置配置文件目录(可以设置多个,优先级根据添加顺序来)
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("./app/config")
		// 设置配置文件
		viper.SetConfigType("yaml")
		viper.SetConfigName("app")
	}
	// 读取环境变量
	viper.AutomaticEnv() // read in environment variables that match
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig: %+v\n", err)
	}
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Printf("%+v\n", appConfig)
}
