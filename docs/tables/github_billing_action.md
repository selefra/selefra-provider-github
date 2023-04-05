# Table: github_billing_action

## Primary Keys 

```
org
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| org | string | √ | √ | `The Github Organization of the resource.` | 
| total_minutes_used | int | X | √ |  | 
| total_paid_minutes_used | float | X | √ |  | 
| included_minutes | int | X | √ |  | 
| minutes_used_breakdown | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


