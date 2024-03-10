# Go OTP Sender and Verifier

This project is a simple Go application for sending and verifying OTP (One-Time Password) using the Gin web framework. It leverages the go-verify-otp library for OTP functionality.

## Features

- Send SMS OTP
- Verify OTP.

## Getting Started


### Prerequisites

- Go installed on your system.

### Installation



Clone the repository:

```bash
git clone https://github.com/Mayejacob/go-verify-otp.git
```
```bash
go mod tidy

```

## Usage

1. **Run the Application**

To run the application, cd to the **CMD** folder and use the following command:

```bash
go run main.go
```

The server will start on port 8000 by default.
## API Endpoints

- **Send OTP:** `POST /api/v1/otp`
- **Verify OTP:** `POST /api/v1/verify_otp`

## Example Request

- **Send OTP:**

```bash
{
    "phoneNumber": "+234810000"
}
```

- **Verify OTP:**

```bash
{
    "user": {
        "phoneNumber": "+234810000"
    },
    "code": "938985"
}
```

## Example response
```bash
{
    "status": 202,
    "message": "Success",
    "data": "OTP verified successfully"
}
```