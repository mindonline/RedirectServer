# Small and fast redirect service 

Use simple JSON schema to redirect HTTP incoming requests.

## Initialization

All schemas must be written in *schema.json* file.
If you need to specify different listen address or port
set environment variable **HOST_ADDR** or write it in *.env* file, for example:

```
# Listen on 8080 port for all interfaces
HOST_ADDR=0.0.0.0:8080
```

## Example
**schema.json:**

```json
{
  "redirects": [
    {
      "from": "/google",
      "to": "http://google.com"
    }
  ]
}
```
Now try to type in the browser *http://localhost/google* to get redirect to *http://google.com*
