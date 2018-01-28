"""cmd has helper functions for running shell commands."""

import os.path
import subprocess


def clone(host, directory):
    cmd = ''
    if __dir_exists(host, os.path.join(directory, 'butler')):
        cmd = 'cd {}/butler; git pull'.format(directory)
        for _ in ssh(host, cmd): continue
        return False
    else:
        cmd = 'cd {}; git clone https://github.com/jwowillo/butler.git'.format(
                directory)
        for _ in ssh(host, cmd): continue
        return True


def __dir_exists(host, directory):
    try:
        CMD = 'test -d {}'.format(directory)
        for _ in ssh(host, CMD): continue
        return True
    except subprocess.CalledProcessError:
        return False


def build(host, directory):
    CMD = 'cd {}; make butler_gen'.format(directory)
    for _ in ssh(host, CMD): continue


def gen(host, directory):
    CMD = 'cd {}; ./run_gen'.format(directory)
    for _ in ssh(host, CMD): continue


def install_dependencies(host):
    CMD = 'go get -u github.com/jwowillo/gen/cmd/gen_server'
    for _ in ssh(host, CMD): continue


def restart_server(host, directory):
    stop_server(host)
    start_server(host, directory)


def start_server(host, directory):
    if is_server_running(host): return
    CMD = '''
cd {}
nohup ./run_server > /dev/null 2> /dev/null < /dev/null &
'''.format(directory)
    for _ in ssh(host, CMD): continue


def stop_server(host):
    try:
        pid = __server_pid(host)
    except ValueError:
        return
    CMD = 'kill {}'.format(pid)
    for _ in ssh(host, CMD): continue


def is_server_running(host):
    try:
        __server_pid(host)
        return True
    except ValueError:
        return False


def __server_pid(host):
    try:
        return int(list(ssh(host, 'pgrep gen_server'))[-1])
    except subprocess.CalledProcessError:
        raise ValueError('server not running on {}'.format(host))


def ssh(host, cmd):
    """ssh into host and run cmd."""
    return run('ssh {} << EOF\nset -e\n{}\nEOF\n'.format(host, cmd))


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
