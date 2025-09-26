package dto

type UIDQuery struct {
	UID string `form:"uid"  validate:"required,objectid"`
}
