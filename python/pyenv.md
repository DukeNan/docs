###  安装

Mac安装：

```shell
brew install pyenv
```

Linux安装：

```shell
$ git clone https://github.com/pyenv/pyenv.git ~/.pyenv
$ echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bash_profile
$ echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bash_profile
```



### 帮助

```shell
pyenv --help

Usage: pyenv <command> [<args>]

Some useful pyenv commands are:
   --version   Display the version of pyenv
   commands    List all available pyenv commands
   exec        Run an executable with the selected Python version
   global      Set or show the global Python version
   help        Display help for a command
   hooks       List hook scripts for a given pyenv command
   init        Configure the shell environment for pyenv
   install     Install a Python version using python-build
   local       Set or show the local application-specific Python version
   prefix      Display prefix for a Python version
   rehash      Rehash pyenv shims (run this after installing executables)
   root        Display the root directory where versions and shims are kept
   shell       Set or show the shell-specific Python version
   shims       List existing pyenv shims
   uninstall   Uninstall a specific Python version
   version     Show the current Python version and its origin
   version-file   Detect the file that sets the current pyenv version
   version-name   Show the current Python version
   version-origin   Explain how the current Python version is set
   versions    List all Python versions available to pyenv
   whence      List all Python versions that contain the given executable
   which       Display the full path to an executable

See `pyenv help <command>' for information on a specific command.
For full documentation, see: https://github.com/pyenv/pyenv#readme
```

### 查看可安装Python版本

```shell
pyenv install --list
```

### 安装特定版本的python

```shell
pyenv install <version>
```



### 查看pyenv安装的版本

```python
pyenv versions
```

###  安装目录

```shell
~/.pyenv/versions
```

### 创建虚拟环境

```shell
# 进入待安装目录，安装虚拟环境
$ ~/.pyenv/versions/3.7.7/bin/python3.7 -m venv py37
```