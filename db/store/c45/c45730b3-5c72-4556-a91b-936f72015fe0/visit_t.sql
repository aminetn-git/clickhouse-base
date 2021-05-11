ATTACH TABLE _ UUID '4968e9bd-a167-47a5-9c6f-d73495946b4b'
(
    `id` UInt64,
    `duration` Float64,
    `url` String,
    `created` DateTime
)
ENGINE = MergeTree
PRIMARY KEY id
ORDER BY id
SETTINGS index_granularity = 8192
