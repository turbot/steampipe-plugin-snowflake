package snowflake

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_database",
		Description: "Snowflake Database",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeDatabases,
		},
		Columns: []*plugin.Column{
			{Name: "name", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "The name of the region", Type: proto.ColumnType_TIMESTAMP},
			{Name: "is_default", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "is_current", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "origin", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "owner", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "comment", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "options", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "retention_time", Description: "The name of the region", Type: proto.ColumnType_STRING},
		},
	}
}

type Database struct {
	CreatedOn     time.Time `json:"created_on"`
	Name          string    `json:"name"`
	IsDefault     string    `json:"is_default"`
	IsCurrent     string    `json:"is_current"`
	Origin        string    `json:"origin"`
	Owner         string    `json:"owner"`
	Comment       string    `json:"comment"`
	Options       string    `json:"options"`
	RetentionTime string    `json:"retention_time"`
}

//// LIST FUNCTION

func listSnowflakeDatabases(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Error("aws_region.listSnowflakeDatabases", "api.error", "nil")
	db, err := connect(ctx, d)
	defer db.Close()
	if err != nil {
		logger.Error("aws_region.listAwsRegions", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW DATABASES")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var CreatedOn time.Time
		var Name string
		var IsDefault string
		var IsCurrent string
		var Origin string
		var Owner string
		var Comment string
		var Options string
		var RetentionTime string

		err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &Origin, &Owner, &Comment, &Options, &RetentionTime)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, Database{CreatedOn, Name, IsDefault, IsCurrent, Origin, Owner, Comment, Options, RetentionTime})

	}
	return nil, nil
}
