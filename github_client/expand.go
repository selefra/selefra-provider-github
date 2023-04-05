package github_client

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func ExpandOrg() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		cli := client.(*Client)
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)

		for _, org := range cli.Orgs {
			clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
				Client: cli.withOrg(org),
				Task:   task.Clone(),
			})
		}
		return clientTaskContextSlice
	}
}
