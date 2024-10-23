package response

type LoginResponseDto struct {
	AccessToken string		`json:"accessToken"`
	UserId		string 		`json:"userId"`
}

type RegisterResponseDto struct {
	Id			string	`json:"id"`
	FullName 	string	`json:"fullName"`
	Email 		string	`json:"email"`
	Username 	string	`json:"username"`	
	Role 		string	`json:"role"`
}