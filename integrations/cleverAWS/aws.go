package cleveraws

import (
	"strings"

	"github.com/Clever/who-is-who/integrations"
)

var (
	// Index specifies the data for querying with the Global secondary index created for
	// queries on AWS usernames.
	Index = integrations.Index{
		Index: "aws",
		Field: "aws",
	}
)

// AwsService does the computation to form AWS usernames with a first initial and last name.
type AwsService struct{}

// Fill uses the first and last name to form an AWS username.
func (a AwsService) Fill(m integrations.UserMap) (integrations.UserMap, error) {
	for email, user := range m {
		if user.FirstName != "" && user.LastName != "" {
			user.AWS = strings.ToLower(user.FirstName[0:1] + user.LastName)
		}
		m[email] = user
	}
	return m, nil
}