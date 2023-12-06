package cloudsql

import (
	"sync"

	"gorm.io/gorm/schema"
)

// GetTableName is get table name from model
func (sql SQL) GetTableName(modelOfTable interface{}) string {
	table, _ := schema.Parse(modelOfTable, &sync.Map{}, schema.NamingStrategy{})
	return table.Table
}
