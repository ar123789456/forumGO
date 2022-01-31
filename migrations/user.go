package migrations

const Users string = `
CREATE TABLE IF NOT EXISTS "user" (
	"ID"	INTEGER NOT NULL UNIQUE,
	"NicName"	TEXT NOT NULL UNIQUE,
	"Email"	TEXT NOT NULL UNIQUE,
	"Password"	TEXT NOT NULL,
	"UID"	TEXT UNIQUE,
	PRIMARY KEY("ID" AUTOINCREMENT)
);
`
