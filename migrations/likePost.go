package migrations

const LikePost string = `
CREATE TABLE "likes_posts" (
	"id"	INTEGER NOT NULL UNIQUE,
	"id_post"	INTEGER NOT NULL,
	"id_user"	INTEGER NOT NULL,
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	FOREIGN KEY("id_post") REFERENCES "posts"("id")
);
`