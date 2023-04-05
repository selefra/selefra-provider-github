package repositories

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubRepositoriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubRepositoriesGenerator{}

func (x *TableGithubRepositoriesGenerator) GetTableName() string {
	return "github_repositories"
}

func (x *TableGithubRepositoriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubRepositoriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubRepositoriesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"id",
		},
	}
}

func (x *TableGithubRepositoriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 100}}
			for {
				repos, resp, err := c.Github.Repositories.ListByOrg(ctx, c.Org, opts)
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

func (x *TableGithubRepositoriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubRepositoriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("languages_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LanguagesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collaborators_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CollaboratorsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("full_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stargazers_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("language").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_pages").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("blobs_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BlobsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issue_comment_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IssueCommentURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_branch").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mirror_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MirrorURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pulls_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PullsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gitignore_template").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssh_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SSHURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("watchers").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("keys_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeysURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_update_branch").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("forks_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fork").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("use_squash_pr_title_as_default").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("UseSquashPRTitleAsDefault")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("merge_commit_title").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_branch_on_merge").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statuses_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("branches_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BranchesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("git_refs_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GitRefsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("releases_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReleasesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team_id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("TeamID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TagsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_and_analysis").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("contents_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ContentsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("milestones_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MilestonesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("homepage").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("clone_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CloneURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_forking").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscribers_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubscribersURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("git_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GitURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topics").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_template").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubscriptionURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("text_matches").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("contributors_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ContributorsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_discussions").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compare_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CompareURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("watchers_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_squash_merge").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_merge_commit").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assignees_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AssigneesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issues_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IssuesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notifications_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NotificationsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_downloads").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trees_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TreesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("archived").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployments_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeploymentsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("git_tags_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GitTagsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("events_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issue_events_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IssueEventsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TeamsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comments_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CommentsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_repository").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_auto_merge").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("squash_merge_commit_title").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("squash_merge_commit_message").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code_of_conduct").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("svn_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SVNURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_issues").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("archive_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ArchiveURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("forks_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ForksURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("git_commits_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GitCommitsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("open_issues").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("open_issues_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscribers_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("merge_commit_message").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_branch").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pushed_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("PushedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("commits_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CommitsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("merges_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MergesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_init").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_rebase_merge").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("downloads_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DownloadsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_wiki").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stargazers_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StargazersURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_template").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_projects").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hooks_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HooksURL")).Build(),
	}
}

func (x *TableGithubRepositoriesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGithubRepositoryDependabotAlertsGenerator{}),
		table_schema_generator.GenTableSchema(&TableGithubRepositoryDependabotSecretsGenerator{}),
		table_schema_generator.GenTableSchema(&TableGithubReleasesGenerator{}),
	}
}
