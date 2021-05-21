package healthcheck

type TablesSize struct {
	TableSchema string
	TableName   string
	TableRows   int
	TableSize   float64
}
