import { AppBar, Button, Grid, Toolbar } from '@material-ui/core'
import { makeStyles } from '@material-ui/styles'
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/router'
import React, { ReactElement, useContext } from 'react'

import { destroyUserSession } from '../api/rest/session'
import UserContext from '../contexts/UserContext'

const useStyles = makeStyles({
  logo: {
    width: '60px',
  },
})

function Navbar(): ReactElement {
  const { user } = useContext(UserContext)
  const router = useRouter()
  const classes = useStyles()

  const logout = async (): Promise<void> => {
    const res = await destroyUserSession()
    if (!res.error) {
      router.reload()
      return
    }
    router.push('/error', { query: { message: 'Network Error' } })
  }
  return (
    <AppBar position="static" color="transparent" elevation={0}>
      <Toolbar>
        <Grid justify="space-between" container>
          <Grid item className={classes.logo}>
            <Link href="/">
              <Image src="/assets/logo.png" width={500} height={300} alt="logo" />
            </Link>
          </Grid>
          <Grid item>
            {user.loggedIn ? (
              <Button color="secondary" variant="contained" onClick={logout}>
                Logout
              </Button>
            ) : (
              ''
            )}
          </Grid>
        </Grid>
      </Toolbar>
    </AppBar>
  )
}

export default Navbar
