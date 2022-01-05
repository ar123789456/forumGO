package migrations

const Comments string = `
CREATE TABLE IF NOT EXISTS "comments" (
	"id"	INTEGER NOT NULL UNIQUE,
	"text"	TEXT NOT NULL,
	"id_user"	INTEGER NOT NULL,
	"id_post"	INTEGER NOT NULL,
	FOREIGN KEY("id_post") REFERENCES "posts"("id"),
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
`
