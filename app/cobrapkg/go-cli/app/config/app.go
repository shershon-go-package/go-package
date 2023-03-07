/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/10/03 17:09
 */

package config

type AppConfig struct {
	App   app   `yaml:"app"`
	MySql mysql `yaml:"mysql"`
}
type app struct {
	Version string `yaml:"version"`
	Author  string `yaml:"author"`
	Port    string `yaml:"port"`
}
type mysql struct {
	Host     string `yaml:"host"`
	DataBase string `yaml:"data_base"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
