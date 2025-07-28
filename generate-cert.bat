@echo off
echo Creating self-signed SSL certificate for localhost development...
echo.

REM Check if OpenSSL is available
openssl version >nul 2>&1
if errorlevel 1 (
    echo OpenSSL not found. Please install OpenSSL or use Git Bash which includes OpenSSL.
    echo You can download Git for Windows from: https://git-scm.com/download/win
    echo Then run this script from Git Bash instead.
    pause
    exit /b 1
)

REM Create certificates directory if it doesn't exist
if not exist "certs" mkdir certs

REM Generate private key
echo Generating private key...
openssl genrsa -out certs/localhost.key 2048

REM Generate certificate signing request
echo Generating certificate signing request...
openssl req -new -key certs/localhost.key -out certs/localhost.csr -subj "/C=ID/ST=State/L=City/O=Organization/OU=OrgUnit/CN=localhost"

REM Generate self-signed certificate
echo Generating self-signed certificate...
openssl x509 -req -days 365 -in certs/localhost.csr -signkey certs/localhost.key -out certs/localhost.crt

REM Clean up CSR file
del certs\localhost.csr

echo.
echo SSL certificate generated successfully!
echo Certificate: certs/localhost.crt
echo Private Key: certs/localhost.key
echo.
echo To use HTTPS, modify your main.go to use:
echo app.ListenTLS(":8443", "certs/localhost.crt", "certs/localhost.key")
echo.
pause
