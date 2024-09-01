# Todo App

## Usage

An cli application for managing tasks in the terminal.

```
$ todo
```

## Example

The following operations can be performed:

```
$ todo new "My new task"
$ todo list
$ todo complete 1
```

### New

The new command takes task description and creates a new task with an assigned ID

```
$ todo new <description>
```

for example:

```
$ todo new "Tidy my desk"
```

adds a new task with the description of "Tidy my desk"

also using option --due allows to set a due time

### List

This shows a list of all of the **uncompleted** tasks, with an option to show all tasks regardless of whether or not they are completed.

for example:

```
$ todo list
ID    Task                                                Created
1     Tidy up my desk                                     a minute ago
3     Change my keyboard mapping to use escape/control    a few seconds ago
```

or for showing all tasks, using a flag (such as -a or --all)

```
$ todo list -a
ID    Task                                                Created          Status        Due
1     Tidy up my desk                                     2 minutes ago    Incomplete    Not set
2     Write up documentation for new project feature      a minute ago     Completed     Not set
3     Change my keyboard mapping to use escape/control    a minute ago     Incomplete    tomorrow
```


### Complete

It marks a task as done

```
$ todo complete <taskid>
```

### Remove

The command deletes a task from the data store

```
$ todo remove <taskid>
```

## Notable Packages Used

- `encoding/csv` for writing out as a csv file
- `strconv` for turning types into strings and visa versa
- `text/tabwriter` for writing out tab aligned output
- `os` for opening and reading files
- `github.com/spf13/cobra` for the command line interface
- `github.com/mergestat/timediff` for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)
 
### Release

You can find the [releases](https://github.com/FireDrop6000/to-do-cli/releases/tag/v1.0.0) of this todo list on the releases tab of this repo.
