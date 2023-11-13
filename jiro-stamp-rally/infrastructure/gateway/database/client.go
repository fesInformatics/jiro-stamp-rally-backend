package client

type DbClient interface {
	InsertSQL(tablename string, valueMap map[string]any) error
}
