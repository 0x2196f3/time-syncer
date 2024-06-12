# Time-Syncer

#### a time synchronizer over http for linux

it will respond any request with current time
# Deploy Server
```bash
docker run -d --name=time-syncer -p 30080:30080 docker.io/0x2196f3/time-syncer
```

# Install Client
- download the project
- run install.sh
```bash
sudo bash ./install.sh
```
# Sync time for a single time
```bash
curl -s http://192.168.2.2:10001 | xargs -I {} sudo date -s @{}
```
- enter server ip and port

# Environment Variable
| Environment Variable | Usage |
| --- | --- |
| `PORT` | port |
