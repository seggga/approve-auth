package entity

import "github.com/google/uuid"

// UserOpts represents user's properties
type UserOpts struct {
	Login    string    `yaml:"login" bson:"login" json:"login"`
	ID       uuid.UUID `yaml:"uuid" bson:"uuid" json:"uuid"`
	PassHash string    `yaml:"pass-hash" bson:"pass-hash" json:"pass-hash"`
}

// Users struct holds data about users capable of authorization
type Users struct {
	Data map[string]UserOpts `yaml:"users"`
}
