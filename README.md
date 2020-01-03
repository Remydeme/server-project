# Microservices Project



## Dependency 

* Minikube 
* VirtualBox


## USE 
in kubernetes directory 
run : 

- chmod 755 run.sh && ./run.sh # this command will install all the dependency 

- in the terminal : minikube ip  # get the cluster ip adress on your machine 

- go to your browser and run ip:30080/ping

if you get pong as answer this mean that your cluster is running.

## Launsh 

Afin de lancer le serveur il faut :

* minikube start -p sessions // lance votre hypervisor 


* Placer vous dans le dossier kubernetes et lancer run.sh. Cette commande va exécuter un script qui va lancer notre cluster. 

## Config file 




#### Deployement.yml 




apiVersion: apps/v1
kind: Deployment
metadata:
    name: servicego

spec:
 replicas: 2
 selector:
  matchLabels:
   app: servicego
 template:
  metadata:
   labels:
    app: servicego
  spec:
   containers:
   - name: servicego
     image: falconr/user-microservice




Non avons configurer notre cluster de tel sorte à ce que l'on ai deux noeuds contenant chacun un serviceGo. Cela augmente la résilience dans le cas on l'un de nos service 
Go venait à crasher. Notre service Go est identifié par le label "servicego".

Notre image falconr/user-microservice correspond à une image que nous avons compilé est push sur notre server Docker.

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
      metadata:
        labels:
          app: nginx
      spec:
        containers:
        - name: nginx
          image: falconr/nginx-http2


Notre micro-service Nginx dont le rôle est de servir de reverse proxy est en un replicas. Nous utilisons l'image falconr/nginx-http2. 

---

apiVersion: apps/v1
kind: Deployment
metadata:
 name: mangodb 
spec:
 replicas: 1
 selector:
  matchLabels:
   app: mongodb
 template:
   metadata:
    labels:
     app: mongodb
   spec:
    containers:
    - name: mongodb
      image: mongo:3.6-jessie
     # we want to persist the database storage on the disk by mounting the /data/db file on our local disk 
      volumeMounts:
      - name: mongodb-persistent-storage
        mountPath: /data/db
    volumes: # the volume describtion on our local disk 
    - name: mongodb-persistent-storage
    # pointing on the describtion of the volume 
      persistentVolumeClaim:
         claimName: mongo-pv


Il y a également un seul noeud de notre serveur mongoDB. On utilise une image public mongoDB. Le volume mongoDB est configuré pour Amazon web service et non pour fonctionner en local.  







## service 

apiVersion: v1
kind: Service
metadata:
    name: servicego
spec:
    #This will define wich pod are going to be serve by this service
    #A service is an network endpoint for other service or for the outside world 
    
    # selector as the same value as the label of the pod that we  
    selector:
      app: servicego
      
    ports:
    # specify the type of protocol we are using 
      - name: http
        port: 8080
    type: ClusterIP 


Nous exposons le port 8080 de notre microservice Go. Le serveur Go communique uniquement avec les noeuds interne du cluster. 



--- 

apiVersion: v1
kind: Service
metadata:
    name: nginx
    labels:
      app: nginx


spec:

    selector:
      app: nginx
    
    ports:
    - name: http
      port: 80
      nodePort: 30080   # expose port 3008à to the outside world 

    type: NodePort
    
    
Nous exposons le port 30080 de notre reverseProxy. Il s'agit d'un nodePort car nous avons besoins qu'il communique avec l'extérieur du cluster. Le service devient 
accessible à l'adresse IP 30080 depuis l'extérieur du cluster.  

--- 

apiVersion: v1
kind: Service 
metadata:
    name: mongodb
    labels:
      app: mongodb

spec:

    selector:
      app: mongodb
    
    ports:
    - name: htpp
      port: 27017

    type: ClusterIP

Nous exposons le port 27017 de mongodb qui correspond au port par défaut par lequel on communique avec le serveur mongoDB













