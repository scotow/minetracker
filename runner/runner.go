package runner

// A Runner is a way to contact/run a RCON command on the server.
type Runner interface {
	// Run the following command on the RCON server.
	// Should return the result of the command as returned from the server, and an error if one occurred.
	Run(string) (string, error)
}
