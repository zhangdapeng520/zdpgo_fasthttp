package zdpgo_fasthttp

/*
@Time : 2022/6/24 16:08
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description:
*/

type Config struct {
	ReadTimeout         int `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout        int `yaml:"write_timeout" json:"write_timeout"`
	MaxIdleConnDuration int `yaml:"max_idle_conn_duration" json:"max_idle_conn_duration"`
}
