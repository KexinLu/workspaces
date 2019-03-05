package config_model


type Config struct {
	BaseDir string `mapstructure:"workspace_dir" json:"workspace_dir"`
	Projects map[string]Project `mapstructure:"projects" json:"projects"`
}

type Project struct {
	Name string `mapstructure:"name" json:"name"`
	Path string `mapstructure:"path" json:"path"`
	IsGit bool `mapstructure:"git" json:"git"`
}
