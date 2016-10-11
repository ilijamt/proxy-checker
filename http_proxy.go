package main

// HTTPProxy contains all the details about a single job for checking if a proxy works or not
type HTTPProxy struct {
	Host     string
	Port     string
	Username string
	Password string
}
