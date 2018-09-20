package people

type Person struct {
	ID        string   `json:"id,"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

// The Address entity is used to marshall/unmarshall JSON.
type Address struct {
	City  string `json:"city, omitempty"`
	State string `json:"state,omitempty"`
}