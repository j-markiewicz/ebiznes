FROM ubuntu

# Current Ubuntu's apt no longer has python 3.10, install it from source instead
RUN apt update && apt install git gcc make zlib1g zlib1g-dev -y
RUN git clone --depth 1 --branch v3.10.20 https://github.com/python/cpython.git
RUN cd ./cpython && \
	./configure && \
	make && \
	make install

ENTRYPOINT [ "python3" ]
