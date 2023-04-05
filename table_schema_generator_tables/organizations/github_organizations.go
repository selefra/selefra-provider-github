package organizations

import (
	"context"

	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGithubOrganizationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGithubOrganizationsGenerator{}

func (x *TableGithubOrganizationsGenerator) GetTableName() string {
	return "github_organizations"
}

func (x *TableGithubOrganizationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGithubOrganizationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGithubOrganizationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"org",
			"id",
		},
	}
}

func (x *TableGithubOrganizationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*github_client.Client)
			org, _, err := c.Github.Organizations.Get(ctx, c.Org)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- org
			return nil
		},
	}
}

func (x *TableGithubOrganizationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return github_client.ExpandOrg()
}

func (x *TableGithubOrganizationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("public_gists").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_repository_projects").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hooks_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HooksURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_repository_permission").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultRepoPermission")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_allowed_repository_creation_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_create_private_pages").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("events_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_usage").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("two_factor_requirement_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_create_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MembersCanCreateRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dependency_graph_enabled_for_new_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DependencyGraphEnabledForNewRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("blog").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collaborators").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_verified").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issues_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IssuesURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secret_scanning_push_protection_enabled_for_new_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SecretScanningPushProtectionEnabledForNewRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_private_repos").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_email").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_commit_signoff_required").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_create_private_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MembersCanCreatePrivateRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_create_pages").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org").ColumnType(schema.ColumnTypeString).Description("`The Github Organization of the resource.`").
			Extractor(github_client.ExtractorOrg()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("avatar_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AvatarURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("followers").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(github_client.ExtractorGithubDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_create_public_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MembersCanCreatePublicRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_repos").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plan").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_create_internal_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MembersCanCreateInternalRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("twitter_username").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_fork_private_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MembersCanForkPrivateRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_security_enabled_for_new_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AdvancedSecurityEnabledForNewRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NodeID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dependabot_alerts_enabled_for_new_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DependabotAlertsEnabledForNewRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_repository_settings").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultRepoSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_members_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicMembersURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_organization_projects").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repos_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReposURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("following").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owned_private_repos").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_gists").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MembersURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members_can_create_public_pages").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dependabot_security_updates_enabled_for_new_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DependabotSecurityUpdatesEnabledForNewRepos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secret_scanning_enabled_for_new_repositories").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SecretScanningEnabledForNewRepos")).Build(),
	}
}

func (x *TableGithubOrganizationsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGithubOrganizationMembersGenerator{}),
		table_schema_generator.GenTableSchema(&TableGithubOrganizationDependabotAlertsGenerator{}),
		table_schema_generator.GenTableSchema(&TableGithubOrganizationDependabotSecretsGenerator{}),
	}
}
