# Table: github_organization_dependabot_secrets

## Primary Keys 

```
org, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| updated_at | timestamp | X | √ |  | 
| visibility | string | X | √ |  | 
| selected_repositories_url | string | X | √ |  | 
| github_organizations_selefra_id | string | X | X | fk to github_organizations.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 


