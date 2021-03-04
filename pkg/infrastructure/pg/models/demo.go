package models

type Demo struct {
	tableName string `pg:"demos,alias:demo"`
	// 例子id
	DemoId int64 `pg:"pk:demo_id"`
	// 例子名称
	DemoName string
}
