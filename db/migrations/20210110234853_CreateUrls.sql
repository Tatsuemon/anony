-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `urls` (
    `original` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'オリジナルURL',
    `short` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '省略URL',
    `status` int COLLATE utf8mb4_bin NOT NULL COMMENT 'ステータス',
    PRIMARY KEY (`original`),
    INDEX short_index(`short`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `urls`;