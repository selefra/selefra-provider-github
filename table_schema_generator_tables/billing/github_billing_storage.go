package billing

import (
	"context"

	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubBillingStorageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubBillingStorageGenerator{}

func (x *TableGithubBillingStorageGenerator) GetTableName() string {
	return "github_billing_storage"
}

func (x *TableGithubBillingStorageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubBillingStorageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubBillingStorageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
		},
	}
}

func (x *TableGithubBillingStorageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			billing, _, err := c.Github.Billing.GetStorageBillingOrg(ctx, c.Org)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- billing
			return nil
		},
	}
}

func (x *TableGithubBillingStorageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubBillingStorageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("days_left_in_billing_cycle").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("estimated_paid_storage_for_month").ColumnType(schema.ColumnTypeFloat).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("estimated_storage_for_month").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableGithubBillingStorageGenerator) GetSubTables() []*schema.Table {
	return nil
}
