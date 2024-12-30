package infrastructure

import (
	"context"
	"log"

	"github.com/eiei114/go-backend-template/domain"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	Conn *bun.DB
}

func NewUserRepository(Conn *bun.DB) *UserRepository {
	return &UserRepository{Conn: Conn}
}

func (u *UserRepository) AddUser(ctx context.Context, id, authToken, name string) error {
	user := &domain.User{
		Id:        id,
		AuthToken: authToken,
		Name:      name,
		Count:     0,
	}
	_, err := u.Conn.NewInsert().Model(user).Exec(ctx)
	return err
}

func (u *UserRepository) UpdateUser(ctx context.Context, user domain.User) error {
	uuser := &domain.User{
		Id:        user.Id,
		AuthToken: user.AuthToken,
		Name:      user.Name,
		Count:     user.Count,
	}
	_, err := u.Conn.NewUpdate().Model(uuser).Where("id = ?", uuser.Id).Exec(ctx)
	log.Println(uuser.Count)
	return err
}

func (u *UserRepository) DeleteUser(ctx context.Context, id string) error {
	user := &domain.User{Id: id}
	_, err := u.Conn.NewDelete().Model(user).Where("id = ?", id).Exec(ctx)
	return err
}

func (u *UserRepository) GetUserByUserId(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	err := u.Conn.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByAuthToken(ctx context.Context, authToken string) (*domain.User, error) {
	user := new(domain.User)
	err := u.Conn.NewSelect().Model(user).Where("auth_token = ?", authToken).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}