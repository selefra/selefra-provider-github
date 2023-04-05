package external

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/github_client/mocks"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildExternalGroups(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockTeamsService(ctrl)

	var cs *github.ExternalGroupList
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListExternalGroups(gomock.Any(), "testorg", gomock.Any()).AnyTimes().Return(cs, &github.Response{}, nil)
	return github_client.GithubServices{Teams: mock}
}

func TestExternalGroups(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubExternalGroupsGenerator{}), buildExternalGroups, github_client.TestOptions{})
}
