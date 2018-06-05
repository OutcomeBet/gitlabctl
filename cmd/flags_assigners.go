// Copyright © 2018 github.com/devopsctl authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

// assignListProjectOptions assigns the flags' values to gitlab.ListProjectsOptions fields.
// If a flag's default value is not changed by the caller,
// it's value will not be assigned to the associated gitlab.ListProjectsOptions field.
func assignListProjectOptions(cmd *cobra.Command) *gitlab.ListProjectsOptions {
	var opts gitlab.ListProjectsOptions
	if cmd.Flag("archived").Changed {
		opts.Archived = gitlab.Bool(getFlagBool(cmd, "archived"))
	}
	if cmd.Flag("order-by").Changed {
		opts.OrderBy = gitlab.String(getFlagString(cmd, "order-by"))
	}
	if cmd.Flag("sort").Changed {
		opts.Sort = gitlab.String(getFlagString(cmd, "sort"))
	}
	if cmd.Flag("search").Changed {
		opts.Search = gitlab.String(getFlagString(cmd, "search"))
	}
	if cmd.Flag("simple").Changed {
		opts.Simple = gitlab.Bool(getFlagBool(cmd, "simple"))
	}
	if cmd.Flag("owned").Changed {
		opts.Owned = gitlab.Bool(getFlagBool(cmd, "owned"))
	}
	if cmd.Flag("membership").Changed {
		opts.Membership = gitlab.Bool(getFlagBool(cmd, "membership"))
	}
	if cmd.Flag("starred").Changed {
		opts.Starred = gitlab.Bool(getFlagBool(cmd, "starred"))
	}
	if cmd.Flag("statistics").Changed {
		opts.Statistics = gitlab.Bool(getFlagBool(cmd, "statistics"))
	}
	if cmd.Flag("visibility").Changed {
		v := getFlagVisibility(cmd)
		opts.Visibility = v
	}
	if cmd.Flag("with-issues-enabled").Changed {
		opts.WithIssuesEnabled = gitlab.Bool(
			getFlagBool(cmd, "with-issues-enabled"))
	}
	if cmd.Flag("with-merge-requests-enabled").Changed {
		opts.WithMergeRequestsEnabled = gitlab.Bool(getFlagBool(cmd,
			"with-merge-requests-enabled"))
	}
	return &opts
}

// assignCreateProjectOptions assigns the flags' values to gitlab.CreateProjectOptions fields.
// If a flag's default value is not changed by the caller,
// it's value will not be assigned to the associated gitlab.CreateProjectOptions field.
func assignCreateProjectOptions(cmd *cobra.Command) (*gitlab.CreateProjectOptions, error) {
	var opts gitlab.CreateProjectOptions

	// default branch is only applied to edit project command
	if f := cmd.Flag("default-branch"); f != nil {
		if cmd.Flag("default-branch").Changed {
			opts.DefaultBranch = gitlab.String(getFlagString(cmd, "default-branch"))
		}
	}

	// change-name is only required when editing a project
	if f := cmd.Flag("change-name"); f != nil {
		if f.Changed {
			opts.Name = gitlab.String(getFlagString(cmd, "change-name"))
		}
	}

	// change-path is only required when editing a project
	if f := cmd.Flag("change-path"); f != nil {
		if f.Changed {
			opts.Path = gitlab.String(getFlagString(cmd, "change-path"))
		}
	}

	// namespace is only required when creating a project
	if f := cmd.Flag("namespace"); f != nil {
		if cmd.Flag("namespace").Changed {
			ns := getFlagString(cmd, "namespace")
			id, err := strconv.Atoi(ns)
			// if not nil take the given number
			if err == nil {
				opts.NamespaceID = &id
			// find the group as string and get it's id
			} else {
				gid, err := getNamespaceID(getFlagString(cmd, "namespace"))
				if err != nil {
					return nil, err
				}
				opts.NamespaceID = gitlab.Int(gid)				
			}
		}
	}

	// common flags for editing and creating a project
	if cmd.Flag("desc").Changed {
		opts.Description = gitlab.String(getFlagString(cmd, "desc"))
	}
	if cmd.Flag("issues-enabled").Changed {
		opts.IssuesEnabled = gitlab.Bool(getFlagBool(cmd, "issues-enabled"))
	}
	if cmd.Flag("merge-requests-enabled").Changed {
		opts.MergeRequestsEnabled = gitlab.Bool(getFlagBool(cmd,
			"merge-requests-enabled"))
	}
	if cmd.Flag("jobs-enabled").Changed {
		opts.JobsEnabled = gitlab.Bool(getFlagBool(cmd, "jobs-enabled"))
	}
	if cmd.Flag("wiki-enabled").Changed {
		opts.WikiEnabled = gitlab.Bool(getFlagBool(cmd, "wiki-enabled"))
	}
	if cmd.Flag("snippets-enabled").Changed {
		opts.SnippetsEnabled = gitlab.Bool(getFlagBool(cmd, "snippets-enabled"))
	}
	if cmd.Flag("resolve-outdated-diff-discussions").Changed {
		opts.ResolveOutdatedDiffDiscussions =
			gitlab.Bool(getFlagBool(cmd, "resolve-outdated-diff-discussions"))
	}
	if cmd.Flag("container-registry-enabled").Changed {
		opts.ContainerRegistryEnabled =
			gitlab.Bool(getFlagBool(cmd, "container-registry-enabled"))
	}
	if cmd.Flag("shared-runners-enabled").Changed {
		opts.SharedRunnersEnabled =
			gitlab.Bool(getFlagBool(cmd, "shared-runners-enabled"))
	}
	if cmd.Flag("visibility").Changed {
		opts.Visibility = getFlagVisibility(cmd)
	}
	if cmd.Flag("public-jobs").Changed {
		opts.PublicJobs = gitlab.Bool(getFlagBool(cmd, "public-jobs"))
	}
	if cmd.Flag("only-allow-merge-if-pipeline-succeeds").Changed {
		opts.OnlyAllowMergeIfPipelineSucceeds =
			gitlab.Bool(
				getFlagBool(cmd, "only-allow-merge-if-pipeline-succeeds"))
	}
	if cmd.Flag("only-allow-merge-if-discussion-are-resolved").Changed {
		opts.OnlyAllowMergeIfAllDiscussionsAreResolved =
			gitlab.Bool(
				getFlagBool(cmd, "only-allow-merge-if-discussion-are-resolved"))
	}
	if cmd.Flag("merge-method").Changed {
		opts.MergeMethod = getFlagMergeMethod(cmd)
	}
	if cmd.Flag("lfs-enabled").Changed {
		opts.LFSEnabled = gitlab.Bool(getFlagBool(cmd, "lfs-enabled"))
	}
	if cmd.Flag("request-access-enabled").Changed {
		opts.RequestAccessEnabled =
			gitlab.Bool(getFlagBool(cmd, "request-access-enabled"))
	}
	if cmd.Flag("tag-list").Changed {
		p := new([]string)
		*p = getFlagStringSlice(cmd, "tag-list")
		opts.TagList = p
	}
	return &opts, nil
}

// assignListUsersOptions assigns the flags' values to gitlab.ListUsersOptions fields.
// If a flag's default value is not changed by the caller,
// it's value will not be assigned to the associated gitlab.ListProjectsOptions field.
func assignListUsersOptions(cmd *cobra.Command) (*gitlab.ListUsersOptions, error) {
	var opts gitlab.ListUsersOptions
	if cmd.Flag("active").Changed {
		opts.Active = gitlab.Bool(getFlagBool(cmd, "active"))
	}
	if cmd.Flag("blocked").Changed {
		opts.Blocked = gitlab.Bool(getFlagBool(cmd, "blocked"))
	}
	if cmd.Flag("search").Changed {
		opts.Search = gitlab.String(getFlagString(cmd, "search"))
	}
	if cmd.Flag("username").Changed {
		opts.Username = gitlab.String(getFlagString(cmd, "username"))
	}
	if cmd.Flag("external-uid").Changed {
		opts.ExternalUID = gitlab.String(getFlagString(cmd, "external-uid"))
	}
	if cmd.Flag("provider").Changed {
		opts.Provider = gitlab.String(getFlagString(cmd, "provider"))
	}
	if cmd.Flag("created-before").Changed {
		ts := getFlagString(cmd, "created-before")
		t, err := newTimeFromString(ts)
		if err != nil {
			return nil, fmt.Errorf("failed parsing time %s, "+
				"got error: %v", ts, err)
		}
		opts.CreatedBefore = t
	}
	if cmd.Flag("created-after").Changed {
		ts := getFlagString(cmd, "created-after")
		t, err := newTimeFromString(ts)
		if err != nil {
			return nil, fmt.Errorf("failed parsing time %s, "+
				"got error: %v", ts, err)
		}
		opts.CreatedAfter = t
	}
	if cmd.Flag("order-by").Changed {
		opts.OrderBy = gitlab.String(getFlagString(cmd, "order-by"))
	}
	if cmd.Flag("sort").Changed {
		opts.Sort = gitlab.String(getFlagString(cmd, "sort"))
	}
	return &opts, nil
}

// assignListGroupOptions assigns the flags' values to gitlab.ListGroupsOptions fields.
// If a flag's default value is not changed by the caller,
// it's value will not be assigned to the associated gitlab.ListGroupsOptions field.
func assignListGroupOptions(cmd *cobra.Command) *gitlab.ListGroupsOptions {
	var opts gitlab.ListGroupsOptions
	if cmd.Flag("all-available").Changed {
		opts.AllAvailable = gitlab.Bool(getFlagBool(cmd, "all-available"))
	}
	if cmd.Flag("owned").Changed {
		opts.Owned = gitlab.Bool(getFlagBool(cmd, "owned"))
	}
	if cmd.Flag("statistics").Changed {
		opts.Statistics = gitlab.Bool(getFlagBool(cmd, "statistics"))
	}
	if cmd.Flag("sort").Changed {
		opts.Sort = gitlab.String(getFlagString(cmd, "sort"))
	}
	if cmd.Flag("search").Changed {
		opts.Search = gitlab.String(getFlagString(cmd, "search"))
	}
	if cmd.Flag("order-by").Changed {
		opts.OrderBy = gitlab.String(getFlagString(cmd, "order-by"))
	}
	return &opts
}

// assignCreateGroupOptions assigns the flags' values to gitlab.CreateGroupOptions fields.
// If a flag's default value is not changed by the caller,
// it's value will not be assigned to the associated gitlab.CreateGroupOptions field.
func assignCreateGroupOptions(cmd *cobra.Command) (*gitlab.CreateGroupOptions, error) {
	var opts gitlab.CreateGroupOptions

	// change-name is only required when editing a group
	if f := cmd.Flag("change-name"); f != nil {
		if f.Changed {
			opts.Name = gitlab.String(getFlagString(cmd, "change-name"))
		}
	}

	// change-path is only required when editing a group
	if f := cmd.Flag("change-path"); f != nil {
		if f.Changed {
			opts.Path = gitlab.String(getFlagString(cmd, "change-path"))
		}
	}

	// namespace is only required when creating a new group
	if f := cmd.Flag("namespace"); f != nil {
		if f.Changed {
			ns := getFlagString(cmd, "namespace")
			id, err := strconv.Atoi(ns)
			// if not nil take the given number
			if err == nil {
				opts.ParentID = &id
			// find the group as string and get it's id
			} else {
				gid, err := getGroupID(getFlagString(cmd, "namespace"))
				if err != nil {
					return nil, err
				}
				opts.ParentID = gitlab.Int(gid)	
			}
		}
	}

	if cmd.Flag("desc").Changed {
		opts.Description = gitlab.String(getFlagString(cmd, "desc"))
	}

	if cmd.Flag("visibility").Changed {
		v := getFlagVisibility(cmd)
		opts.Visibility = v
	}
	if cmd.Flag("lfs-enabled").Changed {
		opts.LFSEnabled = gitlab.Bool(getFlagBool(cmd, "lfs-enabled"))
	}
	if cmd.Flag("request-access-enabled").Changed {
		opts.RequestAccessEnabled = gitlab.Bool(
			getFlagBool(cmd, "request-access-enabled"))
	}
	return &opts, nil
}