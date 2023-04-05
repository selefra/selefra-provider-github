# Table: github_organizations

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| public_gists | int | X | √ |  | 
| has_repository_projects | bool | X | √ |  | 
| hooks_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default_repository_permission | string | X | √ |  | 
| members_allowed_repository_creation_type | string | X | √ |  | 
| login | string | X | √ |  | 
| html_url | string | X | √ |  | 
| members_can_create_private_pages | bool | X | √ |  | 
| events_url | string | X | √ |  | 
| disk_usage | int | X | √ |  | 
| two_factor_requirement_enabled | bool | X | √ |  | 
| members_can_create_repositories | bool | X | √ |  | 
| dependency_graph_enabled_for_new_repositories | bool | X | √ |  | 
| url | string | X | √ |  | 
| name | string | X | √ |  | 
| blog | string | X | √ |  | 
| location | string | X | √ |  | 
| email | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| collaborators | int | X | √ |  | 
| is_verified | bool | X | √ |  | 
| issues_url | string | X | √ |  | 
| secret_scanning_push_protection_enabled_for_new_repositories | bool | X | √ |  | 
| description | string | X | √ |  | 
| total_private_repos | int | X | √ |  | 
| billing_email | string | X | √ |  | 
| web_commit_signoff_required | bool | X | √ |  | 
| members_can_create_private_repositories | bool | X | √ |  | 
| members_can_create_pages | bool | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| avatar_url | string | X | √ |  | 
| followers | int | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| members_can_create_public_repositories | bool | X | √ |  | 
| public_repos | int | X | √ |  | 
| plan | json | X | √ |  | 
| members_can_create_internal_repositories | bool | X | √ |  | 
| company | string | X | √ |  | 
| twitter_username | string | X | √ |  | 
| members_can_fork_private_repositories | bool | X | √ |  | 
| advanced_security_enabled_for_new_repositories | bool | X | √ |  | 
| node_id | string | X | √ |  | 
| dependabot_alerts_enabled_for_new_repositories | bool | X | √ |  | 
| default_repository_settings | string | X | √ |  | 
| public_members_url | string | X | √ |  | 
| has_organization_projects | bool | X | √ |  | 
| repos_url | string | X | √ |  | 
| following | int | X | √ |  | 
| owned_private_repos | int | X | √ |  | 
| private_gists | int | X | √ |  | 
| members_url | string | X | √ |  | 
| id | int | X | √ |  | 
| members_can_create_public_pages | bool | X | √ |  | 
| dependabot_security_updates_enabled_for_new_repositories | bool | X | √ |  | 
| secret_scanning_enabled_for_new_repositories | bool | X | √ |  | 


