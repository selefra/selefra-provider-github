package provider

import (
	"github.com/selefra/selefra-provider-github/table_schema_generator"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/billing"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/external"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/hooks"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/installations"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/issues"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/organizations"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/repositories"
	"github.com/selefra/selefra-provider-github/table_schema_generator_tables/teams"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&billing.TableGithubBillingStorageGenerator{}),
		table_schema_generator.GenTableSchema(&billing.TableGithubBillingActionGenerator{}),
		table_schema_generator.GenTableSchema(&billing.TableGithubBillingPackageGenerator{}),
		table_schema_generator.GenTableSchema(&external.TableGithubExternalGroupsGenerator{}),
		table_schema_generator.GenTableSchema(&hooks.TableGithubHooksGenerator{}),
		table_schema_generator.GenTableSchema(&installations.TableGithubInstallationsGenerator{}),
		table_schema_generator.GenTableSchema(&issues.TableGithubIssuesGenerator{}),
		table_schema_generator.GenTableSchema(&organizations.TableGithubOrganizationsGenerator{}),
		table_schema_generator.GenTableSchema(&repositories.TableGithubRepositoriesGenerator{}),
		table_schema_generator.GenTableSchema(&teams.TableGithubTeamsGenerator{}),
	}
}
