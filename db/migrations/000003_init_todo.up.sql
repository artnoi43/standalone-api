CREATE TABLE IF NOT EXISTS "todos" (
    "uuid" TEXT NOT NULL,
    "user_uuid" TEXT NOT NULL,
    "text" TEXT NOT NULL,
    "done" BOOLEAN NOT NULL,
    "created_at" TIMESTAMPTZ,
    "updated_at" TIMESTAMPTZ,

    PRIMARY KEY ( "uuid" ),
    FOREIGN KEY( "user_uuid" ) 
        REFERENCES users( "uuid" )
)