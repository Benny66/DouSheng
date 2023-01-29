// @Author: Ciusyan 2023/1/23
package conf

//=====
// 项目配置汇总
//=====

// Config 将配置文件抽成一个对象
type Config struct {
	App   *app   `toml:"app"`
	Log   *log   `toml:"log"`
	MySQL *mySQL `toml:"mysql"`
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMySQL(),
	}
}

// 防止配置文件在运行时被更改，设置为私有的
var config *Config

// C 获取总的配置对象
func C() *Config {
	return config
}
