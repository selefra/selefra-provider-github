# Table: github_repository_dependabot_alerts

## Primary Keys 

```
org, repository_id, number
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| updated_at | timestamp | X | √ |  | 
| dismissed_comment | string | X | √ |  | 
| fixed_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| dependency | json | X | √ |  | 
| url | string | X | √ |  | 
| number | int | X | √ |  | 
| state | string | X | √ |  | 
| security_vulnerability | json | X | √ |  | 
| html_url | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| dismissed_by | json | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| repository_id | int | X | √ |  | 
| dismissed_reason | string | X | √ |  | 
| github_repositories_selefra_id | string | X | X | fk to github_repositories.selefra_id | 
| security_advisory | json | X | √ |  | 
| dismissed_at | timestamp | X | √ |  | 


