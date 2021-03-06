# workspaces

Workspaces is a simple cli for workspace management.

As software engineers we inevitably have to juggle between different projects. Maybe you have a good filesystem structure to organize your projects, but from time to time, things are scattered everywhere. Or sometimes you need a simple cli to give you the Path of a project in order to pipe it to other command.





## TODO:
```
[ ] add prune
[ ] add tag to projects
[ ] list by tag
[ ] auto-complete

-----
[ ] add docker support
[ ] scratch container to control environment
```

## usage
```
Navigate between your projects with ease

Usage:
  workspaces [flags]
  workspaces [command]

Available Commands:
  alias       Set alias to project
  help        Help about any command
  list        Show all projects in registry
  pick        Add directory to registry 
  remove      remove project from registry
  scan        Show all projects managed by workspaces
  setup       initialize ~/.workspaces folder and ~/.workspaces/config
  wd          show project path

Flags:
  -c, --config string    config file (default is $HOME/.workspaces/config)
  -h, --help             help for workspaces
      --log_dir string   log directory, default to ~/.workspaces/log
  -v, --verbose          -v or --verbose for debug information

Use "workspaces [command] --help" for more information about a command.
```

## bash
to enable easy access for projects, add the following to bashrc
```
alias ws="workspaces"
alias wsl="workspaces ls"
wcd () {
  if [ "$#" -ne 1 ]; then
    workspaces wd
  else
    cd $(workspaces wd $1)
  fi  
}
```
