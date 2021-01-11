
-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `urls` (
    `id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'URL_ID',
    `original` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'オリジナルURL',
    `short` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '省略URL',
    `status` int NOT NULL COMMENT 'ステータス',
    `user_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザーID',
    PRIMARY KEY (`id`),
    FOREIGN KEY fk_user_id (`user_id`) REFERENCES users (`id`),
    INDEX original_index(`original`),
    INDEX short_index(`short`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `urls`;