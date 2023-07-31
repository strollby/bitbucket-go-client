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

type Pullrequest struct {
	Type_ string       `json:"type"`
	Links *interface{} `json:"links,omitempty"`
	// The pull request's unique ID. Note that pull request IDs are only unique within their associated repository.
	Id int32 `json:"id,omitempty"`
	// Title of the pull request.
	Title string `json:"title,omitempty"`
	// User provided pull request text, interpreted in a markup language and rendered in HTML
	Rendered *interface{} `json:"rendered,omitempty"`
	Summary  *interface{} `json:"summary,omitempty"`
	// The pull request's current status.
	State       string               `json:"state,omitempty"`
	Author      *Account             `json:"author,omitempty"`
	Source      *PullrequestEndpoint `json:"source,omitempty"`
	Destination *PullrequestEndpoint `json:"destination,omitempty"`
	MergeCommit *interface{}         `json:"merge_commit,omitempty"`
	// The number of comments for a specific pull request.
	CommentCount int32 `json:"comment_count,omitempty"`
	// The number of open tasks for a specific pull request.
	TaskCount int32 `json:"task_count,omitempty"`
	// A boolean flag indicating if merging the pull request closes the source branch.
	CloseSourceBranch bool     `json:"close_source_branch,omitempty"`
	ClosedBy          *Account `json:"closed_by,omitempty"`
	// Explains why a pull request was declined. This field is only applicable to pull requests in rejected state.
	Reason string `json:"reason,omitempty"`
	// The ISO8601 timestamp the request was created.
	CreatedOn time.Time `json:"created_on,omitempty"`
	// The ISO8601 timestamp the request was last updated.
	UpdatedOn time.Time `json:"updated_on,omitempty"`
	// The list of users that were added as reviewers on this pull request when it was created. For performance reasons, the API only includes this list on a pull request's `self` URL.
	Reviewers []Account `json:"reviewers,omitempty"`
	//         The list of users that are collaborating on this pull request.         Collaborators are user that:          * are added to the pull request as a reviewer (part of the reviewers           list)         * are not explicit reviewers, but have commented on the pull request         * are not explicit reviewers, but have approved the pull request          Each user is wrapped in an object that indicates the user's role and         whether they have approved the pull request. For performance reasons,         the API only returns this list when an API requests a pull request by         id.
	Participants []Participant `json:"participants,omitempty"`
}
