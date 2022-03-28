package snowflake

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func snowflakeColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumns()...)
}

// column definitions for the common columns
func commonColumns() []*plugin.Column {
	return []*plugin.Column{
		// {
		// 	Name:        "region",
		// 	Type:        proto.ColumnType_STRING,
		// 	Hydrate:      plugin.HydrateFunc(getCommonColumns).WithCache(),
		// 	Description: "The Snowflake Region in which the account is located.",
		// },
		{
			Name:        "account",
			Type:        proto.ColumnType_STRING,
			Hydrate:     plugin.HydrateFunc(getCommonColumns).WithCache(),
			Description: "The Snowflake Account ID.",
			Transform:   transform.FromCamel(),
		},
	}
}

// struct to store the common column data
type snowflakeCommonColumnData struct {
	Account, Region string
}

// get columns which are returned with all tables: region and account
func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	var region string
	if config.Region != nil {
		region = *config.Region
	}

	commonData := snowflakeCommonColumnData{
		Account: *config.Account,
		Region:  region,
	}

	return commonData, nil
}
