/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bitbucket

import (
	"time"
)

type Account struct {
	Type_       string      `json:"type"`
	Links       *ModelError `json:"links,omitempty"`
	CreatedOn   time.Time   `json:"created_on,omitempty"`
	DisplayName string      `json:"display_name,omitempty"`
	Username    string      `json:"username,omitempty"`
	Uuid        string      `json:"uuid,omitempty"`
}
