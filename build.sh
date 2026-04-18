#!/bin/bash

echo "Building LoRa Serial Communicator..."

# Build for host (default)
go build -o serialcom .
if [ $? -eq 0 ]; then
    echo "Host build successful: serialcom"
else
    echo "Host build failed!"
    exit 1
fi

# Build for ARM64 (Android/Linux)
echo "Building for ARM64..."
GOOS=linux GOARCH=arm64 go build -o serialcom-arm64 .
if [ $? -eq 0 ]; then
    echo "ARM64 build successful: serialcom-arm64"
else
    echo "ARM64 build failed!"
    exit 1
fi

echo ""
echo "Usage: ./serialcom (Host) or ./serialcom-arm64 (ARM64)"
echo ""
echo "Keyboard shortcuts:"
echo "  Ctrl+S: Connect to serial port"
echo "  Ctrl+D: Disconnect from serial port"
echo "  Ctrl+H: Toggle help"
echo "  Tab: Navigate between fields"
echo "  Space: Toggle preset (in preset field)"
echo "  Enter: Send message (in message field)"
echo "  q: Quit"
