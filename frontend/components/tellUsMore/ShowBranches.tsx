import { useQuery } from '@apollo/client'
import { CircularProgress, FormControl, InputLabel, MenuItem, Select } from '@material-ui/core'
import React, { Dispatch, ReactElement, SetStateAction, useEffect, useState } from 'react'
import { GET_BRANCHES } from '../../GraphQL/Query'

export default function ShowBranches({
  ownerName,
  repoName,
  accountId,
  provider,
  setBranchName,
  setCloneUrl,
}: ShowBranchesInput): ReactElement {
  const [branches, setbranches] = useState([])

  const { loading, error, data } = useQuery(GET_BRANCHES, {
    variables: {
      ownerName,
      repoName,
      accountId,
      provider,
    },
  })

  useEffect(() => {
    // if(error)
    // TODO:pop error
    if (data && data.repository) {
      setbranches(data.repository.branches || [])
      setCloneUrl(data.repository.clone_url)
    }
  }, [data, error])

  const selectBranch = (
    e: React.ChangeEvent<{
      name?: string | undefined
      value: unknown
    }>
  ): void => {
    if (typeof e.target.value === 'string') setBranchName(e.target.value)
    else setBranchName('')
  }
  return (
    <FormControl variant="outlined">
      <InputLabel id="demo-simple-select-outlined-label">Branch Name</InputLabel>
      <Select
        labelId="demo-simple-select-outlined-label"
        id="demo-simple-select-outlined"
        onChange={selectBranch}
        label="Repo Name"
        disabled={!(branches.length > 0)}>
        {loading && (
          <MenuItem value="">
            <CircularProgress color="secondary" />
          </MenuItem>
        )}
        {branches.map((r: string) => (
          <MenuItem key={r} value={r}>
            {r}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  )
}
interface ShowBranchesInput {
  ownerName: string
  repoName: string
  accountId: string
  provider: string
  setBranchName: Dispatch<SetStateAction<string>>
  setCloneUrl: Dispatch<SetStateAction<string>>
}
