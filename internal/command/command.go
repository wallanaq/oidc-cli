package command

import (
	"fmt"
	"log"

	"github.com/wallanaq/oidc-cli/pkg/fileutils"
)

// HandleOutput manages the output of data based on the provided outputFlag.
//
// It supports three modes:
//   - "-": Outputs the data to standard output (stdout).
//   - "": Does nothing, effectively discarding the output.
//   - <filename>: Saves the data to the specified file.
//
// If an error occurs during file saving, it returns an error.
//
// It logs the actions taken, such as writing to stdout, doing nothing, or saving to a file.
func HandleOutput(data []byte, outputFlag string) error {

	switch outputFlag {
	case "-":
		fmt.Println(string(data))

	case "":
		log.Println("No explicit output target specified, HandleOutput doing nothing.")

	default:
		log.Printf("Handling explicit output to file: %s", outputFlag)

		if err := fileutils.Save(outputFlag, data); err != nil {
			return fmt.Errorf("failed to save output to %s: %w", outputFlag, err)
		}

		log.Printf("Output successfully saved to: %s", outputFlag)
	}

	return nil

}
