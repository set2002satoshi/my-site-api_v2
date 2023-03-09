package types

import (
	"mime/multipart"
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
)

type AuditTrail struct {
	Revision  REVISION
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAuditTrail(
	revision int,
	createdAt,
	updatedAt time.Time,
) (*AuditTrail, error) {
	at := new(AuditTrail)
	var err error
	err = errors.Combine(err, at.setRevision(revision))
	err = errors.Combine(err, at.setCreatedAt(createdAt))
	err = errors.Combine(err, at.setUpdatedAt(updatedAt))
	if err != nil {
		return new(AuditTrail), err
	}
	return at, err
}

func (at *AuditTrail) GetRevision() REVISION {
	return at.Revision
}

func (at *AuditTrail) GetCreatedAt() time.Time {
	return at.CreatedAt
}

func (at *AuditTrail) GetUpdatedAt() time.Time {
	return at.UpdatedAt
}

func (at *AuditTrail) setRevision(revision int) error {
	r, err := NewREVISION(revision)
	if err != nil {
		return err
	}
	at.Revision = r
	return nil
}

func (at *AuditTrail) CountUpRevision(currentNum REVISION) error {
	if at.Revision != currentNum {
		return errors.Add(errors.NewCustomError(), errors.EN0004)
	}
	if err := at.setRevision(int(currentNum) + 1); err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0005, err.Error())
	}
	return nil
}

func (at *AuditTrail) setCreatedAt(createdAt time.Time) error {
	at.CreatedAt = createdAt
	return nil
}

func (at *AuditTrail) setUpdatedAt(updatedAt time.Time) error {
	at.UpdatedAt = updatedAt
	return nil
}

type ImageTypeFileOrURL struct {
	ImgFile      multipart.File
	ImgURL       string
	ImgKey       string
	DataTypeFlag bool
}

func NewImageTypeFileOrURL(
	imgFile multipart.File,
	imgURL,
	imgKey string,
	dataTypeFlag bool,
) (*ImageTypeFileOrURL, error) {
	i := new(ImageTypeFileOrURL)

	var err error
	err = errors.Combine(err, i.setImgFile(imgFile))
	err = errors.Combine(err, i.setImgURL(imgURL))
	err = errors.Combine(err, i.setImgKey(imgKey))
	err = errors.Combine(err, i.setDataTypeFlag(dataTypeFlag))

	if err != nil {
		return new(ImageTypeFileOrURL), err
	}
	if err := i.validation(); err != nil {
		return new(ImageTypeFileOrURL), err
	}
	return i, nil

}

func (i *ImageTypeFileOrURL) GetImgFile() multipart.File {
	return i.ImgFile
}

func (i *ImageTypeFileOrURL) GetImgURL() string {
	return i.ImgURL
}

func (i *ImageTypeFileOrURL) GetImgKey() string {
	return i.ImgKey
}

func (i *ImageTypeFileOrURL) GetDataTypeFlag() bool {
	return i.DataTypeFlag
}

func (i *ImageTypeFileOrURL) setImgFile(imgFile multipart.File) error {
	i.ImgFile = imgFile
	return nil
}

func (i *ImageTypeFileOrURL) setImgURL(imgURL string) error {
	i.ImgURL = imgURL
	return nil
}

func (i *ImageTypeFileOrURL) setImgKey(key string) error {
	i.ImgKey = key
	return nil
}

func (i *ImageTypeFileOrURL) setDataTypeFlag(dataTypeFlag bool) error {
	i.DataTypeFlag = dataTypeFlag
	return nil
}

func (i *ImageTypeFileOrURL) validation() error {
	if i.DataTypeFlag {
		if i.ImgURL == "" || i.ImgKey == "" {
			return errors.NewCustomError(errors.EN0006)
		}
	}
	if !i.DataTypeFlag {
		if i.ImgFile == nil {
			return errors.NewCustomError(errors.EN0006)
		}
	}

	return nil
}
