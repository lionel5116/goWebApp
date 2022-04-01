
#refernce : https://docs.docker.com/engine/reference/builder/#:~:text=A%20Dockerfile%20is%20a%20text,command%2Dline%20instructions%20in%20succession.

#The FROM instruction initializes a new build stage and sets the Base Image for subsequent instructions
FROM golang:1.17

#The WORKDIR instruction sets the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions
#docker copies over the main.go file to this directory
WORKDIR /go/src/app

#The COPY instruction copies new files or directories from <src> and adds them to the filesystem of the container at the path <dest>.
COPY . .

#The RUN instruction will execute any commands in a new layer on top of the current image and commit the results.
#this command says make a director on the root /cmd
RUN mkdir /cmd

#this command says run the go build command, -o output the build file (the binary is called app) to the /cmd directory
#docker / go looks at the /go/src/app/ directory (the WORKDIR) for the main.go file to build
RUN go build -o /cmd/app /go/src/app/main.go

#An ENTRYPOINT allows you to configure a container that will run as an executable.
#the entry point is the command that docker uses to run the app, below the 'app' is the actual binary (build file)
#i:e ENTRYPOINT["Node","server.js"]    #this says, run node server.js
ENTRYPOINT ["/cmd/app"]