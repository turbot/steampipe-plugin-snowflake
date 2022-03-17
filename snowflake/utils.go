package snowflake

import (
	"context"
	"database/sql"

	_ "github.com/snowflakedb/gosnowflake"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

//// TRANSFORM FUNCTIONS

func valueFromNullable(_ context.Context, d *transform.TransformData) (interface{}, error) {
	switch item := d.Value.(type) {
	case sql.NullString:
		if !item.Valid {
			return nil, nil
		}
		return item.String, nil
	case sql.NullBool:
		if !item.Valid {
			return nil, nil
		}
		return item.Bool, nil
	case sql.NullByte:
		if !item.Valid {
			return nil, nil
		}
		return item.Byte, nil
	case sql.NullFloat64:
		if !item.Valid {
			return nil, nil
		}
		return item.Float64, nil
	case sql.NullInt16:
		if !item.Valid {
			return nil, nil
		}
		return item.Int16, nil
	case sql.NullInt32:
		if !item.Valid {
			return nil, nil
		}
		return item.Int32, nil
	case sql.NullInt64:
		if !item.Valid {
			return nil, nil
		}
		return item.Int64, nil
	case sql.NullTime:
		if !item.Valid {
			return nil, nil
		}
		return item.Time, nil
	}
	return nil, nil
}
