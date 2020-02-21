package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

const (
	club int = iota
	contest
	event
	public
)

type ApplicationsDetail struct {
	ID               int             `gorm:"type:int(11) AUTO_INCREMENT;primary_key" json:"-"`
	ApplicationID    uuid.UUID       `gorm:"type:char(36);not null" json:"-"`
	UpdateUserTrapID User            `gorm:"embedded;embedded_prefix:update_user_" json:"update_user"`
	Type             ApplicationType `gorm:"embedded" json:"type"`
	Title            string          `gorm:"type:text;not null" json:"title"`
	Remarks          string          `gorm:"type:text;not null" json:"remarks"`
	Amount           int             `gorm:"type:int(11);not null" json:"amount"`
	PaidAt           time.Time       `gorm:"type:timestamp" json:"paid_at"`
	UpdatedAt        time.Time       `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ApplicationType struct {
	Type int `gorm:"tinyint(4);not null"`
}

func (ty ApplicationType) MarshalJSON() ([]byte, error) {
	switch ty.Type {
	case club:
		return json.Marshal("club")
	case contest:
		return json.Marshal("contest")
	case event:
		return json.Marshal("event")
	case public:
		return json.Marshal("public")
	}
	return nil, errors.New("unknown application type")
}

func (ty *ApplicationType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	typeInt, err := GetApplicationTypeFromString(str)
	if err != nil {
		return err
	}

	ty.Type = typeInt
	return nil
}

func GetApplicationTypeFromString(str string) (int, error) {
	switch str {
	case "club":
		return club, nil
	case "contest":
		return contest, nil
	case "event":
		return event, nil
	case "public":
		return public, nil
	}

	return 0, errors.New("unknown application type")
}

func GetApplicationType(str string) (ApplicationType, error) {
	typeInt, err := GetApplicationTypeFromString(str)
	if err != nil {
		return ApplicationType{}, err
	}

	return ApplicationType{Type:typeInt}, nil
}

func (det *ApplicationsDetail) GiveIsUserAdmin(admins []string) {
	if det == nil {
		return
	}

	det.UpdateUserTrapID.GiveIsUserAdmin(admins)
}

func GetApplicationsDetail(id int, giveAdmin bool) (ApplicationsDetail, error) {
	var detail ApplicationsDetail
	err := db.First(&detail, ApplicationsDetail{ID: id}).Error
	if err != nil {
		return ApplicationsDetail{}, err
	}

	if giveAdmin {
		admins, err := GetAdministratorList()
		if err != nil {
			return ApplicationsDetail{}, err
		}
		detail.GiveIsUserAdmin(admins)
	}

	return detail, nil
}

func CreateApplicationsDetail(applicationId uuid.UUID, updateUserTrapID string, typ ApplicationType, title string, remarks string, amount int, paidAt time.Time) (ApplicationsDetail, error) {
	detail := ApplicationsDetail{
		ApplicationID:    applicationId,
		UpdateUserTrapID: User{TrapId: updateUserTrapID},
		Type:             typ,
		Title:            title,
		Remarks:          remarks,
		Amount:           amount,
		PaidAt:           paidAt,
	}

	err := db.Create(&detail).Error
	if err != nil {
		return ApplicationsDetail{}, err
	}

	return detail, nil
}

func PutApplicationsDetail(currentDetailId int, updateUserTrapID string, typ *ApplicationType, title *string, remarks *string, amount *int, paidAt *time.Time) (ApplicationsDetail, error) {
	var detail ApplicationsDetail
	err := db.Find(&detail, ApplicationsDetail{ID: currentDetailId}).Error
	if err != nil {
		return ApplicationsDetail{}, err
	}

	detail.ID = 0 // zero value of int is 0
	detail.UpdateUserTrapID.TrapId = updateUserTrapID
	if typ != nil {
		detail.Type = *typ
	}

	if title != nil {
		detail.Title = *title
	}

	if remarks != nil {
		detail.Remarks = *remarks
	}

	if amount != nil {
		detail.Amount = *amount
	}

	if paidAt != nil {
		detail.PaidAt = *paidAt
	}

	err = db.Create(&detail).Error
	if err != nil {
		return ApplicationsDetail{}, err
	}

	return detail, nil
}
