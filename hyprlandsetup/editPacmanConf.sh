#!/bin/bash

# Path to pacman.conf file; adjust as necessary for your system.
CONF_FILE_PATH="/etc/pacman.conf"

if [ ! -f "$CONF_FILE_PATH" ]; then
    echo "The configuration file was not found."
    exit 1
fi

# Use sed to edit a specific line in the file. 
# Adjust the pattern and replacement string as needed for your editing purposes.
sed -i 's/#Color/Color/g' $CONF_FILE_PATH
sed -i 's/#ParallelDownloads = 5/ParallelDownloads = 20/g' $CONF_FILE_PATH

echo "pacman.conf edited successfully"