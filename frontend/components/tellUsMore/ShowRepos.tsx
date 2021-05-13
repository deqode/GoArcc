import { useQuery } from '@apollo/client'
import { CircularProgress, FormControl, InputLabel, MenuItem, Select } from '@material-ui/core'
import React, { Dispatch, ReactElement, SetStateAction, useEffect, useState } from 'react'
import { GET_REPOSITORIES } from '../../GraphQL/Query'

export default function ShowRepos({
  userId,
  accountId,
  setCurrentRepo,
  provider,
}: {
  userId: string
  accountId: string
  setCurrentRepo: Dispatch<SetStateAction<string>>
  provider: string
}): ReactElement {
  const { loading, error, data } = useQuery(GET_REPOSITORIES, {
    variables: {
      provider: provider,
      userId,
      accountId,
    },
  })

  const [repos, setRepos] = useState([])

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
    if (typeof e.target.value === 'string') setCurrentRepo(e.target.value)
    else setCurrentRepo('')
  }

  return (
    <FormControl variant="outlined">
      <InputLabel id="demo-simple-select-outlined-label">Repo Name</InputLabel>
      <Select
        labelId="demo-simple-select-outlined-label"
        id="demo-simple-select-outlined"
        label="Repo Name"
        disabled={!(repos.length > 0)}
        onChange={selectRepo}
        autoWidth>
        {loading && (
          <MenuItem value="">
            <CircularProgress color="secondary" />
          </MenuItem>
        )}
        {repos.map((r: { name: string }) => (
          <MenuItem key={r.name} value={r.name}>
            {r.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  )
}