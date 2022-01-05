package migrations

const TagPosts string = `
CREATE TABLE  IF NOT EXISTS "tag_posts" (
	"id"	INTEGER NOT NULL UNIQUE,
	"id_tag"	INTEGER NOT NULL,
	"id_post"	INTEGER NOT NULL,
	FOREIGN KEY("id_post") REFERENCES "posts"("id"),
	FOREIGN KEY("id_tag") REFERENCES "tags"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
`
