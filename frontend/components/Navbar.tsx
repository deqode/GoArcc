import { AppBar, Button, Grid, Toolbar } from '@material-ui/core'
import { makeStyles } from '@material-ui/styles'

import Link from 'next/link'
import { useRouter } from 'next/router'
import Image from 'next/image'

import React, { ReactElement, useContext } from 'react'
import { destroyUserSession } from '../api/rest/session'
import UserContext from '../contexts/UserContext'

const useStyles = makeStyles({
  root: {
    width: '60px',
  },
})

function Navbar(): ReactElement {
  const { user } = useContext(UserContext)
  const router = useRouter()
  const classes = useStyles()
  return (
    <AppBar position="static" color="transparent">
      <Toolbar>
        <Grid justify="space-between" container>
          <Grid item className={classes.root}>
            <Link href="/">
              <Image src="/assets/logo.png" width={500} height={500} alt="logo" />
            </Link>
          </Grid>
          <Grid item>
            {user.loggedIn ? (
              <Button
                onClick={async () => {
                  const res = await destroyUserSession()
                  if (!res.error) {
                    router.reload()
                  } else {
                    router.push('/error', { query: { message: 'Network Error' } })
                  }
                }}
                color="inherit">
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
