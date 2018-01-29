"""cmd has helper functions for running commands."""

import os.path
import subprocess


def clone(host, directory):
    """
    clone butler onto host in directory.

    Will pull if the directory already exists.
    """
    if __dir_exists(host, os.path.join(directory, 'butler')):
        __finish(ssh(host, 'cd {}/butler; git pull'.format(directory)))
        return False
    else:
        tmpl = 'cd {}; git clone https://github.com/jwowillo/butler.git'
        __finish(ssh(host, tmpl.format(directory)))
        return True


def build(host, directory):
    """build butler that is in directory on host."""
    __finish(ssh(host, 'cd {}; make butler_gen'.format(directory)))


def gen(host, directory):
    """gen in butler directory on host."""
    for line in ssh(host, 'cd {}; ./run_gen'.format(directory)):
        print(line)


def install_dependencies(host):
    """install_dependencies for butler on host."""
    cmd = 'go get -u github.com/jwowillo/gen/cmd/gen_server'
    __finish(ssh(host, cmd))


def restart_server(host, directory):
    """restart_server on host in butler directory."""
    stop_server(host)
    start_server(host, directory)


def start_server(host, directory):
    """start_server on host in butler directory."""
    if is_server_running(host): return
    tmpl = 'cd {}; nohup ./run_server > /dev/null 2> /dev/null < /dev/null &'
    __finish(ssh(host, tmpl.format(directory)))


def stop_server(host):
    """stop_server on host."""
    try:
        pid = __server_pid(host)
    except ValueError:
        return
    __finish(ssh(host, 'kill {}'.format(pid)))


def is_server_running(host):
    """is_server_running returns True if the server is running on host."""
    try:
        __server_pid(host)
        return True
    except ValueError:
        return False


def __dir_exists(host, directory):
    """__dir_exists returns True if the directory exists on the host."""
    try:
        __finish(ssh(host, 'test -d {}'.format(directory)))
        return True
    except ValueError:
        return False


def ssh(host, cmd):
    """
    ssh into host and run cmd.

    paramiko imported here so that hosts don't need it installed.
    """
    import paramiko
    global __CLIENT
    if __CLIENT is None:
        name, host = host.split('@')
        __CLIENT = paramiko.SSHClient()
        __CLIENT.set_missing_host_key_policy(paramiko.AutoAddPolicy())
        __CLIENT.connect(host, username=name)
    stdout = __CLIENT.exec_command('set -e\n{}'.format(cmd))[1]
    for line in stdout:
        yield line
    if stdout.channel.recv_exit_status() != 0:
        raise ValueError('failed to execute command')


def run(cmd):
    """run cmd in bash."""
    pipe = subprocess.Popen('set -e\n{}\n'.format(cmd),
                            stdout=subprocess.PIPE, stderr=subprocess.STDOUT,
                            shell=True)
    for line in iter(pipe.stdout.readline, ''):
        line = line.decode('utf-8')
        if line == '': break
        yield line
    pipe.stdout.close()
    code = pipe.wait()
    if code: raise subprocess.CalledProcessError(code, cmd)


def __finish(generator):
    """__finish consumes the generator."""
    for _ in generator: continue


def __server_pid(host):
    """
    server_pid returns the PID of the server on the host.

    Raises a ValueError if the server isn't running.
    """
    try:
        return int(list(ssh(host, 'pgrep gen_server'))[-1])
    except ValueError:
        raise ValueError('server not running on {}'.format(host))


__CLIENT = None
"""__CLIENT is a global ssh client that can be reused.a"""


