# Table: github_external_groups

## Primary Keys 

```
org, group_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| updated_at | timestamp | X | √ |  | 
| teams | json | X | √ |  | 
| members | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| group_id | int | X | √ |  | 
| group_name | string | X | √ |  | 


