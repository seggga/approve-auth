# approve-auth
auth service, a part of task approval microservice system

auth service implements REST API for users to login/logout using JWT.
Service checks validity of JWT tokens passed from other services (task and analytics) via gRPC.  

## User interaciton
User authenticates with basic-authentication (path /login), predefined logins and password-hashes located in config-file. Service replies with jwt-cookie with user's name and a lifetime of 1 minute and a refresh token with user's name valid during 1 hour. After 1 hour user is to be unauthotized.
Path /logout sends empty pair of cookies access and refresh.
Service uses json. 