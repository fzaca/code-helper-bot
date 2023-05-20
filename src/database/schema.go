package database

var instructions = []string{
	"SET FOREIGN_KEY_CHECKS=0;",
	"DROP TABLE IF EXISTS tasks;",
	"SET FOREIGN_KEY_CHECKS=1;",
	`
	CREATE TABLE tasks (
		id INT PRIMARY KEY,
		description TEXT NOT NULL,
		assigned_to TEXT,
		completed BOOLEAN NOT NULL DEFAULT false,
		created_by TEXT NOT NULL,
		updated_by TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP,
		server_id TEXT NOT NULL
    );
	`,
}
