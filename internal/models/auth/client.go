package auth

type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	// OtherDetails can be represented as a map[string]interface{} if the fields are dynamic
	OtherDetails map[string]interface{} `json:"other_details"`
}
