package new

import (
	"fmt"
	"net/url"
)

// Repo represents the configuration for a repository of templates
type Repo struct {
	Name string
	URI  string
}

// AddRepo adds a repository to the local database
func AddRepo(uri string) error {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}
	switch u.Scheme {
	case "https":

	default:
		return fmt.Errorf("unknown scheme: %s", u.Scheme)
	}
	return nil
}
