# wc-tool
WC Tool written as per the specifications in John Crickett's Coding Challenges written in Golang. It supports 5 basic functionalities:
- Count Characters in a file
- Count Words in a file
- Count Bytes in a file
- Count Lines in a file
- Count Lines from Standard Input


## How to run the Project
- Fork the repository and download on your local
- Make sure you have the latest version of go downloaded and installed in your system. If not the same can be done from here https://go.dev/doc/install
- Move to the home directory (containing the `ccwc.go` file)
- Run the program with the command `go run ccwc.go <command line options> <fileName>`
- The Allowed command line options are as follows
  - `-c` : Gives byte count
  - `-l` : Gives line count
  - `-w` : Gives Word count
  - `-m` : Gives char count

 ## Additional Features
 - The Project supports default functionality wherein if the command line argument doesn't match any of the above mentioned 4 options then the application returns the line, word and char count respectively.
 - The Project supports reading text from Standard Input such as `cat test.txt | ccwc -l` to print the number of lines, and similar functionality for byte, char and word count.
