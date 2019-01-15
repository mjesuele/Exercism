package twofer

// ShareWith outputs a nice message about sharing.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
