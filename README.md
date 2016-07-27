# getdns
getdns in Go (work in progress)  

This repo contains very preliminary work on a native Go implementation
of the getdns API. 

It currently uses Miek Gieben's Go DNS library
( https://github.com/miekg/dns ) and returns DNS response structures
populated by that library rather than the response dictionary object
defined in the current getdns API specification. A future iteration
will likely change this.

It only implements the synchronous versions of the 4 query functions,
General(), Address(), Hostname(), and Service(). We are thinking
about the best way to support asynchronous operation in the Go 
implementation. One possibility is to delegate such tasks to the Go
programmer and allow them to use native Go concurrency mechanisms
(e.g. goroutines and channels) to address this need.

It currently only supports stub resolution mode (expected to be the
most common mode of operation).

Items on the todo list include: iterative resolution, support for
OPT parameters, DNSSEC validation, return of validation chain etc.

