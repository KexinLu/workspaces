package config_model


type Config struct {
	BaseDir string `mapstructure:"workspace_dir"`
	Projects []Project `mapstructure:"projects"`
}

type Project struct {
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
	IsGit bool `mapstructure:"git"`
}

