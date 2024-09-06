package handlers

type updateContactRequest struct {
	NewAddress string `json:"new_address" validate:"min=0,max=128"`
	NewName    string `json:"new_name" validate:"alphanum"`
}

type createTemplateRequest struct {
	Text string `json:"text" validate:"required,min=1,max=256"`
}

type signUpRequest struct {
	Name     string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=5,max=64,containsany=$%#@!"`
}

type signInRequest struct {
	Name     string `json:"username,required"`
	Password string `json:"password,required"`
}
