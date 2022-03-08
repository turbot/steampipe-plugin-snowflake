package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableAccountGrant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_account_grant",
		Description: "Lists all the account-level (i.e. global) privileges that have been granted to roles.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeAccountGrants,
		},
		Columns: []*plugin.Column{
			{Name: "name", Description: "An entity to which access can be granted. Unless allowed by a grant, access will be denied.", Type: proto.ColumnType_STRING},
			{Name: "privilege", Description: "A defined level of access to an object.", Type: proto.ColumnType_STRING},
			{Name: "granted_on", Description: "Date and time when the access was granted.", Type: proto.ColumnType_STRING},
			{Name: "granted_to", Description: "Type of the object.", Type: proto.ColumnType_STRING},
			{Name: "grantee_name", Description: "Name of the object role has been granted.", Type: proto.ColumnType_STRING},
			{Name: "granted_by", Description: "Name of the object that granted access on the role.", Type: proto.ColumnType_STRING},
			{Name: "grant_option", Description: "", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

type AccountGrant struct {
	CreatedOn   sql.NullTime   `json:"created_on"`
	Privilege   sql.NullString `json:"privilege"`
	GrantedOn   sql.NullString `json:"granted_on"`
	Name        sql.NullString `json:"name"`
	GrantedTo   sql.NullString `json:"granted_to"`
	GranteeName sql.NullString `json:"grantee_name"`
	GrantOption sql.NullString `json:"grant_option"`
	GrantedBy   sql.NullString `json:"granted_by"`
}

//// LIST FUNCTION

func listSnowflakeAccountGrants(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_account_grant.listSnowflakeAccountGrants", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW GRANTS ON ACCOUNT")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var createdOn sql.NullTime
		var privilege sql.NullString
		var grantedOn sql.NullString
		var name sql.NullString
		var grantedTo sql.NullString
		var granteeName sql.NullString
		var grantOption sql.NullString
		var grantedBy sql.NullString

		err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, AccountGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy})
	}

	for rows.NextResultSet() {
		var createdOn sql.NullTime
		var privilege sql.NullString
		var grantedOn sql.NullString
		var name sql.NullString
		var grantedTo sql.NullString
		var granteeName sql.NullString
		var grantOption sql.NullString
		var grantedBy sql.NullString

		err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, AccountGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy})
	}
	return nil, nil
}
