# Table: github_organization_dependabot_alerts

## Primary Keys 

```
org, number
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| dismissed_at | timestamp | X | √ |  | 
| fixed_at | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| dependency | json | X | √ |  | 
| security_advisory | json | X | √ |  | 
| github_organizations_selefra_id | string | X | X | fk to github_organizations.selefra_id | 
| url | string | X | √ |  | 
| dismissed_by | json | X | √ |  | 
| dismissed_reason | string | X | √ |  | 
| number | int | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_at | timestamp | X | √ |  | 
| dismissed_comment | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| security_vulnerability | json | X | √ |  | 
| html_url | string | X | √ |  | 


