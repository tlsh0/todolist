### cmd app - main app. exec the app
### internal - api routes / handlers / repository 
### postgres db
### migration folder - migrate up and down

endpoints:
- POST /login - login and start session
- POST /logout - end session
- POST /signup - register 
- GET /tasks - list of all tasks of the user
- PATCH /tasks - rearrange the order of the tasks
- POST /tasks - add a new task
- UPDATE /task/{id} - change the status of the task
- DELETE /task/{id} - delete the task completely

ideas for the frontend:
- GET display login page 
- POST login - wait for the backend answer
- GET logout - display the logout page
- GET tasks


docker exec -it name of the docker container running
psql postgres name of the user