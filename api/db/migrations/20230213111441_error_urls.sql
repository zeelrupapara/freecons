-- +migrate Up
CREATE TABLE IF NOT EXISTS errorlinks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    icon_url TEXT,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    )
);


-- +migrate Down
DROP TABLE IF EXISTS errorlinks;