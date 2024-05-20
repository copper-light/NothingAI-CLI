# NothingAI CLI
- Nothing AI 플랫폼의 CLI(Command Line Interface) 프로그램

### build
```bash
$ GOOS=darwin GOARCH=amd64 go build -o nothing
```

### Command
```bash
CLI to Nothing AI

Usage:
  nothing [COMMAND] [flags]

Basic Commands
  create      Create a resource
  describe    Show details of a specific resource
  edit        Edit a resource
  get         List resources

Other Commands
  resources   List types of a resource

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help      help for nothing
  -v, --version   Print version information

Use "nothing [command] --help" for more information about a command.
```

### 지원 리소스의 유형
```bash
$ nothing resources
RESOURCE-TYPE   DESCRIPTION
models          Manage models
datasets        Manage datasets
experiments     Manage experiments
tasks           Manage tasks
```