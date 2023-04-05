package organizations

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubOrganizationMembersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubOrganizationMembersGenerator{}

func (x *TableGithubOrganizationMembersGenerator) GetTableName() string {
	return "github_organization_members"
}

func (x *TableGithubOrganizationMembersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubOrganizationMembersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubOrganizationMembersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"id",
		},
	}
}

func (x *TableGithubOrganizationMembersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			opts := &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 100}}
			for {
				members, resp, err := c.Github.Organizations.ListMembers(ctx, c.Org, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- members
				opts.Page = resp.NextPage
				if opts.Page == resp.LastPage {
					return nil
				}
			}
		},
	}
}

func (x *TableGithubOrganizationMembersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGithubOrganizationMembersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("followers").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_private_repos").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ldap_dn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("following_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FollowingURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("blog").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_repos").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscriptions_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubscriptionsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("github_organizations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to github_organizations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("twitter_username").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_usage").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("text_matches").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gists_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GistsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_gists").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("followers_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FollowersURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gravatar_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GravatarID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspended_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SuspendedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collaborators").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("two_factor_authentication").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("events_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("site_admin").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("following").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bio").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owned_private_repos").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_gists").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organizations_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OrganizationsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("avatar_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvatarURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hireable").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("received_events_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReceivedEventsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("starred_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StarredURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("membership").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					c := client.(*github_client.Client)

					m := result.(*github.User)
					membership, _, err := c.Github.Organizations.GetOrgMembership(ctx, *m.Login, c.Org)
					if err != nil {
						return nil, err
					}
					return membership, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plan").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repos_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReposURL")).Build(),
	}
}

func (x *TableGithubOrganizationMembersGenerator) GetSubTables() []*schema.Table {
	return nil
}
