package billing

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/github_client/mocks"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildPackage(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockBillingService(ctrl)

	var cs *github.PackageBilling
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetPackagesBillingOrg(gomock.Any(), "testorg").AnyTimes().Return(cs, &github.Response{}, nil)
	return github_client.GithubServices{Billing: mock}
}

func TestPackageBillings(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubBillingPackageGenerator{}), buildPackage, github_client.TestOptions{})
}
