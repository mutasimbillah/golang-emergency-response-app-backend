#!/bin/sh
#set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status
set -e

echo "start the app"
#In short, exec "$@" will run the command given by the command line parameters 
#in such a way that the current process is replaced by it (if the exec is able to execute the command at all)
exec "$@"