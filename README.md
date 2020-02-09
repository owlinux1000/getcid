# Get Container ID

## tl;dr

This Docker Image only returns self hostname. 

## How to users

```sh
$ doker run -it -p 80:9191 encry1024:gethostname
$ curl http://localhost
2840cd4acfe0 # Example hostname
```
