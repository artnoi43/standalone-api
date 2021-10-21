CREATE TABLE IF NOT EXISTS "token_infos" (
    "address" TEXT NOT NULL,
    "is_scam" BOOLEAN,
    "pending_scam" BOOLEAN,
    "created_at" TIMESTAMPTZ,
    "updated_at" TIMESTAMPTZ,
    PRIMARY KEY ( "address" )
)