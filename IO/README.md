### File input and output in Golang

## Read and Write a file with `ioutil`.
The technique in this section will use the ioutil.WriteFile() and ioutil.ReadFile()
functions. Note that ioutil.ReadFile() does not implement the io.Reader interface
and therefore is a little restrictive.
```go
    r, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("could not read from source file: %s", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(destFile, r, 0664)
	if err != nil {
		fmt.Printf("could not write to destination file: %s", err)
		os.Exit(1)
	}   
```