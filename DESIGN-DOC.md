# The proposed UX of the app
The features will align closely to this section of the assignment:
* https://github.com/gravitational/careers/blob/main/challenges/fullstack/challenge.md#requirements---level-4

## Wireframes
* [Login.pdf](https://github.com/goteleport-interview/int-fs-nick-2/files/11729812/Login.pdf)
  * Shows the login page
* [Root.pdf](https://github.com/goteleport-interview/int-fs-nick-2/files/11729814/Root.pdf)
  * Shows the table with the root directory contents
* [Root-Reverse.pdf](https://github.com/goteleport-interview/int-fs-nick-2/files/11729813/Root-Reverse.pdf)
  * Shows the table with the root directory contents, but sorted in descending order by name
* [Directory1.pdf](https://github.com/goteleport-interview/int-fs-nick-2/files/11729809/Directory1.pdf)
  * Shows the table if clicked into Directory1 from the root directory
* [Sub-Directory1.pdf](https://github.com/goteleport-interview/int-fs-nick-2/files/11729816/Sub-Directory1.pdf)
  * Shows the table if clicked into Sub-Directory1 from Directory1
* [Sub-Directory1-Filter.pdf](https://github.com/goteleport-interview/int-fs-nick-2/files/11729815/Sub-Directory1-Filter.pdf)
  * Shows the ability to filter table Sub-Directory1 by name

# The proposed API
## RESTful API endpoints

### Get Directory Contents - GET: /directory
Returns a JSON payload with the contents of a provided directory

### Login - POST: /login
Authenticates a user against the in-memory storage solution

### Logout - POST: /logout
Authenticates a user against the in-memory storage solution

### Error codes
* 200 -- successful request
* 401 -- unauthorized access
* 404 -- resource unavailable

# URL structure
## Structure
* https://{root domain}/directory/{directory-path}{query-params}

## Root domain
* Required
* Configurable based on if it is a local instance or if it’s a shared dev, test, or prod environment

## URL Param: directory-path
* Optional
* Purpose: Use to specific the specific path for a directory
* Additional details:
** If a correct value is provided, a new page will be rendered to show the directory’s contents
** If a value is provided but the API response returns no contents, the user will remain on the same page and a warning alert message will be shown in the UI
** If a value is not provided, the user will be taken to the root directory

## URL Param: Page
* Optional
* Purpose: Use this to refer to page if the contents are paginated

## URL Param: Size
* Optional
* Purpose: Use to refer to the number of items that should be shown on the page when paginating the table

## URL Param: Sort
* Optional
* Purpose: Use to refer to the type of sort that will be used client-side

## URL Param: Filter
* Optional
* Purpose: Use to refer to the type of filter that will be used client-side

# Security
## Authentication
I’ll be implementing JWT based authentication
It will be generated and sent to the client upon calling the login endpoint
The client will return the JWT with each HTTP request to be validated to grant access to resources

## TLS setup
First, I will enable my Go HTTP server to work with TLS. Then I will generate a self-signed TLS certificate using the OpenSSL command-line tool by running a command to generate a new 4096-bit RSA key pair and a self-signed certificate that is valid for 365 days. Finally, I will redirect all server traffic from HTTP to HTTPS to ensure secure connections are always used.
## Protection against common web security vulnerabilities

Implement HTTPS: 
* Secure your communication with HTTPS (HTTP over SSL/TLS). This protects the integrity and confidentiality of data between the user's computer and your site.

Cross-Site Request Forgery (CSRF) Protection: 
* Use anti-CSRF tokens and check the HTTP Referer and Origin headers to protect your site against CSRF attacks.

Use Secure Password Policies: 
* Enforce strong password requirements, and always hash and salt passwords using a strong cryptographic algorithm.

Implement Rate Limiting: 
* This can help prevent attacks such as brute-force and DDoS.

Implement Least Privilege Principle: 
* Only give the permissions that are absolutely necessary for a user or a process.

# Implementation details where appropriate
## Session Management
I chose to use JWT for authentication instead of cookie-based sessions. I tend to prefer this over cookie-based auth given it’s stateless, enables cross-domain compatibility, and enables fine-grained access control.

## Pull Requests
1. Backend API with Go
2. Frontend - Directory Browsing and Navigation
3. Frontend - User Authentication and Session Management
4. Final Adjustments and Bug Fixes
