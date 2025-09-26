package models

type PermissionDTO struct {
	Sub    string `bson:"v0"`
	Obj    string `bson:"v1"`
	Act    string `bson:"v2"`
	Attr   string `bson:"v3"`
	Eft    string `bson:"v4"`
	Entity string `bson:"entity"`
}
type CasbinPolicy struct {
	Subject string `json:"subject"`
	Action  string `json:"action"`
	Object  string `json:"object"`
}

type Permission struct {
	Action string `json:"action"`
	Object string `json:"object"`
}

type SubjectWithPermissions struct {
	Subject     string       `json:"subject"`
	Permissions []Permission `json:"permissions"`
}

type CategorizedPermissions struct {
	Subject     string                  `json:"subject"`
	Permissions map[string][]Permission `json:"permissions"` // دسته‌بندی بر اساس category
}
