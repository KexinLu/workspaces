package config_model

import (
	"fmt"
)

type Config struct {
	BaseDir string `mapstructure:"workspace_dir" json:"workspace_dir"`
	Projects map[string]Project `mapstructure:"projects" json:"projects"`
}

type Project struct {
	Name string `mapstructure:"name" json:"name"`
	Alias string `mapstructure:"alias" json:"alias"`
	Path string `mapstructure:"path" json:"path"`
	IsGit bool `mapstructure:"git" json:"git"`
}

func (c *Config) HasProject(id string) bool {
	if _, ok := c.Projects[id]; ok {
		return true
	} else {
		for _, p := range c.Projects {
			if p.Alias == id {
				return true
			}
		}

		return false
	}
}

func (c *Config) GetProject(noa string) *Project {
	if v, ok := c.Projects[noa]; ok {
		return &v
	} else {
		for _, p := range c.Projects {
			if p.Alias == noa {
				return &p
			}
		}

		return &Project{}
	}
}

func (c *Config) RemoveProject(id string) error {
	if _, ok := c.Projects[id]; ok {
		delete(c.Projects, id)
	} else {
		for _, p := range c.Projects {
			if p.Alias == id {
				delete(c.Projects, p.Name)
				return nil
			}
		}
		return fmt.Errorf("project with name or alias not founrd: %s", id)
	}
	return nil
}

func (c *Config) AliasProject(name string, alias string) error {
	if p, ok := c.Projects[name]; ok {
		p.Alias = alias
		c.Projects[name] = p
		return nil
	}
	return fmt.Errorf("project with name not founrd: %s", name)
}
