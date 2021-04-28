import { gql } from '@apollo/client';

export const  GET_REPOSITORIES=gql`
query repositories(
    $userid:String!
    $accountid:String!
    ){
    repositories(
      provider:GITHUB,
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
){
    repository(
      owner_name:$ownerName
      provider:GITHUB
      repo_name:$repoName
      account_id:$accountid
    ){
      branches
    }
  }`