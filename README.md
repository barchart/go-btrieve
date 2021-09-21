# go-btrieve
Golang library for Btrieve / Actian Zen / PSQL
==============================================

Notes:
------
This library will be updated, over time, as needed. While it is currently not fully complete (by far), it does offer the
critical pieces needed to open and access a Btrieve file on a Windows system.

The library supports both Windows 32-bit (WBTRV32.DLL) and Windows 64-bit (w64btrv.dll) architectures. It has been tested with PSQL 9, PSQL 11, Zen 13 - 15. So it _seems_ to work with all versions. 

Please note that 32-bin binaries __cannot__ load the 64-bit dll, and vice-versa. If you need to deploy to a 32-bit system, you must first build the program. Also note that as of go version 12, there _seems_ to be a bug where the 32-bit exe does not properly work on older (Windows 2k3, etc.) systems, but these exe's do seem to work on later versions, such as Win2k19. This is regardless of the build OS source.

To set the OS and architecture targets, you must set 2 environmental variables.
* GOOS=Windows
* GOARCH=386

Example code to open a file (in read only mode), read the first record, and close the file.

```
func example() {
	var (
		posblk     = make([]uint16, 160)
		dataBuffer = make([]byte, 256)
		dataLength = uint32(len(dataBuffer))
		dbfile     = []byte(`file.dat` + "\x00")
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
