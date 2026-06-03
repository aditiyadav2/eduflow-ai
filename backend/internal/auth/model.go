package auth

import "time"

type User struct {
	ID           string    `bson:"_id"`
	Name         string    `bson:"name"`
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"password_hash"`
	Role         string    `bson:"role"`
	TenantID     string    `bson:"tenant_id"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}
