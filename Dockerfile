FROM scratch

COPY cmd/cmd .

EXPOSE 8094

ENTRYPOINT [ "./cmd" ]