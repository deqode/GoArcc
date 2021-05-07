import { AppBar, Button, Grid, Toolbar } from '@material-ui/core'
import Link from 'next/link'
import { useRouter } from 'next/router';

import React, { useContext } from 'react'
import { UserContext } from '../Contexts/UserContext';

function Navbar(): any {
    const { user, } = useContext(UserContext)
    const router = useRouter();
    return (
        <AppBar position="static" color="transparent">
            <Toolbar>
                <Grid justify="space-between" container >
                    <Grid item>
                        <Link href="/">
                            <img src="/assets/logo.png" width="100px" alt="logo" />
                        </Link>
                    </Grid>
                    <Grid item>
                        {user && user.idToken != "" ?
                            <Button onClick={async () => {
                                await fetch(`/api/session/destroy`, {
                                    method: 'POST',
                                    headers: {
                                        'Accept': 'application/json',
                                        'Content-Type': 'application/json'
                                    },
                                })
                                router.push("/")
                            }} color="inherit">Logout</Button>
                            : ""}
                    </Grid>
                </Grid>
            </Toolbar>
        </AppBar >
    )
}

export default Navbar
