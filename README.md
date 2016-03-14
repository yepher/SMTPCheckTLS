# SMTPCheckTLS

CheckTLS is a tool to simplify checking the validity of and SMTP servers TLS setup

When cloning project clone it to a directory names `checktls` so binary will be named appropriately.


## Usage

```
NAME:
   Quck SMTP TLS Checker - ./checktls --n smtp.gmail.com

USAGE:
   checkTLS [global options] command [command options] [arguments...]
   
VERSION:
   0.0.2
   
COMMANDS:
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --port, -p "587"	Optional server port. Default 587
   --hostnames, -n 	One or more SMTP hostnames comma delimited.
   --help, -h		show help
   --version, -v	print the version

```


## Alternative ways to check SMTP TLS configuration

* `openssl s_client -connect SMTP_SERVER:PORT -starttls smtp`
* 
