package migrations

const UserSession string = `
CREATE TABLE IF NOT EXISTS "user_session" (
	"uid"	INTEGER NOT NULL UNIQUE,
	"user_id"	TEXT NOT NULL,
	"extime"	TEXT NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "user"
);`
