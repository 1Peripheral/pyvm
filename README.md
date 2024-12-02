## PYVM : Python virtul environment manager

Simple CLI tool to help you manage your python virtual environments.


### Examples
Creating a virtual environment
```console
user@host:~$ // pyvm create [name] [path]
user@host:~$ pyvm create venv ~/Projects/my_project/virtual_env
```
Deleting a virtual environment
```console
user@host:~$ // pyvm delete [name]
user@host:~$ pyvm create venv 
```


#### Note
All your environments paths are stored in a .pyvm.json file in your
home directory.
