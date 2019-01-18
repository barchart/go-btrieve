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

Example code to open a file (in read only mode), read the first record, and close the file.

```
func example() {
	var (
		posblk     = make([]uint16, 160)
		dataBuffer = make([]byte, 256)
		dataLength = uint32(len(dataBuffer))
		dbfile     = []byte((`file.dat` + "\x00")
		keyBuffer  = make([]byte, 255)
		keyNumber  = int8(-2) // read only mode
		ist        uint16
    	)

	ist, err = btrieve.BTRV(btrieve.B_OPEN, posblk, dataBuffer, dataLength, dbfile, keyNumber)
	if err != nil {
		log.Printf("Error initializing btrieve. %s", err.Error())
	} else if ist != 0 {
		log.Println("Error opening %s. %d", string(dfbile), ist)
	} else {
		log.Printf("Successfully opened %s.", string(dbfile))
		// Success
		keyNumber = int8(0)
		ist, _ = btrieve.BTRV(btrieve.B_GET_FIRST, posblk, dataBuffer, dataLength, keyBuffer, keyNumber)
		fmt.Println("KEY", string(keyBuffer))
        	fmt.Println("DATA", string(dataBuffer))

    		ist, _ = btrieve.BTRV(btrieve.B_CLOSE, posblk, dataBuffer, dataLength, dbfile, keyNumber)
    		if ist != 0 {
	    		log.Printf("Error closing %s. %s", string(dbfile), err.Error())
	    	}

    		btrieve.Release()
	}
}
```

Hope this library helps!
