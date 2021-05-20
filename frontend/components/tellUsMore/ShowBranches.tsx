import { useQuery } from '@apollo/client'
import {
  CircularProgress,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
} from '@material-ui/core'
import React, { Dispatch, ReactElement, SetStateAction, useEffect, useState } from 'react'
import { GET_BRANCHES } from '../../GraphQL/Query'
import { useSelectStyles } from '../../styles/commonStyles'

export default function ShowBranches({
  ownerName,
  repoName,
  accountId,
  provider,
  setBranchName,
  setCloneUrl,
}: ShowBranchesInput): ReactElement {
  const [branches, setbranches] = useState([])
  const classes = useSelectStyles()

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
      setCloneUrl(data.repository.repo_url)
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
    <Grid item xs={12} md={12}>
      {loading ? (
        <CircularProgress />
      ) : (
        <FormControl variant="outlined">
          <InputLabel id="demo-simple-select-outlined-label">Branch Name</InputLabel>
          <Select
            className={classes.root}
            labelId="demo-simple-select-outlined-label"
            id="demo-simple-select-outlined"
            onChange={selectBranch}
            label="Repo Name"
            disabled={!(branches.length > 0)}>
            {branches.map((r: string) => (
              <MenuItem key={r} value={r}>
                {r}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      )}
    </Grid>
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
