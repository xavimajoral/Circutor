
hello:
	echo "hello"

signup:
	curl \
      -X POST \
      http://localhost:1323/signup \
      -H "Content-Type: application/json" \
      -d '{"email":"xorduna@circutor.com","password":"123456"}'

login:
	curl \
	  -X POST \
	  http://localhost:1323/login \
	  -H "Content-Type: application/json" \
	  -d '{"email":"xorduna@circutor.com","password":"123456"}' 2>/dev/null | jq '.Token' -r > token.txt


sites:
	$(eval TOKEN := $(shell cat token.txt))
	curl \
      http://localhost:1323/sites \
      -H "Authorization: Bearer $(TOKEN)"


add-site:
	$(eval TOKEN := $(shell cat token.txt))
	curl \
		-X POST \
		http://localhost:1323/sites \
		-H "Content-Type: application/json" \
		-d '{"LocationId":"market-hall"}' \
      	-H "Authorization: Bearer $(TOKEN)"

docs:
	~/go/bin/swag init -g server.go --parseVendor --parseDependency --parseInternal