package snowflake

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func snowflakeColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumns()...)
}

// column definitions for the common columns
func commonColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "region",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Transform:   transform.FromCamel(),
			Description: "The Snowflake region in which the account is located.",
		},
		{
			Name:        "account",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The Snowflake account ID.",
			Transform:   transform.FromCamel(),
		},
	}
}

// struct to store the common column data
type snowflakeCommonColumnData struct {
	Account, Region string
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getCommonColumnsMemoized = plugin.HydrateFunc(getCommonColumnsUncached).Memoize(memoize.WithCacheKeyFunction(getCommonColumnsCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getCommonColumnsMemoized(ctx, d, h)
}

// Build a cache key for the call to getCommonColumnsCacheKey.
func getCommonColumnsCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getCommonColumns"
	return key, nil
}

func getAccountForConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	acc, err := getCommonColumnsMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}
	return acc.(snowflakeCommonColumnData).Account, nil
}

// get columns which are returned with all tables: region and account
func getCommonColumnsUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	var region string
	if config.Region != nil {
		region = *config.Region
	}

	// us-west-2 is the Snowflake's default region.
	// If it is not available in connection config, default region to us-west-2
	if region == "" {
		region = "us-west-2.aws"
	}

	commonData := snowflakeCommonColumnData{
		Account: *config.Account,
		Region:  region,
	}

	return commonData, nil
}
