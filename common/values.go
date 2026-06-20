package common

type Values struct {
	Address  string   `yaml:"address"`
	Database Database `yaml:"database"`
}

type Database struct {
	Path         string `yaml:"path"`
	Victorialogs string `yaml:"victorialogs"`
}
