#!/bin/bash

# Prompt the user for the server URL
read -p "Enter the server URL: " SERVER_URL

# Add the line to crontab
(crontab -l ; echo "0 0 * * * curl -s $SERVER_URL | xargs -I {} date -s @{}") | crontab -

# Verify that the line has been added to crontab
crontab -l | grep curl

# Sync time now
curl -s $SERVER_URL | xargs -I {} date -s @{}