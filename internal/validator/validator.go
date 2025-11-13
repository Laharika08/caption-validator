
package validator

import (
    "errors"
    "solution/internal/parser"
)

func ValidateCoverage(captions []parser.Caption, tStart, tEnd, requiredCoverage float64) (bool, error) {
    totalDuration := tEnd - tStart
    if totalDuration <= 0 {
        return false, errors.New("invalid time range")
    }

    var covered float64
    for _, c := range captions {
        if c.End > tStart && c.Start < tEnd {
            start := max(c.Start, tStart)
            end := min(c.End, tEnd)
            covered += end - start
        }
    }

    coveragePercent := (covered / totalDuration) * 100
    if coveragePercent < requiredCoverage {
        return false, errors.New("coverage is " + fmt.Sprintf("%.2f%%", coveragePercent))
    }
    return true, nil
}

func max(a, b float64) float64 { if a > b { return a }; return b }
func min(a, b float64) float64 { if a < b { return a }; return b }
