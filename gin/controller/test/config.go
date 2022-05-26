package test

type OkConfig struct {
}

func NewOkConfig() OkInter {
	var okConfig OkInter = &OkConfig{}
	return okConfig
}
