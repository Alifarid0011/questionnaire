package dto

type CheckPermissionDTO struct {
	Sub         string `json:"sub" validate:"required"` //role or user_uid
	Act         string `json:"act" validate:"required"` //  GET, POST, PUT, DELETE
	Obj         string `json:"obj" validate:"required"` // /user/all or /user/:id
	AllowOrDeny string `json:"allow_or_deny" validate:"required,oneof=allow deny"`
}

type GroupingDTO struct {
	Parent string `json:"parent" validate:"required"`
	Child  string `json:"child" validate:"required"`
}

type RolesResponse struct {
	Roles []string `json:"roles"`
}
