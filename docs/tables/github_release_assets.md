# Table: github_release_assets

## Primary Keys 

```
org, repository_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| size | int | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| browser_download_url | string | X | √ |  | 
| uploader | json | X | √ |  | 
| github_releases_selefra_id | string | X | X | fk to github_releases.selefra_id | 
| download_count | int | X | √ |  | 
| id | int | X | √ |  | 
| label | string | X | √ |  | 
| content_type | string | X | √ |  | 
| repository_id | int | X | √ |  | 
| url | string | X | √ |  | 
| name | string | X | √ |  | 
| state | string | X | √ |  | 
| node_id | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 


