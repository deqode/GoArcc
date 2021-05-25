import { ApolloError, useLazyQuery } from '@apollo/client'
import { useEffect, useState } from 'react'

import { GET_REPOSITORIES } from '../../../GraphQL/Query'
import { Repos } from '../ShowRepos'

interface RepoListReturn {
  loading: boolean
  repos: Array<Repos>
  error: ApolloError | undefined
}

const useRepoList = (provider: string, userId: string, accountId: string): RepoListReturn => {
  const [repos, setRepos] = useState<Array<Repos>>([])

  const [fetch, { loading, error, data }] = useLazyQuery(GET_REPOSITORIES)
  useEffect(() => {
    if (provider && userId && accountId) {
      fetch({
        variables: {
          provider,
          userId,
          accountId,
        },
      })
    }
  }, [accountId, fetch, provider, userId])

  useEffect(() => {
    if (data && data.repositories) setRepos(data.repositories.repositories || [])
  }, [data])
  //   TODO: customize error
  return {
    error,
    loading,
    repos,
  }
}

export default useRepoList
