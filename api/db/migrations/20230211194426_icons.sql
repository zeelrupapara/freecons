-- +migrate Up
CREATE TABLE IF NOT EXISTS icons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT,
    name TEXT,
    status INTEGER,
    icon_url TEXT,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    )
);


-- +migrate Down
DROP TABLE IF EXISTS icons;