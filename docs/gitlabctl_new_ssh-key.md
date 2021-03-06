## gitlabctl new ssh-key

Upload or create ssh key for a gitlab user

### Synopsis

Upload or create ssh key for a gitlab user

```
gitlabctl new ssh-key [flags]
```

### Examples

```
# upload a public ssh key for the current user
gitlabctl -f=/path/to/sshkey.pub -t="my ssh key"

# upload ssh key for another user (only for admin)
gitlabctl -f=/path/to/sshkey.pub -u="lebron.james" -t="the GOAT ssh key"

# upload ssh key for another user with user id 23
gitlabctl -f=path/to/sshkey.pub -u="23" -t="the GOAT ssh key"
```

### Options

```
  -h, --help             help for ssh-key
  -f, --keyfile string   Public SSH key file
  -t, --title string     SSH Key's title (default "uploaded by gitlabctl")
  -u, --user string      Upload the ssh key file to the specified user
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

