package factories

type UserFactory struct {
}

// Definition Define the model's default state.
func (f *UserFactory) Definition() map[string]any {
	return map[string]any{
		"Name": "users",
	}
}
