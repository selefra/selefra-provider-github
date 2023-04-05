package organizations

import (
	"context"

	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubOrganizationDependabotSecretsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubOrganizationDependabotSecretsGenerator{}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetTableName() string {
	return "github_organization_dependabot_secrets"
}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"name",
		},
	}
}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)

			secrets, _, err := c.Github.Dependabot.ListOrgSecrets(ctx, c.Org, nil)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- secrets.Secrets

			return nil
		},
	}
}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selected_repositories_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelectedRepositoriesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("github_organizations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to github_organizations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("CreatedAt")).Build(),
	}
}

func (x *TableGithubOrganizationDependabotSecretsGenerator) GetSubTables() []*schema.Table {
	return nil
}
