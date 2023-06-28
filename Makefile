version = "0.5"
org = "devopsrob"

build:
	docker build -t $(org)/fakevault:$(version) .
tag:
	docker tag $(org)/fakevault:$(version) devopsrob/fakevault:$(version)
push:
	docker push $(org)/fakevault:$(version)
deploy: build tag push