package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type logUserModel struct {
	historyUserId types.IDENTIFICATION
	activeUserId  types.IDENTIFICATION
	nickname      string
	email         string
	password      []byte
	icon          string
	roll          types.AccessROLL
	auditTrail    *types.AuditTrail
}

func NewLogUserModel(
	historyUserId int,
	activeIdUserId int,
	nickname string,
	email string,
	password string,
	iconURL string,
	roll string,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*logUserModel, error) {
	lum := new(logUserModel)

	var err error
	err = errors.Combine(err, lum.setHistoryUserId(historyUserId))
	err = errors.Combine(err, lum.setActiveUserId(activeIdUserId))
	err = errors.Combine(err, lum.setNickname(nickname))
	err = errors.Combine(err, lum.setEmail(email))
	err = errors.Combine(err, lum.setPassword(password))
	err = errors.Combine(err, lum.setIcon(iconURL))
	err = errors.Combine(err, lum.setRoll(roll))
	if err != nil {
		return new(logUserModel), err
	}
	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(logUserModel), err
	}
	err = errors.Combine(err, lum.setAuditTrail(com))
	if err != nil {
		return new(logUserModel), err
	}

	return lum, nil

}
func (lum *logUserModel) GetHistoryUserId() types.IDENTIFICATION {
	return lum.historyUserId
}

func (lum *logUserModel) GetActiveUserId() types.IDENTIFICATION {
	return lum.activeUserId
}

func (lum *logUserModel) GetNickname() string {
	return lum.nickname
}

func (lum *logUserModel) GetEmail() string {
	return lum.email
}

func (lum *logUserModel) GetPassword() string {
	return string(lum.password)
}

func (lum *logUserModel) GetIcon() string {
	return lum.icon
}

func (lum *logUserModel) GetRoll() types.AccessROLL {
	return lum.roll
}

func (lum *logUserModel) GetAuditTrail() *types.AuditTrail {
	return lum.auditTrail
}

func (lum *logUserModel) setHistoryUserId(historyUserId int) error {
	i, err := types.NewIDENTIFICATION(historyUserId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	lum.historyUserId = i
	return nil
}

func (lum *logUserModel) setActiveUserId(activeUserId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(activeUserId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	lum.historyUserId = i
	return nil
}

func (lum *logUserModel) setNickname(nickname string) error {
	lum.nickname = nickname
	return nil
}

func (lum *logUserModel) setEmail(email string) error {
	lum.email = email
	return nil
}

func (lum *logUserModel) setPassword(password string) error {
	lum.password = []byte(password)
	return nil
}

func (lum *logUserModel) setIcon(icon string) error {
	lum.icon = icon
	return nil
}

func (lum *logUserModel) setRoll(roll string) error {
	rl, err := types.NewAccessROLL(roll)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0003, err.Error())
	}
	lum.roll = rl
	return nil
}

func (lum *logUserModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	lum.auditTrail = auditTrail
	return nil
}
