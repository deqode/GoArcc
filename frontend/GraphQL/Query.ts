import { gql } from '@apollo/client'

export const GET_REPOSITORIES = gql`
  query repositories($userId: String!, $accountId: String!, $provider: Types_Enum_GitProviders!) {
    repositories(provider: $provider, user_id: $userId, account_id: $accountId) {
      repositories {
        name
      }
    }
  }
`
// TODO : gql file

export const GET_BRANCHES = gql`
  query repository(
    $ownerName: String!
    $repoName: String!
    $accountId: String!
    $provider: Types_Enum_GitProviders!
  ) {
    repository(
      owner_name: $ownerName
      provider: $provider
      repo_name: $repoName
      account_id: $accountId
    ) {
      branches
      clone_url
    }
  }
`

export const GET_OWNER_NAME = gql`
  query vcsConnections($accountId: String!) {
    vcsConnections(account_id: $accountId) {
      vcs_connection {
        user_name
      }
    }
  }
`

export const CLONE_REPOSITORY = gql`
  mutation cloneRepository(
    $provider: Types_Enum_GitProviders!
    $accountId: String!
    $cloneURL: String!
    $branchName: String!
    $ownerName: String!
  ) {
    cloneRepository(
      provider: $provider
      account_id: $accountId
      repository_url: $cloneURL
      branch_name: $branchName
      user_name: $ownerName
    ) {
      run_id
      workflow_id
    }
  }
`

export const CLONNING_STATUS = gql`
  query getCloningStatus($runId: String, $workflowId: String) {
    getCloningStatus(run_id: $runId, workflow_id: $workflowId) {
      status
    }
  }
`
