package actions

import (
	"context"
	"net/url"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubWorkflowsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubWorkflowsGenerator{}

func (x *TableGithubWorkflowsGenerator) GetTableName() string {
	return "github_workflows"
}

func (x *TableGithubWorkflowsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubWorkflowsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubWorkflowsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"id",
		},
	}
}

func (x *TableGithubWorkflowsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 100}}
			for {
				repos, resp, err := c.Github.Repositories.ListByOrg(ctx, c.Org, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, repo := range repos {
					actionOpts := &github.ListOptions{PerPage: 100}
					for {
						workflows, resp, err := c.Github.Actions.ListWorkflows(ctx, *repo.Owner.Login, *repo.Name, actionOpts)
						if err != nil {
							return schema.NewDiagnosticsErrorPullTable(task.Table, err)

						}
						resultChannel <- workflows.Workflows
						opts.Page = resp.NextPage
						if opts.Page == resp.LastPage {
							break
						}
					}
				}
				opts.Page = resp.NextPage
				if opts.Page == resp.LastPage {
					break
				}
			}
			return nil
		},
	}
}

func (x *TableGithubWorkflowsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubWorkflowsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("contents").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*github_client.Client)
					workflow := result.(*github.Workflow)

					parsedUrl, err := url.Parse(*workflow.HTMLURL)
					if err != nil {
						return nil, err
					}

					pathParts := strings.Split(parsedUrl.Path, "/")
					if len(pathParts) < 2 {
						return nil, nil
					}
					owner := pathParts[1]
					repo := pathParts[2]
					ref := pathParts[4]
					path := *workflow.Path
					opts := github.RepositoryContentGetOptions{Ref: ref}

					fileContent, _, _, err := cl.Github.Repositories.GetContents(ctx, owner, repo, path, &opts)
					if err != nil {

						return nil, nil
					}
					content, err := fileContent.GetContent()
					if err != nil {
						return nil, err
					}
					return content, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("badge_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BadgeURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeID")).Build(),
	}
}

func (x *TableGithubWorkflowsGenerator) GetSubTables() []*schema.Table {
	return nil
}
