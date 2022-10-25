package database

import (
	"log"

	"doki/wallet/config"
	"doki/wallet/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectMySQL opens a connection to mysql database with the given config.
func ConnectMySQL(conf config.MySQL) {

	// log.Printf("dsn: %s\n", conf.DSN())
	database, err := gorm.Open(mysql.Open(conf.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("[database]>>> failed to open connection: ", err)
	}
	log.Println("[database]>>> database connected...")
	DB = database
}

func GetDB() *gorm.DB {
	return DB
}

// migrate calls Automigrate function for all models.
// It also drops unused column internally.
func Migrate() {
	db_models := []interface{}{
		&domain.Wallet{},
		&domain.Transaction{},
	}

	log.Println("Current Database: ", DB.Migrator().CurrentDatabase())
	for _, db_model := range db_models {
		if err := DB.AutoMigrate(db_model); err != nil {
			log.Fatal("[database]>>> failed to migrate: ", err)
		}
		dropUnusedColumns(db_model)
	}
	log.Println("[database]>>> database migrated...")
}

func dropUnusedColumns(dst interface{}) {

	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(dst)
	fields := stmt.Schema.Fields
	columns, _ := DB.Debug().Migrator().ColumnTypes(dst)
	for i := range columns {
		found := false
		for j := range fields {
			if columns[i].Name() == fields[j].DBName {
				found = true
				break
			}
		}
		if !found {
			DB.Migrator().DropColumn(dst, columns[i].Name())
		}
	}
}
