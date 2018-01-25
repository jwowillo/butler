"""cmd has helper functions for running shell commands."""

import subprocess


def ssh(host, cmd):
    """ssh into host and run cmd."""
    run('ssh {} << EOF\nset -x\nset -e\n{}\nEOF\n'.format(host, cmd))


def run(cmd):
    """run cmd in bash."""
    pipe = subprocess.Popen('set -e\n{}\n'.format(cmd),
                            stdout=subprocess.PIPE, stderr=subprocess.STDOUT,
                            shell=True)
    for line in iter(pipe.stdout.readline, ''):
        line = line.decode('utf-8')
        if line == '': break
        print(line)
    pipe.stdout.close()
    code = pipe.wait()
    if code: raise subprocess.CalledProcessError(code, cmd)
