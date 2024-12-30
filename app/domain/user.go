package domain

type User struct {
	Id        string `bun:"id"`
	AuthToken string `bun:"auth_token"`
	Name      string `bun:"name"`
	Count     int    `bun:"count"`
}
