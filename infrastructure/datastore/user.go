package datastore

import (
	"context"
	"database/sql"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type userRepository struct {
	conn *sqlx.DB
}

// NewUserRepository create a repository of user.
func NewUserRepository(conn *sqlx.DB) repository.UserRepository {
	return &userRepository{conn: conn}
}

func (r userRepository) FindAll() ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := r.conn.Select(&users, "Select id, name, email FROM users"); err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepository) FindByID(id string) (*model.User, error) {
	user := model.User{}
	if err := r.conn.Get(&user, "Select id, name, email FROM users WHERE id = ?", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) FindByName(name string) (*model.User, error) {
	user := model.User{}
	if err := r.conn.Get(&user, "Select id, name, email FROM users WHERE name = ?", name); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := r.conn.Get(&user, "Select id, name, email FROM users WHERE email = ?", email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Store(ctx context.Context, user *model.User) (*model.User, error) {
	// *sqlx.Tx, *sqlx.DBの両方で使用できるようにinterfaceの指定
	var tx interface {
		Prepare(query string) (*sql.Stmt, error)
	}

	// context.Contextから*sqlx.Txを取得
	tx, ok := GetTx(ctx)
	if !ok {
		tx = r.conn // context.Contextに存在しない場合は, repositoryの*sqlx.DBを使用
	}

	stmt, err := tx.Prepare("INSERT INTO `users` (id, name, email, password) VALUES(?, ?, ?, ?)")

	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.userRepository.Store()")
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			// TODO(Tatseumon): 挙動確認
			err = closeErr
		}
	}()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.EncryptedPass)
	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.userRepository.Store()")
	}

	return user, nil
}

func (r userRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	var tx interface {
		Prepare(query string) (*sql.Stmt, error)
	}

	tx, ok := GetTx(ctx)
	if !ok {
		tx = r.conn
	}

	stmt, err := tx.Prepare("UPDATE `users` SET name = ?, email = ?, password = ? WHERE id = ?")

	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.userRepository.Update()")
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			// TODO(Tatseumon): 挙動確認
			err = closeErr
		}
	}()

	_, err = stmt.Exec(user.Name, user.Email, user.EncryptedPass, user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to datastore.userRepository.Update()")
	}

	return user, nil
}
func (r userRepository) Delete(ctx context.Context, user *model.User) error {
	var tx interface {
		Prepare(query string) (*sql.Stmt, error)
	}

	tx, ok := GetTx(ctx)
	if !ok {
		tx = r.conn
	}

	stmt, err := tx.Prepare("DELETE FROM `users` WHERE id = ?")

	if err != nil {
		return errors.Wrap(err, "failed to datastore.userRepository.Delete()")
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			// TODO(Tatseumon): 挙動確認
			err = closeErr
		}
	}()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to datastore.userRepository.Delete()")
	}

	return nil
}
