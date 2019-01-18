# go-btrieve
Btrieve library for go
=====================================================

Notes:
------
This library will be updated, over time, as needed. While it
is currently not fully complete (by far), it does offer the
critical pieces needed to open and access a Btrieve file on a
Windows system.

The library is meant for 32-bit Btrieve, but can be modified
to use other dlls. Please note that to use the 32-bit system,
you cannot issue a go run command, since go it 64-bit. Instead,
you must build the program first as a 32-bit program. To do this,
you neet to set 2 environmental variables.
* GOOS=Windows
* GOARCh=386

