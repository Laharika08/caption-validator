
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "solution/internal/parser"
    "solution/internal/validator"
    "solution/internal/sender"
)

type ValidationError struct {
    Type        string `json:"type"`
    Description string `json:"description"`
    Details     string `json:"details,omitempty"`
}

func main() {
    // Separate logs from results
    log.SetOutput(os.Stderr)

    tStart := flag.Float64("t_start", 0, "Start time in seconds")
    tEnd := flag.Float64("t_end", 0, "End time in seconds")
    coverage := flag.Float64("coverage", 80, "Required coverage percentage")
    endpoint := flag.String("endpoint", "", "Validation endpoint URL")
    flag.Parse()

    if flag.NArg() < 1 {
        log.Println("No captions file provided")
        os.Exit(1)
    }

    filePath := flag.Arg(0)

    // Detect and parse captions
    captions, format, err := parser.ParseCaptions(filePath)
    if err != nil {
        log.Println("Error parsing captions:", err)
        os.Exit(1)
    }

    // Validate coverage
    coverageValid, coverageErr := validator.ValidateCoverage(captions, *tStart, *tEnd, *coverage)
    var failures []ValidationError
    if !coverageValid {
        failures = append(failures, ValidationError{
            Type: "caption_coverage",
            Description: "Coverage below required percentage",
            Details: coverageErr.Error(),
        })
    }

    // Send captions text to endpoint
    langValid, langErr := sender.SendCaptions(captions, *endpoint)
    if !langValid {
        failures = append(failures, ValidationError{
            Type: "incorrect_language",
            Description: "Language validation failed",
            Details: langErr.Error(),
        })
    }

    if len(failures) > 0 {
        for _, f := range failures {
            json.NewEncoder(os.Stdout).Encode(f)
        }
        os.Exit(0)
    }

    // No validation errors
    os.Exit(0)
}
