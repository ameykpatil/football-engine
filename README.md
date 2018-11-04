# Football-Engine
Application to serve the details of footballers and teams. 

## Problem Statement
Given API endpoint:  
`https://vintagemonster.onefootball.com/api/teams/en/%team_id%.json`   
(`%team_id%` must be replaced with an unsigned integer).  

Using the API endpoint, find the following teams by name:  
`Germany, England, France, Spain, Manchester Utd, Arsenal, Chelsea, Barcelona, Real Madrid, FC Bayern Munich`  

Extract all the players from the given teams and render to stdout the information about players alphabetically ordered by name. Each player entry should contain the following information:   
`full name; age; list of teams`  

Output Example:  
```
Alexander Mustermann; 25; France, Manchester Utd
Brad Exampleman; 30; Arsenal
```

## Instructions
Clone this repository.  
Go to the directory where repository is cloned.   
(`$GOPATH/src/github.com/ameykpatil/football-engine`)

## Run Tests
`/bin/sh ./check.sh`  
`check.sh` is a file created which runs multiple checks such as `fmt`, `vet`, `lint` & `test`  

## Run Service
`go install`   
Check if the service is running by hitting `http://localhost:4000/ping`  
You should get `message` as `pong`  

## Print & Get players of the given Teams
Hit following url in browser  
`http://localhost:4000/players/fetch`  

## Notes
This problem can be solved without creating http server as well but from the perspective of extending it to retrive other information http or web server looked like a better approach.  
Application make use of concurrency to fetch `Team` information from external API in parallel.  
A constant `maxConcurrency` controls the concurrency factor or number of Go routines. We can change it to increase or decrease the concurrecny.  
Go routines are synchronized with main routine with the help of `WaitGroup`.  
To avoid concurrent access to map storing `Player` details, `Mutex` has been used.      
Tests are written for utilities.  
