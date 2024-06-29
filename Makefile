.PHONY: 
	running-user-api-gateway \
	running-account-service \
	watching-proto \
	running-in-background-watching-proto \
	running-in-background-user-api-gateway \
	running-in-background-account-service \
	all-in-background \

running-watching-proto:
	cd shared && make watching-proto

running-user-api-gateway:
	cd user-api-gateway && make running

running-account-service:
	cd accounts-svc && make running

running-in-background-watching-proto:
	cd shared && make watching-proto-in-background

running-in-background-user-api-gateway:
	cd user-api-gateway && make running-in-background

running-in-background-account-service:
	cd accounts-svc && make running-in-background

all-in-background: 
	make running-in-background-watching-proto 
	make running-in-background-user-api-gateway 
	make running-in-background-account-service