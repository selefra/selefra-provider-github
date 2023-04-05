package repositories

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubRepositoryDependabotAlertsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubRepositoryDependabotAlertsGenerator{}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetTableName() string {
	return "github_repository_dependabot_alerts"
}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"repository_id",
			"number",
		},
	}
}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			repo := task.ParentRawResult.(*github.Repository)

			alerts, _, err := c.Github.Dependabot.ListRepoAlerts(ctx, c.Org, *repo.Name, nil)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- alerts

			return nil
		},
	}
}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dismissed_comment").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fixed_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("FixedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dependency").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_vulnerability").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dismissed_by").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_id").ColumnType(schema.ColumnTypeInt).
			Extractor(github_client.ExtractorParentField("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dismissed_reason").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("github_repositories_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to github_repositories.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_advisory").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dismissed_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("DismissedAt")).Build(),
	}
}

func (x *TableGithubRepositoryDependabotAlertsGenerator) GetSubTables() []*schema.Table {
	return nil
}
