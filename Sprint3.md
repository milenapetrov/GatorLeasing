# Sprint 3

#### Backend

- Documentation located at gator-leasing-server/docs
- Unit testing
    - OK cases for each handler (GetAllLeases, PostLease, PutLease, DeleteLease) except GetPaginatedLeases
    - OK and repository error cases for GetAllLeases; OK cases for CreateLease, EditLease, DeleteLease 
- Lease API (add pagination endpoint)
    - Can select number of elements per page
    - Returns token to get next page
    - Allows sorting of leases
- Struct validation
    - Validate lease name 3 character or more, max 20 characters
    - Validate lease rent, required, min 0.01
    - Validate start date before end data
    - Validate lease ID on Edit and Delete
    - Require all data for address object


#### Frontend

- Cypress Unit Testing:
    - Mounts each component, imports lease and user authentication services for components that send requests to either service 
    - Get commands used for simulating input and validating existence of types
    - Button clicks simulated
    - Home: Checks if it can mount; if search input is string; if search button can click; if post button can click
    - Profile: Checks if it can mount; if first name, last name, phone number, email, & address input is string; if update button can click
    - App: checks if each button in navigation bar can be accessed and clicked
- Navigation between pages:
    - Top navigation bar separates existing functionalities between different browsers. 
    - Directs to home page when title is clicked
    - Login/logout
    - My leases displays post information
    - Home page post a lease direct to post creation
    - Working on fixing bugs that occur when making pages only accessible if the user is logged in
    - Added picture to home page, working on how the website will look
- Post a sublease: 
    - A post request is sent to the backend when a post is created
    - Can only create a post if logged in
    - From input for adding all values requiored for a least
    - Input requesting the cost of parking only appears if user indicates that their lease includes parking
- Profile: 
    - Form input created
    - Includes all user characteristics like name, email, & phone number 
-Display leases:
    -Implements a grid to display lease information and apply paging from the backend
    -Syncs rows to properties in a lease and displays based on the field defined in the column
    -Fields manipulated to represent more readable data
    -Will eventually be able to update each post after clicking the button in the grid

#### Documentation

- Documentation pdf: https://github.com/milenapetrov/GatorLeasing/blob/main/gator-leasing-server/docs/api-documentation.pdf
