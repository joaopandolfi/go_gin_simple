package configurations

type Configurations struct {
	Port string
}

var Config Configurations

func Load() {
	Config = Configurations{
		Port: ":8090",
	}
}
