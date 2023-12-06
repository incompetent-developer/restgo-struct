package cloudsql

import (
	"database/sql"

	"gorm.io/gorm"
)

type SQL struct {
	Host             string `json:"host"`
	Port             string `json:"port"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	DatabaseName     string `json:"database_name"`
	CredentialPath   string `json:"credential_path"`
	IsModeLogEnabled bool
	Client           *gorm.DB
	ClientDB         *sql.DB
}
