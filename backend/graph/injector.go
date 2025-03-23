package graph

import (
	"log"
	"server/database"
)

func (r *Resolver) databaseBuilder() {
	if r.db == nil {
		err := database.Connect()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		db, err := database.GetPool()
		if err != nil {
			log.Fatalf("Failed to get database pool: %v", err)
		}
		r.db = db
	}
}
