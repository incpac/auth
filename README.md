# auth
A very simple micro-service for generating JWT tokens. It uses a CouchBase database to store users. The idea for this service is not to provide user management, however, it can create users and reset their passwords.

## Build

Install dependencies and compile

```
go get github.com/incpac/auth
go build github.com/incpac/auth

```

## Usage

### Create a user

    auth create
    
You will be asked to provide user information

### Start the webserver

    auth server --key secret

### Authenticate a user and create a token

    curl -XPOST -u username:password 'http://localhost:3000/authenticate'

### Verify a token

    curl -XPOST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ' 'http://localhost:3000/verify'


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'added some feature`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


## License 

```text
The MIT License (MIT)

Copyright (c) 2017 Thomas Claridge

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```