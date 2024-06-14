#!/bin/bash

# Check if a server URL is provided as an argument
if [ -z "$1" ]; then
    read -p "Enter the server URL: " SERVER_URL
else
    SERVER_URL=$1
fi

# Add the line to crontab
(crontab -l ; echo "0 0 * * * curl -s $SERVER_URL | xargs -I {} date -s @{}") | crontab -

# Verify that the line has been added to crontab
crontab -l | grep curl

# Sync time now
curl -s $SERVER_URL | xargs -I {} date -s @{}