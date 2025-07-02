# simplefin

Basic library to interact with [simplefin.org](https://www.simplefin.org/)

Usage

`go get github.com/jazzboME/simplefin@v0.1.0`

## Working with a local json file
```
	accts, err := SfinFile("local.json")

```

## Connecting to a simplefin endpoint

You'll need to supply a SimpleFin URL, see [simplefile documentation](https://www.simplefin.org/protocol.html#claim-the-access-url)

Assuming the URL is stored in an environment variable ACC_URL:

```
	access_url := os.Getenv("ACC_URL")
	accts, err := SfinURL(access.url)
```

In both cases `accts` will be a struct designed to decode the simplefin json, see [description of simplefin data](https://www.simplefin.org/protocol.html#data)

This is all very simple right now. 

### To Do:
  Add more explicit error conditions.
  Test!

### Long Term To Do:
  Create routines to go through the entire Token -> URL process.  
