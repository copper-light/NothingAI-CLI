# NothingAI CLI
- Nothing AI 플랫폼의 CLI(Command Line Interface) 프로그램

### build
```bash
$ GOOS=darwin GOARCH=amd64 go build -o nothing
```

### Command
```bash
Usage: nothing [COMMNAD] [RESOURCE]

Common Commands:
  get          List resources
  create       Create a new resource
  delete       Delete resources
  edit         Edit the resource information  
  exec         Execute a command in a experimant
  logs         Fetch the logs of a task
                    
Resoruce Commands:
  model        Manage model
  dataset      Manage dataset
  experimant   Manage experiment
  task         Manage task
```