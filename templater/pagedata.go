package templater

type pageData struct {
	Name     string      // The name of the page
	Message  string      // The message to display to a user
	Username string      // The username of the user (if logged in)
	Data     interface{} // Injected data
}
