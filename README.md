
# User Service

User service for registration/login of user. It allows to block/unblock other users and update their profiles.

# Features
- User can do registration using email/phone number/google account with otp.
- For login also otp sent to registered details.
- User can update their profiles.
- User can block/unblock another users.

# Tech Stack 
- GO 1.21
- CockroachDB
- Dbmate
- JWT (json web token)
- Google Authentication (OAuth 2.0)
- Twilio 
- SMTP

## Run Locally

Prerequisites you need to set up on your local computer:

- [Golang](https://go.dev/doc/install)
- [Cockroach](https://www.cockroachlabs.com/docs/releases/)
- [Dbmate](https://github.com/amacneil/dbmate#installation)

1. Clone the project

```bash
  git clone https://github.com/KaranLathiya/user.git
  cd user
```

2. Copy the .env.example file to new .config/.env file and set env variables in .env:

```bash
  cp .env.example .config/.env
```

3. Create `.env` file in current directory and update below configurations:
   - Add Cockroach database URL in `DATABASE_URL` variable.
4. Run `dbmate migrate` to migrate database schema.
5. Run `go run cmd/main.go` to run the programme.

# Routing

## For Signup 

To first time signup for new user using email/phone number (get otp)  --POST

    http://localhost:8000/auth/signup

## For Login

To login for user using email/phone number (get otp) --POST

    http://localhost:8000/auth/login

To login for user using google account(code) (get otp) --GET

    http://localhost:8000/auth/google/login?code=

## For Authorization code of google account 

To get code using google account --GET

    http://localhost:8000/auth/google

## For otp verification 

To verify otp for signup/login --GET

    http://localhost:8000/otp/verify

## For UserProfile

To update privacy for user --PUT

    http://localhost:8000/user/profile/privacy

To update basic details for user --PUT

    http://localhost:8000/user/profile/basic
 
To fetch profile details for user --GET

    http://localhost:8000/user/profile/

## For get other UserDetails

To get other UserDetails by userID  --GET

    http://localhost:8000/users/{id}/id
    
To get other UserDetails by username  --GET

    http://localhost:8080/users/{username}/username

## For get all other UserDetails 

To get all other UserDetails  --GET

    http://localhost:8000/users/  

## For Block Functionality

To block other User  --POST

    http://localhost:8000/users/block 

To unblock User by userID  --DELETE

    http://localhost:8000/users/{blocked}/unblock 

To get list of all blocked Users  --GET

    http://localhost:8000/users/block

## Public apis 

To get details of users  --POST

    http://localhost:8000/internal/users/details
    
To get otp on registered details(for delete organization)  --POST

    http://localhost:8000/internal/user/otp

To verify otp for delete organization --POST

    http://localhost:8000/internal/otp/verify

