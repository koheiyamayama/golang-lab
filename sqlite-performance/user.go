package main

import "context"

type User struct {
	ID    int
	Name  string
	Email string
}

func NewRandomUser() *User {
	return &User{
		Name:  randomString(),
		Email: randomString() + "@example.com",
	}
}

func (s *sqlite) GetUsers(ctx context.Context, limit int, offset int) ([]*User, error) {
	rows, err := s.db.QueryContext(ctx, "select id, name, email from users limit ? offset ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		u := &User{}
		err := rows.Scan(&u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *sqlite) GetUser(ctx context.Context, id int) (*User, error) {
	row := s.db.QueryRowContext(ctx, "select id, name, email from users where id = ?", id)
	u := &User{}
	err := row.Scan(&u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *sqlite) CreateUser(ctx context.Context, u *User) error {
	_, err := s.db.ExecContext(ctx, "insert into users (name, email) values (?, ?)", u.Name, u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlite) DeleteUser(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, "delete from users where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
