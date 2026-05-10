# Input Validation
---
## Overview
Yes you can use regex, but should you?
---
## Origin Validation
- IP address verification
- Access key (token) authentication
- Verify source of the request
---
## Size Validation
- Is message reasonably long?
- Are all fields reasonably long as well?
- Check if input exceeds acceptable limits (e.g., > 1000 bytes)
---
## Lexical Validation
Content uses right characters and encoding?
- Normalize encoding, reject invalid chars
- Allow only certain character categories (i.e.: Only latin chars, uppercase letters etc.)
- Allow only certain chars (i.e.: ' for names for our Irish friends but not .)
- Validate UTF-8 encoding
- Normalize unicode bytes using composing representation
---
## Syntax Validation
Is the format right?
- JSON double property check
- JSON extra property during deserialization
- Email address should adhere to standard
- Use `JSON.Valid()` to verify structure
---
## Parsing
Implement JSON unmarshal using various libraries
- Stock encoding/json
- goccy/go-json
- json-iterator/go
- GoJay
- segmentio/encoding/json
---
## Semantic Validation
Does it make sense?
- Is price of good positive?
- Is this address actually exist?
- Is this filetype make sense?
- Is this email controlled by the user?
- Is cryptographic controls applied strong and passes validation?
- Is this url points to public or private resource?
- Is this user exist in the DB?
- Uniqueness check
- Dereferencing
- TOCTOU locking