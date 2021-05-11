import { Dispatch, SetStateAction, useState } from 'react'

const useGetTellUsMoreState = (): InitTellUSMore => {
  const [currentRepo, setCurrentRepo] = useState('')
  const [branchName, setBranchName] = useState('')
  const [cloneUrl, setCloneUrl] = useState('')
  const [cloneData, setCloneData] = useState({ runId: '', workflowId: '' })
  return {
    currentRepo,
    setCurrentRepo,
    branchName,
    setBranchName,
    cloneUrl,
    setCloneUrl,
    cloneData,
    setCloneData,
  }
}

interface InitTellUSMore {
  currentRepo: string
  setCurrentRepo: Dispatch<SetStateAction<string>>
  branchName: string
  setBranchName: Dispatch<SetStateAction<string>>
  cloneUrl: string
  setCloneUrl: Dispatch<SetStateAction<string>>
  cloneData: {
    runId: string
    workflowId: string
  }
  setCloneData: Dispatch<
    SetStateAction<{
      runId: string
      workflowId: string
    }>
  >
}

export default useGetTellUsMoreState
