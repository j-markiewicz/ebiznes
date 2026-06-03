FROM gradle:jdk21

WORKDIR /app
COPY . .

RUN ./gradlew build
CMD ["./gradlew", "run"]
