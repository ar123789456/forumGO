package migrations

const CategoryPost string = `
CREATE TABLE IF NOT EXISTS "category_posts" (
	"id"	INTEGER NOT NULL UNIQUE,
	"id_category"	INTEGER NOT NULL,
	"id_post"	INTEGER NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_post") REFERENCES "posts"("id"),
	FOREIGN KEY("id_category") REFERENCES "categories"("id")
);
`
