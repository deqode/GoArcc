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

import { GET_REPOSITORIES } from '../../GraphQL/Query'
import { useSelectStyles } from '../../styles/commonStyles'

const ShowRepos = ({
  userId,
  accountId,
  setCurrentRepo,
  provider,
}: {
  userId: string
  accountId: string
  setCurrentRepo: Dispatch<SetStateAction<string>>
  provider: string
}): ReactElement => {
  const { loading, error, data } = useQuery(GET_REPOSITORIES, {
    variables: {
      provider: provider,
      userId,
      accountId,
    },
  })
  interface Repos {
    name: string
  }
  const [repos, setRepos] = useState<Array<Repos>>([])
  const classes = useSelectStyles()

  useEffect(() => {
    // if(error)
    // TODO:pop error
    if (data && data.repositories) setRepos(data.repositories.repositories || [])
  }, [data, error])

  const selectRepo = (
    e: React.ChangeEvent<{
      name?: string | undefined
      value: unknown
    }>
  ): void => {
    setCurrentRepo(e.target.value as string)
  }

  return (
    <Grid item xs={12} md={12}>
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
    </Grid>
  )
}

export default ShowRepos
