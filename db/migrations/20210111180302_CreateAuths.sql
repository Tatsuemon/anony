
-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `auths` (
    `user_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザーID',
    `token` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'アクセスTOKEN',
    `status` int NOT NULL COMMENT 'ステータス',
    PRIMARY KEY (`user_id`),
    FOREIGN KEY fk_user_id (`user_id`) REFERENCES users (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `auths`;