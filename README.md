# SMTPCheckTLS

Tool to simplify checking the validity of and SMTP servers TLS setup

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


## Alternative SMTP TLS Check


* `openssl s_client -connect SMTP_SERVER:PORT -starttls smtp`