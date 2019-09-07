package domain

type DarAccountBody struct {
	UserContext DarAccountUserContext
}

type DarAccountUserContext struct {
	ProfileId string `json:"profileId"`
}