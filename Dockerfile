FROM golang:alpine3.16 AS build

# coping the source project to the image
COPY . /src/app/Covid19_Vaccination_Register

WORKDIR /src/app/Covid19_Vaccination_Register
# building the base application
RUN go build -o /bin/cvregister cmd/cvregister/main.go
# building the application's managment layer
RUN go build -o /bin/cvr_management cmd/management/main.go

# new image creation only with the necesary artifacts
FROM scratch
COPY --from=build ["/bin/cvregister", "/bin/cvr_management", "/bin/"]
ENTRYPOINT ["/bin/cvregister"]
