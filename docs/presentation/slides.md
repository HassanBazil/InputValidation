# Input Validation
---
## Overview
Input validation is used to solve the problem that complex code's correctness is hard to verify. Parsers are really complex so you want to protect them using a much simpler code.

- Don't fix data if security matters.
- Yes you can use regex, but should you? No, regex is hard you want simple.
- Validate all sources (file, config etc.)
- Trade-off especially at semantic level cost of validation vs risk
- Validate inputs that you processs/parse skip what you don't process
- Know when to be careful
    - complex input is where you can embed context several layer deeply

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
---
## Demo 1
AI https://embracethered.com/blog/posts/2024/claude-hidden-prompt-injection-ascii-smuggling/
---
### Demo 2
See code...