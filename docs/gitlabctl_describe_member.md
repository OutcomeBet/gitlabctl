## gitlabctl describe member

Describe a member by specifying the username and source

### Synopsis

Describe a member by specifying the username and source

```
gitlabctl describe member [flags]
```

### Examples

```
# describe a member from a group
gitlabctl describe member john.smith --from-group Group1 -o json

# describe a member from a project
gitlabctl describe member john.smith --from-project Group1/Project1 -o yaml
```

### Options

```
  -G, --from-group string     Use a group as the target namespace when performing the command
  -P, --from-project string   Use a project as the target namespace when performing the command
  -h, --help                  help for member
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl describe](gitlabctl_describe.md)	 - Describe a gitlab resource

