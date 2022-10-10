## File input and output in Golang
### Read a file with `bufio.NewScanner`.
```go
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening %s: %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Err() != nil {
			fmt.Printf("error reading file %s", err)
			os.Exit(1)
		}
		fmt.Println(line)
	}
```
### Read and Write a file with `ioutil`.
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
### Read exact bytes from buffer with `io.ReadFull` and Write to file with `io.WriteString`.
```go
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening %s: %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	buf := make([]byte, 8)
	if _, err := io.ReadFull(f, buf); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
			fmt.Println(err)
		}
	}

	io.WriteString(os.Stdout, string(buf))
```
### Write to a file with `fmt`.
```go
destination, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
		os.Exit(1)
	}
	defer destination.Close()

	fmt.Fprintf(destination, "[%s]:", filename)
	fmt.Fprintf(destination, "Using fmt.Fprintf in %s\n", filename)
```
### Copy an entire file at once with `io`.
```go
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	return io.Copy(destination, source)
}

func main() {
    ...
	nBytes, err := Copy(sourceFile, destinationFile)
	if err != nil {
		fmt.Printf("The copy operation failed %q\n", err)
	} else {
		fmt.Printf("Copied %d bytes!\n", nBytes)
	}
}
```
### Write a byte slice to a file with `os`.
```go
    aByteSlice := []byte("text here")
	os.WriteFile(filename, aByteSlice, 0644)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	anotherByteSlice := make([]byte, 100)
	n, err := f.Read(anotherByteSlice)
	if err == io.EOF {
		fmt.Println("EOF:", err)
	}
	fmt.Printf("Read %d bytes: %s", n, anotherByteSlice)
```