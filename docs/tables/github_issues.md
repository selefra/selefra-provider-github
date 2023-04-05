# Table: github_issues

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| url | string | X | √ |  | 
| events_url | string | X | √ |  | 
| repository | json | X | √ |  | 
| assignees | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| number | int | X | √ |  | 
| locked | bool | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| closed_by | json | X | √ |  | 
| pull_request | json | X | √ |  | 
| reactions | json | X | √ |  | 
| text_matches | json | X | √ |  | 
| body | string | X | √ |  | 
| author_association | string | X | √ |  | 
| assignee | json | X | √ |  | 
| comments_url | string | X | √ |  | 
| milestone | json | X | √ |  | 
| active_lock_reason | string | X | √ |  | 
| title | string | X | √ |  | 
| comments | int | X | √ |  | 
| html_url | string | X | √ |  | 
| labels | json | X | √ |  | 
| labels_url | string | X | √ |  | 
| user | json | X | √ |  | 
| node_id | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| id | int | X | √ |  | 
| state_reason | string | X | √ |  | 
| state | string | X | √ |  | 
| closed_at | timestamp | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| repository_url | string | X | √ |  | 


