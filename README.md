# Google Cloud Platform PubSub Tutorial
## Requirements
Before start up this project, you need to launch local GCP emulators.  
Check [this one](https://github.com/morimolymoly/gcp-emulators) out!
## Usage
```
dep ensure
docker-compose build
# publish messages
docker-compose up pub
# subscribe messages
docker-compose up sub
```
