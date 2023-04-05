# Table: github_installations

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| events | string_array | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| has_multiple_single_files | bool | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| id | int | X | √ |  | 
| html_url | string | X | √ |  | 
| single_file_name | string | X | √ |  | 
| node_id | string | X | √ |  | 
| app_id | int | X | √ |  | 
| access_tokens_url | string | X | √ |  | 
| permissions | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| app_slug | string | X | √ |  | 
| repository_selection | string | X | √ |  | 
| single_file_paths | string_array | X | √ |  | 
| suspended_at | timestamp | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| suspended_by | json | X | √ |  | 
| target_id | int | X | √ |  | 
| account | json | X | √ |  | 
| repositories_url | string | X | √ |  | 
| target_type | string | X | √ |  | 


