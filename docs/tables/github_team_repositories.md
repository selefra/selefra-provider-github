# Table: github_team_repositories

## Primary Keys 

```
org, team_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| size | int | X | √ |  | 
| forks_url | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| parent | json | X | √ |  | 
| languages_url | string | X | √ |  | 
| compare_url | string | X | √ |  | 
| name | string | X | √ |  | 
| mirror_url | string | X | √ |  | 
| organization | json | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| allow_update_branch | bool | X | √ |  | 
| topics | string_array | X | √ |  | 
| gitignore_template | string | X | √ |  | 
| issue_comment_url | string | X | √ |  | 
| ssh_url | string | X | √ |  | 
| has_issues | bool | X | √ |  | 
| statuses_url | string | X | √ |  | 
| open_issues_count | int | X | √ |  | 
| template_repository | json | X | √ |  | 
| issues_url | string | X | √ |  | 
| teams_url | string | X | √ |  | 
| github_teams_selefra_id | string | X | X | fk to github_teams.selefra_id | 
| has_downloads | bool | X | √ |  | 
| permissions | json | X | √ |  | 
| archived | bool | X | √ |  | 
| collaborators_url | string | X | √ |  | 
| notifications_url | string | X | √ |  | 
| labels_url | string | X | √ |  | 
| milestones_url | string | X | √ |  | 
| master_branch | string | X | √ |  | 
| html_url | string | X | √ |  | 
| allow_rebase_merge | bool | X | √ |  | 
| branches_url | string | X | √ |  | 
| full_name | string | X | √ |  | 
| description | string | X | √ |  | 
| merge_commit_title | string | X | √ |  | 
| merge_commit_message | string | X | √ |  | 
| has_discussions | bool | X | √ |  | 
| allow_auto_merge | bool | X | √ |  | 
| commits_url | string | X | √ |  | 
| stargazers_count | int | X | √ |  | 
| subscription_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| is_template | bool | X | √ |  | 
| assignees_url | string | X | √ |  | 
| fork | bool | X | √ |  | 
| network_count | int | X | √ |  | 
| open_issues | int | X | √ |  | 
| has_wiki | bool | X | √ |  | 
| delete_branch_on_merge | bool | X | √ |  | 
| private | bool | X | √ |  | 
| svn_url | string | X | √ |  | 
| allow_forking | bool | X | √ |  | 
| tags_url | string | X | √ |  | 
| homepage | string | X | √ |  | 
| contributors_url | string | X | √ |  | 
| disabled | bool | X | √ |  | 
| contents_url | string | X | √ |  | 
| id | int | X | √ |  | 
| forks_count | int | X | √ |  | 
| deployments_url | string | X | √ |  | 
| squash_merge_commit_title | string | X | √ |  | 
| security_and_analysis | json | X | √ |  | 
| events_url | string | X | √ |  | 
| owner | json | X | √ |  | 
| code_of_conduct | json | X | √ |  | 
| language | string | X | √ |  | 
| watchers_count | int | X | √ |  | 
| subscribers_url | string | X | √ |  | 
| use_squash_pr_title_as_default | bool | X | √ |  | 
| comments_url | string | X | √ |  | 
| downloads_url | string | X | √ |  | 
| team_id | int | X | √ |  | 
| clone_url | string | X | √ |  | 
| auto_init | bool | X | √ |  | 
| allow_squash_merge | bool | X | √ |  | 
| blobs_url | string | X | √ |  | 
| visibility | string | X | √ |  | 
| has_pages | bool | X | √ |  | 
| license_template | string | X | √ |  | 
| pushed_at | timestamp | X | √ |  | 
| git_url | string | X | √ |  | 
| source | json | X | √ |  | 
| issue_events_url | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| license | json | X | √ |  | 
| keys_url | string | X | √ |  | 
| git_tags_url | string | X | √ |  | 
| watchers | int | X | √ |  | 
| allow_merge_commit | bool | X | √ |  | 
| hooks_url | string | X | √ |  | 
| trees_url | string | X | √ |  | 
| releases_url | string | X | √ |  | 
| subscribers_count | int | X | √ |  | 
| squash_merge_commit_message | string | X | √ |  | 
| has_projects | bool | X | √ |  | 
| git_commits_url | string | X | √ |  | 
| pulls_url | string | X | √ |  | 
| default_branch | string | X | √ |  | 
| archive_url | string | X | √ |  | 
| merges_url | string | X | √ |  | 
| stargazers_url | string | X | √ |  | 
| role_name | string | X | √ |  | 
| node_id | string | X | √ |  | 
| text_matches | json | X | √ |  | 
| url | string | X | √ |  | 
| git_refs_url | string | X | √ |  | 


