FROM alpine:3.13

LABEL com.github.ebc-2in2crc.dango.image=https://github.com/ebc-2in2crc/dango.git

RUN apk update && \
	apk --no-cache add curl && \
	curl --location --remote-name https://github.com/ebc-2in2crc/dango/releases/download/v0.9.0/dango_linux_amd64.zip && \
	apk del curl && \
	unzip dango_linux_amd64.zip dango_linux_amd64/dango && \
	cp ./dango_linux_amd64/dango /usr/local/bin && \
	rm -rf dango_linux_amd64.zip dango_linux_amd64/dango

COPY docker-entrypoint.sh /usr/local/bin

ENTRYPOINT ["docker-entrypoint.sh"]
