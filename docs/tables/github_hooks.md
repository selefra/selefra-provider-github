# Table: github_hooks

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ping_url | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| created_at | timestamp | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| id | int | X | √ |  | 
| events | string_array | X | √ |  | 
| active | bool | X | √ |  | 
| test_url | string | X | √ |  | 
| last_response | json | X | √ |  | 
| config | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| url | string | X | √ |  | 
| type | string | X | √ |  | 
| name | string | X | √ |  | 


