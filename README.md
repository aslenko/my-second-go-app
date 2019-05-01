#cd to src directory
cd to src

#build go (-o is output)
go build -o ../bin/main.exe .

#build docker image
docker build -t my-second-go-app .

#rebuld docker image (remove and reinstall)
docker image rm --force my-second-go-app
docker build -t my-second-go-app

#run docker container in interactive mode and map port 8088 to 8083
docker run -it -p 8083:8088 my-second-go-app

#open browser and navigate to http://localhost:8083/