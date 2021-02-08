package datastore

import (
	"context"
	"database/sql"
	"time"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type anonyURLRepository struct {
	conn *sqlx.DB
}

// READで受け取るときに使用
type anonyURLReadEntity struct {
	ID        string    `json:"id" db:"id"`
	Original  string    `json:"original" db:"original"`
	Short     string    `json:"short" db:"short"`
	Status    int64     `json:"status" db:"status"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func mapAnonyURLReadEntityToAnonyURL(entity anonyURLReadEntity) model.AnonyURL {
	return model.AnonyURL{
		ID:       entity.ID,
		Original: entity.Original,
		Short:    entity.Short,
		Status:   entity.Status,
	}
}

// NewAnonyURLRepository creates a repository
func NewAnonyURLRepository(conn *sqlx.DB) repository.AnonyURLRepository {
	return &anonyURLRepository{conn: conn}
}

func (r anonyURLRepository) FindByID(id string) (*model.AnonyURL, error) {
	ae := anonyURLReadEntity{}
	if err := r.conn.Get(&ae, "SELECT id, original, short, status, user_id, created_at, updated_at FROM urls WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	res := mapAnonyURLReadEntityToAnonyURL(ae)
	return &res, nil
}

func (r anonyURLRepository) FindByUserID(userID string) ([]*model.AnonyURL, error) {
	aes := []anonyURLReadEntity{}
	if err := r.conn.Select(&aes, "SELECT id, original, short, status, user_id, created_at, updated_at FROM urls WHERE user_id = ?", userID); err != nil {
		return nil, err
	}
	res := make([]*model.AnonyURL, len(aes))
	for i, v := range aes {
		tmp := mapAnonyURLReadEntityToAnonyURL(v)
		res[i] = &tmp
	}
	return res, nil
}

func (r anonyURLRepository) FindByUserIDWithStatus(userID string, status int64) ([]*model.AnonyURL, error) {
	aes := []anonyURLReadEntity{}
	if err := r.conn.Select(&aes, "SELECT id, original, short, status, user_id, created_at, updated_at FROM urls WHERE user_id = ? and status = ?", userID, status); err != nil {
		return nil, err
	}
	res := make([]*model.AnonyURL, len(aes))
	for i, v := range aes {
		tmp := mapAnonyURLReadEntityToAnonyURL(v)
		res[i] = &tmp
	}
	return res, nil
}

func (r anonyURLRepository) FindByOriginalInUser(original string, userID string) (*model.AnonyURL, error) {
	ae := anonyURLReadEntity{}
	if err := r.conn.Get(&ae, "SELECT id, original, short, status, user_id, created_at, updated_at FROM urls WHERE original = ? AND user_id = ?", original, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	res := mapAnonyURLReadEntityToAnonyURL(ae)
	return &res, nil
}

func (r anonyURLRepository) FindByAnonyURL(anonyURL string) (*model.AnonyURL, error) {
	ae := anonyURLReadEntity{}
	if err := r.conn.Get(&ae, "SELECT id, original, short, status, user_id, created_at, updated_at FROM urls WHERE short = ?", anonyURL); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	res := mapAnonyURLReadEntityToAnonyURL(ae)
	return &res, nil
}

func (r anonyURLRepository) Save(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error) {
	// *sqlx.Tx, *sqlx.DBの両方で使用できるようにinterfaceの指定
	var tx interface {
		Prepare(query string) (*sql.Stmt, error)
	}

	// context.Contextから*sqlx.Txを取得
	tx, ok := GetTx(ctx)
	if !ok {
		tx = r.conn // context.Contextに存在しない場合は, repositoryの*sqlx.DBを使用
	}

	stmt, err := tx.Prepare("INSERT INTO `urls` (id, original, short, status, user_id) VALUES(?, ?, ?, ?, ?)")

	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.AnonyURLRepository.Save()")
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			// TODO(Tatseumon): 挙動確認
			err = closeErr
		}
	}()

	_, err = stmt.Exec(an.ID, an.Original, an.Short, an.Status, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.AnonyURLRepository.Save()")
	}

	return an, nil
}

func (r anonyURLRepository) UpdateStatus(ctx context.Context, an *model.AnonyURL) (*model.AnonyURL, error) {
	// *sqlx.Tx, *sqlx.DBの両方で使用できるようにinterfaceの指定
	var tx interface {
		Prepare(query string) (*sql.Stmt, error)
	}

	// context.Contextから*sqlx.Txを取得
	tx, ok := GetTx(ctx)
	if !ok {
		tx = r.conn // context.Contextに存在しない場合は, repositoryの*sqlx.DBを使用
	}

	stmt, err := tx.Prepare("UPDATE `urls` SET status = ? WHERE id = ?")

	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.AnonyURLRepository.UpdateStatus()")
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			// TODO(Tatseumon): 挙動確認
			err = closeErr
		}
	}()

	_, err = stmt.Exec(an.Status, an.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.AnonyURLRepository.UpdateStatus()")
	}
	return an, nil
}
