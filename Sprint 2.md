# Sprint 2

#### Backend

- Documentation located at gator-leasing-server/docs
- Unit testing
- User authentication
    - Create user table
    - User endpoints
        - Create a user, etc.
- Add more data to lease
    - Address: string
    - Square footage: int
    - number of Bedrooms: int
    - number of bathrooms: float
    - Posted owner: user
    - images???

#### Frontend

- Cypress Unit Testing:
    - Mounts each component, imports lease and user authentication services for components that send requests to either service. 
    - Get commands used for simulating input and validating existence of types
    - Button clicks simulated
- Login page:
    - User authentication service connected
    - Login button changes to display logout once logged in, vice versa
    - Can only access “my leases” once logged in, trying to access this page prompts the user to log in. 
- Navigation between pages:
    - Top navigation bar separates existing functionalities between different browsers. 
    - Directs to home page when title is clicked
    - Login/logout
- Post a sublease: 
    - A post request is sent to the backend when a post is created. 
    - Refreshing the page shows the new leases added to the database.
    - Once the website design is established, we will make the addition of a post appear without having to refresh the page. 
- Edit a sublease:
    - Will implement the put and delete requests once we decide on a final design for the web page; deciding how the user will edit their post.











