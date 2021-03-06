## gitlabctl describe user

Describe a user by specifying the user id or username

### Synopsis

Describe a user by specifying the user id or username

```
gitlabctl describe user [flags]
```

### Examples

```
# describe a user by username
gitlabctl describe user john.smith

# describe a user with user id (13)
gitlabctl describe user 13
```

### Options

```
  -h, --help   help for user
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl describe](gitlabctl_describe.md)	 - Describe a gitlab resource

