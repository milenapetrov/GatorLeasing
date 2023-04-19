# Sprint 4

#### Backend

- Documentation located at gator-leasing-server/docs
- Unit testing
    - Lease Handler
        - Err cases for each endpoint (service, decode, validator, path paramete)
        - OK cases for GetAllLeases, PostLease, PutLease, DeleteLease, GetPaginatedLeases
    - Lease Service
        - Err cases (repository, mismatched user id, bad filter)
        - OK cases for GetAllLeases, CreateLease, EditLease, DeleteLease, GetPaginatedLeases
    - Tenant User Service
        - Err cases for new tenant user
        - OK case for Get or Create user
    - OK cases for each handler (GetAllLeases, PostLease, PutLease, DeleteLease) except GetPaginatedLeases
    - OK and repository error cases for GetAllLeases; OK cases for CreateLease, EditLease, DeleteLease 
- Lease API (add pagination endpoint)
    - Added 3 New Endpoints
    - GET /leases/{id}
        - get a lease by id
    - GET /myleases
        - get list of all leases created by current user context
    - PUT /myleases/paged
        - get paginated list of all leases created by current user context
        - similar to GET /leases/paged
        - can select number of elements per page
        - return token for next page
        - allows sorting of myleases
    - Can select number of elements per page
    - Returns token to get next page
    - Allows sorting of leases
- Faker
    - Generate fake data for testing/demoing purposes
- Fix database types
- Struct validation
    - Don't require room number for address
- Documentation
    - Use go-swagger package for documentation
    - Code annotations for all endpoints and models


#### Frontend

- Display leases on home page
    - Displays all leases posted by any user on the home page
    - Clicking any of these leases directs the user to a page of that lease’s details
- Click on lease and see full details
    - Use Get lease by ID (GET /lease/{id})
    - Pass this data between components
    - Display the properties of a lease according to the row being clicked
- Display my leases on my leases page
    - Display a grid of leases that only the user who logged in created
- Be able to edit a lease
    - Clicking on a lease from the my leases page allows the user to edit the selected lease
    - The user also haas the option to delete this lease
- Add pagination to all-leases page
    - Grids of leases implement paging
    - Can change the amount of leases held in a page
    - Click to the next page if the first is filled
- Search by name and zip code
    - Using the grid filter
    - The user can search leases on the home page by name, availability, and zip code
- Implemented styling, fonts, placement, images, and color
- Testing
    - <b>{finish}</b>


#### Documentation

- Documentation pdf: https://github.com/milenapetrov/GatorLeasing/blob/main/gator-leasing-server/docs/api-documentation.pdf
