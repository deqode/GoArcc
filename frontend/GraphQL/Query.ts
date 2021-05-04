import { gql } from '@apollo/client';

export const GET_REPOSITORIES = gql`
query repositories(
    $userid:String!
    $accountid:String!
    $provider:Types_Enum_GitProviders!
    ){
    repositories(
      provider:$provider,
      user_id:$userid,
      account_id:$accountid,
    ){repositories{
      name
    }}
  }
`
// todo : gql file

export const GET_BRANCHES = gql`
query repository(
    $ownerName:String!
    $repoName:String!
    $accountid:String!
    $provider:Types_Enum_GitProviders!
){
    repository(
      owner_name:$ownerName
      provider:$provider,
      repo_name:$repoName
      account_id:$accountid
    ){
      branches
      clone_url
    }
  }`

export const VCS_CONNECTIONS = gql`
query VCSConnections(
  $accountid:String!
){
  VCSConnections(
    account_id:$accountid
    ){
    vcs_connection{
            user_name
    }
  }
}`

export const CLONE_REPOSITORY = gql`
mutation cloneRepository(
  $provider:Types_Enum_GitProviders!
  $accountid:String!
  $repositoryurl:String!
  $branchname:String!
  $username:String!
){
  cloneRepository(
    provider:$provider
    account_id:$accountid
    repository_url:$repositoryurl
    branch_name:$branchname
    user_name:$username
  ){
    run_id
    workflow_id
  }
}`