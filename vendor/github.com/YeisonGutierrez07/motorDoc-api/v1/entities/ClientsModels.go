package entities

type Client struct {
	BaseModel
	UserID int64 `json:"user_id" db:"user_id"`
}
