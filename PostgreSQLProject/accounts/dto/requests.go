package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type GetAccountRequest struct {
	Name string `query:"name"`
}

type PatchAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountRequest struct {
	Name    string `json:"name"`
	NameNew string `json:"name_new"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
