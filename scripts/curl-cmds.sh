#!/bin/bash

# These should all fail bisq daemon authentication.
curl -v -X POST    http://localhost:8080/v1/call -H "Content-Type: application/json"
echo ""
curl -v -X POST   http://localhost:8080/v1/call -H "Content-Type: application/json" -d '{"params": "random-command-no-auth-header-1 foo foo \"bar\" fubar and more \"cmd tokens\""}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Content-Type: application/json" -d '{"params": "random-command-no-auth-header-1 foo arg1 arg1-value arg2 \"bar\" arg3 \"cmd tokens\""}'
echo ""

# Add an Authorization:<pwd> header for server authentication.
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json"
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "help"}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "getversion"}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "getbalance"}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "setwalletpassword abc"}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "unlockwallet abc 30"}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "lockwallet"}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "setwalletpassword abc def"}'
echo ""
curl -v -X POST  http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "removewalletpassword def"}'
echo ""
