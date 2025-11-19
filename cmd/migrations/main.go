package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	var (
		flags = flag.NewFlagSet("migrate", flag.ExitOnError)
		dir   = flags.String("dir", "migrations", "directory with migration files")
	)

	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	dsn := "host=localhost user=postgres password=postgres dbname=testovoe port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v\n", err)
	}
	defer db.Close()

	if err := goose.Run(command, db, *dir, args[1:]...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
