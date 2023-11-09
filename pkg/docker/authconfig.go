package docker

// AuthSpec represent the secret of type docker-config
type AuthSpec struct {
	Auths Auths `json:"auths"`
}

type Auths map[string]Credential

type Credential struct {
	// +optional
	Username string `json:"username"`
	// +optional
	Password string `json:"password"`
	// +optional
	Email string `json:"email"`
	// +optional
	Auth string `json:"auth"`
}
