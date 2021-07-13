package healthcheck

// TablesSize specified size of table
// TODO: will be implement in the future
type TablesSize struct {
	TableSchema string
	TableName   string
	TableRows   int
	TableSize   float64
}
