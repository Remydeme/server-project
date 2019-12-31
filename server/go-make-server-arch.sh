#/bin/bash 


# create directory arch 




function build_arch(){

    mkdir api;
    mkdir vendor;
    mkdir public;
    mkdir config;
    mkdir locales;
    mkdir db;

    # create directory that contains all directories related to database 
    # handler, models, and service (Structure)
    cd db;
    mkdir models;
    mkdir handlers;
    mkdir tests;
    touch service.go;

    cd ..;

   # create the makefile 
    touch Makefile;

    # IN the API directory 
    cd api;

    
    mkdir auth;
    mkdir registration;
    mkdir cmd;
    
}

if [ "$1" != "" ]
then
    mkdir ./"$1";
    cd ./"$1";
    build_arch;
fi
