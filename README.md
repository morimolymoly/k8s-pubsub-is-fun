# Google Cloud Platform PubSub Tutorial

## k8s
### Requirements
* minikube
* [gcp-emulator on k8s](https://github.com/morimolymoly/gcp-emulators)
### Usage
First of all, you need to launch minikube and gcp-emulator!!
```
# use minikube for kubectl
make use-minikube
# deploy pub
make deploy-pub
# deploy sub
make deploy-sub
```

## docker-compose
### Requirements
Before start up this project, you need to launch local GCP emulators.  
Check [this one](https://github.com/morimolymoly/gcp-emulators) out!
### Usage
```
dep ensure
docker-compose build
# publish messages
docker-compose up pub
# subscribe messages
docker-compose up sub
```

# LICENSE
MIT
