# URL to request
URL="http://localhost:8090/image"

# Make the curl request
curl -X POST -F "file=@test/testdata/ss.png" -k "$URL"