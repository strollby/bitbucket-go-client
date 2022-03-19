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

type PipelineSchedule struct {
	Type_ string `json:"type,omitempty"`
	// The UUID identifying the schedule.
	Uuid string `json:"uuid"`
	// Whether the schedule is enabled.
	Enabled bool               `json:"enabled"`
	Target  *PipelineRefTarget `json:"target,omitempty"`
	// The cron expression that the schedule applies.
	CronPattern string `json:"cron_pattern,omitempty"`
	// The timestamp when the schedule was created.
	CreatedOn *time.Time `json:"created_on"`
	// The timestamp when the schedule was updated.
	UpdatedOn *time.Time `json:"updated_on"`
}
