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
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type BranchRestrictionsApiService service

/*
BranchRestrictionsApiService List branch restrictions
Returns a paginated list of all branch restrictions on the repository.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.
 * @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.
 * @param optional nil or *BranchRestrictionsApiRepositoriesWorkspaceRepoSlugBranchRestrictionsGetOpts - Optional Parameters:
     * @param "Kind" (optional.String) -  Branch restrictions of this type
     * @param "Pattern" (optional.String) -  Branch restrictions applied to branches of this pattern
     * @param "Page" (optional.Int32) -  page
@return PaginatedBranchrestrictions
*/

type BranchRestrictionsApiRepositoriesWorkspaceRepoSlugBranchRestrictionsGetOpts struct {
	Kind    optional.String
	Pattern optional.String
	Page    optional.Int32
}

func (a *BranchRestrictionsApiService) RepositoriesWorkspaceRepoSlugBranchRestrictionsGet(ctx context.Context, repoSlug string, workspace string, localVarOptionals *BranchRestrictionsApiRepositoriesWorkspaceRepoSlugBranchRestrictionsGetOpts) (PaginatedBranchrestrictions, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue PaginatedBranchrestrictions
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/branch-restrictions"
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Kind.IsSet() {
		localVarQueryParams.Add("kind", parameterToString(localVarOptionals.Kind.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pattern.IsSet() {
		localVarQueryParams.Add("pattern", parameterToString(localVarOptionals.Pattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v PaginatedBranchrestrictions
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
BranchRestrictionsApiService Delete a branch restriction rule
Deletes an existing branch restriction rule.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id The restriction rule&#x27;s id
  - @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.
  - @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.
*/
func (a *BranchRestrictionsApiService) RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete(ctx context.Context, id string, repoSlug string, workspace string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/branch-restrictions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
BranchRestrictionsApiService Get a branch restriction rule
Returns a specific branch restriction rule.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id The restriction rule&#x27;s id
  - @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.
  - @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.

@return Branchrestriction
*/
func (a *BranchRestrictionsApiService) RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet(ctx context.Context, id string, repoSlug string, workspace string) (Branchrestriction, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Branchrestriction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/branch-restrictions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Branchrestriction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
BranchRestrictionsApiService Update a branch restriction rule
Updates an existing branch restriction rule.  Fields not present in the request body are ignored.  See [&#x60;POST&#x60;](/cloud/bitbucket/rest/api-group-branch-restrictions/#api-repositories-workspace-repo-slug-branch-restrictions-post) for details.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The new version of the existing rule
  - @param id The restriction rule&#x27;s id
  - @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.
  - @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.

@return Branchrestriction
*/
func (a *BranchRestrictionsApiService) RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut(ctx context.Context, body Branchrestriction, id string, repoSlug string, workspace string) (Branchrestriction, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Branchrestriction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/branch-restrictions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Branchrestriction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
BranchRestrictionsApiService Create a branch restriction rule
Creates a new branch restriction rule for a repository.  &#x60;kind&#x60; describes what will be restricted. Allowed values include: &#x60;push&#x60;, &#x60;force&#x60;, &#x60;delete&#x60;, &#x60;restrict_merges&#x60;, &#x60;require_tasks_to_be_completed&#x60;, &#x60;require_approvals_to_merge&#x60;, &#x60;require_default_reviewer_approvals_to_merge&#x60;, &#x60;require_no_changes_requested&#x60;, &#x60;require_passing_builds_to_merge&#x60;, &#x60;require_commits_behind&#x60;, &#x60;reset_pullrequest_approvals_on_change&#x60;, &#x60;smart_reset_pullrequest_approvals&#x60;, &#x60;reset_pullrequest_changes_requested_on_change&#x60;, &#x60;require_all_dependencies_merged&#x60;, &#x60;enforce_merge_checks&#x60;, and &#x60;allow_auto_merge_when_builds_pass&#x60;.  Different kinds of branch restrictions have different requirements:  * &#x60;push&#x60; and &#x60;restrict_merges&#x60; require &#x60;users&#x60; and &#x60;groups&#x60; to be   specified. Empty lists are allowed, in which case permission is   denied for everybody.  The restriction applies to all branches that match. There are two ways to match a branch. It is configured in &#x60;branch_match_kind&#x60;:  1. &#x60;glob&#x60;: Matches a branch against the &#x60;pattern&#x60;. A &#x60;&#x27;*&#x27;&#x60; in    &#x60;pattern&#x60; will expand to match zero or more characters, and every    other character matches itself. For example, &#x60;&#x27;foo*&#x27;&#x60; will match    &#x60;&#x27;foo&#x27;&#x60; and &#x60;&#x27;foobar&#x27;&#x60;, but not &#x60;&#x27;barfoo&#x27;&#x60;. &#x60;&#x27;*&#x27;&#x60; will match all    branches. 2. &#x60;branching_model&#x60;: Matches a branch against the repository&#x27;s    branching model. The &#x60;branch_type&#x60; controls the type of branch    to match. Allowed values include: &#x60;production&#x60;, &#x60;development&#x60;,    &#x60;bugfix&#x60;, &#x60;release&#x60;, &#x60;feature&#x60; and &#x60;hotfix&#x60;.  The combination of &#x60;kind&#x60; and match must be unique. This means that two &#x60;glob&#x60; restrictions in a repository cannot have the same &#x60;kind&#x60; and &#x60;pattern&#x60;. Additionally, two &#x60;branching_model&#x60; restrictions in a repository cannot have the same &#x60;kind&#x60; and &#x60;branch_type&#x60;.  &#x60;users&#x60; and &#x60;groups&#x60; are lists of users and groups that are except from the restriction. They can only be configured in &#x60;push&#x60; and &#x60;restrict_merges&#x60; restrictions. The &#x60;push&#x60; restriction stops a user pushing to matching branches unless that user is in &#x60;users&#x60; or is a member of a group in &#x60;groups&#x60;. The &#x60;restrict_merges&#x60; stops a user merging pull requests to matching branches unless that user is in &#x60;users&#x60; or is a member of a group in &#x60;groups&#x60;. Adding new users or groups to an existing restriction should be done via &#x60;PUT&#x60;.  Note that branch restrictions with overlapping matchers is allowed, but the resulting behavior may be surprising.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The new rule
  - @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.
  - @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.

@return Branchrestriction
*/
func (a *BranchRestrictionsApiService) RepositoriesWorkspaceRepoSlugBranchRestrictionsPost(ctx context.Context, body Branchrestriction, repoSlug string, workspace string) (Branchrestriction, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Branchrestriction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/branch-restrictions"
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v Branchrestriction
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
