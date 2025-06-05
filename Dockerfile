FROM golang

ARG TARGETPLATFORM

WORKDIR /app
COPY . /app/

RUN go mod download

# split format linux/amd64 into linux amd64 and pass as args to go to compile
RUN TARGETPLATFORM=$TARGETPLATFORM OUTPUT=/budget-backend sh gobuild.sh
RUN rm -rf gobuild.sh
RUN chmod a+x /budget-backend

EXPOSE 80
ENV GIN_MODE=release
ENV WEBPORT=80

CMD ["/budget-backend"]