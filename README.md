<<<<<<< HEAD

# Caption Validator (Beginner-Friendly Guide)

This project validates caption files (WebVTT or SRT) for coverage and language.

## Why This Exists
Captions help make videos accessible and searchable. This tool ensures captions:
1. Cover enough of the video timeline.
2. Are in the correct language (English - en-US).

## How It Works (Human Logic)
- **Step 1**: Read the caption file and figure out its format (WebVTT or SRT).
- **Step 2**: Calculate how much of the video timeline has captions.
- **Step 3**: Send all caption text to a web service to check the language.
- **Step 4**: Print any problems as JSON. If everything is fine, print nothing.

## Requirements
- Docker installed OR Go installed.

## Build and Run with Docker
```bash
# Build the Docker image
docker build -t caption-validator .

# Run the validator
docker run --rm caption-validator   --t_start 0 --t_end 120 --coverage 80   --endpoint http://example.com captions.srt
```

## Run Locally with Go
```bash
# Initialize Go module
go mod init caption-validator

# Tidy dependencies
go mod tidy

# Run the program
go run cmd/main.go --t_start 0 --t_end 120 --coverage 80 --endpoint http://example.com captions.srt
```

## Exit Codes
- `0`: Validation failures or success.
- `1`: Program error (e.g., wrong file type).

## TODOs
- Implement full parsing for WebVTT and SRT.
- Add unit tests in `tests/`.
- Improve error messages and logging.
=======
# caption-validator
>>>>>>> a753cf84e5592df9f4271c4aaf3dde3255f6dcf9
