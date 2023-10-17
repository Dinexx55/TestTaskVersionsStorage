
# Versions-Storage Project

## Start
You can simply start the application using "docker-compose up". 
You don't need to set any envs, config files are committed on purpose,
so you can start the app easily ;D.
This repo was created specifically for easy start. 
You can look through the history of development in the following repos:
https://github.com/Dinexx55/Gateway_Service
https://github.com/Dinexx55/Storage_Service

## Usage

Mind the 8081 port!

- `POST /auth/login`

body:
{
  "login": "example_user",
  "password": "example_password"
}

available users in mockRepo:

{
  "login": "user1",
  "password": "password1"
}
{
  "login": "user2",
  "password": "password2"
}
{
  "login": "user3",
  "password": "password3"
}

- `POST /storage/store`

body:
{
    "name": "Example Store",
    "address": "Karaganda, Lenina, 143",                  
    "owner_name": "John, Doe",                 
    "opening_time": "2013-12-12 12:33:56",                  
    "closing_time": "2013-12-12 12:33:56"                   
}

address format:        "city, street, house"
owner_name format:     "surname, name"
opening_time format:   "YYYY-MM-DD HH:MM:SS"
closing_time format:   "YYYY-MM-DD HH:MM:SS"

- `POST /storage/store/:id/version`

body:
{
    "owner_name": "John, Doe",                          
    "opening_time": "2013-12-12 12:33:56",             
    "closing_time": "2013-12-12 12:33:56"                
}

owner_name format:     "surname, name"
opening_time format:   "YYYY-MM-DD HH:MM:SS"
closing_time format:   "YYYY-MM-DD HH:MM:SS"

- `DELETE /storage/store/:id`
- `DELETE /storage/store/:id/version/:versionId`
- `GET /storage/store/:id`
- `GET /storage/store/:id/history`
- `GET /storage/store/:id/version/:versionId`
    
