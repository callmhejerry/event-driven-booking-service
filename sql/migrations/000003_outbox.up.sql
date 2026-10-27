CREATE TABLE IF NOT EXISTS outbox(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aggregate_type TEXT NOT NULL, -- booking
    aggregate_id UUID NOT NULL,
    event_type TEXT NOT NULL, -- booking.create
    payload JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ
);


CREATE INDEX idx_outbox_unprocessed 
    ON outbox (created_at) 
    WHERE processed_at IS NULL;