package billing

import (
	"context"

	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubBillingPackageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubBillingPackageGenerator{}

func (x *TableGithubBillingPackageGenerator) GetTableName() string {
	return "github_billing_package"
}

func (x *TableGithubBillingPackageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubBillingPackageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubBillingPackageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
		},
	}
}

func (x *TableGithubBillingPackageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			billing, _, err := c.Github.Billing.GetPackagesBillingOrg(ctx, c.Org)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- billing
			return nil
		},
	}
}

func (x *TableGithubBillingPackageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubBillingPackageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("included_gigabytes_bandwidth").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_gigabytes_bandwidth_used").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_paid_gigabytes_bandwidth_used").ColumnType(schema.ColumnTypeInt).Build(),
	}
}

func (x *TableGithubBillingPackageGenerator) GetSubTables() []*schema.Table {
	return nil
}
