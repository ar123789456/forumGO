package migrations

const LikeComment string = `
CREATE TABLE IF NOT EXISTS "likes_comment" (
	"id_comment"	INTEGER NOT NULL,
	"id_user"	INTEGER NOT NULL,
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	FOREIGN KEY("id_comment") REFERENCES "comments"("id")
);
`
