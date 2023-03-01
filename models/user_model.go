package models

import (
	"mime/multipart"
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"golang.org/x/crypto/bcrypt"
)

type activeUserModel struct {
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
) (*activeUserModel, error) {
	aum := new(activeUserModel)

	var err error
	err = errors.Combine(err, aum.setUserId(userId))
	err = errors.Combine(err, aum.setNickname(nickname))
	err = errors.Combine(err, aum.setEmail(email))
	err = errors.Combine(err, aum.setPassword(password))
	err = errors.Combine(err, aum.setRoll(roll))
	if err != nil {
		return new(activeUserModel), err
	}

	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(activeUserModel), err
	}

	err = errors.Combine(err, aum.setAuditTrail(com))
	if err != nil {
		return new(activeUserModel), err
	}

	images, err := types.NewImageTypeFileOrURL(iconFile, iconURL, iconImageFlag)
	if err != nil {
		return new(activeUserModel), err
	}

	err = errors.Combine(err, aum.setIcon(images))
	if err != nil {
		return new(activeUserModel), err
	}

	if err := aum.Validation(); err != nil {
		return new(activeUserModel), err
	}

	return aum, nil
}

func (aum *activeUserModel) GetUserId() types.IDENTIFICATION {
	return aum.userId
}

func (aum *activeUserModel) GetNickname() string {
	return aum.nickname
}

func (aum *activeUserModel) GetEmail() string {
	return aum.email
}

func (aum *activeUserModel) GetPassword() string {
	return string(aum.password)
}

func (aum *activeUserModel) GetIcon() *types.ImageTypeFileOrURL {
	return aum.icon
}

func (aum *activeUserModel) GetRoll() types.AccessROLL {
	return aum.roll
}

func (aum *activeUserModel) GetAuditTrail() *types.AuditTrail {
	return aum.auditTrail
}

func (aum *activeUserModel) setUserId(userId int) error {
	i, err := types.NewIDENTIFICATION(userId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	aum.userId = i
	return nil
}

func (aum *activeUserModel) setNickname(nickname string) error {
	aum.nickname = nickname
	return nil
}

func (aum *activeUserModel) setEmail(email string) error {
	aum.email = email
	return nil
}

func (aum *activeUserModel) setPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	aum.password = pass
	return nil

}

func (aum *activeUserModel) setIcon(icon *types.ImageTypeFileOrURL) error {
	aum.icon = icon
	return nil
}

func (aum *activeUserModel) setRoll(roll string) error {
	rl, err := types.NewAccessROLL(roll)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0003, err.Error())
	}
	aum.roll = rl
	return nil
}

func (aum *activeUserModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	aum.auditTrail = auditTrail
	return nil
}

func (aum *activeUserModel) Validation() error {
	return nil
}
