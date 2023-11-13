FROM scratch
COPY hoomy /usr/bin/hoomy
ENTRYPOINT [ "/usr/bin/hoomy" ]