package storage

import "github.com/seggga/approve-auth/internal/entity"

// UserStorage represents storage with basic operations
type UserStorage interface {
	ReadUser(login string) (*entity.UserOpts, error)
}
