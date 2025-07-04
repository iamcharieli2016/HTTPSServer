package model

// ServiceRequest 服务请求结构体
type ServiceRequest struct {
	Params struct {
		ColumnComment string `json:"columnComment"`
		ColumnType    string `json:"columnType"`
		ColumnName    string `json:"columnName"`
		TableComment  string `json:"tableComment"`
		TableName     string `json:"tableName"`
		TableSchema   string `json:"tableSchema"`
		DBType        string `json:"dbType"`
	} `json:"params"`
	ServiceID    string `json:"serviceId"`
	ShowCount    string `json:"showCount"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	UserID       string `json:"userId"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// TableMetadata 数据库表结构信息
type TableMetadata struct {
	TableSchema   string `json:"tableSchema"`
	TableName     string `json:"tableName"`
	TableComment  string `json:"tableComment"`
	ColumnName    string `json:"columnName"`
	ColumnType    string `json:"columnType"`
	ColumnComment string `json:"columnComment"`
	DBType        string `json:"dbType"`
}
