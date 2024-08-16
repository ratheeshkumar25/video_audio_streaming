# Simple Audio/Video Streaming Server

## Overview
This project is a basic audio/video streaming server built using Golang. It serves files from a specified directory over HTTP, making it easy to stream content like audio files.

## Features
- **File Serving:** Serves files from the `Files` directory.
- **CORS Support:** Adds a CORS header to allow cross-origin requests.
- **Simple Setup:** Minimal setup required to get the server running.

## Technologies Used
- **Programming Language:** Golang
- **Standard Library:** `net/http` for HTTP server functionalities

## Getting Started

### Prerequisites
- Go 1.20 or later
- A directory named `Files` in the project root containing the audio/video files to be served

### Running the Server
1. **Clone the repository:**
   ```bash
   git clone https://github.com/ratheeshkumar25/video_audio_streaming.git
   cd video_audio_streaming
