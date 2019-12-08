#!/bin/bash

# Env Var
GOLANG=1.12.5
KIND=0.5.1
K8S=1.11.1
ARCH=amd64

# Setup
# Move working dirctory
cd /tmp

# Setup kubectl
echo "setup kubectl"

curl -OL https://dl.k8s.io/v$K8S/kubernetes-server-linux-$ARCH.tar.gz
tar xf kubernetes-server-linux-$ARCH.tar.gz
sudo mv /tmp/kubernetes/server/bin/kubectl /usr/bin

# Setup golang
echo "setup golang"

curl -OL https://dl.google.com/go/go$GOLANG.linux-$ARCH.tar.gz
sudo tar -C /usr/local -xzf go$GOLANG.linux-$ARCH.tar.gz

mkdir $HOME/go
echo 'export GOPATH=$HOME/go' >> /home/vagrant/.bashrc
echo 'export PATH=$PATH:$GOPATH:/usr/local/go/bin' >> /home/vagrant/.bashrc

# Setup Docker
echo "setup docker"

sudo apt install -y apt-transport-https
sudo apt install -y ca-certificates
sudo apt install -y software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg >> /tmp/docker_gpg_key
sudo apt-key add /tmp/docker_gpg_key
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable edge"
sudo apt update -y
sudo apt install -y docker-ce
sudo usermod -g docker vagrant
sudo /bin/systemctl restart docker.service

# Setup Kind
echo "setup kind"

curl -Lo /tmp/kind https://github.com/kubernetes-sigs/kind/releases/download/v$KIND/kind-linux-$ARCH
chmod +x /tmp/kind
sudo mv /tmp/kind /usr/local/bin

# add new  group Docker
newgrp docker