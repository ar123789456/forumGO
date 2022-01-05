package migrations

const LikePost string = `
CREATE TABLE IF NOT EXISTS "likes_posts" (
	"id_post"	INTEGER NOT NULL,
	"id_user"	INTEGER NOT NULL,
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	FOREIGN KEY("id_post") REFERENCES "posts"("id")
);
`
