Magic Link
===

Setup an endpoint to receive a GET/POST request, requesting the provisioning of a magic link.

- Web server
- HTTP Endpoint GET /authsomething?app=abc&type=magic&email=alice@example.com
- HTTP Endpoint POST /authsomething { app: abc, type: magic, email: alice@example.com }
- Email Service Provider Integration (AWS? Sendgrid? Other?)

Should the payloads be signed? HMAC or JWT. JWKS (rfc7517)

Rate Limiting requests should also be considered.
An app may not intentionally spam the magiclink endpoint, but faulty queues and connections can cause issues. Additionally, bad-actors can craft the url and attempt dos attacks, or masking other attacks with notification fatigue. Exponential backoff might be a good approach. Most users will request once or twice (maybe three times when pressured - logging in during a meeting and the email's a little behind).

Links in the email should stay active for a few minutes then self expire. Maybe that's the rate limiting? A link exists for 5 minutes, and you can't have more than 3 out at a time. Request rate limiting can be done separately which would be more globally for the service as a whole, not only Magic Links.

Schemas

tbl Application: id, name (more later)
tbl Authentication: id, id_app, type, redirect
tbl Account: id, id_app, email_address, (more later)
tbl Tokens: id, id_app, id_account, date_created, date_expires, consumed, token_string
tbl Keys: id, id_app, date_created, date_expires, enc_private_key, public_key

Token should be signed, keys can be stored (encrypted) in the database with the master key passed in through envar (not in .env files). This would allow implementing JWKS.

Flow:

User makes request to login to client app with their email.
Client app uses its private key to generate an HMAC for the request to Auth1 endpoint
Auth1 endpoint validates the request
Auth1 generates a token for the user
Auth1 emails token to the user
User receives email with token, and clicks url (with token)
Auth1 validates the token
Auth1 generates a User JWT
Auth1 consumes the token (expires manually, not time based)
Auth1 redirects the user to the client app's configured redirection, including the User JWT
Client app receives the JWT
Client app optionally validates the JWT against the JWKS endpoint on Auth1
Client app completes its logic
User is logged in / Authenticated.

---

2025-04-19 20:54	Created task
2025-04-19 21:00	Updated task
2025-04-25 21:24	Updated task
2025-04-25 21:50	Updated task
2025-04-25 22:07	Updated task
2025-04-25 22:16	Updated task