.PHONY: use-minikube
use-minikube:
	kubectl config use-context minikube

.PHONY: deploy-pub
deploy-pub:
	skaffold dev -f skaffold/pub.yaml

.PHONY: delete-pub
delete-pub:
	skaffold delete -f skaffold/pub.yaml

.PHONY: deploy-sub
deploy-sub:
	skaffold dev -f skaffold/sub.yaml

.PHONY: delete-sub
delete-sub:
	skaffold delete -f skaffold/sub.yaml

.PHONY: delete
delete: delete-sub delete-pub

.PHONY: reload
reload: delete deploy
