FROM scratch

CMD ["/example"]
EXPOSE 2222

ARG REVISION=#000
ENV REVISION ${REVISION}

# Always add the binary last to maximize caching
ADD example /example
