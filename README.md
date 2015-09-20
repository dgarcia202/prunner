# prunner
postman collection runner

## usage
```
prunner <flags> -source=https://www.getpostman.com/collections/a6079cfeb3c2ed5aeaef
```

Source flag can be a valid getpostman collection url or a path+filename containing same information.

Flags can be:

-concise: Only output result of API calls.

-export: Save the contents of a url source to a local file named 'source.json'.

