import { useMutation } from '@apollo/client'
import { Button } from '@material-ui/core'
import React, { Dispatch, ReactElement, SetStateAction, useEffect } from 'react'
import { CLONE_REPOSITORY } from '../../GraphQL/Query'
export default function CloneRepository({
  cloneURL,
  branchName,
  accountId,
  provider,
  ownerName,
  setCloneData,
}: CloneRepositoryInput): ReactElement {
  const [callCloneRepo, clonedData] = useMutation(CLONE_REPOSITORY)
  const cloneRepo = () => {
    callCloneRepo({
      variables: {
        cloneURL,
        branchName,
        accountId,
        provider,
        ownerName,
      },
    })
  }

  useEffect(() => {
    if (!clonedData.loading && !clonedData.error && clonedData.data) {
      setCloneData({
        runId: clonedData.data.cloneRepository.run_id,
        workflowId: clonedData.data.cloneRepository.workflow_id,
      })
    }
  }, [clonedData])

  return (
    <Button variant="contained" color="primary" onClick={cloneRepo}>
      Go Fetch
    </Button>
  )
}

interface CloneRepositoryInput {
  cloneURL: string
  branchName: string
  accountId: string
  provider: string
  ownerName: string
  setCloneData: Dispatch<
    SetStateAction<{
      runId: string
      workflowId: string
    }>
  >
}
