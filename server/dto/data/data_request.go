package datadto

type CreateDataRequest struct {
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
}
