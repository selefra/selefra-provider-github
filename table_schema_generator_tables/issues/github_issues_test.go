package issues

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/github_client/mocks"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildIssues(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockIssuesService(ctrl)

	var cs github.Issue
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Repository = &github.Repository{ID: &someId}

	mock.EXPECT().ListByOrg(gomock.Any(), "testorg", gomock.Any()).AnyTimes().Return(
		[]*github.Issue{&cs}, &github.Response{}, nil)

	return github_client.GithubServices{Issues: mock}
}

func TestIssues(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubIssuesGenerator{}), buildIssues, github_client.TestOptions{})
}
