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
    $accountId:String!
    $cloneURL:String!
    $branchName:String!
    $userName:String!
){
cloneRepository(
    provider:$provider
    account_id:$accountId
    repository_url:$cloneURL
    branch_name:$branchName
    user_name:$userName
  ){
    run_id
    workflow_id
     }
  }

`

export const CLONNING_STATUS=gql`
query
 getCloningStatus(
   $runId: String,
   $workflowId: String
   ){
  getCloningStatus(
    run_id: $runId,
     workflow_id: $workflowId){
    status
  }
}`