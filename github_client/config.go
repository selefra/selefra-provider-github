package github_client

type Config struct {
	AccessToken string   `yaml:"access_token,omitempty" mapstructure:"access_token"`
	Orgs        []string `yaml:"orgs,omitempty" mapstructure:"orgs"`
}
