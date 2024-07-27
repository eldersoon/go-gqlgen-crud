// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type IUser interface {
	IsIUser()
	GetID() string
	GetName() string
	GetSetor() string
}

type InputProject struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type Mutation struct {
}

type Project struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Leader *User  `json:"leader"`
}

type Query struct {
}

type ResponseCreateProject struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	UserID  string `json:"userId"`
	Message string `json:"message"`
}

type ResponseCreateUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Setor   string `json:"setor"`
	Message string `json:"message"`
}

func (ResponseCreateUser) IsIUser()              {}
func (this ResponseCreateUser) GetID() string    { return this.ID }
func (this ResponseCreateUser) GetName() string  { return this.Name }
func (this ResponseCreateUser) GetSetor() string { return this.Setor }

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Setor    string `json:"setor"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) IsIUser()              {}
func (this User) GetID() string    { return this.ID }
func (this User) GetName() string  { return this.Name }
func (this User) GetSetor() string { return this.Setor }

type UserInput struct {
	Name     string `json:"name"`
	Setor    string `json:"setor"`
	Username string `json:"username"`
	Password string `json:"password"`
}
