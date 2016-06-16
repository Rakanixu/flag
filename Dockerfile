FROM scratch
WORKDIR /
ADD srv /srv
ENTRYPOINT [ "/flag" ]
