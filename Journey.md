# 2025.1: Making JWT in Go

***++What is JWT?++***

Json Web Tokens is a compact, URL-safe way to transmit information between parties as a JSON object. How It Works:
- Client Loggs In: Server Creates a JWT and sends it back
- Client stores jwt: (usually in local storage or a cookie)
- Client Makes Requests: Attaches the JWT in headers (`Authorization: Bearer <token>`)
- Server verifies JWT: If valid, allows access.

---

# JWT Structure:

Jwts have 3 parts seperated by dots `.`

`header.payload.signature`: Example: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzLCJleHAiOjE2OTAwMDAwfQ.V4rLLk5LZb2pK8Y0mDqF9PoQ6rQ2r9c8KSl-fyq2QUk`

==Header (Base64-encoded JSON):== This specifies the hashing algorithm 
```json
~ Header

{
    "alg": "HS256",
    "typ": "JWT"
}

```

==Payload (Base64-encoded JSON):== This stores user data and the expiration time 
```json
~ Payload

{
    "id": 123,
    "exp": 16900000
}

```

==Signature (Used To Verify authenticity):==  Ensures that the JWT wasn't tampered with.
```
~ signature

HMAC_SHA256(base64(header) + "." + base64(payload), secret)
```

- The `secret` is a private key only the server knows.
- All clients get the same `secret`
