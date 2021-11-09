FROM "gitlab.ar.bsch:4567/santander-tecnologia/docker-base-images/golang:v14"

COPY rates .

EXPOSE 8080
EXPOSE 8081

CMD [ "./rates" ]
