# SimpleCMServer

## Auth service

Implements OAuth2 

1. User opens Client, requests for resources. Client redirects him/her to login.
2. After login it received Authorization Code.
3. Client uses the Code to request for Access Token. 
4. Access token can be used to access resources in server.
