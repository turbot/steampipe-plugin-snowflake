package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/snowflakedb/gosnowflake"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*sql.DB, error) {
	config := GetConfig(d.Connection)
	connectionString := fmt.Sprintf("%s:%s@%s", *config.User, *config.Password, *config.Account)
	db, err := sql.Open("snowflake", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
