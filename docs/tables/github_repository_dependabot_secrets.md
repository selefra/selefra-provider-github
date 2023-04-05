# Table: github_repository_dependabot_secrets

## Primary Keys 

```
org, repository_id, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| repository_id | int | X | √ |  | 
| visibility | string | X | √ |  | 
| selected_repositories_url | string | X | √ |  | 
| name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| github_repositories_selefra_id | string | X | X | fk to github_repositories.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


