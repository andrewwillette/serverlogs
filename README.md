# serverlogs
Quickly access logs hosted on a remote server using the SSH protocol.

## Example Usage
```
serverlogs -privatekey="/Users/jimbob/.ssh/id_rsa" -ip="2.245.6.188" -user="ubuntu" -loglocation="/home/ubuntu/server.log"
```

## Configuration
- environment variables: `SERVER_LOG_PRIVATE_KEY`, `SERVER_LOG_IP`, `SERVER_LOG_USER`, `SERVER_LOG_LOG_LOCATION`
- CLI arguments, see `serverlogs -h` for more info
