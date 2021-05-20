import { Dispatch, SetStateAction, useState } from 'react'

interface Ids {
  runId: string
  workflowId: string
}
const useGetTellUsMoreState = (): InitTellUSMore => {
  const [currentRepo, setCurrentRepo] = useState<string>('')
  const [branchName, setBranchName] = useState<string>('')
  const [cloneUrl, setCloneUrl] = useState<string>('')
  const [cloneData, setCloneData] = useState<Ids>({ runId: '', workflowId: '' })
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
  setCloneData: Dispatch<SetStateAction<Ids>>
}

export default useGetTellUsMoreState
