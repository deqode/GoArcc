import { Dispatch, SetStateAction, useReducer, useState } from 'react'

interface Ids {
  runId: string
  workflowId: string
}
export const initialState: State = {
  branchName: '',
  cloneUrl: '',
  currentRepo: '',
}

interface State {
  branchName: string
  cloneUrl: string
  currentRepo: string
}

export interface Action {
  type: 'branchName' | 'cloneUrl' | 'currentRepo'
  value: string
}
// TODO:init State
const init = (): State => {
  return initialState
}

const tellUsMoreStateReducer = (state: State, action: Action): State => ({
  ...state,
  [action.type]: action.value,
})
const useGetTellUsMoreState = (): InitTellUSMore => {
  const [state, dispatch] = useReducer(tellUsMoreStateReducer, initialState, init)
  const [cloneData, setCloneData] = useState<Ids>({ runId: '', workflowId: '' })
  return {
    state,
    dispatch,
    cloneData,
    setCloneData,
  }
}

interface InitTellUSMore {
  state: State
  dispatch: Dispatch<Action>
  cloneData: {
    runId: string
    workflowId: string
  }
  setCloneData: Dispatch<SetStateAction<Ids>>
}

export default useGetTellUsMoreState
