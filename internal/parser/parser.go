
package parser

import (
    "errors"
    "regexp"
    "strconv"
    "strings"
)

type Caption struct {
    Start float64
    End   float64
    Text  string
}

// ParseCaptions detects format and parses captions
func ParseCaptions(filePath string) ([]Caption, string, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, "", err
    }
    content := string(data)

    if strings.HasPrefix(content, "WEBVTT") {
        return parseWebVTT(content), "webvtt", nil
    } else if strings.Contains(content, "-->") {
        return parseSRT(content), "srt", nil
    }
    return nil, "", errors.New("unsupported file format")
}

// parseWebVTT parses WebVTT captions
func parseWebVTT(content string) []Caption {
    lines := strings.Split(content, "
")
    var captions []Caption
    timePattern := regexp.MustCompile(`(\d{2}:\d{2}:\d{2}\.\d{3}) --> (\d{2}:\d{2}:\d{2}\.\d{3})`)
    var currentText []string
    var start, end float64

    for _, line := range lines {
        if timePattern.MatchString(line) {
            if len(currentText) > 0 {
                captions = append(captions, Caption{Start: start, End: end, Text: strings.Join(currentText, " ")})
                currentText = []string{}
            }
            matches := timePattern.FindStringSubmatch(line)
            start = timeToSeconds(matches[1])
            end = timeToSeconds(matches[2])
        } else if line != "" && !strings.HasPrefix(line, "WEBVTT") {
            currentText = append(currentText, line)
        }
    }
    if len(currentText) > 0 {
        captions = append(captions, Caption{Start: start, End: end, Text: strings.Join(currentText, " ")})
    }
    return captions
}

// parseSRT parses SRT captions
func parseSRT(content string) []Caption {
    blocks := strings.Split(content, "

")
    var captions []Caption
    timePattern := regexp.MustCompile(`(\d{2}:\d{2}:\d{2},\d{3}) --> (\d{2}:\d{2}:\d{2},\d{3})`)

    for _, block := range blocks {
        lines := strings.Split(block, "
")
        if len(lines) < 2 {
            continue
        }
        if timePattern.MatchString(lines[1]) {
            matches := timePattern.FindStringSubmatch(lines[1])
            start := timeToSeconds(strings.Replace(matches[1], ",", ".", 1))
            end := timeToSeconds(strings.Replace(matches[2], ",", ".", 1))
            text := strings.Join(lines[2:], " ")
            captions = append(captions, Caption{Start: start, End: end, Text: text})
        }
    }
    return captions
}

// timeToSeconds converts HH:MM:SS.mmm to seconds
func timeToSeconds(ts string) float64 {
    parts := strings.Split(ts, ":")
    h, _ := strconv.Atoi(parts[0])
    m, _ := strconv.Atoi(parts[1])
    secParts := strings.Split(parts[2], ".")
    s, _ := strconv.Atoi(secParts[0])
    ms, _ := strconv.Atoi(secParts[1])
    return float64(h*3600+m*60+s) + float64(ms)/1000.0
}
