package migrate

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

func Migrate(db *sql.DB, schema string) {
	fmt.Println("migrating...")
	CurrentDB := DB{
		Connection: db,
		Schema: Schema{
			Tables: make(map[string]Table),
		},
	}

	// Temporary In Memory DB - Based on the schema.sql file
	CleanDB := NewDB([]string{"file:clean.sqlite3?mode=memory", schema})
	defer CleanDB.Close()

	// Apply schema changes (create tables/indices, drop tables/indices)
	CurrentDB.ApplySchemaChanges(CleanDB)

	// 1. Disable foreign keys
	CurrentDB.DisableForeignKeys()

	for tableName, table := range CurrentDB.findAlteredTables(CleanDB) {
		fmt.Println("found altered table " + tableName)

		// 2. Start transaction
		tx, err := CurrentDB.Connection.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// 3. Define create table statement with new name
		tableNameNew := tableName + "_new"

		// 4. Create new tables
		createTable(tx, tableName, tableNameNew, table)

		// 5. Transfer table contents to new table
		migrateContent(tx, CleanDB, &CurrentDB, tableName, tableNameNew)

		// 6. Drop old table
		dropTable(tx, tableName)

		// 7. Rename new table to old table
		renameTable(tx, tableName, tableNameNew)

		// 8. Use CREATE INDEX, CREATE TRIGGER, and CREATE VIEW to reconstruct indexes, triggers, and views associated with table X. Perhaps use the old format of the triggers, indexes, and views saved from step 3 above as a guide, making changes as appropriate for the alteration.
		createIndicesOnTable(tx, tableName, CleanDB)

		// 9. If any views refer to table X in a way that is affected by the schema change, then drop those views using DROP VIEW and recreate them with whatever changes are necessary to accommodate the schema change using CREATE VIEW.

		// 10. If foreign key constraints were originally enabled then run PRAGMA foreign_key_check to verify that the schema change did not break any foreign key constraints.
		err = CurrentDB.Exec("PRAGMA foreign_key_check")
		if err != nil {
			log.Fatal(err)
		}

		// 11. End transaction
		err = tx.Commit()
	}

	// 12. Enable foreign keys again
	err := CurrentDB.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("migration complete.")
}

func createIndicesOnTable(tx *sql.Tx, tableName string, CleanDB *DB) {
	indices := CleanDB.GetSchema().GetTableIndices(tableName)
	for _, index := range indices {
		_, err := tx.Exec(index.SQL)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func renameTable(tx *sql.Tx, tableName string, tableNameNew string) {
	fmt.Println("renaming table " + tableName)
	_, err := tx.Exec("ALTER TABLE "+tableNameNew+" RENAME TO "+tableName, tableNameNew, tableName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("renamed " + tableNameNew + " to " + tableName)
}

func createTable(tx *sql.Tx, tableName string, tableNameNew string, table Table) {
	fmt.Println("creating table " + tableNameNew)
	stmt := strings.Replace(table.SQL, tableName, tableNameNew, 1)

	fmt.Println(stmt)
	_, err := tx.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func dropTable(tx *sql.Tx, tableName string) {
	fmt.Println("dropping " + tableName)
	_, err := tx.Exec(`DROP TABLE IF EXISTS ` + tableName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dropped " + tableName)
}

func migrateContent(tx *sql.Tx, CleanDB *DB, CurrentDB *DB, tableName string, tableNameNew string) {
	fmt.Println("migrating content from " + tableName + " to " + tableNameNew + "...")
	intersection := Intersect(
		CleanDB.GetColumns(tableName),
		CurrentDB.GetColumns(tableName),
	)
	if len(intersection) == 0 {
		return
	}

	cols := strings.Join(intersection[:], ", ")
	_, err := tx.Exec("INSERT INTO " + tableNameNew + " (" + cols + ") SELECT " + cols + " FROM " + tableName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted " + tableNameNew)
}
