# Table: github_billing_package

## Primary Keys 

```
org
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| included_gigabytes_bandwidth | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| org | string | √ | √ | `The Github Organization of the resource.` | 
| total_gigabytes_bandwidth_used | int | X | √ |  | 
| total_paid_gigabytes_bandwidth_used | int | X | √ |  | 


