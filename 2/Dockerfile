FROM sbtscala/scala-sbt:eclipse-temurin-25.0.1_8_1.12.8_3.8.3

RUN curl -sSL https://ngrok-agent.s3.amazonaws.com/ngrok.asc \
	| tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null \
	&& echo "deb https://ngrok-agent.s3.amazonaws.com bookworm main" \
	| tee /etc/apt/sources.list.d/ngrok.list \
	&& apt update \
	&& apt install ngrok
RUN ngrok config add-authtoken $NGROK_TOKEN

COPY . .

EXPOSE 8080
CMD ["/bin/sh", "./run.sh"]
