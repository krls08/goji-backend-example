package authz

func FetchUser(name string) user {
	// TODO mogoose

	return user{
		Name: name,
		Age:  30,
	}
}
