import { ApolloError, useLazyQuery } from '@apollo/client'
import { useState, useEffect } from 'react'

import { GET_BRANCHES } from '../../../GraphQL/Query'

interface BranchListReturn {
  loading: boolean
  branches: Array<string>
  cloneUrl: string
  error: ApolloError | undefined
}

const useBranchList = (
  ownerName: string,
  repoName: string,
  accountId: string,
  provider: string
): BranchListReturn => {
  const [branches, setBranches] = useState<Array<string>>([])
  const [cloneUrl, setCloneUrl] = useState('')

  const [fetch, { loading, error, data }] = useLazyQuery(GET_BRANCHES)

  useEffect(() => {
    // TODO:return error
    if (data && data.repository) {
      setBranches(data.repository.branches || [])
      setCloneUrl(data.repository.repo_url)
    }
  }, [data, error])

  useEffect(() => {
    if (ownerName && repoName && accountId && provider) {
      fetch({
        variables: {
          ownerName,
          repoName,
          accountId,
          provider,
        },
      })
    }
  }, [ownerName, repoName, accountId, provider, fetch])
  //   TODO: customize error

  return { loading, branches, cloneUrl, error }
}

export default useBranchList
