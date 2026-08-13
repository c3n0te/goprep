url testing (on windows):
curl.exe -X POST 127.0.0.1:8000/shorten -H "Content-Type: application/json" -d "{\`"url\`": \`"https://www.example.com/some/long/url\`"}"
