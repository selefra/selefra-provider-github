# Table: github_teams

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| members_count | int | X | √ |  | 
| html_url | string | X | √ |  | 
| members_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| id | int | X | √ |  | 
| url | string | X | √ |  | 
| slug | string | X | √ |  | 
| repos_count | int | X | √ |  | 
| repositories_url | string | X | √ |  | 
| parent | json | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| permission | string | X | √ |  | 
| permissions | json | X | √ |  | 
| privacy | string | X | √ |  | 
| node_id | string | X | √ |  | 
| organization | json | X | √ |  | 
| ldap_dn | string | X | √ |  | 


