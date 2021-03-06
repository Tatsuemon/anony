
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
    `id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザーID',
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザー名',
    `email` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'メールアドレス',
    `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '暗号化されたパスワード',
    `created_at` DATETIME COLLATE utf8mb4_bin NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME COLLATE utf8mb4_bin NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE name_index (`name`),
    UNIQUE email_index (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;
