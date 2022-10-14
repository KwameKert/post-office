package core

import (
	"time"
)

type Pagination struct {
	CurrentPage  int   `json:"current_page,omitempty"`
	NextPage     int   `json:"next_page,omitempty"`
	PreviousPage int   `json:"previous_page,omitempty"`
	Count        int64 `json:"count"`
}

type Meta struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Message    string      `json:"message"`
}

type Response struct {
	Error bool `json:"error"`
	Code  int  `json:"code"`
	Meta  Meta `json:"meta"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateAppRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateAppRequest struct {
	Id          string `json:"id" bson:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateDomainRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ModuleId    string `json:"module_id"`
}

type UpdateDomainRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ModuleId    string `json:"module_id"`
}

type CreateModuleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AppId       string `json:"app_id"`
}

type TempModule struct {
	Id          string             `bson:"_id"`
	App         []UpdateAppRequest `bson:"app"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type ModuleResponse struct {
	Id          string           `json:"id"`
	App         UpdateAppRequest `json:"app"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Status      string           `json:"status"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type DomainResponse struct {
	Id          string              `json:"id"`
	Module      UpdateModuleRequest `json:"module"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Status      string              `json:"status"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type TempDomain struct {
	Id          string                `bson:"_id"`
	Module      []UpdateModuleRequest `bson:"module"`
	Name        string                `bson:"name"`
	Description string                `bson:"description"`
	Status      string                `bson:"status"`
	CreatedAt   time.Time             `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time             `bson:"updated_at" json:"updated_at"`
}

type UpdateModuleRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AppId       string `json:"app_id"`
}

type CreateLogRequest struct {
	Data     string `json:"data"`
	DomainId string `json:"domain_id"`
	Action   string `json:"action"`
	UserId   string `json:"user_id"`
}

type CreateIssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

type CreatePaymentRequest struct {
	From        int    `json:"from"`
	To          int    `json:"to"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

type SchemaLoginResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
}

type SearchRequest struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	Text     string `json:"text"`
	DomainId string `json:"domain_id"`
	AppId    string `json:"app_id"`
	ModuleId string `json:"module_id"`
	UserId   string `json:"user_id"`
	Action   string `json:"action"`
}
