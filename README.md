# gh-pages-internal
Serve static files directly from git repos

## Environment variables (optional)

- `GIT_TOKEN`: 
    - A token in the form `host:token` for example `github.com:1234`
- `FS_TYPE`: 
    - File system type, one of `os`, `memory`, or `gcs`
- `GCS_BUCKET`:
    - The bucket name that objects should be cached to, or the cache file if FS_TYPE not `gcs`
- `BASE`:
    - The base for the files to be fetched, can be `github.com/myorg` or `github.com` for example
        - if  `github.com/myorg`, the url for the website will be `<host>/repo/path` 
        - if `github.com`, the url for the website will be `<host>org/repo/path`
        
## Caching in gcs and getting around the github api rate limit
- Rate limiting will occur with github api tokens (eventually), because of this you can choose to cache assets in a GCS bucket by setting the `GCS_BUCKET` environment variable
- To remove cache (could be a ci step), you can request to delete a path:
```bash
 curl --location --request DELETE 'http://<host>/<resource>/'
```
- This will delete from gcs cache so the next request of `resource` will be a fresh fetch from github
