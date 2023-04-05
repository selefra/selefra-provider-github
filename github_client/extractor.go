package github_client

import (
	"context"
	"time"

	"github.com/selefra/selefra-provider-github/constants"

	"github.com/google/go-github/v48/github"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
)

func ExtractorOrg() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return client.(*Client).Org, nil
	})
}

func ExtractorParentField(field string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return column_value_extractor.StructSelector(field).Extract(ctx, clientMeta, client, task, row, column, task.ParentRawResult)
	})
}

func ExtractorGithubDateTime(col string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		if reflect_util.IsNil(result) {
			return nil, nil
		}
		value, err := column_value_extractor.StructSelector(col).Extract(ctx, clientMeta, client, task, row, column, result)
		if err != nil && err.HasError() {
			return nil, err
		}
		switch v := value.(type) {
		case *github.Timestamp:
			return v.Time, nil
		case *time.Time:
			return v, nil
		case github.Timestamp:
			return v.Time, nil
		case time.Time:
			return v, nil
		default:
			return nil, schema.NewDiagnostics().AddErrorMsg(constants.InvalidtypeT, v)
		}
	})
}
