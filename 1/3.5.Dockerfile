FROM ubuntu:24.04

RUN apt update && apt install openjdk-8-jre openjdk-8-jdk kotlin -y
