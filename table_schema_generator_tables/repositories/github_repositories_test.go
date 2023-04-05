package repositories

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
	mock.EXPECT().ListRepoAlerts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.DependabotAlert{&alert}, &github.Response{}, nil)

	var secret github.Secret
	if err := faker.FakeObject(&secret); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListRepoSecrets(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&github.Secrets{TotalCount: 1, Secrets: []*github.Secret{&secret}}, &github.Response{}, nil)

	return mock
}

func buildRepositiories(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockRepositoriesService(ctrl)

	var cs github.Repository
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Parent = &github.Repository{ID: &someId}
	cs.TemplateRepository = &github.Repository{ID: &someId}
	cs.Source = &github.Repository{ID: &someId}

	mock.EXPECT().ListByOrg(gomock.Any(), "testorg", gomock.Any()).AnyTimes().Return(
		[]*github.Repository{&cs}, &github.Response{}, nil)

	var release github.RepositoryRelease
	if err := faker.FakeObject(&release); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListReleases(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).AnyTimes().Return(
		[]*github.RepositoryRelease{&release}, &github.Response{}, nil)

	var releaseAsset github.ReleaseAsset
	if err := faker.FakeObject(&releaseAsset); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListReleaseAssets(gomock.Any(), "testorg", gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		[]*github.ReleaseAsset{&releaseAsset}, &github.Response{}, nil)

	dependabot := buildDependabot(t, ctrl)

	return github_client.GithubServices{
		Dependabot:   dependabot,
		Repositories: mock,
	}
}

func TestRepos(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubRepositoriesGenerator{}), buildRepositiories, github_client.TestOptions{})
}
