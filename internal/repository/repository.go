package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/badprinter/remember/internal/config"
	_ "github.com/lib/pq"
)

type Repository struct {
	SQL *sql.DB
}

func CreateRepository(cfg *config.RepositoryCfg) *Repository {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.HOST, cfg.PORT, cfg.USER, cfg.PASSWORD, cfg.DB_NAME, cfg.SSL_MODE)
	// psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
	//	cfg.USER, cfg.PASSWORD, cfg.HOST, cfg.PORT, cfg.DB_NAME, cfg.SSL_MODE)
	log.Println("Database connect URL:", psqlconn)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalln("Create Repository", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Ping SQL", err)
	}
	log.Println("Success connect postgres.")
	return &Repository{
		SQL: db,
	}
}
