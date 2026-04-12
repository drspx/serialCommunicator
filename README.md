# LoRa Serial Communicator

A TUI application built with Go and BubbleTea for communicating with LoRa modules via serial port (FTDI adapter).

## Features

- **Serial Communication**: Connect to LoRa modules via serial port (FTDI adapter)
- **AT Command Presets**: Configure LoRa parameters with predefined presets
  - Short Range & Fast: SF7, 125kHz, CR5, 14dBm
  - Long Range & Slow: SF12, 125kHz, CR8, 20dBm
- **Chat Interface**: Send and receive messages with real-time chat display
- **Complete Message Sending**: Messages sent as whole packets for optimal LoRa performance
- **TUI Navigation**: Easy navigation with keyboard shortcuts

## Installation

1. Install Go (version 1.24+)
2. Clone/download this repository
3. Build the application:

```bash
chmod +x build.sh run.sh
./build.sh
```

## Usage

Run the application:
```bash
./run.sh
```

Or directly:
```bash
go run .
```

## Keyboard Shortcuts

- **Tab/Shift+Tab**: Cycle through input fields
- **Space**: Toggle preset (when preset field is focused)
- **Enter**: Send message (when message field is focused)
- **Backspace**: Delete character (when message field is focused)
- **Ctrl+S**: Connect to serial port
- **Ctrl+D**: Disconnect from serial port
- **Ctrl+H**: Toggle help
- **q**: Quit application

## Default Settings

- Serial Port: `/dev/ttyUSB0` (common for FTDI adapters)
- Baud Rate: `9600`

## LoRa AT Commands Sent

When switching presets, the application sends these AT commands:

### Short Range & Fast
```
AT+SF=7
AT+BW=125000
AT+CR=5
AT+TP=14
```

### Long Range & Slow
```
AT+SF=12
AT+BW=125000
AT+CR=8
AT+TP=20
```

## Requirements

- Go 1.24 or higher
- LoRa module configured to accept AT commands
- FTDI USB-to-serial adapter
- Serial port access (typically `/dev/ttyUSB0` on Linux)

## Project Structure

- `main.go` - Application entry point
- `model.go` - BubbleTea model and UI logic
- `serial.go` - Serial communication implementation
- `lora_presets.go` - LoRa preset configurations
- `build.sh` - Build script
- `run.sh` - Run script

## Notes

- The application sends messages as complete packets for optimal LoRa performance
- Make sure your LoRa module is configured to accept AT commands
- Adjust serial port in the configuration if not using `/dev/ttyUSB0`
- Baud rate is fixed at 9600 - configure your LoRa module accordingly# serialCommunicator
