package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename!")
		os.Exit(1)
	}
	filename := os.Args[1]

	aByteSlice := []byte(`The country's roots lay in the October Revolution of 1917, when the Bolsheviks, 
under the leadership of Vladimir Lenin, overthrew the Russian Provisional Government 
that had earlier replaced the House of Romanov of the Russian Empire. The Bolshevik 
victory established the Russian Soviet Republic, the world's first constitutionally 
guaranteed socialist state. Persisting internal tensions escalated into the Russian 
Civil War. By 1922, the Bolsheviks under Vladimir Lenin had emerged victorious, 
forming the Soviet Union.`)

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
}
