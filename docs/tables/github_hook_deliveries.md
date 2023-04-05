# Table: github_hook_deliveries

## Primary Keys 

```
org, hook_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| response | string | X | √ |  | 
| delivered_at | timestamp | X | √ |  | 
| event | string | X | √ |  | 
| request | string | X | √ |  | 
| id | int | X | √ |  | 
| guid | string | X | √ |  | 
| redelivery | bool | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| hook_id | int | X | √ | `Hook ID` | 
| duration | float | X | √ |  | 
| status | string | X | √ |  | 
| status_code | int | X | √ |  | 
| action | string | X | √ |  | 
| installation_id | int | X | √ |  | 
| repository_id | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| github_hooks_selefra_id | string | X | X | fk to github_hooks.selefra_id | 


