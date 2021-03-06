## gitlabctl edit group

Update a group by specifying the group id or path and using flags for fields to modify

### Synopsis

Update a group by specifying the group id or path and using flags for fields to modify

```
gitlabctl edit group [flags]
```

### Examples

```
# edit a group
gitlabctl edit group GroupAZ --desc="Updated group"

# edit a subgroup
gitlabctl edit group GroupX/GroupZ --desc="Updated group"

# edit a group with id (23)
gitlabctl edit group 23 --visibility="public"
```

### Options

```
      --desc string              The description of the resource
  -h, --help                     help for group
      --lfs-enabled              Enable LFS
      --name string              New group name
      --path string              New group path
      --request-access-enabled   Enable request access
      --visibility string        public, internal or private (default "private")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource

