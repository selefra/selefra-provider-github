package provider

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
	"os"
	"strings"

	"github.com/selefra/selefra-provider-github/constants"
	"github.com/selefra/selefra-provider-github/github_client"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Github,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var githubConfig github_client.Config
				err := config.Unmarshal(&githubConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}

				if githubConfig.AccessToken == "" {
					githubConfig.AccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
				}

				if githubConfig.AccessToken == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing access_token in configuration")
				}

				if len(githubConfig.Orgs) == 0 {
					orgData := os.Getenv("GITHUB_ORGS")

					var orgList []string

					if orgData != "" {
						orgList = strings.Split(orgData, ",")
					}

					githubConfig.Orgs = orgList
				}

				if len(githubConfig.Orgs) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing orgs in configuration")
				}

				clients, err := github_client.NewClients(githubConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# Required. Personal Access Token
# access_token: <YOUR_ACCESS_TOKEN_HERE>
# Required. List of organizations to extract from
# orgs:
#  - google
#  - aws`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var githubConfig github_client.Config
				err := config.Unmarshal(&githubConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_4,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
