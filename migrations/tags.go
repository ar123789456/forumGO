package migrations

const Tags string = `
CREATE TABLE IF NOT EXISTS "tags" (
	"id"	INTEGER NOT NULL UNIQUE,
	"title"	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("id" AUTOINCREMENT)
);
`
