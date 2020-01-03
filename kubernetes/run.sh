echo "Brew install minikube"
#installation de einikube 
brew install minikube 
#installation de kubectl
echo "Brew install kubectl"
brew install kubectl 
#installation de virtualBox 
echo "Brew install virtualboxe"
brew install virtualbox 
#script de création du cluster 
minikube delete 
minikube start
kubectl apply -f deployement.yml
kubectl apply -f services.yml

#cehck if the server is running using 
curl http://192.168.99.101:30080/ping/
