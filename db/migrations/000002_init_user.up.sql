CREATE TABLE IF NOT EXISTS "users" (
    "uuid" TEXT NOT NULL,
    "username" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    PRIMARY KEY ( "username" )
)