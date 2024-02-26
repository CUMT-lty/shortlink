package config

// 这个后期要从配置文件中读

type Common struct {
	CommonMysql CommonMysql `mapstructure:"common-mysql"`
}

type CommonMysql struct {
	DSN string `mapstructure:"dsn"`
	//DSN string `toml:"dsn"`
}

//type CommonRedis struct {
//}
