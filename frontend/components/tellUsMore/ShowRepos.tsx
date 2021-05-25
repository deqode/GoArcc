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

import { Action } from './hooks/useGetTellUsMoreState'
import useRepoList from './hooks/useRepoList'

export interface Repos {
  name: string
}
const ShowRepos = ({
  userId,
  accountId,
  setCurrentRepo,
  provider,
}: {
  userId: string
  accountId: string
  setCurrentRepo: Dispatch<Action>
  provider: string
}): ReactElement => {
  const { loading, repos } = useRepoList(provider, userId, accountId)

  const classes = useSelectStyles()

  const selectRepo = (
    e: React.ChangeEvent<{
      name?: string | undefined
      value: unknown
    }>
  ): void => {
    setCurrentRepo({
      type: 'currentRepo',
      value: e.target.value as string,
    })
  }

  return (
    <Grid item xs={12} md={12}>
      {loading ? (
        <CircularProgress />
      ) : (
        <FormControl variant="outlined">
          <InputLabel id="demo-simple-select-outlined-label">Repo Name</InputLabel>
          <Select
            className={classes.root}
            labelId="demo-simple-select-outlined-label"
            id="demo-simple-select-outlined"
            label="Repo Name"
            disabled={!(repos.length > 0)}
            onChange={selectRepo}>
            {loading && (
              <MenuItem value="">
                <CircularProgress color="secondary" />
              </MenuItem>
            )}
            {repos.map((repo: Repos) => (
              <MenuItem key={repo.name} value={repo.name}>
                {repo.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      )}
    </Grid>
  )
}

export default ShowRepos
