package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type HistoryUserModel struct {
	Id           types.IDENTIFICATION
	activeUserId types.IDENTIFICATION
	nickname     string
	email        string
	password     []byte
	icon         *types.ImageTypeFileOrURL
	roll         types.AccessROLL
	auditTrail   *types.AuditTrail
}

func NewHistoryUserModel(
	id int,
	activeUserId int,
	nickname string,
	email string,
	password string,
	iconURL string,
	imgKey string,
	roll string,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*HistoryUserModel, error) {
	lum := new(HistoryUserModel)

	var err error
	err = errors.Combine(err, lum.setId(id))
	err = errors.Combine(err, lum.setActiveUserId(activeUserId))
	err = errors.Combine(err, lum.setNickname(nickname))
	err = errors.Combine(err, lum.setEmail(email))
	err = errors.Combine(err, lum.setPassword(password))
	err = errors.Combine(err, lum.setRoll(roll))
	if err != nil {
		return new(HistoryUserModel), err
	}
	images, err := types.NewImageTypeFileOrURL(nil, iconURL, imgKey, true)
	if err != nil {
		return new(HistoryUserModel), err
	}
	err = errors.Combine(err, lum.setIcon(images))
	if err != nil {
		return new(HistoryUserModel), err
	}

	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(HistoryUserModel), err
	}
	err = errors.Combine(err, lum.setAuditTrail(com))
	if err != nil {
		return new(HistoryUserModel), err
	}

	return lum, nil

}
func (lum *HistoryUserModel) GetId() types.IDENTIFICATION {
	return lum.Id
}

func (lum *HistoryUserModel) GetActiveUserId() types.IDENTIFICATION {
	return lum.activeUserId
}

func (lum *HistoryUserModel) GetNickname() string {
	return lum.nickname
}

func (lum *HistoryUserModel) GetEmail() string {
	return lum.email
}

func (lum *HistoryUserModel) GetPassword() string {
	return string(lum.password)
}

func (lum *HistoryUserModel) GetIcon() *types.ImageTypeFileOrURL {
	return lum.icon
}

func (lum *HistoryUserModel) GetRoll() types.AccessROLL {
	return lum.roll
}

func (lum *HistoryUserModel) GetAuditTrail() *types.AuditTrail {
	return lum.auditTrail
}

func (lum *HistoryUserModel) setId(Id int) error {
	i, err := types.NewIDENTIFICATION(Id)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	lum.Id = i
	return nil
}

func (lum *HistoryUserModel) setActiveUserId(activeUserId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(activeUserId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	lum.activeUserId = i
	return nil
}

func (lum *HistoryUserModel) setNickname(nickname string) error {
	lum.nickname = nickname
	return nil
}

func (lum *HistoryUserModel) setEmail(email string) error {
	lum.email = email
	return nil
}

func (lum *HistoryUserModel) setPassword(password string) error {
	lum.password = []byte(password)
	return nil
}

func (lum *HistoryUserModel) setIcon(icon *types.ImageTypeFileOrURL) error {
	lum.icon = icon
	return nil
}

func (lum *HistoryUserModel) setRoll(roll string) error {
	rl, err := types.NewAccessROLL(roll)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0003, err.Error())
	}
	lum.roll = rl
	return nil
}

func (lum *HistoryUserModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	lum.auditTrail = auditTrail
	return nil
}
