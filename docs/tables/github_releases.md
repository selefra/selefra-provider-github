# Table: github_releases

## Primary Keys 

```
org, repository_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| discussion_category_name | string | X | √ |  | 
| tarball_url | string | X | √ |  | 
| author | json | X | √ |  | 
| name | string | X | √ |  | 
| body | string | X | √ |  | 
| prerelease | bool | X | √ |  | 
| generate_release_notes | bool | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| url | string | X | √ |  | 
| html_url | string | X | √ |  | 
| assets_url | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| repository_id | int | X | √ |  | 
| tag_name | string | X | √ |  | 
| zipball_url | string | X | √ |  | 
| node_id | string | X | √ |  | 
| github_repositories_selefra_id | string | X | X | fk to github_repositories.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| target_commitish | string | X | √ |  | 
| make_latest | string | X | √ |  | 
| upload_url | string | X | √ |  | 
| assets | json | X | √ |  | 
| draft | bool | X | √ |  | 
| id | int | X | √ |  | 
| published_at | timestamp | X | √ |  | 


