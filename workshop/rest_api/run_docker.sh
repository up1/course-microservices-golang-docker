docker image build -t demo:0.0.1
docker container run -d -name api -p 9999:8080 demo:0.0.1