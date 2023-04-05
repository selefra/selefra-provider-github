# Table: github_team_members

## Primary Keys 

```
org, team_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| events_url | string | X | √ |  | 
| repos_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| team_id | int | X | √ |  | 
| total_private_repos | int | X | √ |  | 
| collaborators | int | X | √ |  | 
| two_factor_authentication | bool | X | √ |  | 
| site_admin | bool | X | √ |  | 
| disk_usage | int | X | √ |  | 
| gists_url | string | X | √ |  | 
| organizations_url | string | X | √ |  | 
| gravatar_id | string | X | √ |  | 
| company | string | X | √ |  | 
| location | string | X | √ |  | 
| email | string | X | √ |  | 
| subscriptions_url | string | X | √ |  | 
| followers | int | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| suspended_at | timestamp | X | √ |  | 
| starred_url | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| membership | json | X | √ |  | 
| id | int | X | √ |  | 
| bio | string | X | √ |  | 
| role_name | string | X | √ |  | 
| plan | json | X | √ |  | 
| text_matches | json | X | √ |  | 
| type | string | X | √ |  | 
| owned_private_repos | int | X | √ |  | 
| private_gists | int | X | √ |  | 
| received_events_url | string | X | √ |  | 
| blog | string | X | √ |  | 
| public_repos | int | X | √ |  | 
| public_gists | int | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| permissions | json | X | √ |  | 
| followers_url | string | X | √ |  | 
| github_teams_selefra_id | string | X | X | fk to github_teams.selefra_id | 
| node_id | string | X | √ |  | 
| avatar_url | string | X | √ |  | 
| name | string | X | √ |  | 
| following_url | string | X | √ |  | 
| twitter_username | string | X | √ |  | 
| following | int | X | √ |  | 
| url | string | X | √ |  | 
| login | string | X | √ |  | 
| html_url | string | X | √ |  | 
| hireable | bool | X | √ |  | 
| ldap_dn | string | X | √ |  | 


