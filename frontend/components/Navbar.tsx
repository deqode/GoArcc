import { useQuery } from '@apollo/client';
import { AppBar, Button, Grid, makeStyles, Toolbar, Typography } from '@material-ui/core'
import Link from 'next/link'
import { useRouter } from 'next/router';

import React, { useContext } from 'react'
import { UserContext } from '../Contexts/UserContext';

function Navbar() {
    const { user ,removeUser} = useContext(UserContext)
    const router = useRouter();

    return (
        <AppBar position="static" color="transparent">
            <Toolbar>
                <Grid justify="space-between" container >
                    <Grid item>
                        <Link href="/">
                            <img src="/assets/logo.png" width="100px" />
                        </Link>
                    </Grid>
                    <Grid item>
                        {user.idToken != "" ?
                            <Button onClick={()=>{removeUser();router.push("/")}} color="inherit">Logout</Button>
                            :""}
                    </Grid>
                </Grid>
            </Toolbar>
        </AppBar >
    )
}

export default Navbar
