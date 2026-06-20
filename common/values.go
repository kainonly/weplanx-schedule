package common

type Values struct {
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
}

type Database struct {
	Path         string `yaml:"path"`
	Victorialogs string `yaml:"victorialogs"`
}
