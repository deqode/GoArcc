import { AppBar, Button, Grid, Toolbar } from '@material-ui/core'
import Link from 'next/link'
import { useRouter } from 'next/router'

import React, { useContext } from 'react'
import { destroyUserSession } from '../api/rest/session'
import UserContext from '../contexts/UserContext'

function Navbar(): any {
  const { user } = useContext(UserContext)
  const router = useRouter()
  return (
    <AppBar position="static" color="transparent">
      <Toolbar>
        <Grid justify="space-between" container>
          <Grid item>
            <Link href="/">
              <img src="/assets/logo.png" width="100px" alt="logo" />
            </Link>
          </Grid>
          <Grid item>
            {user.loggedIn ? (
              <Button
                onClick={async () => {
                  const res = await destroyUserSession()
                  if (!res.error) {
                    router.reload()
                  }
                  // else{}
                  //  todo:redirect to errorpage
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
