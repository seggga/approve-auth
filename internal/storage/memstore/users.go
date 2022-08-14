// loadYaml reads users login/pass-hash data from yaml config file
// yaml-file format is supposed to be like:
//
// users:
// 	user_1:
//   login: "user_1"
//   uuid: "f700be7b-d722-42a3-8f2e-6e90d3425104"
//   pass-hash: "1234"
//

package memstore

import (
	"fmt"
	"sync"

	"github.com/seggga/approve-auth/internal/entity"
)

// Store ...
type Store struct {
	sync.RWMutex
	users entity.Users
}

// New reads user data (login, ID, pass-hash) from given yaml-file
func New(users map[string]entity.UserOpts) (*Store, error) {
	return &Store{
		users: entity.Users{
			Data: users,
		},
	}, nil
}

// ReadUser extracts a user by login
func (s *Store) ReadUser(login string) (*entity.UserOpts, error) {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	u, ok := s.users.Data[login]
	if !ok {
		return nil, fmt.Errorf("user with login %s was not found", login)
	}

	return &u, nil
}
