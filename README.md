# check-shib3idp-login

[![Build Status](https://travis-ci.org/KULeuven-CCIS/check-shib3idp-login.svg?branch=master)](https://travis-ci.org/KULeuven-CCIS/check-shib3idp-login)

A Nagios/Icinga plugin to check an end-to-end user/pass login to a Shibboleth Idp3 instance. The program can also be used standalone.

## Usage

See the help page:

```
$ shibdl -h
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
  -u <user>, --user <user>			 Username
  -p <password>, --pass <password>   Password
  -P, --prompt                       Prompt for password
  -d <directory>, --dir <directory>  Directory to safe file (optional)
  -f <file>, --file <file>           Full path of filename (optional)
  -h, --help                         Show this help screen
  -v, --version                      Show the version message

```
