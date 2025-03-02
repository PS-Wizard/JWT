Sun Mar  2 07:39:54 PM +0545 2025
---

# Making JWT implementation from scratch

[See The Journey](./Journey.md)

# TODO List for Extending JWT Implementation

## 1. Refactor Code for Better Organization
- [x] Create separate functions for:
  - Header generation
  - Payload generation
  - Signature generation
  - JWT creation (combining header, payload, signature)
- [ ] Move code into separate files:
  - `header.go`: For header-related functionality.
  - `payload.go`: For payload-related functionality.
  - `signature.go`: For signature creation.
  - `jwt.go`: For the main JWT structure and logic.

## 2. Implement Claim Validation
- [ ] Create a `claims.go` file to manage claim validation (e.g., `exp`, `iat`, `sub`, etc.).
  - [ ] Implement function to check if the token has expired (`exp`).
  - [ ] Implement function to validate other claims like `iss`, `sub`, `aud`.

## 3. Improve Signature Handling
- [ ] Support multiple signing algorithms (e.g., RSA, ECDSA).
  - [ ] Implement an RSA signing method (e.g., using `crypto/rsa`).
  - [ ] Implement an ECDSA signing method (e.g., using `crypto/ecdsa`).
  - [ ] Create an interface to handle the signing process dynamically based on the algorithm.

## 4. Enhance Error Handling
- [ ] Implement proper error handling for all major operations (e.g., invalid header, invalid signature).
  - [ ] Create custom error types for different JWT errors.
  - [ ] Return detailed error messages when validation fails.

## 5. Add Token Expiration Handling
- [ ] Implement token expiration (`exp` claim).
  - [ ] Check if the token is expired during verification.
  - [ ] Optionally add refresh token functionality to allow re-authentication.

## 6. Add Unit Tests
- [ ] Write unit tests for each function (header creation, payload creation, signature generation).
  - [ ] Use Go's testing framework (`testing` package).
  - [ ] Ensure each edge case is tested (e.g., invalid inputs, expired tokens).

## 7. Implement JWT Parsing
- [ ] Add a function to parse a JWT from a string and extract header, payload, and signature.
- [ ] Implement a method to verify a token's signature during parsing.

## 8. (Optional) Implement JWT Claims Struct
- [ ] Create a `JWTClaims` struct to hold the decoded claims.
  - [ ] Store information like `id`, `exp`, `iat`, `sub`, etc.
- [ ] Add a function to map JWT claims into the struct for easy access.

## 9. Documentation
- [ ] Add comments to all functions and code blocks explaining what they do.
- [ ] Write a `README.md` file to explain how to use the JWT package, with example code.

## 10. (Optional) Add JWT Parsing for RSA Keys
- [ ] Implement parsing for public/private keys if you plan to use RSA or ECDSA signatures.
- [ ] Add the ability to verify tokens with public keys.


