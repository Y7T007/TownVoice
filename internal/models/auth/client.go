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

func SaveClient(user *Client) error {

	// Save the client to the database (or in this case, IPFS)
	// For now, we will just print the client to the console
	println("Saving client to the database:")
	println("ID:", user.ID)
	println("Name:", user.Name)
	println("Email:", user.Email)
	println("City:", user.City)
	println("Address:", user.Address)
	return nil
}

func AuthenticateClient(user *Client) interface{} {
	// Authenticate the client
	// For now, we will just print the client to the console
	println("Authenticating client:")
	println("ID:", user.ID)
	println("Name:", user.Name)
	println("Email:", user.Email)
	println("City:", user.City)
	println("Address:", user.Address)
	return true

}
