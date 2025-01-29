# URL Crashing Tool

This Go program is designed to send a large number of GET requests to a specified URL, potentially causing a denial-of-service (DoS) attack.  It constructs URLs by appending random strings to the base URL provided by the user.  **Use this tool responsibly and only on systems you own or have explicit permission to test.**  Misuse can have serious legal consequences.

## How it Works

1. **User Input:** The program prompts the user to enter the target URL.
2. **Random String Generation:** It generates random strings of 100 characters using a combination of uppercase and lowercase letters.
3. **URL Construction:**  The random string is appended to the base URL provided by the user.
4. **Concurrent Requests:** The program launches 1000 concurrent goroutines (10 loops of 100). Each goroutine sends a GET request to the constructed URL.
5. **Error Handling:** The program prints any errors encountered during the GET requests.
6. **Synchronization:** A `sync.WaitGroup` is used to ensure that the program waits for all goroutines to complete before exiting.

## Usage

1. **Clone the repository (or copy the code):**  If you've cloned, navigate to the project directory.
2. **Build:** `go build`
3. **Run:** `./url-crashing-tool` (or the name you gave the binary)
4. **Enter URL:** The program will prompt you to enter the target URL.  **Be extremely careful what URL you enter.**

## Disclaimer

This tool is provided for educational and testing purposes only.  The author is not responsible for any misuse of this program.  Using this tool against systems without explicit permission is illegal and unethical.  **Do not use this tool to attack any website or server without proper authorization.**  The program's effectiveness in causing a DoS attack is not guaranteed and depends on various factors, including the target server's capacity and network conditions.  Use at your own risk.