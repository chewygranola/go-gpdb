package main

var (
	programName    = "gpdb"
	programVersion = "3.6"
)

func main() {
	// Execute the cobra CLI & run the program
	rootCmd.Execute()
}
