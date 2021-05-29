package repository

import (
	"database/sql"
	"github.com/vusalalishov/manpass/internal/db"
	"time"

	"github.com/vusalalishov/manpass/internal/model"
)

type CredentialRepository interface {
	Save(*model.Credential) (int64, error)
	Update(int64, *model.Credential) error
	Get(int64) (*model.Credential, error)
}

type credentialRepository struct {
	db *sql.DB
}

func ProvideCredentialRepository(db *sql.DB) CredentialRepository {
	return &credentialRepository{
		db,
	}
}

func InjectCredRepository() (CredentialRepository, error) {
	dbInjected, err := db.InjectDb()
	if err != nil {
		return nil, err
	}
	return ProvideCredentialRepository(dbInjected), nil
}

func (r *credentialRepository) Save(cred *model.Credential) (int64, error) {
	stmt, err := r.db.Prepare("insert into credential (title, login, password, updated_at) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(cred.Title, cred.Login, cred.Password, cred.UpdatedAt)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *credentialRepository) Update(id int64, cred *model.Credential) error {
	stmt, err := r.db.Prepare("update credential set title = ?, login = ?, password = ?, updated_at = date('now') where id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cred.Title, cred.Login, cred.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *credentialRepository) Get(id int64) (*model.Credential, error) {
	stmt, err := r.db.Prepare("select * from credential where id = ?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	var (
		title string
		login string
		password string
		updatedAt time.Time
	)
	err = row.Scan(nil, &title, &login, &password, &updatedAt)
	if err != nil {
		return nil, err
	}
	return &model.Credential{
		Id: id,
		Title: title,
		Login: login,
		Password: password,
		UpdatedAt: updatedAt,
	}, nil
}
