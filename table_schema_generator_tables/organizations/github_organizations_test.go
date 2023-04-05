package organizations

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"
	"github.com/selefra/selefra-provider-github/github_client/mocks"
	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildDependabot(t *testing.T, ctrl *gomock.Controller) github_client.DependabotService {
	mock := mocks.NewMockDependabotService(ctrl)

	var alert github.DependabotAlert
	if err := faker.FakeObject(&alert); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListOrgAlerts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.DependabotAlert{&alert}, &github.Response{}, nil)

	var secret github.Secret
	if err := faker.FakeObject(&secret); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListOrgSecrets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&github.Secrets{TotalCount: 1, Secrets: []*github.Secret{&secret}}, &github.Response{}, nil)

	return mock
}

func buildOrganizations(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs *github.Organization
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Get(gomock.Any(), gomock.Any()).AnyTimes().Return(cs, &github.Response{}, nil)

	var u github.User
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		[]*github.User{&u}, &github.Response{}, nil)

	var m github.Membership
	if err := faker.FakeObject(&m); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetOrgMembership(gomock.Any(), *u.Login, gomock.Any()).AnyTimes().Return(
		&m, &github.Response{}, nil)

	dependabot := buildDependabot(t, ctrl)

	return github_client.GithubServices{
		Dependabot:    dependabot,
		Organizations: mock,
	}
}

func TestOrganizations(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubOrganizationsGenerator{}), buildOrganizations, github_client.TestOptions{})
}
