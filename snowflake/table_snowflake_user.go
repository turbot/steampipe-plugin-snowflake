package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_user",
		Description: "Snowflake User",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeUsers,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "login_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "email", Type: proto.ColumnType_STRING, Description: ""},

			{Name: "mins_to_unlock", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "days_to_expiry", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "disabled", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "must_change_password", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "snowflake_lock", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "default_warehouse", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "default_namespace", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "default_role", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "default_secondary_roles", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ext_authn_duo", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ext_authn_uid", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "mins_to_bypass_mfa", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "last_success_login", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "expires_at_time", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "locked_until_time", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "has_password", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "has_rsa_public_key", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

type User struct {
	Name                  sql.NullString `json:"name"`
	CreatedOn             sql.NullTime   `json:"created_on"`
	LoginName             sql.NullString `json:"login_name"`
	DisplayName           sql.NullString `json:"display_name"`
	FirstName             sql.NullString `json:"first_name"`
	LastName              sql.NullString `json:"last_name"`
	Email                 sql.NullString `json:"email"`
	MinsToUnlock          sql.NullString `json:"mins_to_unlock"`
	DaysToExpiry          sql.NullString `json:"days_to_expiry"`
	Comment               sql.NullString `json:"comment"`
	Disabled              sql.NullString `json:"disabled"`
	MustChangePassword    sql.NullString `json:"must_change_password"`
	SnowflakeLock         sql.NullString `json:"snowflake_lock"`
	DefaultWarehouse      sql.NullString `json:"default_warehouse"`
	DefaultNamespace      sql.NullString `json:"default_namespace"`
	DefaultRole           sql.NullString `json:"default_role"`
	DefaultSecondaryRoles sql.NullString `json:"default_secondary_roles"`
	ExtAuthnDuo           sql.NullString `json:"ext_authn_duo"`
	ExtAuthnUid           sql.NullString `json:"ext_authn_uid"`
	MinsToBypassMFA       sql.NullString `json:"mins_to_bypass_mfa"`
	Owner                 sql.NullString `json:"owner"`
	LastSuccessLogin      sql.NullTime   `json:"last_success_login"`
	ExpiresAtTime         sql.NullTime   `json:"expires_at_time"`
	LockedUntilTime       sql.NullTime   `json:"locked_until_time"`
	HasPassword           sql.NullString `json:"has_password"`
	HasRSAPublicKey       sql.NullString `json:"has_rsa_public_key"`
}

//// LIST FUNCTION

func listSnowflakeUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("aws_region.listAwsRegions", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW USERS")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var Name sql.NullString
		var CreatedOn sql.NullTime
		var LoginName sql.NullString
		var DisplayName sql.NullString
		var FirstName sql.NullString
		var LastName sql.NullString
		var Email sql.NullString
		var MinsToUnlock sql.NullString
		var DaysToExpiry sql.NullString
		var Comment sql.NullString
		var Disabled sql.NullString
		var MustChangePassword sql.NullString
		var SnowflakeLock sql.NullString
		var DefaultWarehouse sql.NullString
		var DefaultNamespace sql.NullString
		var DefaultRole sql.NullString
		var DefaultSecondaryRoles sql.NullString
		var ExtAuthnDuo sql.NullString
		var ExtAuthnUid sql.NullString
		var MinsToBypassMFA sql.NullString
		var Owner sql.NullString
		var LastSuccessLogin sql.NullTime
		var ExpiresAtTime sql.NullTime
		var LockedUntilTime sql.NullTime
		var HasPassword sql.NullString
		var HasRSAPublicKey sql.NullString

		err = rows.Scan(&Name, &CreatedOn, &LoginName, &DisplayName, &FirstName, &LastName, &Email, &MinsToUnlock, &DaysToExpiry, &Comment, &Disabled, &MustChangePassword, &SnowflakeLock, &DefaultWarehouse, &DefaultNamespace, &DefaultRole, &DefaultSecondaryRoles, &ExtAuthnDuo, &ExtAuthnUid, &MinsToBypassMFA, &Owner, &LastSuccessLogin, &ExpiresAtTime, &LockedUntilTime, &HasPassword, &HasRSAPublicKey)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, User{Name, CreatedOn, LoginName, DisplayName, FirstName, LastName, Email, MinsToUnlock, DaysToExpiry, Comment, Disabled, MustChangePassword, SnowflakeLock, DefaultWarehouse, DefaultNamespace, DefaultRole, DefaultSecondaryRoles, ExtAuthnDuo, ExtAuthnUid, MinsToBypassMFA, Owner, LastSuccessLogin, ExpiresAtTime, LockedUntilTime, HasPassword, HasRSAPublicKey})
	}

	for rows.NextResultSet() {
		for rows.Next() {
			var Name sql.NullString
			var CreatedOn sql.NullTime
			var LoginName sql.NullString
			var DisplayName sql.NullString
			var FirstName sql.NullString
			var LastName sql.NullString
			var Email sql.NullString
			var MinsToUnlock sql.NullString
			var DaysToExpiry sql.NullString
			var Comment sql.NullString
			var Disabled sql.NullString
			var MustChangePassword sql.NullString
			var SnowflakeLock sql.NullString
			var DefaultWarehouse sql.NullString
			var DefaultNamespace sql.NullString
			var DefaultRole sql.NullString
			var DefaultSecondaryRoles sql.NullString
			var ExtAuthnDuo sql.NullString
			var ExtAuthnUid sql.NullString
			var MinsToBypassMFA sql.NullString
			var Owner sql.NullString
			var LastSuccessLogin sql.NullTime
			var ExpiresAtTime sql.NullTime
			var LockedUntilTime sql.NullTime
			var HasPassword sql.NullString
			var HasRSAPublicKey sql.NullString

			err = rows.Scan(&Name, &CreatedOn, &LoginName, &DisplayName, &FirstName, &LastName, &Email, &MinsToUnlock, &DaysToExpiry, &Comment, &Disabled, &MustChangePassword, &SnowflakeLock, &DefaultWarehouse, &DefaultNamespace, &DefaultRole, &DefaultSecondaryRoles, &ExtAuthnDuo, &ExtAuthnUid, &MinsToBypassMFA, &Owner, &LastSuccessLogin, &ExpiresAtTime, &LockedUntilTime, &HasPassword, &HasRSAPublicKey)
			if err != nil {
				return nil, err
			}

			d.StreamListItem(ctx, User{Name, CreatedOn, LoginName, DisplayName, FirstName, LastName, Email, MinsToUnlock, DaysToExpiry, Comment, Disabled, MustChangePassword, SnowflakeLock, DefaultWarehouse, DefaultNamespace, DefaultRole, DefaultSecondaryRoles, ExtAuthnDuo, ExtAuthnUid, MinsToBypassMFA, Owner, LastSuccessLogin, ExpiresAtTime, LockedUntilTime, HasPassword, HasRSAPublicKey})
		}
	}
	defer db.Close()
	return nil, nil
}
