FROM alpine
ADD myTest /myTest
ENTRYPOINT [ "/myTest" ]
