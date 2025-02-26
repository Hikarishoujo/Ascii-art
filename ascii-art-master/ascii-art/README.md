## ASCII Art Generator

### Overview
This script takes a single string input and converts it into ASCII art using the  fond standard. It supports multiple lines and handles invalid characters gracefully.

## Features
- Converts input text into ASCII art based on the standard fond.
- Handles multi-line inputs using `\n`.
- Detects and prevents out-of-range ASCII characters.
- Maintains the format of new lines.

## Usage
### Running the Script
Run the script using:
```sh
$ go run main.go "Your text here"
```

### Example
Input:
```sh
$ go run main.go "Hello"
```
Output:
```
 _   _      _ _       
| | | |    | | |      
| |_| | ___| | | ___   
|  _  |/ _ \ | |/ _ \   
| | | |  __/ | | (_) |   
\_| |_/\___|_|_|\___/    
```

## File Structure
- `main.go` - The main script.
- `standard.txt` - ASCII font template used for character conversion.
- `utils` - Folder containing the functions used in `main.go`.
-- `validateinput.txt` - Function that reads, validates and saves user's input alongside all ascii-art characters of the standard fond.
-- `charmanipultion.go` - Function tht correlates the requested characters with the ascii-art characters of the standard fond.
-- `exhibitart.go` - Functions that handles printing the requested output.


## How It Works
1. Reads the input text from the command line.
2. Splits the fond file into character sections.
3. Constructs a 2D table mapping ASCII characters to their art representations.
4. Iterates through the input and constructs the ASCII art output.
5. Prints the final output.

## Error Handling
- If no argument is provided, the script prints: `Please enter a valid number of arguments.`
- If an invalid ASCII character is found, it prints: `Found an Invalid Ascii Character.`

## Requirements
- Go installed on your system.
- `standard.txt` file in the same directory as `main.go`.

## Notes
- The script supports printable ASCII characters (from ASCII 32 to ASCII 126).
- Ensure that the fond file follows the correct format (each character is defined in an 8-line block, separated by `\n\n`).

## License
This project is open-source. Feel free to modify and enhance it!

