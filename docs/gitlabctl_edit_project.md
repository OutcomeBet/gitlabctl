## gitlabctl edit project

Edit a project by specifying the project id or path and using flags for fields to modify

### Synopsis

Edit a project by specifying the project id or path and using flags for fields to modify

```
gitlabctl edit project [flags]
```

### Examples

```
# update a project by path
gitlabctl edit project ProjectX --desc="A go project"
gitlabctl edit project GroupX/ProjectX --merge-method=rebase_merge 

# update a project with id (23)
gitlabctl edit project 3 --desc="A go project"
```

### Options

```
      --ci-config-path string                         The path to CI config file
      --container-registry-enabled                    Enable container registry for this project
      --default-branch string                         The default branch (default "master")
      --desc string                                   The description of the resource
  -h, --help                                          help for project
      --issues-enabled                                Enable issues (default true)
      --jobs-enabled                                  Enable jobs (default true)
      --lfs-enabled                                   Enable LFS
      --merge-method string                           Set the merge method used. (available: 'merge', 'rebase_merge', 'ff') (default "merge")
      --merge-requests-enabled                        Enable merge requests (default true)
      --name string                                   New project name
      --only-allow-merge-if-discussion-are-resolved   Set whether merge requests can only be merged when all the discussions are resolved
      --only-allow-merge-if-pipeline-succeeds         Set whether merge requests can only be merged with successful jobs
      --path string                                   New project path
      --printing-merge-request-link-enabled           Show link to create/view merge request when pushing from the command line (default true)
      --public-jobs                                   If true, jobs can be viewed by non-project-members
      --request-access-enabled                        Enable request access
      --resolve-outdated-diff-discussions             Automatically resolve merge request diffs discussions on lines changed with a push
      --shared-runners-enabled                        Enable shared runners for this project
      --snippets-enabled                              Enable snippets (default true)
      --tag-list strings                              The list of tags for a project; put array of tags, that should be finally assigned to a project.
                                                      Example: --tag-list='tag1,tag2'
      --visibility string                             public, internal or private (default "private")
      --wiki-enabled                                  Enable wiki (default true)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource

