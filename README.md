# WRY

## Description
WRY is a centralized authentication microservice designed to serve multiple existing and future applications, such as datall.com.mx.

The name WRY carries a double meaning: it stands for Who r You as the core question of identity verification, while also playing on the definition of the word wry, bringing an ironic and mocking touch to identity and access control.

The main goal is to decouple and centralize authentication logic, preventing each satellite application from having to handle its own user management, sessions, and token signing.

## General Authentication Flow
1. A client application sends a login request with username, password, and the target application identifier (app_id).
2. WRY validates user credentials and retrieves the roles or permissions assigned to that user for the specified app.
3. WRY generates and signs a JWT token using an asymmetric RSA private key.
4. The token is returned to the client application, which attaches it to subsequent requests.
5. Consumer applications can verify token authenticity autonomously without querying WRY on every request, by fetching the public key exposed via the JWKS endpoint.

## PostgreSQL Integration
Currently, the project runs a base in-memory setup for testing and signing tokens. The upcoming stage includes integrating PostgreSQL to manage:
- User directory and password hashes.
- Authorized application registry (app_id).
- User roles and permissions per application.
- Access logs and audit events.

## Available Endpoints

*POST /api/v1/login*
Receives user credentials and returns the signed access token.

*GET /.well-known/jwks.json*
Exposes the public key set in JWKS format for decentralized token verification.
