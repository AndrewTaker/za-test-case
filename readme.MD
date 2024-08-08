You can pick any `Bearer` token from `scheme.sql` dump.

1. run `./infra/make-container.sh` from project root directory.

2. list endpoint

```bash
curl \
    --request GET \
    --url "http://localhost:33333/api/v1/news/list?limit=5&offset=0" \
    --header "Authorization: Bearer 5EiFJtQpXarwy0hK7SdlTAkUbO59wmDg"
```

3. edit enpoint
```bash
curl \
    --request POST \
    --url "http://localhost:33333/api/v1/news/edit/1" \
    --header "content-type: application/json" \
    --header "Authorization: Bearer 5EiFJtQpXarwy0hK7SdlTAkUbO59wmDg" \
    --data '
    {
        "title": "new_title",
        "content": "new_content",
        "categories": [1,2,3,4,5,6,7]
    }
    '
```