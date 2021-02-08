package datastore

import (
	"database/sql"

	"github.com/Tatsuemon/anony/usecase/dto"
	"github.com/Tatsuemon/anony/usecase/queryservice"
	"github.com/jmoiron/sqlx"
)

type userAnonyURLAccessor struct {
	conn *sqlx.DB
}

// NewUserAnonyURLAccessor create a accessor
func NewUserAnonyURLAccessor(conn *sqlx.DB) queryservice.UserAnonyURLAccessor {
	return &userAnonyURLAccessor{conn: conn}
}

func (a userAnonyURLAccessor) CountAnonyURLByUser(userID string) (*dto.AnonyURLCountByUser, error) {
	res := dto.AnonyURLCountByUser{}
	q := `
	SELECT name, email, COUNT(user_id) AS count_urls, COUNT(status=1 or null) AS count_active_urls
	FROM users
	INNER JOIN urls ON users.id = urls.user_id
	WHERE user_id = ?
	GROUP BY (user_id)
	`

	if err := a.conn.Get(&res, q, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
