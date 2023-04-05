package issues

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubIssuesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubIssuesGenerator{}

func (x *TableGithubIssuesGenerator) GetTableName() string {
	return "github_issues"
}

func (x *TableGithubIssuesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubIssuesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubIssuesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"id",
		},
	}
}

func (x *TableGithubIssuesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			opts := &github.IssueListOptions{
				Filter: "all",
				State:  "all",
				ListOptions: github.ListOptions{
					Page:    0,
					PerPage: 100,
				},
			}
			for {
				issues, resp, err := c.Github.Issues.ListByOrg(ctx, c.Org, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- issues
				opts.Page = resp.NextPage
				if opts.Page == resp.LastPage {
					break
				}
			}
			return nil
		},
	}
}

func (x *TableGithubIssuesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubIssuesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("events_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assignees").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locked").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("closed_by").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pull_request").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PullRequestLinks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reactions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("text_matches").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("body").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("author_association").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assignee").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comments_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CommentsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("milestone").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_lock_reason").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comments").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_reason").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("closed_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClosedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RepositoryURL")).Build(),
	}
}

func (x *TableGithubIssuesGenerator) GetSubTables() []*schema.Table {
	return nil
}
