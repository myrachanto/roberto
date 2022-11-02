package sqldbconn

import (
	httperrors "github.com/myrachanto/erroring"
	job "github.com/myrachanto/roberto/src/jobs/models"
	synca "github.com/myrachanto/roberto/src/synca/models"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// centralRepo
var (
	CentralRepo centralRepo = centralRepo{}
	DbType      string      = "mysql"
	DbName      string      = "synca"
	DbUsername  string      = "root"
	DbPassword  string      = "Golang@456"
	// DbPort        string      = "3306"
)

type centralRepo struct {
}

func init() {
	GormDB, err := gorm.Open(mysql.New(mysql.Config{
		// DSN: sdb.DbUsername + ":" + sdb.DbPassword + "@tcp(mysql_database:3306)/" + sdb.DbNameDefault + "?charset=utf8&parseTime=True&loc=Local", // data source name
		// DSN:                      sdb.DbUsername + ":" + sdb.DbPassword + "@tcp(host.docker.internal:3306)/" + sdb.DbNameDefault + "?charset=utf8&parseTime=True&loc=Local",
		DSN:                      DbUsername + ":" + DbPassword + "@tcp(127.0.0.1:3306)/" + DbName + "?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:        256,                                                                                                         // default size for string fields
		DisableDatetimePrecision: true,                                                                                                        // disable datetime precision, which not supported before MySQL 5.6
		// DontSupportRenamecentral: true, // drop & create when rename central, rename central not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return
	}
	GormDB.AutoMigrate(&job.Job{})
	GormDB.AutoMigrate(&synca.Synca{})
}
func (centralRepo *centralRepo) Getconnected() (*gorm.DB, httperrors.HttpErr) {
	GormDB, err := gorm.Open(mysql.New(mysql.Config{
		// DSN: sdb.DbUsername + ":" + sdb.DbPassword + "@tcp(mysql_database:3306)/" + sdb.DbNameDefault + "?charset=utf8&parseTime=True&loc=Local", // data source name
		// DSN:                      sdb.DbUsername + ":" + sdb.DbPassword + "@tcp(host.docker.internal:3306)/" + sdb.DbNameDefault + "?charset=utf8&parseTime=True&loc=Local",
		DSN:                      DbUsername + ":" + DbPassword + "@tcp(127.0.0.1:3306)/" + DbName + "?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:        256,                                                                                                         // default size for string fields
		DisableDatetimePrecision: true,                                                                                                        // disable datetime precision, which not supported before MySQL 5.6
		// DontSupportRenamecentral: true, // drop & create when rename central, rename central not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, httperrors.NewNotFoundError("Something went wrong connecting to the --db!")
	}
	return GormDB, nil
}
func (centralRepo *centralRepo) DbClose(GormDB *gorm.DB) {
	sqlDB, err := GormDB.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}
