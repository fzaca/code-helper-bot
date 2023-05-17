package database

var instructions = []string{
	"SET FOREIGN_KEY_CHECKS=0;",
	"DROP TABLE IF EXISTS todo;",
	"SET FOREIGN_KEY_CHECKS=1;",
	`
	CREATE TABLE tasks (
		id INT PRIMARY KEY AUTO_INCREMENT,
		created_by INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		description TEXT NOT NULL,
		completed BOOLEAN NOT NULL
    );
	`,
}
