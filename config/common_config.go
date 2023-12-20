package config

type Common struct {
	CommonMysql CommonMysql `mapstructure:"common-mysql"`
}

type CommonMysql struct {
	DSN string `mapstructure:"dsn"`
	//DSN string `toml:"dsn"`
}

//type CommonRedis struct {
//}
