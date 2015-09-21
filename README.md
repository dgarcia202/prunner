# prunner
postman collection runner

## usage

With postman collection url:
```
prunner <flags> -source=https://www.getpostman.com/collections/a6079cfeb3c2ed5aeaef
```

With json file containing collection info:
```
prunner <flags> -source=\data\mycollection.json
```

Source flag can be a valid getpostman collection url or a path+filename containing same information.

Flags can be:
* concise: Only output result of API calls.
* export: Save the contents of a url source to a local file named 'source.json'.

## todo
* execution iterations.
* optional execution of requests in parallel.
* ordering request execution as is in the collection definition.
* somehow stablish expected results for requests and present them properly.

