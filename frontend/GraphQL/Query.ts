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
    }
  }`

export const VCS_CONNECTIONS = gql`
query VCSConnections(
  $accountid:String!
  $provider:Types_Enum_GitProviders!
){
  VCSConnections(
    account_id:$accountid
    provider:$provider,
    ){
    vcs_connection{
            user_name
    }
  }
}`