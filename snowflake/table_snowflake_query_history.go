package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/account-usage/query_history.html
// TODO Implement quals to on time, session, user, warehouse
func tableSnowflakeQueryHistory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_query_history",
		Description: "This Account Usage view can be used to query Snowflake query history by various dimensions (time range, session, user, warehouse, etc.) within the last 365 days (1 year).",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeQueryHistory,
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "query_id", Type: proto.ColumnType_STRING, Description: "The statement's unique id."},
			{Name: "query_text", Type: proto.ColumnType_STRING, Description: "Text of the SQL statement."},
			// {Name: "database_id", Type: proto.ColumnType_INT, Description: ""},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "Database that was in use at the time of the query."},
			// {Name: "schema_id", Type: proto.ColumnType_INT, Description: ""},
			{Name: "schema_name", Type: proto.ColumnType_STRING, Description: "Schema that was in use at the time of the query."},
			{Name: "query_type", Type: proto.ColumnType_STRING, Description: "DML, query, etc. If the query is currently running, or the query failed, then the query type may be UNKNOWN."},
			{Name: "session_id", Type: proto.ColumnType_INT, Description: "Session that executed the statement."},
			{Name: "user_name", Type: proto.ColumnType_STRING, Description: "User who issued the query."},
			{Name: "role_name", Type: proto.ColumnType_STRING, Description: "Role that was active in the session at the time of the query."},
			// {Name: "warehouse_id", Type: proto.ColumnType_INT, Description: ""},
			{Name: "warehouse_name", Type: proto.ColumnType_STRING, Description: "Warehouse that the query executed on, if any."},
			{Name: "warehouse_size", Type: proto.ColumnType_STRING, Description: "Size of the warehouse when this statement executed."},
			{Name: "warehouse_type", Type: proto.ColumnType_STRING, Description: "Type of the warehouse when this statement executed."},
			{Name: "cluster_number", Type: proto.ColumnType_INT, Description: "The cluster (in a multi-cluster warehouse) that this statement executed on."},
			{Name: "query_tag", Type: proto.ColumnType_STRING, Description: "Query tag set for this statement through the QUERY_TAG session parameter."},
			{Name: "execution_status", Type: proto.ColumnType_STRING, Description: "Execution status for the query: resuming_warehouse, running, queued, blocked, success, failed_with_error, or failed_with_incident."},
			{Name: "error_code", Type: proto.ColumnType_STRING, Description: "Error code, if the query returned an error."},
			{Name: "error_message", Type: proto.ColumnType_STRING, Description: "Error message, if the query returned an error"},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "Statement start time"},
			{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "Statement end time."},
			{Name: "total_elapsed_time", Type: proto.ColumnType_INT, Description: "Elapsed time (in milliseconds)."},
			// {Name: "bytes_scanned", Type: proto.ColumnType_INT, Description: ""},
			{Name: "percentage_scanned_from_cache", Type: proto.ColumnType_DOUBLE, Description: "The percentage of data scanned from the local disk cache. The value ranges from 0.0 to 1.0. Multiply by 100 to get a true percentage."},
			// {Name: "bytes_written", Type: proto.ColumnType_INT, Description: ""},
			// {Name: "bytes_written_to_result", Type: proto.ColumnType_INT, Description: ""},
			// {Name: "bytes_read_from_result", Type: proto.ColumnType_INT, Description: ""},
			{Name: "rows_produced", Type: proto.ColumnType_INT, Description: "Number of rows produced by this statement."},
			{Name: "rows_inserted", Type: proto.ColumnType_INT, Description: "Number of rows inserted by the query."},
			{Name: "rows_updated", Type: proto.ColumnType_INT, Description: "Number of rows updated by the query."},
			{Name: "rows_deleted", Type: proto.ColumnType_INT, Description: "Number of rows deleted by the query."},
			{Name: "rows_unloaded", Type: proto.ColumnType_INT, Description: "Number of rows unloaded during data export."},
			// {Name: "bytes_deleted", Type: proto.ColumnType_INT, Description: ""},
			{Name: "partitions_scanned", Type: proto.ColumnType_INT, Description: "Number of micro-partitions scanned."},
			{Name: "partitions_total", Type: proto.ColumnType_INT, Description: "Total micro-partitions of all tables included in this query."},
			{Name: "bytes_spilled_to_local_storage", Type: proto.ColumnType_INT, Description: "Volume of data spilled to local disk."},
			{Name: "bytes_spilled_to_remote_storage", Type: proto.ColumnType_INT, Description: "Volume of data spilled to remote disk."},
			{Name: "bytes_sent_over_the_network", Type: proto.ColumnType_INT, Description: "Volume of data sent over the network."},
			{Name: "compilation_time", Type: proto.ColumnType_INT, Description: "Compilation time (in milliseconds)."},
			{Name: "execution_time", Type: proto.ColumnType_INT, Description: "Execution time (in milliseconds)."},
			{Name: "queued_provisioning_time", Type: proto.ColumnType_INT, Description: "Time (in milliseconds) spent in the warehouse queue, waiting for the warehouse compute resources to provision, due to warehouse creation, resume, or resize."},
			{Name: "queued_repair_time", Type: proto.ColumnType_INT, Description: "Time (in milliseconds) spent in the warehouse queue, waiting for compute resources in the warehouse to be repaired."},
			{Name: "queued_overload_time", Type: proto.ColumnType_INT, Description: "Time (in milliseconds) spent in the warehouse queue, due to the warehouse being overloaded by the current query workload."},
			{Name: "transaction_blocked_time", Type: proto.ColumnType_INT, Description: "Time (in milliseconds) spent blocked by a concurrent DML."},
			{Name: "outbound_data_transfer_cloud", Type: proto.ColumnType_STRING, Description: "Target cloud provider for statements that unload data to another region and/or cloud."},
			{Name: "outbound_data_transfer_region", Type: proto.ColumnType_STRING, Description: "Target region for statements that unload data to another region and/or cloud."},
			// {Name: "outbound_data_transfer_bytes", Type: proto.ColumnType_INT, Description: ""},
			{Name: "inbound_data_transfer_cloud", Type: proto.ColumnType_STRING, Description: "Source cloud provider for statements that load data from another region and/or cloud."},
			{Name: "inbound_data_transfer_region", Type: proto.ColumnType_STRING, Description: "Source region for statements that load data from another region and/or cloud."},
			// {Name: "inbound_data_transfer_bytes", Type: proto.ColumnType_INT, Description: ""},
			{Name: "list_external_files_time", Type: proto.ColumnType_INT, Description: "Time (in milliseconds) spent listing external files."},
			{Name: "credits_used_cloud_services", Type: proto.ColumnType_DOUBLE, Description: "Number of credits used for cloud services in the hour."},
			// {Name: "release_version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "external_function_total_invocations", Type: proto.ColumnType_INT, Description: "The aggregate number of times that this query called remote services."},
			{Name: "external_function_total_sent_rows", Type: proto.ColumnType_INT, Description: "The total number of rows that this query sent in all calls to all remote services."},
			{Name: "external_function_total_received_rows", Type: proto.ColumnType_INT, Description: "The total number of rows that this query received from all calls to all remote services."},
			{Name: "external_function_total_sent_bytes", Type: proto.ColumnType_INT, Description: "The total number of bytes that this query sent in all calls to all remote services."},
			{Name: "external_function_total_received_bytes", Type: proto.ColumnType_INT, Description: "The total number of bytes that this query received from all calls to all remote services."},
			{Name: "query_load_percent", Type: proto.ColumnType_INT, Description: "The approximate percentage of active compute resources in the warehouse for this query execution."},
			{Name: "is_client_generated_statement", Type: proto.ColumnType_BOOL, Description: "Whether the query was client-generated."},
		}),
	}
}

type QueryHistory struct {
	QueryId                            sql.NullString  `json:"QUERY_ID"`
	QueryText                          sql.NullString  `json:"QUERY_TEXT"`
	DatabaseId                         sql.NullInt64   `json:"DATABASE_ID"`
	DatabaseName                       sql.NullString  `json:"DATABASE_NAME"`
	SchemaId                           sql.NullInt64   `json:"SCHEMA_ID"`
	SchemaName                         sql.NullString  `json:"SCHEMA_NAME"`
	QueryType                          sql.NullString  `json:"QUERY_TYPE"`
	SessionId                          sql.NullInt64   `json:"SESSION_ID"`
	UserName                           sql.NullString  `json:"USER_NAME"`
	RoleName                           sql.NullString  `json:"ROLE_NAME"`
	WarehouseId                        sql.NullInt64   `json:"WAREHOUSE_ID"`
	WarehouseName                      sql.NullString  `json:"WAREHOUSE_NAME"`
	WarehouseSize                      sql.NullString  `json:"WAREHOUSE_SIZE"`
	WarehouseType                      sql.NullString  `json:"WAREHOUSE_TYPE"`
	ClusterNumber                      sql.NullInt64   `json:"CLUSTER_NUMBER"`
	QueryTag                           sql.NullString  `json:"QUERY_TAG"`
	ExecutionStatus                    sql.NullString  `json:"EXECUTION_STATUS"`
	ErrorCode                          sql.NullString  `json:"ERROR_CODE"`
	ErrorMessage                       sql.NullString  `json:"ERROR_MESSAGE"`
	StartTime                          sql.NullTime    `json:"START_TIME"`
	EndTime                            sql.NullTime    `json:"END_TIME"`
	TotalElapsedTime                   sql.NullInt64   `json:"TOTAL_ELAPSED_TIME"`
	BytesScanned                       sql.NullInt64   `json:"BYTES_SCANNED"`
	PercentageScannedFromCache         sql.NullFloat64 `json:"PERCENTAGE_SCANNED_FROM_CACHE"`
	BytesWritten                       sql.NullInt64   `json:"BYTES_WRITTEN"`
	BytesWrittenToResult               sql.NullInt64   `json:"BYTES_WRITTEN_TO_RESULT"`
	BytesReadFromResult                sql.NullInt64   `json:"BYTES_READ_FROM_RESULT"`
	RowsProduced                       sql.NullInt64   `json:"ROWS_PRODUCED"`
	RowsInserted                       sql.NullInt64   `json:"ROWS_INSERTED"`
	RowsUpdated                        sql.NullInt64   `json:"ROWS_UPDATED"`
	RowsDeleted                        sql.NullInt64   `json:"ROWS_DELETED"`
	RowsUnloaded                       sql.NullInt64   `json:"ROWS_UNLOADED"`
	BytesDeleted                       sql.NullInt64   `json:"BYTES_DELETED"`
	PartitionsScanned                  sql.NullInt64   `json:"PARTITIONS_SCANNED"`
	PartitionsTotal                    sql.NullInt64   `json:"PARTITIONS_TOTAL"`
	BytesSpilledToLocalStorage         sql.NullInt64   `json:"BYTES_SPILLED_TO_LOCAL_STORAGE"`
	BytesSpilledToRemoteStorage        sql.NullInt64   `json:"BYTES_SPILLED_TO_REMOTE_STORAGE"`
	BytesSentOverTheNetwork            sql.NullInt64   `json:"BYTES_SENT_OVER_THE_NETWORK"`
	CompilationTime                    sql.NullInt64   `json:"COMPILATION_TIME"`
	ExecutionTime                      sql.NullInt64   `json:"EXECUTION_TIME"`
	QueuedProvisioningTime             sql.NullInt64   `json:"QUEUED_PROVISIONING_TIME"`
	QueuedRepairTime                   sql.NullInt64   `json:"QUEUED_REPAIR_TIME"`
	QueuedOverloadTime                 sql.NullInt64   `json:"QUEUED_OVERLOAD_TIME"`
	TransactionBlockedTime             sql.NullInt64   `json:"TRANSACTION_BLOCKED_TIME"`
	OutboundDataTransferCloud          sql.NullString  `json:"OUTBOUND_DATA_TRANSFER_CLOUD"`
	OutboundDataTransferRegion         sql.NullString  `json:"OUTBOUND_DATA_TRANSFER_REGION"`
	OutboundDataTransferBytes          sql.NullInt64   `json:"OUTBOUND_DATA_TRANSFER_BYTES"`
	InboundDataTransferCloud           sql.NullString  `json:"INBOUND_DATA_TRANSFER_CLOUD"`
	InboundDataTransferRegion          sql.NullString  `json:"INBOUND_DATA_TRANSFER_REGION"`
	InboundDataTransferBytes           sql.NullInt64   `json:"INBOUND_DATA_TRANSFER_BYTES"`
	ListExternalFilesTime              sql.NullInt64   `json:"LIST_EXTERNAL_FILES_TIME"`
	CreditsUsedCloudServices           sql.NullFloat64 `json:"CREDITS_USED_CLOUD_SERVICES"`
	ReleaseVersion                     sql.NullString  `json:"RELEASE_VERSION"`
	ExternalFunctionTotalInvocations   sql.NullInt64   `json:"EXTERNAL_FUNCTION_TOTAL_INVOCATIONS"`
	ExternalFunctionTotalSentRows      sql.NullInt64   `json:"EXTERNAL_FUNCTION_TOTAL_SENT_ROWS"`
	ExternalFunctionTotalReceivedRows  sql.NullInt64   `json:"EXTERNAL_FUNCTION_TOTAL_RECEIVED_ROWS"`
	ExternalFunctionTotalSentBytes     sql.NullInt64   `json:"EXTERNAL_FUNCTION_TOTAL_SENT_BYTES"`
	ExternalFunctionTotalReceivedBytes sql.NullInt64   `json:"EXTERNAL_FUNCTION_TOTAL_RECEIVED_BYTES"`
	QueryLoadPercent                   sql.NullInt64   `json:"QUERY_LOAD_PERCENT"`
	IsClientGeneratedStatement         sql.NullBool    `json:"IS_CLIENT_GENERATED_STATEMENT"`
}

// QueryHistoryCol returns a reference for a column of a QueryHistory
func QueryHistoryCol(colname string, item *QueryHistory) interface{} {
	switch colname {
	case "QUERY_ID":
		return &item.QueryId
	case "QUERY_TEXT":
		return &item.QueryText
	case "DATABASE_ID":
		return &item.DatabaseId
	case "DATABASE_NAME":
		return &item.DatabaseName
	case "SCHEMA_ID":
		return &item.SchemaId
	case "SCHEMA_NAME":
		return &item.SchemaName
	case "QUERY_TYPE":
		return &item.QueryType
	case "SESSION_ID":
		return &item.SessionId
	case "USER_NAME":
		return &item.UserName
	case "ROLE_NAME":
		return &item.RoleName
	case "WAREHOUSE_ID":
		return &item.WarehouseId
	case "WAREHOUSE_NAME":
		return &item.WarehouseName
	case "WAREHOUSE_SIZE":
		return &item.WarehouseSize
	case "WAREHOUSE_TYPE":
		return &item.WarehouseType
	case "CLUSTER_NUMBER":
		return &item.ClusterNumber
	case "QUERY_TAG":
		return &item.QueryTag
	case "EXECUTION_STATUS":
		return &item.ExecutionStatus
	case "ERROR_CODE":
		return &item.ErrorCode
	case "ERROR_MESSAGE":
		return &item.ErrorMessage
	case "START_TIME":
		return &item.StartTime
	case "END_TIME":
		return &item.EndTime
	case "TOTAL_ELAPSED_TIME":
		return &item.TotalElapsedTime
	case "BYTES_SCANNED":
		return &item.BytesScanned
	case "PERCENTAGE_SCANNED_FROM_CACHE":
		return &item.PercentageScannedFromCache
	case "BYTES_WRITTEN":
		return &item.BytesWritten
	case "BYTES_WRITTEN_TO_RESULT":
		return &item.BytesWrittenToResult
	case "BYTES_READ_FROM_RESULT":
		return &item.BytesReadFromResult
	case "ROWS_PRODUCED":
		return &item.RowsProduced
	case "ROWS_INSERTED":
		return &item.RowsInserted
	case "ROWS_UPDATED":
		return &item.RowsUpdated
	case "ROWS_DELETED":
		return &item.RowsDeleted
	case "ROWS_UNLOADED":
		return &item.RowsUnloaded
	case "BYTES_DELETED":
		return &item.BytesDeleted
	case "PARTITIONS_SCANNED":
		return &item.PartitionsScanned
	case "PARTITIONS_TOTAL":
		return &item.PartitionsTotal
	case "BYTES_SPILLED_TO_LOCAL_STORAGE":
		return &item.BytesSpilledToLocalStorage
	case "BYTES_SPILLED_TO_REMOTE_STORAGE":
		return &item.BytesSpilledToRemoteStorage
	case "BYTES_SENT_OVER_THE_NETWORK":
		return &item.BytesSentOverTheNetwork
	case "COMPILATION_TIME":
		return &item.CompilationTime
	case "EXECUTION_TIME":
		return &item.ExecutionTime
	case "QUEUED_PROVISIONING_TIME":
		return &item.QueuedProvisioningTime
	case "QUEUED_REPAIR_TIME":
		return &item.QueuedRepairTime
	case "QUEUED_OVERLOAD_TIME":
		return &item.QueuedOverloadTime
	case "TRANSACTION_BLOCKED_TIME":
		return &item.TransactionBlockedTime
	case "OUTBOUND_DATA_TRANSFER_CLOUD":
		return &item.OutboundDataTransferCloud
	case "OUTBOUND_DATA_TRANSFER_REGION":
		return &item.OutboundDataTransferRegion
	case "OUTBOUND_DATA_TRANSFER_BYTES":
		return &item.OutboundDataTransferBytes
	case "INBOUND_DATA_TRANSFER_CLOUD":
		return &item.InboundDataTransferCloud
	case "INBOUND_DATA_TRANSFER_REGION":
		return &item.InboundDataTransferRegion
	case "INBOUND_DATA_TRANSFER_BYTES":
		return &item.InboundDataTransferBytes
	case "LIST_EXTERNAL_FILES_TIME":
		return &item.ListExternalFilesTime
	case "CREDITS_USED_CLOUD_SERVICES":
		return &item.CreditsUsedCloudServices
	case "RELEASE_VERSION":
		return &item.ReleaseVersion
	case "EXTERNAL_FUNCTION_TOTAL_INVOCATIONS":
		return &item.ExternalFunctionTotalInvocations
	case "EXTERNAL_FUNCTION_TOTAL_SENT_ROWS":
		return &item.ExternalFunctionTotalSentRows
	case "EXTERNAL_FUNCTION_TOTAL_RECEIVED_ROWS":
		return &item.ExternalFunctionTotalReceivedRows
	case "EXTERNAL_FUNCTION_TOTAL_SENT_BYTES":
		return &item.ExternalFunctionTotalSentBytes
	case "EXTERNAL_FUNCTION_TOTAL_RECEIVED_BYTES":
		return &item.ExternalFunctionTotalReceivedBytes
	case "QUERY_LOAD_PERCENT":
		return &item.QueryLoadPercent
	case "IS_CLIENT_GENERATED_STATEMENT":
		return &item.IsClientGeneratedStatement
	default:
		panic("unknown column " + colname)
	}
}

//// LIST FUNCTION

func listSnowflakeQueryHistory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_query_history.listSnowflakeQueryHistory", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "select * from SNOWFLAKE.ACCOUNT_USAGE.QUERY_HISTORY;")
	if err != nil {
		logger.Error("snowflake_query_history.listSnowflakeQueryHistory", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Error("snowflake_query_history.listSnowflakeQueryHistory", "get_columns.error", err)
		return nil, err
	}

	for rows.Next() {
		queryHistory := QueryHistory{}
		// make references for the cols with the aid of QueryHistoryCol
		cols := make([]interface{}, len(columns))
		for i, col := range columns {
			cols[i] = QueryHistoryCol(col, &queryHistory)
		}

		err = rows.Scan(cols...)
		if err != nil {
			logger.Error("snowflake_query_history.listSnowflakeQueryHistory", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, queryHistory)
	}

	for rows.NextResultSet() {
		for rows.Next() {
			queryHistory := QueryHistory{}
			// make references for the cols with the aid of QueryHistoryCol
			cols := make([]interface{}, len(columns))
			for i, col := range columns {
				cols[i] = QueryHistoryCol(col, &queryHistory)
			}

			err = rows.Scan(cols...)
			if err != nil {
				logger.Error("snowflake_query_history.listSnowflakeQueryHistory", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, queryHistory)
		}
	}
	return nil, nil
}
