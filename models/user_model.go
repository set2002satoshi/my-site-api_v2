package models

import (
	"mime/multipart"
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"golang.org/x/crypto/bcrypt"
)

type ActiveUserModel struct {
	userId     types.IDENTIFICATION
	nickname   string
	email      string
	password   []byte
	icon       *types.ImageTypeFileOrURL
	roll       types.AccessROLL
	auditTrail *types.AuditTrail
}

func NewActiveUserModel(
	userId int,
	nickname string,
	email string,
	password string,
	iconFile *multipart.FileHeader,
	iconURL string,
	iconImageFlag bool,
	roll string,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*ActiveUserModel, error) {
	aum := new(ActiveUserModel)

	var err error
	err = errors.Combine(err, aum.setUserId(userId))
	err = errors.Combine(err, aum.setNickname(nickname))
	err = errors.Combine(err, aum.setEmail(email))
	err = errors.Combine(err, aum.setPassword(password))
	err = errors.Combine(err, aum.setRoll(roll))
	if err != nil {
		return new(ActiveUserModel), err
	}

	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(ActiveUserModel), err
	}

	err = errors.Combine(err, aum.setAuditTrail(com))
	if err != nil {
		return new(ActiveUserModel), err
	}

	images, err := types.NewImageTypeFileOrURL(iconFile, iconURL, iconImageFlag)
	if err != nil {
		return new(ActiveUserModel), err
	}

	err = errors.Combine(err, aum.setIcon(images))
	if err != nil {
		return new(ActiveUserModel), err
	}

	if err := aum.Validation(); err != nil {
		return new(ActiveUserModel), err
	}

	return aum, nil
}

func (aum *ActiveUserModel) GetUserId() types.IDENTIFICATION {
	return aum.userId
}

func (aum *ActiveUserModel) GetNickname() string {
	return aum.nickname
}

func (aum *ActiveUserModel) GetEmail() string {
	return aum.email
}

func (aum *ActiveUserModel) GetPassword() string {
	return string(aum.password)
}

func (aum *ActiveUserModel) GetIcon() *types.ImageTypeFileOrURL {
	return aum.icon
}

func (aum *ActiveUserModel) GetRoll() types.AccessROLL {
	return aum.roll
}

func (aum *ActiveUserModel) GetAuditTrail() *types.AuditTrail {
	return aum.auditTrail
}

func (aum *ActiveUserModel) setUserId(userId int) error {
	i, err := types.NewIDENTIFICATION(userId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	aum.userId = i
	return nil
}

func (aum *ActiveUserModel) setNickname(nickname string) error {
	aum.nickname = nickname
	return nil
}

func (aum *ActiveUserModel) setEmail(email string) error {
	aum.email = email
	return nil
}

func (aum *ActiveUserModel) setPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	aum.password = pass
	return nil

}

func (aum *ActiveUserModel) setIcon(icon *types.ImageTypeFileOrURL) error {
	aum.icon = icon
	return nil
}

func (aum *ActiveUserModel) setRoll(roll string) error {
	rl, err := types.NewAccessROLL(roll)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0003, err.Error())
	}
	aum.roll = rl
	return nil
}

func (aum *ActiveUserModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	aum.auditTrail = auditTrail
	return nil
}

func (aum *ActiveUserModel) Validation() error {
	return nil
}
