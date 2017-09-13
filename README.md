# shibdl

[![Build Status](https://travis-ci.org/nxadm/shibdl.svg?branch=master)](https://travis-ci.org/nxadm/shibdl)

Download files secured by a Shibboleth IdP.

Release can be fount at [releases](https://github.com/nxadm/shibdl/releases).

## Usage

See the help page:

```console
$ shibdl -h
shibdl, vx.x.x
Download files secured by a Shibboleth IdP.
Code, bugs and feature requests: https://github.com/nxadm/shibdl.
Author: Claudio Ramirez <pub.claudio@gmail.com>.
        _       _       _       _       _       _       _       _
     _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
   *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
    // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\

Usage:
  shibdl <URL> -u user [-p password | -P ] [-f file | -d directory] [-l]
  shibdl -h | --help
  shibdl -v | --version

Options:
  <URL>                              URL to download
  -u <user>, --user <user>           Username
  -p <password>, --pass <password>   Password
  -P, --prompt                       Prompt for password
  -d <directory>, --dir <directory>  Directory to safe file (optional)
  -f <file>, --file <file>           Full path of filename (optional)
  -h, --help                         Show this help screen
  -v, --version                      Show the version message

```
