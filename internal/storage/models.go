package storage

import "gitlab.com/g6834/team9/auth/internal/entity"

// UserStorage represents storage with basic operations
type UserStorage interface {
	ReadUser(login string) (*entity.UserOpts, error)
}
