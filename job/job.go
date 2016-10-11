package job

import "net/url"

// Detail  contains all the details about a single job for checking if a proxy works or not
type Detail struct {
	Host     string
	Username string
	Password string
}

func (h Detail) ToURL() (rsp *url.URL, err error) {
	rsp, err = url.Parse(h.Host)
	if len(h.Username) > 0 && len(h.Password) > 0 {
		rsp.User = url.UserPassword(h.Username, h.Password)
	}
	return rsp, err
}
