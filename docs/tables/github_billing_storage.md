# Table: github_billing_storage

## Primary Keys 

```
org
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| org | string | √ | √ | `The Github Organization of the resource.` | 
| days_left_in_billing_cycle | int | X | √ |  | 
| estimated_paid_storage_for_month | float | X | √ |  | 
| estimated_storage_for_month | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


