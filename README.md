# UNIX Shell in 50 lines of Go

This experiment involved trying to build a simple UNIX-like shell in Golang.
Currently, we are able to execute commands which have their own binaries. Other 
than that, for built-in commands within the shell we have added support for 
`cd` and `exit` only. This project can be extended to add support for others.

## Few Ideas
1. It would be good if we could go down to the system-call level and handle 
the implementation of `cd` and `exit`
2. It would be good if we could avoid using the `exec` package and implement the
internals ourselves. These would include
  - Writing to stdout
  - Using stderr for errors
  - Checking for number of args
  - Making system-calls

## Next Steps
Rebuild the entire project in C with near-zero abstractions.

