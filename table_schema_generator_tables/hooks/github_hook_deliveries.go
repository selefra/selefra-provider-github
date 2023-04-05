package hooks

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubHookDeliveriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubHookDeliveriesGenerator{}

func (x *TableGithubHookDeliveriesGenerator) GetTableName() string {
	return "github_hook_deliveries"
}

func (x *TableGithubHookDeliveriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubHookDeliveriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubHookDeliveriesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"hook_id",
			"id",
		},
	}
}

func (x *TableGithubHookDeliveriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			id := *task.ParentRawResult.(*github.Hook).ID

			c := client.(*github_client.Client)
			opts := &github.ListCursorOptions{PerPage: 100}

			for {
				deliveries, resp, err := c.Github.Organizations.ListHookDeliveries(ctx, c.Org, id, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- deliveries

				opts.Cursor = resp.NextPageToken
				if resp.NextPageToken == "" {
					return nil
				}
			}
		},
	}
}

func (x *TableGithubHookDeliveriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGithubHookDeliveriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("response").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					delivery := result.(*github.HookDelivery)
					return delivery.Response.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivered_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("DeliveredAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					delivery := result.(*github.HookDelivery)
					return delivery.Request.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redelivery").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hook_id").ColumnType(schema.ColumnTypeInt).Description("`Hook ID`").
			Extractor(github_client.ExtractorParentField("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("duration").ColumnType(schema.ColumnTypeFloat).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_code").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("action").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("installation_id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("InstallationID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("RepositoryID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("github_hooks_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to github_hooks.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGithubHookDeliveriesGenerator) GetSubTables() []*schema.Table {
	return nil
}
