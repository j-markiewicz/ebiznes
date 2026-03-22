FROM ubuntu:24.04

RUN apt update && apt install openjdk-8-jre openjdk-8-jdk kotlin -y

RUN apt update && apt install default-jdk curl -y
ARG GRADLE_VERSION=9.4.1
RUN curl -L https://services.gradle.org/distributions/gradle-${GRADLE_VERSION}-bin.zip -o ./gradle-${GRADLE_VERSION}-bin.zip
RUN mkdir /opt/gradle && \
	unzip -d /opt/gradle ./gradle-${GRADLE_VERSION}-bin.zip && \
	ls /opt/gradle/gradle-${GRADLE_VERSION}
ENV PATH=$PATH:/opt/gradle/gradle-9.4.1/bin

WORKDIR /workdir
RUN gradle && gradle init --type kotlin-application --dsl groovy
COPY ./build.gradle ./app/build.gradle
COPY ./HelloWorld.kt ./app/src/main/kotlin/org/example/App.kt
RUN rm ./app/src/test/kotlin/org/example/AppTest.kt
RUN gradle build

CMD [ "java", "-jar", "/workdir/app/build/libs/app.jar", "-class", "AppKt" ]
