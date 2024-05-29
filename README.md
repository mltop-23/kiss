![Alt text](/foto/structure.png?raw=true "Optional Title")
![Alt text](/foto/BD.png?raw=true "Optional Title")



Contains the main API server code. main.go: Initializes the server, sets up routes, and starts the API. handlers: Houses HTTP request handlers for specific functionalities. auth.go: Handles authentication and authorization requests. users.go: Handles user management requests (create, read, update, delete). dishes.go: Handles dish management requests (create, read, update, delete). 2. config:

Stores configuration settings for the application. config.go: Defines and loads configuration values for database connection, authentication, etc. 3. database:

Encapsulates database interaction logic. models.go: Defines data structures for entities that represent database tables. repository.go: Implements data access methods for interacting with the database (CRUD operations). 4. main.go (optional):

Provides an entry point for testing purposes. You can use this file to test your handlers and database interactions separately. 5. README.md:

Contains project documentation and instructions. Provide a brief overview of the project, setup instructions, and usage guidelines.