package hooks

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/github_client/mocks"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildHooks(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Hook
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	cs.Config = make(map[string]any)
	cs.LastResponse = make(map[string]any)
	mock.EXPECT().ListHooks(gomock.Any(), "testorg", gomock.Any()).AnyTimes().Return([]*github.Hook{&cs}, &github.Response{}, nil)

	var hd *github.HookDelivery
	if err := faker.FakeObject(&hd); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListHookDeliveries(gomock.Any(), "testorg", *cs.ID, gomock.Any()).AnyTimes().Return([]*github.HookDelivery{hd}, &github.Response{}, nil)
	return github_client.GithubServices{Organizations: mock}
}

func TestHooks(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubHooksGenerator{}), buildHooks, github_client.TestOptions{})
}
