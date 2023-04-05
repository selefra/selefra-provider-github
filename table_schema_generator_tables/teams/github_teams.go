package teams

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubTeamsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubTeamsGenerator{}

func (x *TableGithubTeamsGenerator) GetTableName() string {
	return "github_teams"
}

func (x *TableGithubTeamsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubTeamsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubTeamsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"id",
		},
	}
}

func (x *TableGithubTeamsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			opts := &github.ListOptions{PerPage: 100}
			for {
				repos, resp, err := c.Github.Teams.ListTeams(ctx, c.Org, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- repos
				opts.Page = resp.NextPage
				if opts.Page == resp.LastPage {
					break
				}
			}
			return nil
		},
	}
}

func (x *TableGithubTeamsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubTeamsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("members_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MembersURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("slug").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repos_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repositories_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoriesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("privacy").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ldap_dn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LDAPDN")).Build(),
	}
}

func (x *TableGithubTeamsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGithubTeamMembersGenerator{}),
		table_schema_generator.GenTableSchema(&TableGithubTeamRepositoriesGenerator{}),
	}
}
