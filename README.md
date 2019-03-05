# workspaces
A CLI to manage your workspaces

## usage
  workspaces [flags]
  workspaces [command]

Available Commands:
  help        Help about any command
  list        Show all projects managed by workspaces
  pick        Add directory to managed projects
  scan        Show all projects managed by workspaces
  setup       initialize ~/.workspaces folder and ~/.workspaces/config
  wd          show project path

Flags:
  -c, --config string    config file (default is $HOME/.workspaces/config)
  -h, --help             help for workspaces
      --log_dir string   log directory, default to ~/.workspaces/log
  -v, --verbose          -v or --verbose for debug information

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
