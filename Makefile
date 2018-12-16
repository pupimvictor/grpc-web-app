

start-local-env:
    minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.10.0 \
        --extra-config=controller-manager.ClusterSigningCertFile="/var/lib/localkube/certs/ca.crt"
        --extra-config=controller-manager.ClusterSigningKeyFile="/var/lib/localkube/certs/ca.key"
        --vm-driver=virtualbox

    eval $(minikube docker-env)

clean-local-env:
    minikube delete
    rm -rf ~/.minikube