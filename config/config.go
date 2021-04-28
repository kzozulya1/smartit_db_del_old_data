package config

// Configuration is main app config
type Configuration struct {
	ListenAddr    string `envconfig:"LISTEN_ADDR"      default:"0.0.0.0:2021"`
	DBConn        string `envconfig:"DB_CONN"          default:"postgres://postgres:postgres@127.0.0.1:5432/mesh_group?sslmode=disable"`
	DBSQLQueryLog bool   `envconfig:"DB_SQL_QUERY_LOG" default:"false"`

	TableRecordsLifeTimeDays int `envconfig:"TABLE_RECORDS_LIFETIME_DAYS" default:"30"`     // записи ранее 30 дней от роду не удаляются!
	TableRecordsDelBatchSize int `envconfig:"TABLE_RECORDS_DEL_BATCHSIZE" default:"163800"` // ачка записей для удаления - за раз удаляем только  16380 записей и ...
	TableRecordsRemovePause  int `envconfig:"TABLE_RECORDS_REMOVE_PAUSE" default:"800"`     // делаем паузу в 800 мс (для уменьшения нагрузки на pg)
}
