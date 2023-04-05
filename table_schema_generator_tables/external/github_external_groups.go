package external

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubExternalGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubExternalGroupsGenerator{}

func (x *TableGithubExternalGroupsGenerator) GetTableName() string {
	return "github_external_groups"
}

func (x *TableGithubExternalGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubExternalGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubExternalGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"group_id",
		},
	}
}

func (x *TableGithubExternalGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			opts := &github.ListExternalGroupsOptions{
				ListOptions: github.ListOptions{
					Page:    0,
					PerPage: 100,
				},
			}
			for {
				groups, resp, err := c.Github.Teams.ListExternalGroups(ctx, c.Org, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- groups.Groups
				opts.Page = resp.NextPage
				if opts.Page == resp.LastPage {
					break
				}
			}
			return nil
		},
	}
}

func (x *TableGithubExternalGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubExternalGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("GroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableGithubExternalGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
