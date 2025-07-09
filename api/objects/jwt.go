package objects

struct Validation {
	name 		string `json:"name"`
	status 	string `json:"status"`
	value 	string `json:"value"`
}

struct VerifiedPayload {
	UserID 	string 		`json:"user_id"`
	Validations []Validation 	`json:"validations"`
	Site 		number 		`json:"site"`
	ExpiresAt 	Date 			`json:"expires_at"`
	IP 		string 		`json:"ip_address"`
}