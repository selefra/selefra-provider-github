# Table: github_organization_members

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| followers | int | X | √ |  | 
| total_private_repos | int | X | √ |  | 
| ldap_dn | string | X | √ |  | 
| following_url | string | X | √ |  | 
| login | string | X | √ |  | 
| blog | string | X | √ |  | 
| public_repos | int | X | √ |  | 
| subscriptions_url | string | X | √ |  | 
| role_name | string | X | √ |  | 
| github_organizations_selefra_id | string | X | X | fk to github_organizations.selefra_id | 
| location | string | X | √ |  | 
| twitter_username | string | X | √ |  | 
| disk_usage | int | X | √ |  | 
| text_matches | json | X | √ |  | 
| name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| gists_url | string | X | √ |  | 
| public_gists | int | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| followers_url | string | X | √ |  | 
| gravatar_id | string | X | √ |  | 
| suspended_at | timestamp | X | √ |  | 
| collaborators | int | X | √ |  | 
| two_factor_authentication | bool | X | √ |  | 
| events_url | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| id | int | X | √ |  | 
| node_id | string | X | √ |  | 
| site_admin | bool | X | √ |  | 
| url | string | X | √ |  | 
| permissions | json | X | √ |  | 
| company | string | X | √ |  | 
| email | string | X | √ |  | 
| following | int | X | √ |  | 
| bio | string | X | √ |  | 
| type | string | X | √ |  | 
| owned_private_repos | int | X | √ |  | 
| private_gists | int | X | √ |  | 
| organizations_url | string | X | √ |  | 
| avatar_url | string | X | √ |  | 
| html_url | string | X | √ |  | 
| hireable | bool | X | √ |  | 
| received_events_url | string | X | √ |  | 
| starred_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| membership | json | X | √ |  | 
| plan | json | X | √ |  | 
| repos_url | string | X | √ |  | 


