#Service for search words from given letters

##API
###Find words from given letters
```
curl -v 'http://127.0.0.1:8080/?letters=wordOrLetters'
```

###Load dict
```
curl -vX POST 'http://127.0.0.1:8080/' \
 -H 'Content-Type: application/json' \
 --data '["word", "test", "tree"]'
```