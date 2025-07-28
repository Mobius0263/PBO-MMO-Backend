# CoEmotion API - HTTP and HTTPS Support

## Current Setup

Your Swagger documentation now supports both `http` and `https` schemes. Here's how to use them:

### Development (Current Setup)
- **HTTP**: `http://localhost:8080` ✅ **Currently Running**
- **Swagger UI**: `http://localhost:8080/swagger/index.html`

### Using Swagger UI

1. **HTTP Scheme (Recommended for Development)**:
   - Select `http` from the scheme dropdown in Swagger UI
   - All API requests will work normally
   - No certificate warnings

2. **HTTPS Scheme (For Testing Production Setup)**:
   - Select `https` from the scheme dropdown in Swagger UI
   - ⚠️ **Will fail** because server is running on HTTP only
   - This is expected behavior in development

## Setting Up HTTPS for Local Development

If you want to test HTTPS locally, follow these steps:

### Option 1: Generate SSL Certificate (Recommended)

1. **Run the certificate generator**:
   ```bash
   # For Windows Command Prompt
   generate-cert.bat
   
   # Or using Git Bash (if OpenSSL not available in CMD)
   openssl genrsa -out certs/localhost.key 2048
   openssl req -new -key certs/localhost.key -out certs/localhost.csr -subj "/C=ID/ST=State/L=City/O=Organization/OU=OrgUnit/CN=localhost"
   openssl x509 -req -days 365 -in certs/localhost.csr -signkey certs/localhost.key -out certs/localhost.crt
   rm certs/localhost.csr
   ```

2. **Use the HTTPS-enabled version**:
   ```bash
   # Backup current main.go
   cp main.go main-http-only.go
   
   # Use the HTTPS version
   cp main-https-example.go.txt main.go
   
   # Run server
   go run main.go
   ```

3. **Access your API**:
   - HTTP: `http://localhost:8080`
   - HTTPS: `https://localhost:8443` (will show security warning - click "Advanced" → "Proceed")

### Option 2: Production Deployment

For production deployment with proper SSL certificates:

1. **Get SSL certificates from a Certificate Authority** (Let's Encrypt, etc.)
2. **Update host in Swagger annotations**:
   ```go
   //	@host		yourdomain.com
   ```
3. **Configure your server** with proper certificates
4. **Deploy** to your hosting provider

## Swagger Schemes Explanation

The `@schemes http https` annotation in your Swagger documentation means:

- **HTTP**: For development and testing
- **HTTPS**: For production deployment with SSL certificates

Both schemes appear in the Swagger UI dropdown, but only the scheme your server actually supports will work.

## Current File Structure

```
├── main.go                     # HTTP-only version (current)
├── main-https-example.go.txt   # HTTPS-enabled version example
├── generate-cert.bat           # SSL certificate generator
├── certs/                      # SSL certificates (when generated)
│   ├── localhost.crt
│   └── localhost.key
└── docs/
    ├── swagger.json            # Contains both http and https schemes
    └── swagger.yaml            # Contains both http and https schemes
```

## Troubleshooting

### "Failed to fetch" Error with HTTPS
- **Cause**: Server is running on HTTP but Swagger is trying HTTPS
- **Solution**: Use HTTP scheme in Swagger UI for development

### Browser Security Warning with HTTPS
- **Cause**: Self-signed certificate
- **Solution**: Click "Advanced" → "Proceed to localhost (unsafe)" in browser

### Port Already in Use
- **Solution**: Stop existing server process or use different port:
  ```bash
  PORT=3000 go run main.go
  ```

## Production Considerations

1. **Use proper SSL certificates** (not self-signed)
2. **Update CORS origins** for your production domain
3. **Set proper environment variables**:
   ```bash
   PORT=443
   HTTPS_PORT=443
   CORS_ORIGINS=https://yourdomain.com
   ```
4. **Update Swagger host**:
   ```go
   //	@host		yourdomain.com
   ```
