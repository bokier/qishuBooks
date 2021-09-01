package setting

var Conf = new(Configs)

type Configs struct {
	*AppConfig   `mapstructure:"app"`
	*MysqlConfig `mapstructure:"mysql"`
}

type AppConfig struct {
	Version string `mapstructure:"version"`
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	MaxOpenCons int    `mapstructure:"max_open_cons"`
	MaxIdleCons int    `mapstructure:"max_idle_cons"`
}
