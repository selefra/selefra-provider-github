package github_client

import (
	"context"
	"testing"

	"github.com/selefra/selefra-provider-github/constants"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) GithubServices, _ TestOptions) {
	testProvider := newTestProvider(t, table, builder)
	config := constants.Testtest
	test_helper.RunProviderPullTables(testProvider, config, constants.Constants_1, constants.Constants_2)
}

func newTestProvider(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) GithubServices) *provider.Provider {
	return &provider.Provider{
		Name:      constants.Github,
		Version:   constants.V,
		TableList: []*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				return []any{&Client{
					Github: builder(t, gomock.NewController(t)),
					Orgs:   []string{constants.Testorg},
				}}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `##  Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
#accounts:
#  - access_token: "xxx" # access_token
#    orgs: #  org list 
#      - "xxx"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_3,
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
