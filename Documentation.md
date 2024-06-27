# PROJECT STRUCTURE

## ROLES ->

1. main.go-> ENTRY POINT OF THE APPLICATION WHERE THE SERVER IS INITIALIZED AND STARTED

2. db-> MANAGES THE DATABASE CONNECTION

3. models-> DEFINING THE DATA MODLES AND THEIR ASSOCIATED DATABASE OPERATIONS 

4. controllers-> HANDLES THE INCOMING HTTP REQUESTS , PROCESSESS THEM AND RETURNS RESPONSES

5. routes-> DEFINES THE APPLICATION ROUTES

6. middleware-> MIDDLEWARE FUNCTION THAT HANDLES TASKS 

7. utils-> CONTAINS UTILITY FUNCTIONS AND HELPERS


## SEQUENCE TO FOLLOW->

1. Initializing a Project with a go module

2. Setting up Database Connection to handle database initialization  and  table creation

3. Define Models containing different files to define structure and methods to intract with the database

4. Create Controllers for different files to handle the logic for each type of request

5. Set Up Routes by linking them to their respective controller function 

6. Implement Middleware to Protect Routes

7. Add Utilities

8. Set up main entry point by initializing servers , setting up routes and start the server 

