package migrations

const Posts string = `
CREATE TABLE "posts" (
	"id"	INTEGER NOT NULL UNIQUE,
	"title"	TEXT NOT NULL UNIQUE,
	"content"	TEXT NOT NULL UNIQUE,
	"create_at"	TEXT NOT NULL,
	"update_at"	TEXT NOT NULL,
	"id_user"	INTEGER NOT NULL,
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
`
