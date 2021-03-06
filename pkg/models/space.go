package models

import (
	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap/zapcore"
	"time"
)

type Space struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`                        // unique space identifier
	Name        string        `bson:"name" json:"name" validate:"required"` // space name
	Description string        `bson:"description" json:"description"`       // space description
	IsActive    bool          `bson:"is_active" json:"is_active"`           // is space active
	CreatedAt   time.Time     `bson:"created_at" json:"-"`                  // date of create space
	UpdatedAt   time.Time     `bson:"updated_at" json:"-"`                  // date of update space
}

type SpaceForm struct {
	Name        string `bson:"name" json:"name" validate:"required"` // space name
	Description string `bson:"description" json:"description"`       // space description
	IsActive    bool   `bson:"is_active" json:"is_active"`           // is space active
}

func (s *Space) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("id", s.Id.String())
	enc.AddString("name", s.Name)
	enc.AddString("description", s.Name)
	enc.AddBool("isActive", s.IsActive)

	return nil
}

func (s *SpaceForm) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("Name", s.Name)
	enc.AddString("Description", s.Description)
	enc.AddBool("IsActive", s.IsActive)

	return nil
}
