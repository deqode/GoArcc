import { gql } from '@apollo/client';

export const  GET_REPOSITORIES=gql`
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


export const GET_BRANCHES=gql`
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