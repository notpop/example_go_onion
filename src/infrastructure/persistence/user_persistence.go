package persistence
import (
	"database/sql"
	"example_onion/domain/entity"
)

type UserPersistence interface {
	FindByID(id int) (*entity.User, error)
}

type userPersistence struct {
	db *sql.DB
}

func (p *userPersistence) FindByID(id int) (*entity.User, error) {
	row := p.db.QueryRow("SELECT id, name, age, first_name, last_name, email FROM users WHERE id = ?", id)
	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserPersistence(db *sql.DB) UserPersistence {
	return &userPersistence{db: db}
}
