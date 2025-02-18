failing-request:
	curl --location 'http://localhost:8080/graphql' \
    --header 'Content-Type: application/json' \
    --data-raw '{"query":"mutation CreateUser($$type: UserType!) {\n    createUser(input: {type: $$type}) {\n        __typename\n    }\n}","variables":{}}' ; echo

success-request:
	curl --location 'http://localhost:8080/graphql' \
    --header 'Content-Type: application/json' \
    --data-raw '{"query":"mutation CreateUser($$input: NewUser!) {\n    createUser(input: $$input) {\n        __typename\n    }\n}","variables":{"input":{}}}' ; echo

