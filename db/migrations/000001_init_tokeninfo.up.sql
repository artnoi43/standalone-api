CREATE TABLE IF NOT EXISTS "token_infos" (
    "address" TEXT NOT NULL,
    "is_scam" BOOLEAN,
    "pending_scam" BOOLEAN,
    PRIMARY KEY ( "address" )
)