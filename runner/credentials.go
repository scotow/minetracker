package runner

// The Credentials struct is used to connect to the RCON server.
type Credentials struct {
	// The hostname where the RCON server is running.
	Hostname string

	// The port where the RCON server is listening.
	Port int

	// The password needed to connect to the RCON server.
	// If no password is required, set an empty string.
	Password string
}
