# hello-cf-stream
This is a sample for Cloudflare Streams, which includes the following:

- Markdown to generate sample HTML and instructions on how to use it
- Golang client for generating tokens to use with signed URLs.

## Simple sample
### Create HTML
```html
<html>
<head>
    <title>My First PHP Page</title>
</head>
<body>
    <iframe src="https://customer-{customer-id}.cloudflarestream.com/{video-id}/iframe"
        width="1000" height="600" title="Example Stream video" frameBorder="0"
        allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowFullScreen>
    </iframe>
</body>
</html>
```

### Run localhost server
```bash
npx http-server -p 8090
```

## Sample with Signed URL
### Obtain a key
```bash
API_TOKEN=<your_api_token>
ACCOUNT_ID=<your_account_id>
curl -X POST -H "Authorization: Bearer $API_TOKEN" "https://api.cloudflare.com/client/v4/accounts/$ACCOUNT_ID/stream/keys" : keys.json
```

### Generate JWT
```bash
go run main.go <your_video_id>
```

## Render video
Render video using generated JWT
