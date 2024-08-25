package types

// 阿里 oss 配置
type AliOssConfig struct {
	AccessUrl string `mapstructure:"accessUrl"`
	Endpoint  string `mapstructure:"endpoint"`
	Ak        string `mapstructure:"ak"`
	Sk        string `mapstructure:"sk"`
	Bucket    string `mapstructure:"bucket"`
}
