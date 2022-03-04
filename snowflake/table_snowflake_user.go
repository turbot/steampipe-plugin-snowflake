package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/snowflakedb/gosnowflake"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
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
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the snowflake user."},
			{Name: "login_name", Type: proto.ColumnType_STRING, Description: "Login name of the user."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email address of the user"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the user in Snowflake."},
			{Name: "has_password", Type: proto.ColumnType_BOOL, Description: "Whether the user has password."},
			{Name: "has_rsa_public_key", Type: proto.ColumnType_STRING, Description: "Whether the user has RSA public key."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the user was created."},

			{Name: "custom_landing_page_url", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("CUSTOM_LANDING_PAGE_URL"), Description: "Snowflake Support is allowed to use the user or account."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment associated to user in the dictionary."},
			{Name: "custom_landing_page_url_flush_next_ui_load", Type: proto.ColumnType_BOOL, Hydrate: DescribeUser, Transform: transform.FromField("CUSTOM_LANDING_PAGE_URL_FLUSH_NEXT_UI_LOAD"), Description: "The timestamp on which the last non-null password was set for the user. Default to null if no password has been set yet."},
			{Name: "days_to_expiry", Type: proto.ColumnType_STRING, Description: "User record will be treated as expired after specified number of days."},
			{Name: "default_namespace", Type: proto.ColumnType_STRING, Description: "Default database namespace prefix for this user."},
			{Name: "default_role", Type: proto.ColumnType_STRING, Description: "Primary principal of user session will be set to this role."},
			{Name: "default_secondary_roles", Type: proto.ColumnType_STRING, Description: "The secondary roles will be set to all roles provided here."},
			{Name: "default_warehouse", Type: proto.ColumnType_STRING, Description: "Default warehouse for this user."},
			{Name: "disabled", Type: proto.ColumnType_STRING, Description: "Whether the user is disabled."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display name of the user."},
			{Name: "expires_at_time", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "ext_authn_duo", Type: proto.ColumnType_STRING, Description: "Whether Duo Security is enabled as second factor authentication."},
			{Name: "ext_authn_uid", Type: proto.ColumnType_STRING, Description: "External authentication ID of the user."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "First name of the user."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "Last name of the user."},
			{Name: "last_success_login", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "locked_until_time", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "mins_to_bypass_mfa", Type: proto.ColumnType_STRING, Description: "Temporary bypass MFA for the user for a specified number of minutes."},
			{Name: "mins_to_bypass_network_policy", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("MINS_TO_BYPASS_NETWORK_POLICY"), Description: "Temporary bypass network policy on the user for a specified number of minutes."},
			{Name: "mins_to_unlock", Type: proto.ColumnType_STRING, Description: "Temporary lock on the user will be removed after specified number of minutes."},
			{Name: "must_change_password", Type: proto.ColumnType_STRING, Description: "User must change the password."},
			{Name: "password_last_set_time", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("PASSWORD_LAST_SET_TIME"), Description: "The timestamp on which the last non-null password was set for the user. Default to null if no password has been set yet."},
			{Name: "rsa_public_key", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("RSA_PUBLIC_KEY"), Description: "RSA public key of the user."},
			{Name: "rsa_public_key_2", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("RSA_PUBLIC_KEY_2"), Description: "Second RSA public key of the user."},
			{Name: "rsa_public_key_2_fp", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("RSA_PUBLIC_KEY_2_FP"), Description: "Fingerprint of user's second RSA public key."},
			{Name: "rsa_public_key_fp", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("RSA_PUBLIC_KEY_FP"), Description: "Fingerprint of user's RSA public key."},
			{Name: "snowflake_lock", Type: proto.ColumnType_STRING, Description: "Whether the user or account is locked by Snowflake."},
			{Name: "snowflake_support", Type: proto.ColumnType_STRING, Hydrate: DescribeUser, Transform: transform.FromField("SNOWFLAKE_SUPPORT"), Description: "Snowflake Support is allowed to use the user or account."},
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
		logger.Error("snowflake_user.listSnowflakeUsers", "connnection.error", err)
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

func DescribeUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var userName string
	if h.Item != nil {
		userName = h.Item.(User).Name.String
	} else {
		userName = d.KeyColumnQualString("name")
	}

	plugin.Logger(ctx).Info("snowflake_user.DescribeUser", "USER NAME", userName)

	if userName == "" {
		return nil, nil
	}

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("snowflake_user.DescribeUser", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("DESCRIBE USER %s", userName))
	if err != nil {
		plugin.Logger(ctx).Info("snowflake_user.DescribeUser", fmt.Sprintf("query_error for user %s", userName), fmt.Errorf("%#v", err))
		if err.(*gosnowflake.SnowflakeError) != nil {
			plugin.Logger(ctx).Info("snowflake_user.DescribeUser", fmt.Sprintf("query_error for user %s", userName), err.(*gosnowflake.SnowflakeError).Error())
			return nil, nil
		}
		return nil, err
	}
	userProperties := map[string]string{}
	for rows.Next() {
		var property sql.NullString
		var value sql.NullString
		var defaultval sql.NullString
		var description sql.NullString

		err = rows.Scan(&property, &value, &defaultval, &description)
		if err != nil {
			return nil, err
		}
		userProperties[property.String] = value.String
	}
	return userProperties, nil
}
