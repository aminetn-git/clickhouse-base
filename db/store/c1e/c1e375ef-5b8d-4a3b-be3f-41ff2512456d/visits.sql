ATTACH TABLE _ UUID 'c9908d64-e1b7-4377-a6ee-cb97c589f74e'
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
