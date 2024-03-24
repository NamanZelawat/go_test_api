# URL to request
URL="http://127.0.0.1:52114/image"

# Make the curl request
curl -X POST -F "file=@test/testdata/ss.png" -k "$URL"