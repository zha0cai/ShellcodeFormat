Create raw shellcode file.

Usage: ShellcodeFormat.exe [Options]

Options:
  -file string
         path to the binary file
        
  -lang string
        output programming language format (go, c, csharp, java, rust, python, ruby) (default "c")

Example:
 ShellcodeFormat.exe -file=path/to/payload.bin -lang=go
