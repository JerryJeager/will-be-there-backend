package invitees

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/JerryJeager/will-be-there-backend/service"
	"github.com/google/uuid"
)

type Invitee struct {
	service.BaseModel
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" binding:"required" gorm:"uniqueIndex:idx_email_event_id"`
	Status    Status    `json:"status" binding:"required"`
	PlusOnes  *PlusOnes `json:"plus_ones"`
	EventID   uuid.UUID `json:"event_id" binding:"required" gorm:"uniqueIndex:idx_email_event_id"`
	Message   string    `json:"message"`
}
type InviteeByEmail struct {
	service.BaseModel
	Email     string    `json:"email" binding:"required" gorm:"uniqueIndex:idx_email_event_id"`
	EventID   uuid.UUID `json:"event_id" binding:"required" gorm:"uniqueIndex:idx_email_event_id"`
}

type Status string

type NewStatus struct {
	Status
}

const (
	APPROVED  Status = "approved"
	PENDING   Status = "pending"
	ATTENDING Status = "attending"
	REJECTED  Status = "rejected"
)

type Invitees []Invitee

type PlusOne struct {
	Name string
}

type PlusOnes []PlusOne

func (c PlusOnes) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *PlusOnes) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		// Unmarshal JSON data into the Values
		return json.Unmarshal(v, c)
	default:
		return fmt.Errorf("unsupported type for Values: %T", v)
	}
}

func (o *Invitee) MarshalJSON() ([]byte, error) {

	invitee := map[string]interface{}{
		"id":         o.ID,
		"first_name": o.FirstName,
		"last_name":  o.LastName,
		"email":      o.Email,
		"status":     o.Status,
		"event_id":   o.EventID,
		"plus_ones":  o.PlusOnes,
		"message":    o.Message,
	}

	return json.Marshal(invitee)
}

func IsValidStatus(status Status) error {
	switch status {
	case APPROVED, PENDING, ATTENDING, REJECTED:
		return nil
	default:
		return errors.New("invalid status")
	}
}
