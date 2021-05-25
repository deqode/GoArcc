import {
  CircularProgress,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
} from '@material-ui/core'
import React, { Dispatch, ReactElement } from 'react'

import { useSelectStyles } from '../../styles/commonStyles'

import useBranchList from './hooks/useBranchList'
import { Action } from './hooks/useGetTellUsMoreState'

interface ShowBranchesInput {
  ownerName: string
  repoName: string
  accountId: string
  provider: string
  set: Dispatch<Action>
}

const ShowBranches = ({
  ownerName,
  repoName,
  accountId,
  provider,
  set,
}: ShowBranchesInput): ReactElement => {
  const { loading, branches, cloneUrl /*error*/ } = useBranchList(
    ownerName,
    repoName,
    accountId,
    provider
  )
  const classes = useSelectStyles()

  const selectBranch = (
    e: React.ChangeEvent<{
      name?: string | undefined
      value: unknown
    }>
  ): void => {
    set({
      type: 'branchName',
      value: e.target.value as string,
    })
    set({
      type: 'cloneUrl',
      value: cloneUrl,
    })
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
            {branches.map((name: string) => (
              <MenuItem key={name} value={name}>
                {name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      )}
    </Grid>
  )
}

export default ShowBranches
