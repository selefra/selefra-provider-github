package installations

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/github_client/mocks"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildInstallations(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Installation
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	total := 1
	mock.EXPECT().ListInstallations(gomock.Any(), "testorg", gomock.Any()).AnyTimes().Return(
		&github.OrganizationInstallations{TotalCount: &total, Installations: []*github.Installation{&cs}}, &github.Response{}, nil)

	return github_client.GithubServices{Organizations: mock}
}

func TestInstallations(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubInstallationsGenerator{}), buildInstallations, github_client.TestOptions{})
}
