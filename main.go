package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Define command line arguments
	filePath := flag.String("file", "", "path to the binary file")
	lang := flag.String("lang", "c", "output programming language format (go, c, csharp, java, rust, python, ruby)")

	// Customize the usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [Options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("\nExample:")
		fmt.Printf("  %s -file=path/to/binary -lang=go\n", os.Args[0])
	}

	// Parse command line arguments
	flag.Parse()

	if *filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Read the binary file
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	// Print the byte array in the selected language format
	length := len(data)
	switch *lang {
	case "go":
		fmt.Println("shellcode_buf length: ", length)
		fmt.Printf("shellcode_buf := []byte{ \n")
		printFormattedBytes(data, "go")
		fmt.Print("}\n")
	case "c":
		fmt.Println("shellcode_buf[] length: ", length)
		fmt.Printf("unsigned char shellcode_buf[] = {\n")
		printFormattedBytes(data, "c")
		fmt.Print("};\n")
	case "csharp":
		fmt.Println("shellcode_buf length: ", length)
		fmt.Printf("byte[] shellcode_buf = new byte[] {\n")
		printFormattedBytes(data, "csharp")
		fmt.Print("};\n")
	case "java":
		fmt.Printf("byte[] shellcode_buf = new byte[%d] {\n", length)
		printFormattedBytes(data, "java")
		fmt.Print("};\n")
	case "rust":
		fmt.Printf("let shellcode_buf: [u8; %d] = [\n", length)
		printFormattedBytes(data, "rust")
		fmt.Print("];\n")
	case "python":
		fmt.Println("shellcode_buf length: ", length)
		fmt.Printf("shellcode_buf = b\"")
		printFormattedBytes(data, "python")
		fmt.Print("\"\n")
	case "ruby":
		fmt.Println("shellcode_buf length: ", length)
		fmt.Printf("shellcode_buf = \"")
		printFormattedBytes(data, "ruby")
		fmt.Print("\"\n")
	default:
		fmt.Printf("Unsupported programming language format: %s\n", *lang)
	}
}

// printFormattedBytes prints the data in the specified language format
func printFormattedBytes(data []byte, lang string) {
	switch lang {
	case "go", "c", "csharp", "java", "rust":
		for i, b := range data {
			if i == 0 {
				fmt.Print("    ")
			}
			if i%12 == 0 && i != 0 {
				fmt.Print("\n    ")
			}
			fmt.Printf("0x%02x, ", b)
		}
		fmt.Print("\n")
	case "python", "ruby":
		for _, b := range data {
			fmt.Printf("\\x%02x", b)
		}
	default:
		fmt.Printf("Unsupported programming language format: %s\n", lang)
	}
}
