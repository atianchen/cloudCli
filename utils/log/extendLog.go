package log

type EsConfig struct {
	Adress string
	Port   string
}

func NewEsConfig() *EsConfig {
	entry := &EsConfig{
		Adress: "",
		Port:   "",
	}
	return entry
}
